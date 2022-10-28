package utils

import (
	"gorm.io/gorm"
)

/*
分页
https://gorm.io/zh_CN/docs/scopes.html#%E5%88%86%E9%A1%B5
*/
func Paginate(page int, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * size
		return db.Offset(offset).Limit(size)
	}
}
