package requests

import "mime/multipart"

type Nhan_vien_create struct {
	Hinh_anh 		*multipart.FileHeader 	`form:"hinh_anh" binding:"required" swaggerignore:"true"`
	Ten_dang_nhap 	string					`form:"ten_dang_nhap" binding:"required"`
	Ho_ten 			string					`form:"ho_ten" binding:"required"`
	Email 			string 					`form:"email" binding:"required"`
	Dien_thoai 		string					`form:"dien_thoai" binding:"required"`
	Dia_chi 		string 					`form:"dia_chi" binding:"required"`
	Chuc_vu 		int 					`form:"chuc_vu" binding:"required"`
}

type Nhan_vien_update struct {
	Id 				int 					`form:"id" binding:"required"`
	Hinh_anh 		*multipart.FileHeader 	`form:"hinh_anh" binding:"omitempty" swaggerignore:"true"`
	Ten_dang_nhap 	string					`form:"ten_dang_nhap" binding:"required"`
	Ho_ten 			string					`form:"ho_ten" binding:"required"`
	Email 			string 					`form:"email" binding:"required"`
	Dien_thoai 		string					`form:"dien_thoai" binding:"required"`
	Dia_chi 		string 					`form:"dia_chi" binding:"required"`
	Chuc_vu 		int 					`form:"chuc_vu" binding:"required"`
}

type Nhan_vien_delete struct {
	Id 				int 					`form:"id" binding:"required"`
}