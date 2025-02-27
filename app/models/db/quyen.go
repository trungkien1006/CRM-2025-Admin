package db

import "gorm.io/gorm"

type Quyen struct {
	gorm.Model

	Chuc_vu_id   int
	Chuc_nang_id int
}