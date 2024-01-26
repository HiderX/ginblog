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

func GetCateArt(id int, pageSize int, pageNum int) ([]Article, int, int64) {
	var cateArt []Article
	var offset int
	var total int64
	if pageSize == -1 {
		offset = -1
	} else {
		offset = (pageNum - 1) * pageSize
	}
	err := Db.Preload("Category").Where("cid = ?", id).Limit(pageSize).Offset(offset).Find(&cateArt).Error
	Db.Model(&cateArt).Count(&total)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.ERROR, 0
	}
	return cateArt, utils.SUCCESS, total
}

func GetArtInfo(id int) (Article, int) {
	var art Article
	err := Db.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return art, utils.ERROR_ART_NOT_EXIST
	}
	return art, utils.SUCCESS
}

func GetArt(pageSize int, pageNum int) ([]Article, int, int64) {
	var articles []Article
	var offset int
	var total int64
	if pageSize == -1 {
		offset = -1
	} else {
		offset = (pageNum - 1) * pageSize
	}
	err := Db.Preload("Category").Limit(pageSize).Offset(offset).Find(&articles).Error
	Db.Model(&articles).Count(&total)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.ERROR, 0
	}
	return articles, utils.SUCCESS, total

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
