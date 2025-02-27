package db

import "gorm.io/gorm"

type Nhan_vien struct {
	gorm.Model

	Ten_dang_nhap string
	Mat_khau      string
	Ho_ten        string
	Email         string
	Dien_thoai    string
	Dia_chi       string
	Avatar        string
	Chuc_vu_id    int
}