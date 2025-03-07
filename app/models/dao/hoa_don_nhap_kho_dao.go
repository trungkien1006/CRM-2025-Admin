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

func CreateImportInvoice(req *requests.Hoa_don_nhap_kho_create, res *responses.Hoa_don_nhap_kho_create) error {
	var chi_tiet_hoa_don_nhap_kho []db.Chi_tiet_hoa_don_nhap_kho

	//bat dau transaction
	tx := helpers.GormDB.Begin()

	//tao 1 danh sach doi tuong cthdnk neu co danh sach tu client truyen len
	if len(req.Chi_tiet_hoa_don_nhap_kho) > 0 {
		//tao tham so date 1 phan cua sku
		var date = helpers.GetCurrentTimeVN().Format(datetime.Date)

		//tao 2 ds chua cac id de kiem tra ton tai hay chua
		var ds_san_pham_id 	[]int64
		var ds_ctsp_id		[]int64

		for _, value := range req.Chi_tiet_hoa_don_nhap_kho {
			//tao tham so counter 1 phan cua sku
			var counter int64

			if err := tx.Table("chi_tiet_hoa_don_nhap_kho").
				Where("ctsp_id = ?", value.Ctsp_id).
				Where("DATE(created_at) = ?", date).
				Count(&counter).Error;
			err != nil {
				return errors.New("khong the tim chi tiet hoa don nhap kho: " + err.Error())
			}

			//them tung cthdnk vao ds
			chi_tiet_hoa_don_nhap_kho = append(chi_tiet_hoa_don_nhap_kho, db.Chi_tiet_hoa_don_nhap_kho{
				San_pham_id: value.San_pham_id,
				Ctsp_id: value.Ctsp_id,
				Sku: helpers.GenerateSKU(value.Upc, value.Ctsp_id, counter),
				So_luong: value.So_luong,
				Don_vi_tinh: value.Don_vi_tinh,
				Ke: value.Ke,
				Gia_nhap: value.Gia_nhap,
				Gia_ban: value.Gia_ban,
				Chiet_khau: value.Chiet_khau,
				Thanh_tien: value.Thanh_tien,
				La_qua_tang: value.La_qua_tang,
			})
		}

		//kiem tra cac id cua san pham va ctsp co ton tai hay khong
		var count int64 = 0

		if err := tx.Debug().Table("san_pham").Where("id IN ?", ds_san_pham_id).Count(&count).Error; err != nil {
			tx.Rollback()
			return errors.New("loi khi kiem tra san pham: " + err.Error())
		}

		if count != int64(len(ds_san_pham_id)) {
			tx.Rollback()
			return errors.New("co id san pham khong ton tai")
		}

		count = 0

		if err := tx.Debug().Table("chi_tiet_san_pham").Where("id IN ?", ds_ctsp_id).Count(&count).Error; err != nil {
			tx.Rollback()
			return errors.New("loi khi kiem tra chi tiet san pham: " + err.Error())
		}

		if count != int64(len(ds_ctsp_id)) {
			tx.Rollback()
			return errors.New("co id chi tiet san pham khong ton tai")
		}
	}

	//tao doi tuong hdnk
	var hoa_don_nhap_kho = db.Hoa_don_nhap_kho{
		Nha_phan_phoi_id: req.Nha_phan_phoi_id,	
		Kho_id: req.Kho_id,
		Ngay_nhap: req.Ngay_nhap,
		Tong_tien: req.Tong_tien,

		Chi_tiet_hoa_don_nhap_kho: chi_tiet_hoa_don_nhap_kho,
	}

	//insert hdnk
	if err := tx.Debug().Session(&gorm.Session{FullSaveAssociations: true}).Create(&hoa_don_nhap_kho).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the tao hoa don nhap kho: " + err.Error())
	}

	//cap nhat lai chi tiet san pham: so luong, gia nhap, gia ban
	if len(hoa_don_nhap_kho.Chi_tiet_hoa_don_nhap_kho) > 0 {
		var ids_Ctsp []uint
		var chi_tiet_san_pham []db.Chi_tiet_san_pham

		for _, value := range hoa_don_nhap_kho.Chi_tiet_hoa_don_nhap_kho {
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

		//loc qua ds ctsp va cap nhat lai thong tin
		for idx, value := range hoa_don_nhap_kho.Chi_tiet_hoa_don_nhap_kho {
			chi_tiet_san_pham[idx].So_luong = chi_tiet_san_pham[idx].So_luong + value.So_luong
			chi_tiet_san_pham[idx].Gia_nhap = value.Gia_nhap * ((100 - value.Chiet_khau) / 100)
			chi_tiet_san_pham[idx].Gia_ban = value.Gia_ban
		}

		//update ctsp: bulk update
		if err := tx.Debug().Table("chi_tiet_san_pham").Updates(&chi_tiet_san_pham).Error; err != nil {
			tx.Rollback()
			return errors.New("khong the cap nhat chi tiet san pham")
		}
	}

	//commit transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	res.Hoa_don_nhap_kho = hoa_don_nhap_kho

	return nil
}