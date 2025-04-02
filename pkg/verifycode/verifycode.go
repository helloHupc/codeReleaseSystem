package verifycode

import (
	"codeReleaseSystem/pkg/config"
	"codeReleaseSystem/pkg/helpers"
	"codeReleaseSystem/pkg/logger"
	"codeReleaseSystem/pkg/mail"
	"codeReleaseSystem/pkg/redis"
	"time"

	"github.com/gin-gonic/gin"
)

func SendEmailVerifyCode(email string) bool {
	// 生成验证码
	code := generateCode()

	// 发送邮件
	emailSent := mail.NewMailer().Send(mail.Email{
		From: mail.From{
			Address: config.GetString("mail.from.address"),
			Name:    config.GetString("mail.from.name"),
		},
		To:      []string{email},
		Subject: "验证码",
		Text:    []byte("您的验证码是：" + code),
		Html:    []byte("<h1>您的验证码是：" + code + "</h1>"),
	})
	// 发送邮件失败
	if !emailSent {
		logger.ErrorJSON("Mail", "邮件发送失败", "邮件发送失败，请检查邮箱地址或邮件配置")
		return false
	}

	// 存储验证码到缓存
	redisKey := config.GetString("verifycode.email_prefix") + email
	redis.Redis.Set(redisKey, code, time.Duration(config.GetInt("verifycode.expire_time"))*time.Minute)

	logger.DebugJSON("Mail", "邮件发送成功", gin.H{
		"email": email,
		"code":  code,
	})
	return true
}

// 生成验证码
func generateCode() string {
	// 生成6位数字验证码
	code := helpers.GenerateRandomNumber(config.GetInt("verifycode.code_length"))
	return code
}
