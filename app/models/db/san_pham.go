package db

import "gorm.io/gorm"

type San_pham struct {
	gorm.Model

	Ten                   string	`json:"ten"`
	Upc                   string	`json:"upc"`
	Loai_san_pham_id      int		`json:"loai_san_pham_id"`
	Hinh_anh              string	`json:"hinh_anh"`
	Don_vi_tinh_id        int		`json:"don_vi_tinh"`
	Vat                   float32	`json:"vat"`
	Mo_ta                 string	`json:"mo_ta"`
	Trang_thai            int		`json:"trang_thai"`
	Loai_giam_gia_id      int		`json:"loai_giam_gia_id"`
	Thoi_gian_bao_hanh_id int		`json:"thoi_gian_bao_hanh_id"`

	Chi_tiet_san_pham []Chi_tiet_san_pham `json:"chi_tiet_san_pham" gorm:"foreignKey:san_pham_id"`
	Loai_san_pham      Loai_san_pham      `json:"loai_san_pham" gorm:"constraint:OnDelete:CASCADE;"`
	Don_vi_tinh        Don_vi_tinh        `json:"don_vi_tinh" gorm:"constraint:OnDelete:CASCADE;"`
	Loai_giam_gia      Loai_giam_gia      `json:"loai_giam_gia" gorm:"constraint:OnDelete:CASCADE;"`
	Thoi_gian_bao_hanh Thoi_gian_bao_hanh `json:"thoi_gian_bao_hanh" gorm:"constraint:OnDelete:CASCADE;"`
}