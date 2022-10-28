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
