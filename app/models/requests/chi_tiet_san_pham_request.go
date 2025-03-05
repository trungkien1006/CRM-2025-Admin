package requests

import "mime/multipart"

type Chi_tiet_san_pham_request struct {
	Ten_phan_loai 	string                	`form:"ten_phan_loai" binding:"required"`
	Hinh_anh      	*multipart.FileHeader 	`form:"file" binding:"omitempty" swaggerignore:"true"`
	Trang_thai    	int                   	`form:"trang_thai" binding:"required"`
}

type Chi_tiet_san_pham_get_by_product_id struct {
	Product_id 		int						`uri:"product_id" binding:"required"`
}