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

	if tableName == "san_pham" {
		query.Joins("LEFT JOIN loai_san_pham ON loai_san_pham.id = san_pham.loai_san_pham_id").
			Joins("LEFT JOIN don_vi_tinh ON don_vi_tinh.id = san_pham.don_vi_tinh_id").
			Joins("LEFT JOIN thoi_gian_bao_hanh ON thoi_gian_bao_hanh.id = san_pham.thoi_gian_bao_hanh_id").
			Joins("LEFT JOIN loai_giam_gia ON loai_giam_gia.id = san_pham.loai_giam_gia_id")
			// Joins("LEFT JOIN nha_phan_phoi ON nha_phan_phoi.id = san_pham.nha_phan_phoi_id")
	} else if tableName == "nhan_vien" {
		query.Joins("LEFT JOIN chuc_vu ON chuc_vu.id = nhan_vien.chuc_vu_id")
	}else if tableName == "hoa_don_nhap_kho" {
		query.Joins("LEFT JOIN nha_phan_phoi ON nha_phan_phoi.id = hoa_don_nhap_kho.nha_phan_phoi_id").
			Joins("LEFT JOIN kho ON kho.id = hoa_don_nhap_kho.kho_id")
	}

	if req.Filters != "" {
		var filters []requests.FilterStruc

		err := json.Unmarshal([]byte(req.Filters), &filters)
		if err != nil {
			return err
		}

		helpers.Filter(query, filters)
	}

	var totalRecord int64 = 0

	if err := query.Where(tableName + ".deleted_at IS NULL").Count(&totalRecord).Error; err != nil {
		return err
	}

	if req.Limit != 0 && req.Page != 0 {
		query.Offset((req.Page - 1) * req.Limit).Limit(req.Limit)
	}

	if(tableName == "san_pham"){
		query.Select("san_pham.*, loai_san_pham.ten as loai_san_pham, don_vi_tinh.ten as don_vi_tinh, loai_giam_gia.ten as loai_giam_gia, thoi_gian_bao_hanh.ten as thoi_gian_bao_hanh, loai_san_pham.id as loai_san_pham_id, don_vi_tinh.id as don_vi_tinh_id, loai_giam_gia.id as loai_giam_gia_id, thoi_gian_bao_hanh.id as thoi_gian_bao_hanh_id")
	} else if tableName == "nhan_vien" {
		query.Select("nhan_vien.*, chuc_vu.ten as chuc_vu")
	} else if tableName == "hoa_don_nhap_kho" {
		query.Select("hoa_don_nhap_kho.*, kho.ten as kho, nha_phan_phoi.ten as nha_phan_phoi")
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