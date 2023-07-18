package orm

import "gorm.io/gorm"

func Where(wheres map[string]map[string]string) func(db *gorm.DB) *gorm.DB {
	var (
		andWhere, orWhere = map[string]interface{}{}, map[string]interface{}{}
		likeWhere         string
		timeWhere         string
	)

	for key, value := range wheres {
		//组装where
		if key == "AND" {
			for key2, value2 := range value {
				andWhere[key2] = value2
			}
		}
		if key == "OR" {
			for key2, value2 := range value {
				orWhere[key2] = value2
			}
		}
		if key == "LIKE" {
			for key2, value2 := range value {
				likeWhere += key2 + ` LIKE "` + value2 + `%" OR `
			}
			// 去除尾部OR
			likeWhere = likeWhere[:len(likeWhere)-3]
		}
		if key == "TIME" {
			for _, value2 := range value {
				timeWhere = value2
			}
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(andWhere).Where(likeWhere).Where(timeWhere).Or(orWhere)
	}
}
