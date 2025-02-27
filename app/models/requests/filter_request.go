package requests

type FilterStruc struct {
	Field     string `json:"field"`
	Condition string `json:"condition"`
	Value     string `json:"value"`
}

type Filter struct {
	Filters string `json:"filters" binding:"omitempty"`
	Sort    string `json:"sort" binding:"omitempty"`
	Order   string `json:"order" binding:"omitempty,oneif:asc desc"`
	Page    int    `json:"page" binding:"required"`
	Limit   int    `json:"limit" binding:"required"`
}