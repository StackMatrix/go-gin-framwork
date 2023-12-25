package response

import (
	"net/http"
	"os"
	"rat_server/global"

	"github.com/gin-gonic/gin"
)

// Response 响应结构体
type Response struct {
	ErrorCode int         `json:"status"` // 自定义错误码
	Data      interface{} `json:"data"`   // 数据
	Message   string      `json:"msg"`    // 信息
}

// Success 响应成功 status 为 200 表示成功
func Success(c *gin.Context, data interface{}, msg string) {
	c.JSON(http.StatusOK, Response{
		200,
		data,
		msg,
	})
}

// Fail 响应失败 status 不为 1 表示失败
func Fail(c *gin.Context, status int, msg string) {
	c.JSON(http.StatusOK, Response{
		status,
		nil,
		msg,
	})
}

// FailByError 失败响应 返回自定义错误的错误码、错误信息
func FailByError(c *gin.Context, error global.CustomError) {
	Fail(c, error.ErrorCode, error.ErrorMsg)
}

// ValidateFail 请求参数验证失败
func ValidateFail(c *gin.Context, msg string) {
	Fail(c, global.Errors.ValidateError.ErrorCode, msg)
}

// BusinessFail 业务逻辑失败
func BusinessFail(c *gin.Context, msg string) {
	Fail(c, global.Errors.BusinessError.ErrorCode, msg)
}

// TokenFail 令牌鉴权
func TokenFail(c *gin.Context) {
	FailByError(c, global.Errors.TokenError)
}

// ServerError 中间件
func ServerError(c *gin.Context, err interface{}) {
	msg := "Internal Server Error"
	// 非生产环境显示具体错误信息
	if global.App.Config.App.Env != "production" && os.Getenv(gin.EnvGinMode) != gin.ReleaseMode {
		if _, ok := err.(error); ok {
			msg = err.(error).Error()
		}
	}
	c.JSON(http.StatusInternalServerError, Response{
		http.StatusInternalServerError,
		nil,
		msg,
	})
	c.Abort()
}
