package responses

type Dang_nhap struct {
	Token 		string		`json:"token"`
	Ds_quyen 	[]string	`json:"ds_quyen"`
}