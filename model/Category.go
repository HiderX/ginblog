package model

import (
	"errors"
	"ginblog/utils"
	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

func CheckCategory(name string) int {
	var category Category
	Db.Select("id").Where("name = ?", name).First(&category)
	if category.ID > 0 {
		return utils.ERROR_CATEGORY_USED
	}
	return utils.SUCCESS
}

func CreateCategory(data *Category) int {
	code := CheckCategory(data.Name)
	if code == utils.ERROR_CATEGORY_USED {
		return utils.ERROR_CATEGORY_USED
	}
	err = Db.Create(&data).Error
	if err != nil {
		return utils.ERROR
	}
	return utils.SUCCESS
}

func GetCategories(pageSize int, pageNum int) ([]Category, int64) {
	var categories []Category
	var offset int
	var total int64
	if pageSize == -1 {
		offset = -1
	} else {
		offset = (pageNum - 1) * pageSize
	}
	err = Db.Limit(pageSize).Offset(offset).Find(&categories).Error
	Db.Model(&categories).Count(&total)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0
	}
	return categories, total
}

func EditCategory(id int, data *Category) int {
	var category Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	Db.Select("id").Where("id = ?", id).First(&category)
	if category.ID <= 0 {
		return utils.ERROR_CATEGORY_NOT_EXIST
	}
	code := CheckCategory(data.Name)
	if code == utils.ERROR_CATEGORY_USED {
		return utils.ERROR_CATEGORY_USED
	}
	err = Db.Model(&category).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return utils.ERROR
	}
	return utils.SUCCESS
}

func DeleteCategory(id int) int {
	var category Category
	Db.Select("id").Where("id = ?", id).First(&category)
	if category.ID <= 0 {
		return utils.ERROR_CATEGORY_NOT_EXIST
	}
	err = Db.Delete(&category).Error
	if err != nil {
		return utils.ERROR
	}
	return utils.SUCCESS
}
