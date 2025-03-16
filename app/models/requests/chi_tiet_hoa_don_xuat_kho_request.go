package requests

type Chi_tiet_hoa_don_xuat_kho_create struct {
	San_pham_id			int					`json:"san_pham_id" binding:"required"`
	Ctsp_id				int					`json:"ctsp_id" binding:"required"`
	Sku					string				`json:"sku" binding:"required"`
	Don_vi_tinh			string				`json:"don_vi_tinh" binding:"required"`
	So_luong_ban		int					`json:"so_luong_ban" binding:"required"`
	Gia_ban				float32				`json:"gia_ban" binding:"required"`
	Chiet_khau			float32				`json:"chiet_khau" binding:"required"`
	Thanh_tien			float32				`json:"thanh_tien" binding:"required"`
	Gia_nhap			float32				`json:"gia_nhap" binding:"required"`
	Loi_nhuan			float32				`json:"loi_nhuan" binding:"required"`
	La_qua_tang			int					`json:"la_qua_tang" binding:"required"`
}