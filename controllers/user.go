package controllers

import (
	"fmt"
	"ginStudy/dao"
	"ginStudy/models"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type WebUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserApi struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
}

type UserController struct {
}

func (u UserController) Register(c *gin.Context) {
	// username := c.DefaultPostForm("username", "")
	// password := c.DefaultPostForm("password", "")
	var webuser WebUser
	if err := c.ShouldBindJSON(&webuser); err != nil {
		ReturnError(c, 4001, "参数错误")
		return
	}
	username := webuser.Username
	password := webuser.Password
	fmt.Println("username:", username, "password:", password)
	// TODO: save user to database

	if dao.Db == nil {
		fmt.Println("数据库连接为空")
		ReturnError(c, 5001, "数据库连接为空")
		return
	}

	fmt.Println("try to save user to database")
	user, _ := models.GetUserInfoByUsername(username)
	fmt.Println("has queried")
	if user.ID != 0 {
		ReturnError(c, 4001, "用户名已存在")
		return
	}
	fmt.Println("username is available")
	_, err2 := models.AddUser(username, EncryMd5(password))
	if err2 != nil {
		ReturnError(c, 5001, "保存用户信息失败,请联系管理员")
		return
	}
	fmt.Println("save user to database success")
	ReturnSuccess(c, 0, "注册成功", nil, 0)
}

func (u UserController) Login(c *gin.Context) {
	var webuser WebUser
	if err := c.ShouldBindJSON(&webuser); err != nil {
		ReturnError(c, 4001, "参数错误")
		return
	}
	username := webuser.Username
	password := webuser.Password
	fmt.Println("username:", username, "password:", password)
	// TODO: check user login
	user, _ := models.GetUserInfoByUsername(username)
	if user.ID == 0 {
		ReturnError(c, 4001, "用户名不存在")
		fmt.Println("username not exist")
		return
	}
	if user.Password != EncryMd5(password) {
		ReturnError(c, 4001, "用户名或密码错误")
		fmt.Println("username or password error")
		return
	}

	// TODO: set session
	fmt.Println("login success")
	session := sessions.Default(c)
	session.Set("login"+strconv.Itoa(user.ID), user.ID)
	session.Save()

	data := UserApi{Id: user.ID, UserName: user.Username}
	ReturnSuccess(c, 0, "登录成功", data, 1)
	fmt.Println("login success")
}

func init() {
	fmt.Println("init user controller")

}
