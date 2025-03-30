package bootstrap

import (
	"codeReleaseSystem/pkg/captcha"
)

// SetupCaptcha 初始化验证码
func SetupCaptcha() {
	store := captcha.NewRedisStore()
	captcha.New(store)
}
