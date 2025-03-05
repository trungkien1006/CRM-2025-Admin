package requests

import (
	"mime/multipart"
)

type Loai_san_pham_create struct {
	Ten   string `form:"ten" binding:"required"`
	Hinh_anh *multipart.FileHeader `form:"hinh_anh" binding:"required" swaggerignore:"true"`
}

type Loai_san_pham_update struct {
	Id    int    `form:"id" binding:"required"`
	Ten   string `form:"ten" binding:"required"`
	Hinh_anh *multipart.FileHeader `form:"hinh_anh" binding:"omitempty" swaggerignore:"true"`
}

type Loai_san_pham_delete struct {
	Id int `uri:"id" binding:"required"`
}