package dao

import (
	"admin-v1/app/helpers"
	"admin-v1/app/models/db"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"errors"
)

func CreateImportInvoice(req *requests.Hoa_don_nhap_kho_create, res *responses.Hoa_don_nhap_kho_create) error {
	var chi_tiet_hoa_don_nhap_kho []db.Chi_tiet_hoa_don_nhap_kho

	tx := helpers.GormDB.Begin()

	if len(req.Chi_tiet_hoa_don_nhap_kho) > 0 {
		var date = helpers.GetCurrentTimeVN().Format("2006-01-02")

		for _, value := range req.Chi_tiet_hoa_don_nhap_kho {
			var counter int64

			if err := tx.Table("chi_tiet_hoa_don_nhap_kho").
				Where("ctsp_id = ?", value.Ctsp_id).
				Where("DATE(created_at) = ?", date).
				Count(&counter).Error;
			err != nil {
				return errors.New("khong the tim chi tiet hoa don nhap kho: " + err.Error())
			}

			chi_tiet_hoa_don_nhap_kho = append(chi_tiet_hoa_don_nhap_kho, db.Chi_tiet_hoa_don_nhap_kho{
				Hoa_don_id: value.Hoa_don_id,
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
	}

	var hoa_don_nhap_kho = db.Hoa_don_nhap_kho{
		Nha_phan_phoi_id: req.Nha_phan_phoi_id,	
		Kho_id: req.Kho_id,
		Ngay_nhap: req.Ngay_nhap,
		Tong_tien: req.Tong_tien,

		Chi_tiet_hoa_don_nhap_kho: chi_tiet_hoa_don_nhap_kho,
	}

	if err := tx.Debug().Create(&hoa_don_nhap_kho).Error; err != nil {
		return errors.New("khong the tao hoa don nhap kho: " + err.Error())
	}

	if len(hoa_don_nhap_kho.Chi_tiet_hoa_don_nhap_kho) > 0 {
		var ids_Ctsp []uint
		var chi_tiet_san_pham []db.Chi_tiet_san_pham

		for _, value := range hoa_don_nhap_kho.Chi_tiet_hoa_don_nhap_kho {
			ids_Ctsp = append(ids_Ctsp, value.ID)
		}

		if err := tx.Debug().Table("chi_tiet_san_pham").
			Where("ID IN ?", ids_Ctsp).
			Find(&chi_tiet_san_pham).Error; 
		err != nil {
			return errors.New("khong the tim thay chi tiet san pham")
		}

		for idx, value := range hoa_don_nhap_kho.Chi_tiet_hoa_don_nhap_kho {
			chi_tiet_san_pham[idx].So_luong = chi_tiet_san_pham[idx].So_luong + value.So_luong
			chi_tiet_san_pham[idx].Gia_nhap = value.Gia_nhap * ((100 - value.Chiet_khau) / 100)
			chi_tiet_san_pham[idx].Gia_ban = value.Gia_ban
		}

		if err := tx.Debug().Table("chi_tiet_san_pham").Updates(&chi_tiet_san_pham).Error; err != nil {
			return errors.New("khong the cap nhat chi tiet san pham")
		}
	}

	res.Hoa_don_nhap_kho = hoa_don_nhap_kho

	return nil
}