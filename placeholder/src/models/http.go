package models

type Response struct {
	Data       any             `json:"data,omitempty"`
	Pagination *Pagination     `json:"pagination,omitempty"`
	Errors     []ErrorResponse `json:"errors,omitempty"`
}

type ErrorResponse struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
	Tag     string `json:"tag,omitempty"`
}

type Pagination struct {
	PerPage   int               `json:"per_page,omitempty" query:"per_page"`
	Page      int               `json:"page,omitempty" query:"page"`
	SortBy    string            `json:"sort_by,omitempty" query:"sort_by"`
	SortOrder string            `json:"sort_order,omitempty" query:"sort_order"`
	Total     int64             `json:"total,omitempty"`
	Search    map[string]string `json:"search,omitempty"`
}
