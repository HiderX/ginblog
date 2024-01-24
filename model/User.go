package model

import (
	"encoding/base64"
	"errors"
	"ginblog/utils"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int;DEFAULT:1" json:"role"`
}

func CreateUser(data *User) int {
	data.Password = ScryptPw(data.Password)
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

func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 43, 54, 65, 76, 87, 98}
	key, err := scrypt.Key([]byte(password), salt, 1<<15, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	str := base64.StdEncoding.EncodeToString(key)
	return str
}
