package dao

import (
	"admin-v1/app/helpers"
	"admin-v1/app/models/db"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"errors"
)

func CreateRoleExec(req *requests.Chuc_vu_create, res *responses.Chuc_vu_create) error {
	if result := helpers.GormDB.Debug().
		Table("chuc_vu").
		Where("ten = ?", req.Ten).
		First(&res.Chuc_vu);
	result.RowsAffected > 0 {
		return errors.New("chuc vu da ton tai")
	}
	
	var chuc_vu = db.Chuc_vu{
		Ten: req.Ten,
	}

	if err := helpers.GormDB.Debug().Create(&chuc_vu).Error; err != nil {
		return errors.New("khong the tao chuc vu: " + err.Error())
	}

	res.Chuc_vu = chuc_vu

	return nil
}

func UpdateRoleExec(req *requests.Chuc_vu_update) error {
	var chuc_vu db.Chuc_vu
	var chuc_vu_temp db.Chuc_vu

	tx := helpers.GormDB.Begin()

	if result := tx.Debug().
		Table("chuc_vu").
		Where("ten = ?", req.Ten).
		First(&chuc_vu_temp);
	result.RowsAffected > 0 {
		return errors.New("ten chuc vu da ton tai")
	}

	if result := tx.Debug().
		Table("chuc_vu").
		Where("id = ?", req.Id).
		First(&chuc_vu);
	result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("chuc vu khong ton tai")
	}

	chuc_vu.Ten = req.Ten

	if err := tx.Model(&chuc_vu).Debug().Updates(&chuc_vu).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the cap nhat chuc vu: " + err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	return nil
}

func DeleteRoleExec(req *requests.Chuc_vu_delete) error {
	var chuc_vu db.Chuc_vu

	if result := helpers.GormDB.Debug().
		Table("chuc_vu").
		Where("id = ?", req.Id).
		First(&chuc_vu);
	result.RowsAffected == 0 {
		return errors.New("chuc vu khong ton tai")
	}

	if err := helpers.GormDB.Model(&chuc_vu).Debug().Update("deleted_at", helpers.GetCurrentTimeVN().String()).Error; err != nil {
		return errors.New("khong the xoa chuc vu: " + err.Error())
	}

	return nil
}