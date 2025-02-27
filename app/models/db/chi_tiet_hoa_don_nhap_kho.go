package db

import "gorm.io/gorm"

type Chi_tiet_hoa_don_nhap_kho struct {
	gorm.Model

	Hoa_don_id  int
	San_pham_id int
	Ctsp_id     int
	Sku         string
	So_luong    string
	Don_vi_tinh string
	Ke          string
	Gia_nhap    string
	Gia_ban     string
	Chiet_khau  string
	Thanh_tien  string
	La_qua_tang int
}