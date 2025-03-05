package db

import "gorm.io/gorm"

type Hoa_don_nhap_kho struct {
	gorm.Model

	Nha_phan_phoi_id 			int							`json:"nha_phan_phoi_id"`
	Kho_id 						int							`json:"kho_id"`
	Ngay_nhap 					string						`json:"ngay_nhap"`
	Tong_tien 					float32						`json:"tong_tien"`

	Nha_phan_phoi 				Nha_phan_phoi				`json:"nha_phan_phoi" gorm:"constraint:OnDelete:CASCADE"`
	Kho							Kho							`json:"kho" gorm:"constraint:OnDelete:CASCADE"`
	Chi_tiet_hoa_don_nhap_kho	[]Chi_tiet_hoa_don_nhap_kho	`json:"chi_tiet_hoa_don_nhap_kho" gorm:"foreignKey:hoa_don_id"`
}