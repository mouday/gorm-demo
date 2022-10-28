package service

import (
	"fmt"
	"time"

	"github.com/mouday/gorm-demo/models"
)

/*
* 创建记录
 */
func CreateRow() {
	db, _ := GetDb()

	user := models.UserModel{
		Name:     "Tom",
		Age:      20,
		Birthday: time.Now(),
	}

	// 通过数据的指针来创建
	db.Create(&user)
	// INSERT INTO `user`
	// (`create_time`,`update_time`,`delete_time`,`name`,`age`,`birthday`)
	// VALUES
	// ('2022-10-28 14:02:50.256','2022-10-28 14:02:50.256',NULL,'Tom',20,'2022-10-28 14:02:50.254')

	// 返回插入数据的主键
	fmt.Printf("user.ID: %v\n", user.ID)
	// user.ID: 1
}

// 指定的字段创建记录
func SelectCreateRow() {
	db, _ := GetDb()

	user := models.UserModel{
		Name:     "Tom",
		Age:      20,
		Birthday: time.Now(),
	}

	// 通过数据的指针来创建
	// Omit 忽略
	db.Select("Name", "Age").Create(&user)
	// INSERT INTO `user`
	// (`create_time`,`update_time`,`name`,`age`)
	// VALUES
	// ('2022-10-28 14:07:21.884','2022-10-28 14:07:21.884','Tom',20)
}

// 批量创建
func CreateInBatches() {
	db, _ := GetDb()

	users := []models.UserModel{{
		Name:     "Tom",
		Age:      20,
		Birthday: time.Now(),
	}, {
		Name:     "Jack",
		Age:      21,
		Birthday: time.Now(),
	}}

	batchSize := 100
	db.CreateInBatches(users, batchSize)

	// INSERT INTO `user`
	// (`create_time`,`update_time`,`delete_time`,`name`,`age`,`birthday`)
	// VALUES
	// ('2022-10-28 14:11:20.253','2022-10-28 14:11:20.253',NULL,'Tom',20,'2022-10-28 14:11:20.252'),
	// ('2022-10-28 14:11:20.253','2022-10-28 14:11:20.253',NULL,'Jack',21,'2022-10-28 14:11:20.252')
}

// 通过map创建
func CreateByMap() {
	db, _ := GetDb()

	user := map[string]interface{}{
		"Name": "jinzhu",
		"Age":  23,
	}

	db.Model(&models.UserModel{}).Create(user)
	//  INSERT INTO `user` (`age`,`name`) VALUES (23,'jinzhu')
}
