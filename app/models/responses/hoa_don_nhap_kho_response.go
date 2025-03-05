package responses

import "admin-v1/app/models/db"

type Hoa_don_nhap_kho_filter struct {
	Nha_phan_phoi 		string				`json:"nha_phan_phoi"`
	Kho 				string				`json:"kho"`
	Ngay_nhap 			string				`json:"ngay_nhap"`
	Tong_tien 			float32				`json:"tong_tien"`
}

type Hoa_don_nhap_kho_create struct {
	Hoa_don_nhap_kho 	db.Hoa_don_nhap_kho `json:"hoa_don_nhap_kho"`
}