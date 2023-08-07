package orm

import "gorm.io/gorm"

// Paginate
//  @Description: 公共分页方法
//  @param page 页码
//  @param pageSize 当前页数量
//  @return func(db *gorm.DB) *gorm.DB gorm连接对象

func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
