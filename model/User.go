package model

import (
	"errors"
	"ginblog/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int;DEFAULT:1" json:"role"`
}

func CreateUser(data *User) int {
	err := Db.Create(&data).Error
	if err != nil {
		return utils.ERROR
	}
	return utils.SUCCESS
}

func CheckUser(username string) int {
	var user User
	Db.Select("id").Where("username = ?", username).First(&user)
	if user.ID > 0 {
		return utils.ERROR_USERNAME_USED
	}
	return utils.SUCCESS
}

func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err := Db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return users
}

func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err := Db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return utils.ERROR
	}
	return utils.SUCCESS
}

func DeleteUser(id int) int {
	var user User
	err := Db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return utils.ERROR
	}
	return utils.SUCCESS
}
