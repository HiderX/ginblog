package model

import (
	"errors"
	"ginblog/utils"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title    string   `gorm:"type:varchar(100);not null" json:"title"`
	Category Category `gorm:"foreignkey:Cid"`
	Cid      int      `gorm:"type:int;not null" json:"cid"`
	Desc     string   `gorm:"type:varchar(200)" json:"desc"`
	Content  string   `gorm:"type:longtext" json:"content"`
	Img      string   `gorm:"type:varchar(100)" json:"img"`
}

func CreateArt(data *Article) int {
	err := Db.Create(&data).Error
	if err != nil {
		return utils.SUCCESS
	}
	return utils.ERROR
}

func GetArt(pageSize int, pageNum int) []Article {
	var articles []Article
	var offset int
	if pageSize == -1 {
		offset = -1
	} else {
		offset = (pageNum - 1) * pageSize
	}
	err := Db.Limit(pageSize).Offset(offset).Find(&articles).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return articles

}

func EditArt(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	Db.Select("id").Where("id = ?", id).First(&art)
	if art.ID <= 0 {
		return utils.ERROR_ART_NOT_EXIST
	}
	err := Db.Model(&art).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return utils.ERROR
	}
	return utils.SUCCESS
}

func DeleteArt(id int) int {
	var art Article
	Db.Select("id").Where("id = ?", id).First(&art)
	if art.ID <= 0 {
		return utils.ERROR_ART_NOT_EXIST
	}
	err := Db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return utils.ERROR
	}
	return utils.SUCCESS
}
