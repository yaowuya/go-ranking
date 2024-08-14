package router

import (
	"github.com/gin-gonic/gin"
	"go-ranking/controller"
	"go-ranking/pkg/logger"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)
	user := r.Group("/user")
	{
		user.GET("/info/:id", controller.UserController{}.GetUserInfo)
		user.GET("/list/test", controller.UserController{}.GetList)
		user.GET("/list/:id", controller.UserController{}.GetUserList)
		user.POST("/add/", controller.UserController{}.AddUser)
		user.POST("/update/", controller.UserController{}.UpdateUser)
		user.POST("/delete/", controller.UserController{}.DeleteUser)
	}

	order := r.Group("/order")
	{
		order.POST("/list", controller.OrderController{}.GetList)
	}
	return r
}
