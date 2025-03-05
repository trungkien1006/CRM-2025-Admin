package requests

type Nha_phan_phoi_create struct {
	Ten        	string	`json:"ten" binding:"required"`
	Dia_chi    	string	`json:"dia_chi" binding:"required"`
	Dien_thoai 	string	`json:"dien_thoai" binding:"required"`
	Email      	string	`json:"email" binding:"required"`

	San_pham_id	[]int 	`json:"san_pham_id" binding:"omitempty"`
}	

type Nha_phan_phoi_update struct {
	Id 			int		`json:"id" binding:"required"`
	Ten        	string	`json:"ten" binding:"required"`
	Dia_chi    	string	`json:"dia_chi" binding:"required"`
	Dien_thoai 	string	`json:"dien_thoai" binding:"required"`
	Email      	string	`json:"email" binding:"required"`

	San_pham_id	[]int 	`json:"san_pham_id" binding:"omitempty"`
}

type Nha_phan_phoi_delete struct {
	Id			int		`uri:"id" binding:"required"`
}