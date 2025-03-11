package db

import "gorm.io/gorm"

type Ton_kho struct {
	gorm.Model

	San_pham_id					int							`json:"san_pham_id"`
	Ctsp_id						int							`json:"ctsp_id"`
	Sku 						string						`json:"sku"`
	So_luong_ton				int							`json:"so_luong_ton"`

	San_pham					San_pham					`json:"san_pham"`
	Chi_tiet_san_pham			Chi_tiet_san_pham			`json:"chi_tiet_san_pham"`
}