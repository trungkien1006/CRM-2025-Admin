package requests

type Hoa_don_xuat_kho_create struct {
	So_hoa_don					int							`json:"so_hoa_don" binding:"omitempty"`
	Ma_hoa_don					string						`json:"ma_hoa_don" binding:"omitempty"`
	Khach_hang_id				int							`json:"khach_hang_id" binding:"required"`
	Nhan_vien_sale_id			int							`json:"nhan_vien_sale_id" binding:"required"`
	Nhan_vien_giao_hang_id		int							`json:"nhan_vien_giao_hang_id" binding:"required"`
	Ngay_xuat					string						`json:"ngay_xuat" binding:"required"`
	Tong_tien					float32						`json:"tong_tien" binding:"required"`
	Vat							float32						`json:"vat" binding:"required"`
	Thanh_tien					float32						`json:"thanh_tien" binding:"required"`
	Tra_truoc 					float32						`json:"tra_truoc" binding:"required"`
	Con_lai						float32						`json:"con_lai" binding:"omitempty"`
	Tong_gia_nhap				float32						`json:"tong_gia_nhap" binding:"required"`
	Loi_nhuan					float32						`json:"loi_nhuan" binding:"required"`
	Ghi_chu						string						`json:"ghi_chu" binding:"omitempty"`
	Da_giao_hang				int							`json:"da_giao_hang" binding:"required"`
	Loai_chiet_khau				int							`json:"loai_chiet_khau" binding:"required"`
	Gia_tri_chiet_khau			string						`json:"gia_tri_chiet_khau" binding:"required"`

	Chi_tiet_hoa_don_xuat_kho 	[]Chi_tiet_hoa_don_xuat_kho_create 	`json:"chi_tiet_hoa_don_xuat_kho" binding:"required"`
}