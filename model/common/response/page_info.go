package response

import "math"

// PageInfo
// @Description: 分页信息
type PageInfo struct {
	Page       int   `json:"page" form:"page"`
	PageSize   int   `json:"page_size" form:"page_size"`
	TotalPages int   `json:"total_pages" form:"total_pages"`
	TotalCount int64 `json:"total_count" form:"total_count"`
}

// GetPageInfo
//
//	@Description: 获取分页信息
//	@param info 页码信息
//	@param page 页码
//	@param pageSize 分页数量

func (receiver PageInfo) GetPageInfo(info *PageInfo, page int, pageSize int) {
	info.Page = page
	info.PageSize = pageSize
	info.TotalPages = receiver.CalcPages(info.TotalCount, pageSize)
}

// CalcPages
//
//	@Description: 计算总页数
//	@param count 当前总数据量
//	@param pageSize 分页数量
//	@return int 总页数

func (receiver PageInfo) CalcPages(count int64, pageSize int) int {
	return int(math.Ceil(float64(count) / float64(pageSize)))

}
