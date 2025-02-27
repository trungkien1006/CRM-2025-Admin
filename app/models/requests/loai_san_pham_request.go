package requests

import (
	"mime/multipart"
)

type Loai_san_pham_create struct {
	Ten   string `form:"ten" binding:"required"`
	Image *multipart.FileHeader `form:"file" binding:"required"`
}

type Loai_san_pham_update struct {
	Id    int    `form:"id" binding:"required"`
	Ten   string `form:"ten" binding:"required"`
	Image *multipart.FileHeader `form:"file" binding:"omitempty"`
}

type Loai_san_pham_delete struct {
	Id int `uri:"id" binding:"required"`
}