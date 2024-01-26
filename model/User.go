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
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2"`
}

func CreateUser(data *User) int {
	code := CheckUser(data.Username)
	if code == utils.ERROR_USERNAME_USED {
		return utils.ERROR_USERNAME_USED
	}
	data.Password = ScryptPw(data.Password)
	err = Db.Create(&data).Error
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

func GetUsers(pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var offset int
	var total int64
	if pageSize == -1 {
		offset = -1
	} else {
		offset = (pageNum - 1) * pageSize
	}
	err = Db.Limit(pageSize).Offset(offset).Find(&users).Error
	Db.Model(&users).Count(&total)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0
	}
	return users, total
}

func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	Db.Select("id").Where("id = ?", id).First(&user)
	if user.ID <= 0 {
		return utils.ERROR_USER_NOT_EXIST
	}
	code := CheckUser(data.Username)
	if code == utils.ERROR_USERNAME_USED {
		return utils.ERROR_USERNAME_USED
	}
	err = Db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return utils.ERROR
	}
	return utils.SUCCESS
}

func DeleteUser(id int) int {
	var user User
	Db.Select("id").Where("id = ?", id).First(&user)
	if user.ID <= 0 {
		return utils.ERROR_USER_NOT_EXIST
	}
	err = Db.Where("id = ?", id).Delete(&user).Error
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

func CheckLogin(username string, password string) int {
	var user User
	Db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return utils.ERROR_USER_NOT_EXIST
	}
	if ScryptPw(password) != user.Password {
		return utils.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return utils.ERROR_USER_NO_RIGHT
	}
	return utils.SUCCESS
}
