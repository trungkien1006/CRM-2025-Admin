package db

import "gorm.io/gorm"

type Chuc_vu struct {
	gorm.Model

	Ten  	string	

	Chuc_nang []Chuc_nang	`gorm:"many2many:quyen;"`
}