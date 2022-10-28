package models

import (
	"time"

	"gorm.io/gorm"
)

// 自定义模型基类
// https://gorm.io/zh_CN/docs/models.html
type BaseModel struct {
	ID         uint           `gorm:"primaryKey"`
	CreateTime time.Time      `gorm:"autoCreateTime"`
	UpdateTime time.Time      `gorm:"autoUpdateTime"`
	DeleteTime gorm.DeletedAt `gorm:"index"`
}
