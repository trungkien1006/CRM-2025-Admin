package requests

type Quyen_read struct {
	Chuc_vu_id int	`query:"chuc_vu_id" binding:"required"`	
}

type Quyen_modify struct {
	Chuc_vu_id 	int 				`json:"chuc_vu_id" binding:"required"`
	Quyen 		[]Quyen_modify_data	`json:"quyen" binding:"required"`
}

type Quyen_modify_data struct {
	Quyen_id	int	`json:"id" binding:"required"`
	Active		int `json:"active" binding:"required"`
}