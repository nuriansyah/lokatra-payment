package dto

type Metadata struct {
	Strategy  string `json:"strategy,omitempty"`
	PageSize  int    `json:"pageSize"`
	Page      int    `json:"page,omitempty"`
	TotalPage int    `json:"totalPage,omitempty"`
	TotalData int    `json:"totalData,omitempty"`

	NextCursor interface{} `json:"nextCursor,omitempty"`
	PrevCursor interface{} `json:"prevCursor,omitempty"`
	HasNext    bool        `json:"hasNext"`
	HasPrev    bool        `json:"hasPrev"`
	HasMore    bool        `json:"hasMore"`
}
