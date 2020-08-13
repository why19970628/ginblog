package routes

import (
	"ginblog/middleware"
	"ginblog/utils"
	"ginblog/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Log())
	r.Use(gin.Recovery())

	RouterV1 := r.Group("/api/v1")
	{
		RouterV1.POST("user/add", v1.AddUser)
		RouterV1.GET("users", v1.GetUsers)

	}

	return r

}
