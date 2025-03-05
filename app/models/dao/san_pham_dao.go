package dao

import (
	"admin-v1/app/helpers"
	"admin-v1/app/models/db"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"errors"
	"os"
)

func CreateProductExec(req *requests.San_pham_create, res *responses.San_pham_create) error {
	if result := helpers.GormDB.Debug().
		Table("san_pham").
		Where("upc = ?", req.Upc).
		First(&res.San_pham);
	result.RowsAffected > 0 {
		return errors.New("loai san pham da ton tai")
	}
	
	var chi_tiet_san_pham_arr []db.Chi_tiet_san_pham

	if len(req.Chi_tiet_san_pham) > 0 {
		for _, value := range req.Chi_tiet_san_pham {
			chi_tiet_san_pham_arr = append(chi_tiet_san_pham_arr, db.Chi_tiet_san_pham{
				Ten_phan_loai: value.Ten_phan_loai,
				Hinh_anh: value.Hinh_anh.Filename,
				Gia_nhap: 0,
				Gia_ban: 0,
				So_luong: 0,
				Trang_thai: value.Trang_thai,
				Khong_phan_loai: 0,
			})
		}
	} else {
		chi_tiet_san_pham_arr = append(chi_tiet_san_pham_arr, db.Chi_tiet_san_pham{
			Ten_phan_loai: "",
			Hinh_anh: "",
			Gia_nhap: 0,
			Gia_ban: 0,
			So_luong: 0,
			Trang_thai: 0,
			Khong_phan_loai: 1,
		})
	}

	var san_pham = db.San_pham{
		Ten: req.Ten,
		Upc: req.Upc,
		Loai_san_pham_id: req.Loai_san_pham_id,
		Hinh_anh: req.Hinh_anh.Filename,
		Don_vi_tinh_id: req.Don_vi_tinh_id,
		Vat: req.Vat,
		Mo_ta: req.Mo_ta,
		Trang_thai: req.Trang_thai,
		Loai_giam_gia_id: req.Loai_giam_gia_id,
		Thoi_gian_bao_hanh_id: req.Thoi_gian_bao_hanh_id,

		Chi_tiet_san_pham: chi_tiet_san_pham_arr,
	}

	if err := helpers.GormDB.Debug().Create(&san_pham).Error; err != nil {
		return errors.New("khong the tao loai san pham: " + err.Error())
	}

	res.San_pham = san_pham

	return nil
}

func UpdateProductExec(req *requests.San_pham_update) error {
	var san_pham db.San_pham
	var san_pham_temp db.San_pham

	tx := helpers.GormDB.Begin()

	if result := tx.Debug().
		Table("san_pham").
		Where("upc = ?", req.Upc).
		First(&san_pham_temp);
	result.RowsAffected > 0 {
		return errors.New("ten san pham da ton tai")
	}

	if result := tx.Debug().
		Table("san_pham").
		Preload("Chi_tiet_san_pham").
		Where("id = ?", req.Id).
		First(&san_pham);
	result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("san pham khong ton tai")
	}

	if req.Hinh_anh != nil {
		filePath := "public/images/" + san_pham.Hinh_anh
	
		if _, err := os.Stat(filePath); !os.IsNotExist(err) {
			err := os.Remove(filePath)
			if err != nil {
				return errors.New("loi khi xoa file")
			} 
		}
	}

	for _, value := range san_pham.Chi_tiet_san_pham {
		if value.Hinh_anh != "" {
			filePath := "public/images/" + value.Hinh_anh

			if _, err := os.Stat(filePath); !os.IsNotExist(err) {
				err := os.Remove(filePath)

				if err != nil {
					return errors.New("loi khi xoa file")
				} 
			}
		}

		if err := tx.Model(&value).Debug().Update("deleted_at", helpers.GetCurrentTimeVN().String()).Error; err != nil {
			return errors.New("khong the xoa san pham: " + err.Error())
		}
	}

	var chi_tiet_san_pham_arr []db.Chi_tiet_san_pham

	if len(req.Chi_tiet_san_pham) > 0 {
		for _, value := range req.Chi_tiet_san_pham {
			chi_tiet_san_pham_arr = append(chi_tiet_san_pham_arr, db.Chi_tiet_san_pham{
				Ten_phan_loai: value.Ten_phan_loai,
				Hinh_anh: value.Hinh_anh.Filename,
				Gia_nhap: 0,
				Gia_ban: 0,
				So_luong: 0,
				Trang_thai: value.Trang_thai,
				Khong_phan_loai: 0,
			})
		}
	} else {
		chi_tiet_san_pham_arr = append(chi_tiet_san_pham_arr, db.Chi_tiet_san_pham{
			Ten_phan_loai: "",
			Hinh_anh: "",
			Gia_nhap: 0,
			Gia_ban: 0,
			So_luong: 0,
			Trang_thai: 0,
			Khong_phan_loai: 1,
		})
	}

	san_pham.Ten = req.Ten
	san_pham.Upc = req.Upc
	san_pham.Loai_san_pham_id = req.Loai_san_pham_id
	san_pham.Hinh_anh = req.Hinh_anh.Filename
	san_pham.Don_vi_tinh_id = req.Don_vi_tinh_id
	san_pham.Vat = req.Vat
	san_pham.Mo_ta = req.Mo_ta
	san_pham.Trang_thai = req.Trang_thai
	san_pham.Loai_giam_gia_id = req.Loai_giam_gia_id
	san_pham.Thoi_gian_bao_hanh_id = req.Thoi_gian_bao_hanh_id

	san_pham.Chi_tiet_san_pham = chi_tiet_san_pham_arr

	if err := tx.Model(&san_pham).Debug().Updates(&san_pham).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the cap nhat san pham: " + err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	return nil
}

func DeleteProductExec(req *requests.San_pham_delete) error {
	var san_pham db.San_pham

	if result := helpers.GormDB.Debug().
		Table("san_pham").
		Where("id = ?", req.Id).
		First(&san_pham);
	result.RowsAffected == 0 {
		return errors.New("san pham khong ton tai")
	}

	if err := helpers.GormDB.Model(&san_pham).Debug().Update("deleted_at", helpers.GetCurrentTimeVN().String()).Error; err != nil {
		return errors.New("khong the xoa san pham: " + err.Error())
	}

	return nil
}