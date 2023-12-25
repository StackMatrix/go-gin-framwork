package services

import (
	"errors"
	"rat_server/app/models"
	"rat_server/app/pkg/request"
	"rat_server/global"
	"rat_server/utils"
	"strconv"

	"rat_server/app/pkg/response/auth"

	"github.com/gin-gonic/gin"
)

type userService struct {
}

var UserService = new(userService)

// Register 注册
func (userService *userService) Register(params request.Register) (err error, user models.User) {
	var result = global.App.DB.Where("account = ?", params.Account).Select("id").First(&models.User{})
	if result.RowsAffected != 0 {
		err = errors.New("账号已存在")
		return
	}
	user = models.User{Name: params.Name, Account: params.Account, Password: utils.BcryptMake([]byte(params.Password))}
	err = global.App.DB.Create(&user).Error
	return
}

// Login 登录
func (userService *userService) Login(c *gin.Context, params request.Login) (data interface{}, err error) {
	// 将查找到的帐号保存到结构体中
	user := models.User{}

	err = global.App.DB.Where("account = ?", params.Account).First(&user).Error
	if err != nil || !utils.BcryptMakeCheck([]byte(params.Password), user.Password) {
		err = errors.New("账号不存在或密码错误")
	}

	// 调用 JwtService 服务，颁发 Token
	tokenData, err, _ := JwtService.CreateToken(AppGuardName, &user)
	if err != nil {
		err = errors.New("生成口令失败")
		return
	}

	// 返回的数据
	data = auth.LoginResponseData{
		Jwt: tokenData.AccessToken,
	}

	return
}

// GetUserInfo 获取用户信息
func (userService *userService) GetUserInfo(id string) (err error, user models.User) {
	intId, err := strconv.Atoi(id)
	err = global.App.DB.First(&user, intId).Error
	if err != nil {
		err = errors.New("数据不存在")
	}
	return
}
