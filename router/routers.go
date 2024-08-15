package router

import (
	"github.com/gin-contrib/sessions"
	sessionsredis "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"go-ranking/config"
	"go-ranking/controller"
	"go-ranking/pkg/logger"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)
	store, _ := sessionsredis.NewStore(10, "tcp", config.RedisAddress, "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	user := r.Group("/user")
	{
		user.POST("/login/", controller.UserController{}.Login)
		user.POST("/register/", controller.UserController{}.Register)
	}

	order := r.Group("/order")
	{
		order.POST("/list", controller.OrderController{}.GetList)
	}
	return r
}
