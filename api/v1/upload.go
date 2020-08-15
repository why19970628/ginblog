package v1

import (
	"fmt"
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

//func UpLoad(c *gin.Context) {
//	file, fileHeader, _ := c.Request.FormFile("file")
//	fmt.Println(c.Request.FormFile("file"))
//	url, code := model.UploadFile(fileHeader.Filename, file)
//	c.JSON(http.StatusOK, gin.H{
//		"status":  code,
//		"message": errmsg.GetErrMsg(code),
//		"url":     url,
//	})
//
//}

//上传Oss文件
func UploadOssImages(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(c, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	suffix_map := map[string]string{"BMP": "", "JPG": "", "JPEG": "", "PNG": "", "GIF": ""}
	if _, ok := suffix_map[strings.ToUpper(strings.Trim(path.Ext(file.Filename), "."))]; !ok {
		fmt.Println(c, "请选择正确图片类型")
		return
	}

	file.Filename = fmt.Sprint(time.Now().Unix()) + "_" + file.Filename
	if err := c.SaveUploadedFile(file, file.Filename); err != nil {
		fmt.Println(c, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}
	url ,code := model.AliPostFile(file.Filename)
	// 删除本地文件
	if err := os.Remove(file.Filename); err != nil {
		fmt.Println(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"url":     url,
	})
	return
}

// 上传七牛
//func UpLoadFile(file multipart.File, fileSize int64) (string, int) {
//	putPolicy := storage.PutPolicy{
//		Scope: Bucket,
//	}
//	mac := qbox.NewMac(AccessKey, SecretKey)
//	upToken := putPolicy.UploadToken(mac)
//
//	cfg := storage.Config{
//		Zone:          &storage.ZoneHuadong,
//		UseCdnDomains: false,
//		UseHTTPS:      false,
//	}
//
//	putExtra := storage.PutExtra{}
//
//	formUploader := storage.NewFormUploader(&cfg)
//	ret := storage.PutRet{}
//
//	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
//	if err != nil {
//		return "", errmsg.ERROR
//	}
//	url := ImgUrl + ret.Key
//	return url, errmsg.SUCCSE
//
//}