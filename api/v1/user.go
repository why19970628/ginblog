package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"ginblog/utils/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加用户
func AddUser(c *gin.Context)  {
	var user *model.User

	_ = c.ShouldBindJSON(&user)
	msg, code := validator.Validate(&user)
	if code != errmsg.SUCCSE{
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"sessage": msg,
		})
		return
	}

	code = model.CheckUser(user.Username)
	if code == errmsg.SUCCSE{
		model.CreateUser(user)
	}
	if code == errmsg.ERROR_USERNAME_USED{
		code = errmsg.ERROR_USERNAME_USED
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"sessage": errmsg.GetErrMsg(code),
	})
}

// 查询用户列表
func GetUsers(c *gin.Context){
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, total := model.GetUsers(pageSize, pageNum)
	code := errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})

}


// 编辑用户
func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCSE {
		model.EditUser(id, &data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteUser(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}