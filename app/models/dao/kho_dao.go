package dao

import (
	"admin-v1/app/helpers"
	"admin-v1/app/models/db"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"errors"
)

func CreateWareHouseExec(req *requests.Kho_create, res *responses.Kho_create) error {
	if result := helpers.GormDB.Debug().
		Table("kho").
		Where("ten = ?", req.Ten).
		First(&res.Kho);
	result.RowsAffected > 0 {
		return errors.New("kho da ton tai")
	}
	
	var kho = db.Kho{
		Ten: req.Ten,
		Dia_chi: req.Dia_chi,
	}

	if err := helpers.GormDB.Debug().Create(&kho).Error; err != nil {
		return errors.New("khong the tao kho: " + err.Error())
	}

	res.Kho = kho

	return nil
}

func UpdateWareHouseExec(req *requests.Kho_update) error {
	var kho db.Kho
	var kho_temp db.Kho

	tx := helpers.GormDB.Begin()

	if result := tx.Debug().
		Table("kho").
		Where("ten = ?", req.Ten).
		First(&kho_temp);
	result.RowsAffected > 0 {
		return errors.New("ten kho da ton tai")
	}

	if result := tx.Debug().
		Table("kho").
		Where("id = ?", req.Id).
		First(&kho);
	result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("kho khong ton tai")
	}

	kho.Ten = req.Ten
	kho.Dia_chi = req.Dia_chi

	if err := tx.Model(&kho).Debug().Updates(&kho).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the cap nhat kho: " + err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	return nil
}

func DeleteWareHouseExec(req *requests.Kho_delete) error {
	var kho db.Kho

	if result := helpers.GormDB.Debug().
		Table("kho").
		Where("id = ?", req.Id).
		First(&kho);
	result.RowsAffected == 0 {
		return errors.New("kho khong ton tai")
	}

	if err := helpers.GormDB.Model(&kho).Debug().Update("deleted_at", helpers.GetCurrentTimeVN().String()).Error; err != nil {
		return errors.New("khong the xoa kho: " + err.Error())
	}

	return nil
}