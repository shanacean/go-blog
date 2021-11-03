package model

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go-blog/utils"
	"go-blog/utils/result"
	"mime/multipart"
)

var AccessKey = utils.AccessKey
var SecretKey = utils.SecretKey
var Bucket = utils.Bucket
var Domain = utils.Domain

func UploadFile(file multipart.File, fileSize int64) (string, result.Code) {
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}

	fmt.Println(AccessKey)

	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)

	fmt.Println(upToken)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	putExtra := storage.PutExtra{}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)

	if err != nil {
		fmt.Println(err)
		return "", result.ERROR
	}

	url := Domain + ret.Key

	return url, result.SUCCESS
}
