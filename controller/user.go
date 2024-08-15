package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-ranking/models"
	"strconv"
)

type UserController struct {
}

type UserApi struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func (u UserController) Login(c *gin.Context) {
	// 获取参数信息
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	if username == "" || password == "" {
		ReturnError(c, 4001, "请输入正确的用户名或密码")
		return
	}
	user, _ := models.GetUserInfoByUsername(username)
	if user.Id == 0 || user.Password != EncryptMd5(password) {
		ReturnError(c, 4002, "用户名或密码错误")
		return
	}
	data := UserApi{Id: user.Id, Username: user.Username}
	session := sessions.Default(c)
	session.Set("login:"+strconv.Itoa(user.Id), user.Id)
	err := session.Save()
	if err != nil {
		ReturnError(c, 4003, "session保存失败")
		return
	}
	ReturnSuccess(c, 2000, "登陆成功", data, 1)

}

func (u UserController) Register(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	confirmPassword := c.DefaultPostForm("confirmPassword", "")
	if username == "" || password == "" || confirmPassword == "" {
		ReturnError(c, 4001, "请输入正确的用户名或密码")
		return
	}
	if password != confirmPassword {
		ReturnError(c, 4002, "密码和确认密码不一致")
		return
	}
	user, err := models.GetUserInfoByUsername(username)
	if user.Id != 0 {
		ReturnError(c, 4003, "此用户名已经存在")
		return
	}
	_, err = models.AddUser(username, EncryptMd5(password))
	if err != nil {
		ReturnError(c, 4004, "保存失败")
		return
	}
	ReturnSuccess(c, 2000, "注册成功", user, 1)
}
