package v1

//func UpLoad(c *gin.Context) {
//	file, fileHeader, _ := c.Request.FormFile("file")
//	ossFilePath := fmt.Sprintf("%s", fileHeader.Filename) //完整的oss路径
//	url, code := model.UploadFile(file, ossFilePath)
//	c.JSON(http.StatusOK, gin.H{
//		"status":  code,
//		"message": errmsg.GetErrMsg(code),
//		"url":     url,
//	})
//
//}
