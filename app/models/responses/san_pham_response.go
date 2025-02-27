package responses

import "admin-v1/app/models/db"

type San_pham_filter struct {
	San_pham 	[]db.San_pham `json:"san_pham"`
	Total_Page 		int `json:"total_page"`
}

type San_pham_create struct {
	San_pham 	db.San_pham `json:"san_pham"`
}