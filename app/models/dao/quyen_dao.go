package dao

import (
	"admin-v1/app/helpers"
	"admin-v1/app/models/db"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"errors"
	"strings"

	"gorm.io/gorm"
)

func GetPermissionExec(req *requests.Quyen_read ,res *responses.Quyen_read) error {
	var chuc_nang []db.Chuc_nang

	if err := helpers.GormDB.
		Table("chuc_nang").
		Select("chuc_nang.*, CASE WHEN quyen.chuc_nang_id IS NOT NULL THEN true ELSE false END as thuoc_chuc_vu").
		Joins("LEFT JOIN quyen ON chuc_nang.id = quyen.chuc_nang_id AND quyen.chuc_vu_id = ?", req.Chuc_vu_id).
		Find(&chuc_nang).Error; 
	err != nil {
		return errors.New("khong co chuc nang nao")
	}

	group := make(map[string][]responses.Quyen_sub)

	for _, item := range chuc_nang {
		group[item.Show_in_menu + "+" + item.Loai] = append(group[item.Show_in_menu+"+"+item.Loai], responses.Quyen_sub{
			Id: item.ID,
			Ten: item.Ten,
			Code: item.Code,
			Trang_thai: item.Thuoc_chuc_vu,
		})
	}

	for key, value := range group {
		res.Quyen = append(res.Quyen, responses.Quyen{
			Hien_thi_menu: strings.Split(key, "+")[0],
			Loai_quyen: strings.Split(key, "+")[1],
			Quyen: value,
		})
	}

	return nil
}

func ModifyPermissionExec(req *requests.Quyen_modify) error {
	
	var caseWhenClauses []interface{}
	var ids []interface{}

	for _, quyen := range req.Quyen {
		if quyen.Active == 1 {
			caseWhenClauses = append(caseWhenClauses, quyen.Quyen_id, nil)
		} else {
			caseWhenClauses = append(caseWhenClauses, quyen.Quyen_id, helpers.GetCurrentTimeVN().String())
		}

		ids = append(ids, quyen.Quyen_id)
	}

	if err := helpers.GormDB.Debug().
		Model(&db.Quyen{}).
		Where("chuc_vu_id = ?", req.Chuc_vu_id).
		Where("chuc_nang_id IN ?", ids).
		Updates(map[string]interface{}{
			"deleted_at": gorm.Expr(
				`CASE
					WHEN chuc_nang_id = ? AND active = 1 THEN NULL
					WHEN chuc_nang_id = ? AND active = 0 THEN ?
					ELSE deleted_at
				END`, 
				caseWhenClauses...
			),
		},
	).Error; err != nil {
		return errors.New("khong the thuc hien chinh sua quyen")
	}
	

	return nil
}