package auth

import (
	"codeReleaseSystem/pkg/captcha"
	"codeReleaseSystem/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CaptchaController struct{}

// Generate 生成验证码
func (cc *CaptchaController) Generate(c *gin.Context) {
	// 获取验证码实例
	cp := captcha.New(captcha.NewRedisStore())

	id, b64, err := cp.Generate()
	if err != nil {
		response.Abort500(c, "生成验证码失败")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"captcha_id":    id,
		"captcha_image": b64,
	})
}

// Verify 验证验证码
func (cc *CaptchaController) Verify(c *gin.Context) {
	type captchaVerifyRequest struct {
		CaptchaID     string `json:"captcha_id" binding:"required"`
		CaptchaAnswer string `json:"captcha_answer" binding:"required"`
	}

	var request captchaVerifyRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Abort404(c, "验证码参数错误")
		return
	}

	cp := captcha.New(captcha.NewRedisStore())
	res := cp.Verify(request.CaptchaID, request.CaptchaAnswer)

	response.Result(c, res)
}
