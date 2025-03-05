package dao

import (
	"admin-v1/app/helpers"
	"admin-v1/app/models/db"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"errors"
	"os"
)

func CreateProductTypeExec(req *requests.Loai_san_pham_create, res *responses.Loai_san_pham_create) error {
	if result := helpers.GormDB.Debug().
		Table("loai_san_pham").
		Where("ten = ?", req.Ten).
		First(&res.Loai_san_pham);
	result.RowsAffected > 0 {
		return errors.New("loai san pham da ton tai")
	}
	
	var loai_san_pham = db.Loai_san_pham{
		Ten: req.Ten,
		Hinh_anh: req.Hinh_anh.Filename,
	}

	if err := helpers.GormDB.Debug().Create(&loai_san_pham).Error; err != nil {
		return errors.New("khong the tao loai san pham: " + err.Error())
	}

	res.Loai_san_pham = loai_san_pham

	return nil
}

func UpdateProductTypeExec(req *requests.Loai_san_pham_update) error {
	var loai_san_pham db.Loai_san_pham
	var loai_san_pham_temp db.Loai_san_pham

	tx := helpers.GormDB.Begin()

	if result := tx.Debug().
		Table("loai_san_pham").
		Where("ten = ?", req.Ten).
		Where("ID != ?", req.Id).
		First(&loai_san_pham_temp);
	result.RowsAffected > 0 {
		return errors.New("ten loai san pham da ton tai")
	}

	if result := tx.Debug().
		Table("loai_san_pham").
		Where("id = ?", req.Id).
		First(&loai_san_pham);
	result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("loai san pham khong ton tai")
	}

	if req.Hinh_anh != nil {
		filePath := "public/images/" + loai_san_pham.Hinh_anh
	
		if _, err := os.Stat(filePath); !os.IsNotExist(err) {
			err := os.Remove(filePath)
			if err != nil {
				return errors.New("loi khi xoa file")
			} 
		}

		loai_san_pham.Hinh_anh = req.Hinh_anh.Filename
	}

	loai_san_pham.Ten = req.Ten

	if err := tx.Model(&loai_san_pham).Debug().Updates(&loai_san_pham).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the cap nhat loai san pham: " + err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	return nil
}

func DeleteProductTypeExec(req *requests.Loai_san_pham_delete) error {
	var loai_san_pham db.Loai_san_pham

	if result := helpers.GormDB.Debug().
		Table("loai_san_pham").
		Where("id = ?", req.Id).
		First(&loai_san_pham);
	result.RowsAffected == 0 {
		return errors.New("loai san pham khong ton tai")
	}

	if err := helpers.GormDB.Model(&loai_san_pham).Debug().Update("deleted_at", helpers.GetCurrentTimeVN().String()).Error; err != nil {
		return errors.New("khong the xoa loai san pham: " + err.Error())
	}

	return nil
}