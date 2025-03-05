package responses

import "admin-v1/app/models/db"

type San_pham_filter struct {
	Ten                   		string						`json:"ten"`
	Upc                   		string						`json:"upc"`
	Loai_san_pham     			string						`json:"loai_san_pham"`
	Hinh_anh              		string						`json:"hinh_anh"`
	Don_vi_tinh        			string						`json:"don_vi_tinh"`
	Vat                   		float32						`json:"vat"`
	Mo_ta                 		string						`json:"mo_ta"`
	Trang_thai            		int							`json:"trang_thai"`
	Loai_giam_gia      			string						`json:"loai_giam_gia"`
	Thoi_gian_bao_hanh 			string						`json:"thoi_gian_bao_hanh"`
}

type San_pham_create struct {
	San_pham 	db.San_pham `json:"san_pham"`
}