package dao

import (
	"admin-v1/app/enums/datetime"
	"admin-v1/app/helpers"
	"admin-v1/app/models/db"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"errors"

	"gorm.io/gorm"
)

func CreateExportInvoice(req *requests.Hoa_don_xuat_kho_create, res *responses.Hoa_don_xuat_kho_create) error {
	var chi_tiet_hoa_don_xuat_kho []db.Chi_tiet_hoa_don_xuat_kho

	//bat dau transaction
	tx := helpers.GormDB.Begin()

	var sp_ids 		[]int
	var ctsp_ids	[]int

	//tao 1 danh sach doi tuong cthdnk neu co danh sach tu client truyen len
	if len(req.Chi_tiet_hoa_don_xuat_kho) > 0 {
		for _, value := range req.Chi_tiet_hoa_don_xuat_kho {
			sp_ids = append(sp_ids, value.San_pham_id)
			ctsp_ids = append(ctsp_ids, value.Ctsp_id)

			//them tung cthdxk vao ds
			chi_tiet_hoa_don_xuat_kho = append(chi_tiet_hoa_don_xuat_kho, db.Chi_tiet_hoa_don_xuat_kho{	
				Ctsp_id: uint(value.Ctsp_id),
				Sku: value.Sku,
				Don_vi_tinh: value.Don_vi_tinh,	
				So_luong_ban: value.So_luong_ban,
				Gia_ban: value.Gia_ban,
				Chiet_khau: value.Chiet_khau,
				Thanh_tien: value.Thanh_tien,
				Gia_nhap: value.Gia_nhap,
				Loi_nhuan: value.Loi_nhuan,
				La_qua_tang: value.La_qua_tang,
			})
		}

		//kiem tra cac id cua san pham va ctsp co ton tai hay khong
		var count int64 = 0

		if err := tx.Debug().Table("san_pham").Where("id IN ?", sp_ids).Count(&count).Error; err != nil {
			tx.Rollback()
			return errors.New("loi khi kiem tra san pham: " + err.Error())
		}

		if count != int64(len(sp_ids)) {
			tx.Rollback()
			return errors.New("co id san pham khong ton tai")
		}

		count = 0

		if err := tx.Debug().Table("chi_tiet_san_pham").Where("id IN ?", ctsp_ids).Count(&count).Error; err != nil {
			tx.Rollback()
			return errors.New("loi khi kiem tra chi tiet san pham: " + err.Error())
		}

		if count != int64(len(ctsp_ids)) {
			tx.Rollback()
			return errors.New("co id chi tiet san pham khong ton tai")
		}
	}

	var hdxk_count_in_year int64 = 0

	if err := tx.Table("hoa_don_xuat_kho").
		Where("YEAR(created_at) = ?", helpers.GetCurrentTimeVN().Format(datetime.Year)).
		Count(&hdxk_count_in_year).Error;
	err != nil {
		tx.Rollback()
		return errors.New("loi khi tinh toan so luong hoa don trong nam: " + err.Error())
	}

	//tao doi tuong hdxk
	var hoa_don_xuat_kho = db.Hoa_don_xuat_kho{
		So_hoa_don: int(hdxk_count_in_year),					
		Ma_hoa_don: "HDX" + string(hdxk_count_in_year),
		Khach_hang_id: req.Khach_hang_id,
		Nhan_vien_sale_id: req.Nhan_vien_sale_id,
		Nhan_vien_giao_hang_id: req.Nhan_vien_giao_hang_id,
		Ngay_xuat: req.Ngay_xuat,
		Tong_tien: req.Tong_tien,	
		Vat: req.Vat,	
		Thanh_tien: req.Thanh_tien,				
		Tra_truoc: req.Tra_truoc,
		Con_lai: req.Tong_tien - req.Tra_truoc,
		Tong_gia_nhap: req.Tong_gia_nhap,		
		Loi_nhuan: req.Loi_nhuan,
		Ghi_chu: req.Ghi_chu,
		Da_giao_hang: req.Da_giao_hang,		
		Loai_chiet_khau: req.Loai_chiet_khau,
		Gia_tri_chiet_khau: req.Gia_tri_chiet_khau,

		Chi_tiet_hoa_don_xuat_kho: chi_tiet_hoa_don_xuat_kho,
	}

	//insert hdnk
	if err := tx.Debug().Session(&gorm.Session{FullSaveAssociations: true}).Create(&hoa_don_xuat_kho).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the tao hoa don xuat kho: " + err.Error())
	}

	//cap nhat lai chi tiet san pham: so luong, gia nhap, gia ban
	if len(hoa_don_xuat_kho.Chi_tiet_hoa_don_xuat_kho) > 0 {
		var ids_Ctsp []uint
		var chi_tiet_san_pham []db.Chi_tiet_san_pham

		for _, value := range hoa_don_xuat_kho.Chi_tiet_hoa_don_xuat_kho {
			ids_Ctsp = append(ids_Ctsp, value.ID)
		}

		//lay ra ds ctsp theo ds id
		if err := tx.Debug().Table("chi_tiet_san_pham").
			Where("ID IN ?", ids_Ctsp).
			Find(&chi_tiet_san_pham).Error; 
		err != nil {
			tx.Rollback()
			return errors.New("khong the tim thay chi tiet san pham: " + err.Error())
		}

		//loc qua ds ctsp, cap nhat lai thong tin, tao ds ton kho
		for idx, value := range hoa_don_xuat_kho.Chi_tiet_hoa_don_xuat_kho {
			chi_tiet_san_pham[idx].So_luong = chi_tiet_san_pham[idx].So_luong - value.So_luong_ban
			// chi_tiet_san_pham[idx].Gia_nhap = value.Gia_nhap * ((100 - value.Chiet_khau) / 100)
			// chi_tiet_san_pham[idx].Gia_ban = value.Gia_ban

			//cap nhat so luong ton kho
			var ton_kho db.Ton_kho

			if err := tx.Table("ton_kho").
				Where("sku = ?", value.Sku).
				Find(&ton_kho).Error;
			err != nil {
				tx.Rollback()
				return errors.New("co loi khi truy xuat ton kho theo sku: " + err.Error())
			}

			ton_kho.So_luong_ton = ton_kho.So_luong_ton - value.So_luong_ban

			if err := tx.Save(&ton_kho).Error; err != nil {
				tx.Rollback()
				return errors.New("co loi khi luu ton kho da cap nhat: " + err.Error())
			}
		}

		//update ctsp: bulk update
		if err := tx.Debug().Table("chi_tiet_san_pham").Updates(&chi_tiet_san_pham).Error; err != nil {
			tx.Rollback()
			return errors.New("khong the cap nhat chi tiet san pham: " + err.Error())
		}
	}

	//commit transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	res.Hoa_don_xuat_kho = hoa_don_xuat_kho

	return nil
}