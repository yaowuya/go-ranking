package router

import (
	"github.com/gin-gonic/gin"
	"go-ranking/controller"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()
	user := r.Group("/user")
	{
		user.GET("/info/:id/:name", controller.UserController{}.GetUserInfo)
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
