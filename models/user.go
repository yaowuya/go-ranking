package models

import (
	"go-ranking/dao"
	"time"
)

type User struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	AddTime    int64  `json:"AddTime"`
	UpdateTime int64  `json:"UpdateTime"`
}

func (User) TableName() string {
	return "user"
}

func GetUserInfoByUsername(username string) (User, error) {
	var user User
	err := dao.Db.Where("username =?", username).First(&user).Error
	return user, err
}

func GetUserTest(id int) (User, error) {
	var user User
	err := dao.Db.Where("id =?", id).First(&user).Error
	return user, err
}

func GetUserList(id int) ([]User, error) {
	var users []User
	err := dao.Db.Where("id < ?", id).Find(&users).Error
	return users, err
}

func AddUser(username string, password string) (int, error) {
	user := User{Username: username, Password: password, AddTime: time.Now().Unix(), UpdateTime: time.Now().Unix()}
	err := dao.Db.Create(&user).Error
	return user.Id, err
}

func UpdateUser(id int, username string) {
	dao.Db.Model(&User{}).Where("id = ?", id).Update("username", username)
}

func DeleteUser(id int) error {
	err := dao.Db.Delete(&User{}, id).Error
	return err
}

func GetUserInfo(id int) (User, error) {
	var user User
	err := dao.Db.Where("id = ?", id).First(&user).Error
	return user, err
}
