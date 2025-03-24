package auth

import (
	v1 "codeReleaseSystem/app/http/controllers/api/v1"
	"codeReleaseSystem/app/models/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

// IsPhoneExist 检测手机号是否被注册
func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	// 获取请求参数
	phone := c.Query("phone")

	// 响应
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(phone),
	})
}
