package dto

type Metadata struct {
	PageSize  int `json:"pageSize"`
	Page      int `json:"page"`
	TotalPage int `json:"totalPage"`
	TotalData int `json:"totalData"`
}
