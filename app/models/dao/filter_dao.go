package dao

import (
	"admin-v1/app/helpers"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"encoding/json"
	"math"
)

func FilterExec[T any](req *requests.Filter, res *responses.Filter[T], tableName string) error {
	query := helpers.GormDB.Debug().
		Table(tableName)

	if(tableName == "san_pham"){
		query.Joins("JOIN loai_san_pham ON loai_san_pham.id = san_pham.loai_san_pham_id").
			Joins("JOIN don_vi_tinh ON don_vi_tinh.id = san_pham.don_vi_tinh_id").
			Joins("JOIN thoi_gian_bao_hanh ON thoi_gian_bao_hanh.id = san_pham.thoi_gian_bao_hanh_id").
			Joins("JOIN loai_giam_gia ON loai_giam_gia.id = san_pham.loai_giam_gia_id")
	}

	var filters []requests.FilterStruc

	err := json.Unmarshal([]byte(req.Filters), &filters)
	if err != nil {
		return err
	}

	helpers.Filter(query, filters)

	var totalRecord int64

	if err := query.Count(&totalRecord).Error; err != nil {
		return err
	}

	query.Offset((req.Page - 1) * req.Limit).Limit(req.Limit)

	if(tableName == "san_pham"){
		query.Preload("Loai_san_pham").
			Preload("Don_vi_tinh").
			Preload("Thoi_gian_bao_hanh").
			Preload("Loai_giam_gia").
			Preload("Chi_tiet_san_pham")
	}

	if req.Sort != "" {
		query.Order(req.Sort + " " + req.Order)
	}

	if err := query.Find(&res.Data).Error; err != nil {
		return err
	}

	res.Total_Page = int(math.Ceil(float64(totalRecord) / float64(req.Limit)))

	return nil
}