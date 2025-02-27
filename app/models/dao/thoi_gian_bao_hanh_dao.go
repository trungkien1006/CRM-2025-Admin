package dao

import (
	"admin-v1/app/helpers"
	"admin-v1/app/models/db"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"errors"
)

func CreateWarrantyTimeExec(req *requests.Thoi_gian_bao_hanh_create, res *responses.Thoi_gian_bao_hanh_create) error {
	if result := helpers.GormDB.Debug().
		Table("thoi_gian_bao_hanh").
		Where("ten = ?", req.Ten).
		First(&res.Thoi_gian_bao_hanh);
	result.RowsAffected > 0 {
		return errors.New("thoi gian bao hanh da ton tai")
	}
	
	var thoi_gian_bao_hanh = db.Thoi_gian_bao_hanh{
		Ten: req.Ten,
	}

	if err := helpers.GormDB.Debug().Create(&thoi_gian_bao_hanh).Error; err != nil {
		return errors.New("khong the tao thoi gian bao hanh: " + err.Error())
	}

	res.Thoi_gian_bao_hanh = thoi_gian_bao_hanh

	return nil
}

func UpdateWarrantyTimeExec(req *requests.Thoi_gian_bao_hanh_update) error {
	var thoi_gian_bao_hanh db.Thoi_gian_bao_hanh

	tx := helpers.GormDB.Begin()

	if result := tx.Debug().
		Table("thoi_gian_bao_hanh").
		Where("id = ?", req.Id).
		Where("deleted_at IS NULL").
		First(&thoi_gian_bao_hanh);
	result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("thoi gian bao hanh khong ton tai")
	}

	thoi_gian_bao_hanh.Ten = req.Ten

	if err := tx.Model(&thoi_gian_bao_hanh).Debug().Update("ten", thoi_gian_bao_hanh.Ten).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the cap nhat thoi gian bao hanh: " + err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	return nil
}

func DeleteWarrantyTimeExec(req *requests.Thoi_gian_bao_hanh_delete) error {
	var thoi_gian_bao_hanh db.Thoi_gian_bao_hanh

	if result := helpers.GormDB.Debug().
		Table("thoi_gian_bao_hanh").
		Where("id = ?", req.Id).
		Where("deleted_at IS NULL").
		First(&thoi_gian_bao_hanh);
	result.RowsAffected == 0 {
		return errors.New("thoi gian bao hanh khong ton tai")
	}

	if err := helpers.GormDB.Model(&thoi_gian_bao_hanh).Debug().Update("deleted_at", helpers.GetCurrentTimeVN().String()).Error; err != nil {
		return errors.New("khong the xoa thoi gian bao hanh: " + err.Error())
	}

	return nil
}