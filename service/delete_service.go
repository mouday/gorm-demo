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
