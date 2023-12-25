package app

import (
	"github.com/dgrijalva/jwt-go"
	"rat_server/app/common/request"
	"rat_server/app/common/response"
	"rat_server/app/services"

	"github.com/gin-gonic/gin"
)

// Register 用户注册
func Register(c *gin.Context) {
	var form request.RegisterModel
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	err, user := services.UserService.Register(form)
	if err != nil {
		response.BusinessFail(c, err.Error())
	}

	response.Success(c, user, "注册成功")
}

// Login 用户登录
func Login(c *gin.Context) {
	var form request.LoginModel
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	data, err := services.UserService.Login(c, form)
	if err != nil {
		response.BusinessFail(c, err.Error())
	}

	response.Success(c, data, "登陆成功")
}

// Logout 退出登录
func Logout(c *gin.Context) {
	if err := services.JwtService.JoinBlackList(c.Keys["token"].(*jwt.Token)); err != nil {
		response.BusinessFail(c, "登出失败")
		return
	}

	response.Success(c, nil, "退出成功")
}

// Info 用户信息
func Info(c *gin.Context) {
	err, user := services.UserService.GetUserInfo(c.Keys["id"].(string))
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}

	response.Success(c, user, "查询成功")
}
