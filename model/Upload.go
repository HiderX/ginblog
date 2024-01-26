package model

import (
	"context"
	"ginblog/utils"
	"mime/multipart"
)
import "github.com/qiniu/go-sdk/v7/auth/qbox"
import "github.com/qiniu/go-sdk/v7/storage"

var AccessKey = utils.AccessKey
var SecretKey = utils.SecretKey
var Bucket = utils.Bucket
var QiniuServer = utils.QiniuServer

func UpLoadFile(file multipart.File, filesize int64) (string, int) {
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, filesize, &putExtra)
	if err != nil {
		return "", utils.ERROR
	}
	url := QiniuServer + ret.Key
	return url, utils.SUCCESS
}
