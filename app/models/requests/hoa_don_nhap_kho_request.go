package requests

type Hoa_don_nhap_kho_create struct {
	Nha_phan_phoi_id 			int									`json:"nha_phan_phoi_id" binding:"required"`
	Kho_id 						int									`json:"kho_id" binding:"required"`
	Ngay_nhap 					string								`json:"ngay_nhap" binding:"required"`
	Tong_tien 					float32								`json:"tong_tien" binding:"required"`

	Chi_tiet_hoa_don_nhap_kho 	[]Chi_tiet_hoa_don_nhap_kho_create 	`json:"chi_tiet_hoa_don_nhap_kho" binding:"required"`
}