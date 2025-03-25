package auth

import (
	v1 "codeReleaseSystem/app/http/controllers/api/v1"
	"codeReleaseSystem/app/models/user"
	"codeReleaseSystem/app/requests"
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
	param := requests.SignupPhoneExistRequest{}
	if err := c.ShouldBindJSON(&param); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 表单验证
	errs := requests.ValidateSignupPhoneExist(&param, c)
	if len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": errs,
		})
		return
	}

	// 返回响应
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(param.Phone),
	})
}

func (sc *SignupController) IsEmailExist(c *gin.Context) {
	// 获取请求参数
	param := requests.SignupEmailExistRequest{}
	if err := c.ShouldBindJSON(&param); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"param_error": err.Error(),
		})
		return
	}

	// 表单验证
	errs := requests.ValidateSignupEmailExist(&param, c)
	if len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": errs,
		})
		return
	}

	// 返回响应
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsEmailExist(param.Email),
	})
}
