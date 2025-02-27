package requests

import "mime/multipart"

type Chi_tiet_san_pham_request struct {
	Ten_phan_loai string                `form:"ten_phan_loai" binding:"required"`
	Hinh_anh      *multipart.FileHeader `form:"file" binding:"omitempty"`
	Trang_thai    int                   `form:"trang_thai" binding:"required"`
}