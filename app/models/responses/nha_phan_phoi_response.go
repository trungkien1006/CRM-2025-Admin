package responses

import "admin-v1/app/models/db"

type Nha_phan_phoi_create struct {
	Nha_phan_phoi db.Nha_phan_phoi `json:"nha_phan_phoi"`
}