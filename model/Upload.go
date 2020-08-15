package model

import (
	"fmt"
	"ginblog/utils"
	"ginblog/utils/errmsg"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"path/filepath"
	"time"
)

var OssAccessKey = utils.OssAccessKey
var OssSecretKey = utils.OssSecretKey
var OssBucket = utils.OssBucket
var OssEndPoint = utils.OssEndPoint
var OssFolderPrefix = utils.OssFolderPrefix
var OssUrlPrefixPublic = utils.OssUrlPrefixPublic


func AliPostFile(filename string) (url string, code int) {
	// 创建OSSClient实例。
	client, err := oss.New(OssEndPoint, OssAccessKey, OssSecretKey)
	if err != nil {
		fmt.Println("Error1 : ", err)
		return
	}
	// 获取存储空间。
	bucket, err := client.Bucket(OssBucket)
	if err != nil {
		fmt.Println("Error2 :", err)
		return
	}
	// 上传本地文件。
	folderName := time.Now().Format("2006-01-02")
	// oss路径
	yunFileTmpPath := filepath.Join(OssFolderPrefix, folderName)+"/"+filename
	err = bucket.PutObjectFromFile(yunFileTmpPath, filename)
	if err != nil {
		fmt.Println("Error3 :", err)
		return
	}
	return OssUrlPrefixPublic +"/"+ yunFileTmpPath, errmsg.SUCCSE
}
