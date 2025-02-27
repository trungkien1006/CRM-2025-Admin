package db

import "gorm.io/gorm"

type Chuc_nang struct {
	gorm.Model

	Ten string 
	Code string
	Loai string
	Show_in_menu string
}