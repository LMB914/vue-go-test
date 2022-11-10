package controller

import (
	"bip/app/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)
var users models.User
func UserRegister(c *gin.Context) {
	//获取参数
	users.UserName = c.PostForm("username")
	users.UserPass = c.PostForm("userpass")
	if len(users.UserName) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户名不能为空"})
		return
	}
	count := users.CheckTable(users.UserName)
	if count != 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已经存在"})
		return
	}
	if len(users.UserPass) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码必须大于6位"})
		return
	}
	log.Println(users.UserName, users.UserPass)

	users.InsertTable(&users)
	c.JSON(200, gin.H{
		"msg": "注册成功",
	})
}