package service

import (
	"fmt"

	"github.com/mouday/gorm-demo/models"
	"github.com/mouday/gorm-demo/utils"
)

// 查询单条记录
// 没有找到记录时，它会返回 ErrRecordNotFound 错误
func GetFirst() {
	db, _ := GetDb()

	var user models.UserModel

	// （主键升序）
	db.First(&user)
	// SELECT * FROM `user` WHERE `user`.`delete_time` IS NULL ORDER BY `user`.`id` LIMIT 1

	fmt.Printf("user: %v\n", user)
}

// 查询记录
func GetTake() {
	db, _ := GetDb()

	var user models.UserModel

	// 获取一条记录，没有指定排序字段
	db.Take(&user)
	// SELECT * FROM `user` WHERE `user`.`delete_time` IS NULL LIMIT 1

	fmt.Printf("user: %v\n", user)
}

// 查询记录
func GetLast() {
	db, _ := GetDb()

	var user models.UserModel

	// 获取最后一条记录（主键降序）
	db.Last(&user)
	// SELECT * FROM `user` WHERE `user`.`delete_time` IS NULL ORDER BY `user`.`id` DESC LIMIT 1

	fmt.Printf("user: %v\n", user)
}

// 根据主键检索
func GetFirstByID() {
	db, _ := GetDb()

	var user models.UserModel

	// 根据主键检索
	db.First(&user, 1)
	// SELECT * FROM `user` WHERE `user`.`id` = 1 AND `user`.`delete_time` IS NULL ORDER BY `user`.`id` LIMIT 1

	fmt.Printf("user: %v\n", user)
}

// Find
func GetFind() {
	db, _ := GetDb()

	var users []models.UserModel

	// 根据主键检索
	db.Find(&users, []int{1, 2, 3})
	// SELECT * FROM `user` WHERE `user`.`id` IN (1,2,3) AND `user`.`delete_time` IS NULL

	fmt.Printf("user: %v\n", users)
}

// Where Find
func GetWhereFind() {
	db, _ := GetDb()

	var users []models.UserModel

	db.Where("name = ?", "Tom").Find(&users)
	// SELECT * FROM `user` WHERE name = 'Tom' AND `user`.`delete_time` IS NULL

	fmt.Printf("user: %v\n", users)
}

// Where First
func GetWhereFirst() {
	db, _ := GetDb()

	var user models.UserModel

	db.Where("name = ?", "Tom").First(&user)
	// SELECT * FROM `user` WHERE name = 'Tom' AND `user`.`delete_time` IS NULL ORDER BY `user`.`id` LIMIT 1

	fmt.Printf("user: %v\n", user)
}

// LIKE
func GetWhereLikeFind() {
	db, _ := GetDb()

	var users []models.UserModel

	db.Where("name LIKE ?", "%Tom%").Find(&users)
	// SELECT * FROM `user` WHERE name LIKE '%Tom%' AND `user`.`delete_time` IS NULL

	fmt.Printf("user: %v\n", users)
}

// Struct & Map 条件
func GetWhereStructFind() {
	db, _ := GetDb()

	var users []models.UserModel

	db.Where(&models.UserModel{Name: "jinzhu", Age: 20}).Find(&users)
	// SELECT * FROM `user` WHERE `user`.`name` = 'jinzhu' AND `user`.`age` = 20 AND `user`.`delete_time` IS NULL

	fmt.Printf("user: %v\n", users)
}

// 选择特定字段
func GetSelectFind() {
	db, _ := GetDb()

	var users []models.UserModel

	db.Select("name", "age").Find(&users)
	// SELECT `name`,`age` FROM `user` WHERE `user`.`delete_time` IS NULL

	fmt.Printf("user: %v\n", users)
}

// Order
func GetOrderFind() {
	db, _ := GetDb()

	var users []models.UserModel

	db.Order("age desc").Find(&users)
	// SELECT * FROM `user` WHERE `user`.`delete_time` IS NULL ORDER BY age desc

	fmt.Printf("user: %v\n", users)
}

// Limit & Offset
func GetLimitOffsetFind() {
	db, _ := GetDb()

	var users []models.UserModel

	db.Limit(1).Offset(2).Find(&users)
	// SELECT * FROM `user` WHERE `user`.`delete_time` IS NULL LIMIT 1 OFFSET 2

	fmt.Printf("user: %v\n", users)
}

// 分页 Paginate
func GetPaginateFind() {
	db, _ := GetDb()

	var users []models.UserModel

	db.Scopes(utils.Paginate(2, 10)).Find(&users)
	// SELECT * FROM `user` WHERE `user`.`delete_time` IS NULL LIMIT 10 OFFSET 10

	fmt.Printf("user: %v\n", users)
}

// Pluck
func GetPluckFind() {
	db, _ := GetDb()

	var ages []int64

	db.Model(&models.UserModel{}).Pluck("age", &ages)
	// SELECT `age` FROM `user` WHERE `user`.`delete_time` IS NULL

	fmt.Printf("user: %v\n", ages)
	// user: [20 20 20 20 21 23]
}

// Count
func GetCountFind() {
	db, _ := GetDb()

	var count int64
	db.Model(&models.UserModel{}).Count(&count)
	// SELECT count(*) FROM `user` WHERE `user`.`delete_time` IS NULL

	fmt.Printf("count: %v\n", count)
	// count: 6
}
