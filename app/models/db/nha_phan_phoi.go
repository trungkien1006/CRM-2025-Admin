package db

import "gorm.io/gorm"

type Nha_phan_phoi struct {
	gorm.Model

	Ten        string
	Dia_chi    string
	Dien_thoai string
	Email      string
}