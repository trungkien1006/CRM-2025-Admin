package db

import "gorm.io/gorm"

type Hoa_don_nhap_kho struct {
	gorm.Model

	Nha_phan_phoi_id int
	Kho_id int
	Ngay_nhap string
	Tong_tien float32
}