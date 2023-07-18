package request

type PageInfo struct {
	Page     int `json:"page" form:"page,default=1" binding:"gte=1" label:"页码"`
	PageSize int `json:"page_size" form:"page_size,default=10" binding:"gte=1" label:"分页数量"`
}
