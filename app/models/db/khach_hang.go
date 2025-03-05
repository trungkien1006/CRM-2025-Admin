package db

import "gorm.io/gorm"

type Khach_hang struct {
	gorm.Model

	Ho_ten        string
	Dien_thoai    string
	Dia_chi       string
}