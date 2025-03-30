package config

import (
	"codeReleaseSystem/pkg/app"
	"codeReleaseSystem/pkg/config"
)

func init() {
	config.Add("captcha", func() map[string]interface{} {
		return map[string]interface{}{
			"height":            config.Env("CAPTCHA_HEIGHT", 80),
			"width":             config.Env("CAPTCHA_WIDTH", 240),
			"length":            config.Env("CAPTCHA_LENGTH", 6),
			"maxskew":           config.Env("CAPTCHA_MAX_SKEW", 0.7),
			"dotcount":          config.Env("CAPTCHA_DOT_COUNT", 80),
			"expire_time":       config.Env("CAPTCHA_EXPIRE_TIME", 15),          // 分钟
			"debug_expire_time": config.Env("CAPTCHA_DEBUG_EXPIRE_TIME", 10080), // 分钟
		}
	})
}

// GetCaptchaExpire 获取验证码过期时间(秒)
func GetCaptchaExpire() int {
	if app.IsLocal() {
		return config.GetInt("captcha.debug_expire_time") * 60
	}
	return config.GetInt("captcha.expire_time") * 60
}

// GetCaptchaConfig 获取验证码配置
func GetCaptchaConfig() (height, width, length int, maxskew float64, dotcount int) {
	return config.GetInt("captcha.height"),
		config.GetInt("captcha.width"),
		config.GetInt("captcha.length"),
		config.GetFloat64("captcha.maxskew"),
		config.GetInt("captcha.dotcount")
}
