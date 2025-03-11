package dao

import (
	"admin-v1/app/helpers"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"encoding/json"
	"math"
)

func FilterExec[T any](req *requests.Filter, res *responses.Filter[T], tableName string) error {
	//ghep ten bang can truy van
	query := helpers.GormDB.Debug().
		Table(tableName)

	//loc cac dieu kien, join 1 so bang phu doi voi 1 so trương hơp cu the
	if tableName == "san_pham" {
		query.Joins("LEFT JOIN loai_san_pham ON loai_san_pham.id = san_pham.loai_san_pham_id").
			Joins("LEFT JOIN don_vi_tinh ON don_vi_tinh.id = san_pham.don_vi_tinh_id").
			Joins("LEFT JOIN thoi_gian_bao_hanh ON thoi_gian_bao_hanh.id = san_pham.thoi_gian_bao_hanh_id").
			Joins("LEFT JOIN loai_giam_gia ON loai_giam_gia.id = san_pham.loai_giam_gia_id")
	} else if tableName == "nhan_vien" {
		query.Joins("LEFT JOIN chuc_vu ON chuc_vu.id = nhan_vien.chuc_vu_id")
	}else if tableName == "hoa_don_nhap_kho" {
		query.Joins("LEFT JOIN nha_phan_phoi ON nha_phan_phoi.id = hoa_don_nhap_kho.nha_phan_phoi_id").
			Joins("LEFT JOIN kho ON kho.id = hoa_don_nhap_kho.kho_id")
	}

	//chay ham filter neu co truyen 1 vao chuoi filter duoc ma hoa JSON
	if req.Filters != "" {
		var filters []requests.FilterStruc

		err := json.Unmarshal([]byte(req.Filters), &filters)
		if err != nil {
			return err
		}

		helpers.Filter(query, filters)
	}

	var totalRecord int64 = 0
	
	//lay ra tong so record
	if err := query.Where(tableName + ".deleted_at IS NULL").Count(&totalRecord).Error; err != nil {
		return err
	}

	//phan trang
	if req.Limit != 0 && req.Page != 0 {
		query.Offset((req.Page - 1) * req.Limit).Limit(req.Limit)
	}

	//select 1 so truong du lieu phu hop dau ra
	if tableName == "san_pham" {
		query.Select("san_pham.*, loai_san_pham.ten as loai_san_pham, don_vi_tinh.ten as don_vi_tinh, loai_giam_gia.ten as loai_giam_gia, thoi_gian_bao_hanh.ten as thoi_gian_bao_hanh, loai_san_pham.id as loai_san_pham_id, don_vi_tinh.id as don_vi_tinh_id, loai_giam_gia.id as loai_giam_gia_id, thoi_gian_bao_hanh.id as thoi_gian_bao_hanh_id")
	} else if tableName == "nhan_vien" {
		query.Select("nhan_vien.*, chuc_vu.ten as chuc_vu")
	} else if tableName == "hoa_don_nhap_kho" {
		query.Select("hoa_don_nhap_kho.*, kho.ten as kho, nha_phan_phoi.ten as nha_phan_phoi")
	}

	//sort du lieu
	if req.Sort != "" {
		query.Order(req.Sort + " " + req.Order)
	}

	//chay truy van neu bang can tim khong phai nha phan phoi
	if tableName != "nha_phan_phoi" {
		if err := query.Find(&res.Data).Error; err != nil {
			return err
		}
	} else {
		var Nha_phan_phoi 	[]responses.Nha_phan_phoi_filter
		var Npp_san_pham 	[]responses.Nha_phan_phoi_san_pham_response
		var Npp_ids 		[]int

		//lay ra danh sach nha phan phoi
		if err := query.Find(&Nha_phan_phoi).Error; err != nil {
			return err
		}

		//tao ra danh sach id cua nha phan phoi
		for _, value := range Nha_phan_phoi {
			Npp_ids = append(Npp_ids, value.Id)
		}

		//truy van danh sach sp thuoc nha phan phoi
		if err := helpers.GormDB.Debug().Table("san_pham").
			Joins("JOIN san_pham_nha_phan_phoi as sp_npp ON sp_npp.san_pham_id = san_pham.id").
			Where("sp_npp.nha_phan_phoi_id IN ?", Npp_ids).
			Select("san_pham.ten, san_pham.id, san_pham.upc, sp_npp.nha_phan_phoi_id as nha_phan_phoi_id").
			Find(&Npp_san_pham).Error; 
		err != nil {
			return err
		}

		var Npp_san_pham_group = make(map[int][]responses.Nha_phan_phoi_san_pham_response)

		//map san pham theo id npp
		for _, value := range Npp_san_pham {
			Npp_san_pham_group[value.Nha_phan_phoi_id] = append(Npp_san_pham_group[value.Nha_phan_phoi_id], value)
		}

		//truyen ds san pham theo tung nha phan phoi
		for index, value := range Nha_phan_phoi {
			Nha_phan_phoi[index].Ds_san_pham = Npp_san_pham_group[value.Id]
		}

		//map ket qua vao bien response generic
		res.Data = make([]T, len(Nha_phan_phoi))

		for index, value := range Nha_phan_phoi {
			res.Data[index] = any(value).(T)
		}
	}

	//tinh toan total page
	res.Total_Page = int(math.Ceil(float64(totalRecord) / float64(req.Limit)))

	return nil
}