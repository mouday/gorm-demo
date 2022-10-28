# gorm试用Demo

文档：
- [https://github.com/go-gorm/gorm](https://github.com/go-gorm/gorm)
- [https://gorm.io/zh_CN/](https://gorm.io/zh_CN/)

初始化项目

```bash
go mod init github.com/mouday/gorm-demo
```

安装依赖
```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

## 连接数据库

```go
package service

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

const DB_URL = "root:123456@tcp(127.0.0.1:3306)/data?charset=utf8mb4&parseTime=True&loc=Local"

// 连接数据库
func GetDb() (db *gorm.DB, err error) {

    return gorm.Open(mysql.Open(DB_URL), &gorm.Config{
        // 日志级别
        Logger: logger.Default.LogMode(logger.Info),
    })

}

```

## 定义模型

```go
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

```

```go
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

```

## 表操作

```go
package service

import (
	"fmt"

	"github.com/mouday/gorm-demo/models"
)

// 创建表
func CreateTable() {
	db, _ := GetDb()

	// 为 `User` 创建表
	db.Migrator().CreateTable(&models.UserModel{})

	// CREATE TABLE `user` (
	// `id` bigint unsigned AUTO_INCREMENT,
	// 	`create_time` datetime(3) NULL,
	// 	`update_time` datetime(3) NULL,
	// 	`delete_time` datetime(3) NULL,
	// 	`name` longtext,
	// 	`age` tinyint unsigned,
	// 	`birthday` datetime(3) NULL,
	// 	PRIMARY KEY (`id`),
	// 	INDEX `idx_user_delete_time` (`delete_time`)
	// )
}

// 存在检查
func HasTable() {
	db, _ := GetDb()

	// 检查 `User` 对应的表是否存在
	hasTable := db.Migrator().HasTable(&models.UserModel{})
	// SELECT DATABASE()
	// SELECT SCHEMA_NAME from Information_schema.SCHEMATA where SCHEMA_NAME LIKE 'data%' ORDER BY SCHEMA_NAME='data' DESC,SCHEMA_NAME limit 1
	// SELECT count(*) FROM information_schema.tables WHERE table_schema = 'data' AND table_name = 'user' AND table_type = 'BASE TABLE'

	fmt.Printf("hasTable: %v\n", hasTable)
	// hasTable: true
}

// 删除表
func DropTable() {
	db, _ := GetDb()

	db.Migrator().DropTable(&models.UserModel{})
	// SET FOREIGN_KEY_CHECKS = 0;
	// DROP TABLE IF EXISTS `user` CASCADE
	// SET FOREIGN_KEY_CHECKS = 1;
}

```

## 写入操作

```go
package service

import (
	"fmt"
	"time"

	"github.com/mouday/gorm-demo/models"
)

// 创建记录
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

```

## 更新操作

```go
package service

import (
	"github.com/mouday/gorm-demo/models"
)

func UpdateSingleColumn() {
	db, _ := GetDb()

	// 实例中有主键，可以不写where
	// db.Model(&user).Update("name", "Jack")
	db.Model(&models.UserModel{}).Where("id = ?", 1).Update("name", "Jack")
	// UPDATE `user` SET `name`='Jack',`update_time`='2022-10-28 15:35:00.109' WHERE id = 1 AND `user`.`delete_time` IS NULL
}

func UpdateMoreColumn() {
	db, _ := GetDb()

	// 根据主键更新
	// db.Model(&user).Updates(User{Name: "hello", Age: 18})
	db.Model(&models.UserModel{}).Where("id = ?", 1).Updates(
		models.UserModel{
			Name: "hello",
			Age:  18,
		})
	// UPDATE `user` SET `update_time`='2022-10-28 15:36:59.661',`name`='hello',`age`=18 WHERE id = 1 AND `user`.`delete_time` IS NULL
}

```

## 删除操作

```go
package service

import (
	"github.com/mouday/gorm-demo/models"
)

// 软删除
func SoftDelete() {
	db, _ := GetDb()
	// 开启了
	// 实例中有主键，可以不写where
	// db.Delete(&user)
	// db.Delete(&User{}, 10)
	db.Where("id = ?", 1).Delete(&models.UserModel{})
	// UPDATE `user` SET `delete_time`='2022-10-28 15:45:38.108' WHERE id = 1 AND `user`.`delete_time` IS NULL
}

// 物理删除
func ForceDelete() {
	db, _ := GetDb()
	db.Unscoped().Delete(&models.UserModel{}, 10)
	// DELETE FROM `user` WHERE `user`.`id` = 10
}

```

## 查询操作
```go
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

```