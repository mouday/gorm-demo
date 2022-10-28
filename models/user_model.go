package models

import "time"

// 定义模型
type UserModel struct {
	BaseModel

	Name     string
	Age      uint8
	Birthday time.Time
}

// 自定义表名
func (UserModel) TableName() string {
	return "user"
}
