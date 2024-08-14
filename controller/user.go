package controller

import (
	"github.com/gin-gonic/gin"
	"go-ranking/models"
	"go-ranking/pkg/logger"
	"strconv"
)

type UserController struct {
}

func (u UserController) GetUserInfo(c *gin.Context) {
	idStr := c.Param("id")
	name := c.Param("name")

	id, _ := strconv.Atoi(idStr)
	user, _ := models.GetUserTest(id)
	ReturnSuccess(c, 0, name, user, 1)
}

func (u UserController) GetList(c *gin.Context) {
	logger.Write("日志信息", "user")
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println("捕获异常", err)
	//	}
	//}()
	num1 := 1
	num2 := 0
	num3 := num1 / num2
	ReturnError(c, 404, num3)
}
