package db

import "gorm.io/gorm"

type Chuc_vu struct {
	gorm.Model

	Ten        string
}