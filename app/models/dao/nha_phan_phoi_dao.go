package dao

import (
	"admin-v1/app/helpers"
	"admin-v1/app/models/db"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"errors"
)

func CreateProviderExec(req *requests.Nha_phan_phoi_create, res *responses.Nha_phan_phoi_create) error {
	if result := helpers.GormDB.Debug().
		Table("nha_phan_phoi").
		Where("ten = ?", req.Ten).
		First(&res.Nha_phan_phoi);
	result.RowsAffected > 0 {
		return errors.New("nha phan phoi da ton tai")
	}
	
	var ds_san_pham []db.San_pham

	if result := helpers.GormDB.Debug().
		Table("san_pham").
		Where("id IN ?", req.San_pham_id).
		Find(&ds_san_pham).Error; 
	result != nil {
		return errors.New("tim san pham bi loi")
	}

	var nha_phan_phoi = db.Nha_phan_phoi{
		Ten: req.Ten,
		San_pham: ds_san_pham,
	}

	if err := helpers.GormDB.Debug().Create(&nha_phan_phoi).Error; err != nil {
		return errors.New("khong the tao nha phan phoi: " + err.Error())
	}

	res.Nha_phan_phoi = nha_phan_phoi

	return nil
}

func UpdateProviderExec(req *requests.Nha_phan_phoi_update) error {
	var nha_phan_phoi db.Nha_phan_phoi
	var nha_phan_phoi_temp db.Nha_phan_phoi

	tx := helpers.GormDB.Begin()

	if result := tx.Debug().
		Table("nha_phan_phoi").
		Where("ten = ?", req.Ten).
		First(&nha_phan_phoi_temp);
	result.RowsAffected > 0 {
		return errors.New("ten nha phan phoi da ton tai")
	}

	if result := tx.Debug().
		Table("nha_phan_phoi").
		Where("id = ?", req.Id).
		First(&nha_phan_phoi);
	result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("nha phan phoi khong ton tai")
	}

		var ds_san_pham []db.San_pham

	if result := helpers.GormDB.Debug().
		Table("san_pham").
		Where("id IN ?", req.San_pham_id).
		Find(&ds_san_pham).Error; 
	result != nil {
		return errors.New("tim san pham bi loi")
	}

	nha_phan_phoi.Ten = req.Ten
	nha_phan_phoi.San_pham = ds_san_pham

	if err := tx.Model(&nha_phan_phoi).Debug().Association("San_pham").Replace(ds_san_pham); err != nil {
		tx.Rollback()
		return errors.New("khong the cap nhat san pham cua nha phan phoi: " + err.Error())
	}

	if err := tx.Model(&nha_phan_phoi).Debug().Updates(&nha_phan_phoi).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the cap nhat nha phan phoi: " + err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	return nil
}

func DeleteProviderExec(req *requests.Nha_phan_phoi_delete) error {
	var nha_phan_phoi db.Nha_phan_phoi

	if result := helpers.GormDB.Debug().
		Table("nha_phan_phoi").
		Where("id = ?", req.Id).
		First(&nha_phan_phoi);
	result.RowsAffected == 0 {
		return errors.New("nha phan phoi khong ton tai")
	}

	if err := helpers.GormDB.Model(&nha_phan_phoi).Debug().Update("deleted_at", helpers.GetCurrentTimeVN().String()).Error; err != nil {
		return errors.New("khong the xoa nha phan phoi: " + err.Error())
	}

	return nil
}