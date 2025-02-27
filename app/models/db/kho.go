package db

import "gorm.io/gorm"

type Kho struct {
	gorm.Model

	Ten        string
	Dia_chi    string
}