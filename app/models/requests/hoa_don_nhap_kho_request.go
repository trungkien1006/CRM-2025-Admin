package requests

type Hoa_don_nhap_kho_create struct {
	So_hoa_don					int									`json:"so_hoa_don" binding:"omitempty"`
	Ma_hoa_don					string								`json:"ma_hoa_don" binding:"omitempty"`
	Nha_phan_phoi_id 			int									`json:"nha_phan_phoi_id" binding:"required"`
	Kho_id 						int									`json:"kho_id" binding:"required"`
	Ngay_nhap 					string								`json:"ngay_nhap" binding:"required"`
	Tong_tien 					float32								`json:"tong_tien" binding:"required"`
	Tra_truoc					float32								`json:"tra_truoc" binding:"omitempty"`
	Con_lai						float32								`json:"con_lai" binding:"omitempty"`
	Ghi_chu						string								`json:"ghi_chu" binding:"omitempty"`

	Ds_san_pham_nhap 			[]Chi_tiet_hoa_don_nhap_kho_create 	`json:"ds_san_pham_nhap" binding:"required"`
}