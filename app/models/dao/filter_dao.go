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
	} else if tableName == "hoa_don_xuat_kho" {
		query.Joins("LEFT JOIN khach_hang ON khach_hang.id = hoa_don_nhap_kho.khach_hang_id").
			Joins("LEFT JOIN nhan_vien as nhan_vien_sale ON nhan_vien_sale.id = hoa_don_nhap_kho.nhan_vien_sale_id").
			Joins("LEFT JOIN nhan_vien as nhan_vien_giao_hang ON nhan_vien_giao_hang.id = hoa_don_nhap_kho.nhan_vien_giao_hang_id")
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
	} else if tableName == "hoa_don_xuat_kho" {
		query.Select("hoa_don_xuat_kho.*, nhan_vien_sale.ten as nhan_vien_sale, nhan_vien_giao_hang.ten as nhan_vien_giao_hang, khach_hang.ten as khach_hang")
	}

	//sort du lieu
	if req.Sort != "" {
		query.Order(req.Sort + " " + req.Order)
	}

	//chay truy van select them bang con 
	if tableName == "nha_phan_phoi" {
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
	} else if tableName == "hoa_don_nhap_kho" {
		var Hoa_don_nhap_kho 			[]responses.Hoa_don_nhap_kho_filter
		var Chi_tiet_hoa_don_nhap_kho	[]responses.Chi_tiet_hoa_don_nhap_kho_response
		var Hdnk_ids 					[]int

		//lay ra danh sach hoa don nhap kho
		if err := query.Find(&Hoa_don_nhap_kho).Error; err != nil {
			return err
		}

		//tao ra danh sach id cua hoa don nhap kho
		for _, value := range Hoa_don_nhap_kho {
			Hdnk_ids = append(Hdnk_ids, value.Id)
		}

		//truy van danh sach chi tiet hdnk
		if err := helpers.GormDB.Debug().Table("chi_tiet_hoa_don_nhap_kho as ct_hdnk").
			Joins("JOIN san_pham as sp ON sp.id = ct_hdnk.san_pham_id").
			Where("ct_hdnk.hoa_don_id IN ?", Hdnk_ids).
			Select("ct_hdnk.*, sp.ten as ten_san_pham").
			Find(&Chi_tiet_hoa_don_nhap_kho).Error
		err != nil {
			return err
		}

		var Chi_tiet_hoa_don_nhap_kho_group = make(map[int][]responses.Chi_tiet_hoa_don_nhap_kho_response)

		//map cthdnk theo hoa don id
		for _, value := range Chi_tiet_hoa_don_nhap_kho {
			Chi_tiet_hoa_don_nhap_kho_group[value.Hoa_don_id] = append(Chi_tiet_hoa_don_nhap_kho_group[value.Hoa_don_id], value)
		}

		//truyen ds cthdnk vao tung hoa don nhap kho
		for index, value := range Hoa_don_nhap_kho {
			Hoa_don_nhap_kho[index].Chi_tiet_hoa_don_nhap_kho = Chi_tiet_hoa_don_nhap_kho_group[value.Id]
		}

		//map ket qua vao bien response generic
		res.Data = make([]T, len(Hoa_don_nhap_kho))

		for index, value := range Hoa_don_nhap_kho {
			res.Data[index] = any(value).(T)
		}
	} else if tableName == "hoa_don_xuat_kho" {
		var Hoa_don_xuat_kho 			[]responses.Hoa_don_xuat_kho_filter
		var Chi_tiet_hoa_don_xuat_kho	[]responses.Chi_tiet_hoa_don_xuat_kho_response
		var Hdxk_ids 					[]int

		//lay ra danh sach hoa don xuat kho
		if err := query.Find(&Hoa_don_xuat_kho).Error; err != nil {
			return err
		}

		//tao ra danh sach id cua hoa don xuat kho
		for _, value := range Hoa_don_xuat_kho {
			Hdxk_ids = append(Hdxk_ids, value.Id)
		}

		//truy van danh sach cthdxk thuoc hoa don xuat kho
		if err := helpers.GormDB.Debug().Table("chi_tiet_hoa_don_xuat_kho as ct_hdxk").
			Joins("JOIN san_pham as sp ON sp.id = ct_hdxk.san_pham_id").
			Where("ct_hdxk.hoa_don_id IN ?", Hdxk_ids).
			Select("ct_hdxk.*, sp.ten as ten_san_pham").
			Find(&Chi_tiet_hoa_don_xuat_kho).Error
		err != nil {
			return err
		}

		var Chi_tiet_hoa_don_xuat_kho_group = make(map[int][]responses.Chi_tiet_hoa_don_xuat_kho_response)

		//map cthdxk theo hoa don id
		for _, value := range Chi_tiet_hoa_don_xuat_kho {
			Chi_tiet_hoa_don_xuat_kho_group[value.Hoa_don_id] = append(Chi_tiet_hoa_don_xuat_kho_group[value.Hoa_don_id], value)
		}

		//truyen ds cthdxk theo tung hoa don xuat kho
		for index, value := range Hoa_don_xuat_kho {
			Hoa_don_xuat_kho[index].Chi_tiet_hoa_don_xuat_kho = Chi_tiet_hoa_don_xuat_kho_group[value.Id]
		}

		//map ket qua vao bien response generic
		res.Data = make([]T, len(Hoa_don_xuat_kho))

		for index, value := range Hoa_don_xuat_kho {
			res.Data[index] = any(value).(T)
		}
	} else {
		if err := query.Find(&res.Data).Error; err != nil {
			return err
		}
	}

	//tinh toan total page
	res.Total_Page = int(math.Ceil(float64(totalRecord) / float64(req.Limit)))

	return nil
}