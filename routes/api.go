package routes

import (
	"codeReleaseSystem/app/http/controllers/api/v1/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {
	// v1路由组
	v1 := r.Group("/v1")
	{
		// 注册一个首页目录
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "hello world",
			})
		})

		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			// 判断手机是否已注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			// 判断邮箱是否已注册
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)

			// 验证码相关
			cc := new(auth.CaptchaController)
			authGroup.GET("/captcha", cc.Generate)
			authGroup.POST("/captcha/verify", cc.Verify)
		}
	}

}
