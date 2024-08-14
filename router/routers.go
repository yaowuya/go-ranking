package router

import (
	"github.com/gin-gonic/gin"
	"go-ranking/controller"
	"go-ranking/pkg/logger"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)
	user := r.Group("/user")
	{
		user.GET("/info/:id", controller.UserController{}.GetUserInfo)
		user.GET("/list", controller.UserController{}.GetList)
		user.POST("/add", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "user add")
		})
		user.DELETE("/delete", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "user delete")
		})
	}

	order := r.Group("/order")
	{
		order.POST("/list", controller.OrderController{}.GetList)
	}
	return r
}
