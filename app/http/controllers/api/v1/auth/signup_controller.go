package auth

import (
	v1 "codeReleaseSystem/app/http/controllers/api/v1"
	"codeReleaseSystem/app/models/user"
	"codeReleaseSystem/app/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

// IsPhoneExist 检测手机号是否被注册
func (sc *SignupController) IsPhoneExist(c *gin.Context) {

	param := requests.SignupPhoneExistRequest{}
	if ok, errors := requests.ValidateRequest(c, &param, requests.ValidateSignupPhoneExist); !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errors})
		return
	}

	// 返回响应
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(param.Phone),
	})
}

func (sc *SignupController) IsEmailExist(c *gin.Context) {
	param := requests.SignupEmailExistRequest{}
	if ok, errs := requests.ValidateRequest(c, &param, requests.ValidateSignupEmailExist); !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errs})
		return
	}

	// 返回响应
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsEmailExist(param.Email),
	})
}
