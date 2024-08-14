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

func (u UserController) GetUserList(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	users, err := models.GetUserList(id)
	if err != nil {
		ReturnError(c, 4004, "没有相关数据")
		return
	}
	ReturnSuccess(c, 2000, "获取成功", users, 1)
}

func (u UserController) AddUser(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	id, err := models.AddUser(username)
	if err != nil {
		ReturnError(c, 4002, "保存错误")
	}
	ReturnSuccess(c, 2000, "保存成功", id, 1)
}

func (u UserController) UpdateUser(c *gin.Context) {
	idStr := c.DefaultPostForm("id", "")
	username := c.DefaultPostForm("username", "")
	id, _ := strconv.Atoi(idStr)
	models.UpdateUser(id, username)
	ReturnSuccess(c, 2000, "更新成功", id, 1)
}

func (u UserController) DeleteUser(c *gin.Context) {
	idStr := c.DefaultPostForm("id", "")
	id, _ := strconv.Atoi(idStr)
	err := models.DeleteUser(id)
	if err != nil {
		ReturnError(c, 4003, "删除失败")
	}
	ReturnSuccess(c, 2000, "删除成功", id, 1)
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
