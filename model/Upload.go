package model

import (
	"fmt"
	"ginblog/utils"
	"ginblog/utils/errmsg"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var OssAccessKey = utils.OssAccessKey
var OssSecretKey = utils.OssSecretKey
var OssBucket = utils.OssBucket
var OssEndPoint = utils.OssEndPoint
var OssFolderPrefix = utils.OssFolderPrefix
var OssUrlPrefixPublic = utils.OssUrlPrefixPublic


//上传文件
func UploadFile(localfile, uploadfile string) (string, int) {
	// 创建OSSClient实例。
	client, err := oss.New(OssEndPoint, OssAccessKey, OssSecretKey)

	// <yourObjectName>上传文件到OSS时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	// uploadfile = strings.Trim(uploadfile, "/")
	objectName := fmt.Sprintf("%s/%s", OssFolderPrefix, uploadfile) //完整的oss路径
	// <yourLocalFileName>由本地文件路径加文件名包括后缀组成，例如/users/local/myfile.txt。
	localFileName := localfile

	// 获取存储空间。
	bucket, err := client.Bucket(OssBucket)
	if err != nil {
		return "", errmsg.ERROR
	}
	// 上传文件。
	err = bucket.PutObjectFromFile(objectName, localFileName)
	if err != nil {
		return "", errmsg.ERROR
	}
	resultfile := objectName
	return resultfile, errmsg.SUCCSE
}

//获取文件列表
// func GetFilelist()(list []string,err error){
// 	list = make([]string,100)
// 	client,err := Initserver()
// 	// 获取存储空间。
// 	bucketName := viper.GetString("common.aliyunoss.bucket")
// 	bucket, err := client.Bucket(bucketName)
// 	if err != nil {
// 		return list,err
// 	}
// 	// 列举文件。
// 	marker := ""
// 	for {
// 		lsRes, err := bucket.ListObjects(oss.Marker(marker))
// 		if err != nil {
// 			return list,err
// 		}
// 		// 打印列举文件，默认情况下一次返回100条记录。
// 		for _, object := range lsRes.Objects {
// 			log.Printf("object.Key:%v\n",object.Key)
// 			list = append(list,object.Key)
// 		}
// 		if lsRes.IsTruncated {
// 			marker = lsRes.NextMarker
// 		} else {
// 			break
// 		}
// 	}
// 	return list,err
// }
