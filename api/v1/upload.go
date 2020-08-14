package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
)

func UpLoad(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	ossFilePath := fmt.Sprintf("%s", fileHeader.Filename) //完整的oss路径
	url, code := model.UploadFile(file, ossFilePath)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"url":     url,
	})

}
