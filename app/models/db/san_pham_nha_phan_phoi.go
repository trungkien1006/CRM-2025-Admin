package db

import "gorm.io/gorm"

type San_pham_nha_phan_phoi struct {
	gorm.Model

	Nha_phan_phoi_id int
	San_pham_id      int
}