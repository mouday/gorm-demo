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
