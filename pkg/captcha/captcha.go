package captcha

import (
	"codeReleaseSystem/config"

	"github.com/mojocn/base64Captcha"
)

type Captcha struct {
	store  Store
	driver base64Captcha.Driver
}

func New(store Store) *Captcha {
	// 获取验证码配置
	height, width, length, maxskew, dotcount := config.GetCaptchaConfig()

	// 配置验证码驱动
	driver := base64Captcha.DriverDigit{
		Height:   height,
		Width:    width,
		Length:   length,
		MaxSkew:  maxskew,
		DotCount: dotcount,
	}

	return &Captcha{
		store:  store,
		driver: &driver,
	}
}

// Generate 生成验证码
func (c *Captcha) Generate() (id string, b64 string, err error) {
	// 生成验证码
	cp := base64Captcha.NewCaptcha(c.driver, base64Captcha.DefaultMemStore)
	id, b64, answer, err := cp.Generate()
	if err != nil {
		return "", "", err
	}

	// 存储验证码到Redis
	if err := c.store.Set(id, answer, config.GetCaptchaExpire()); err != nil {
		return "", "", err
	}

	return id, b64, nil
}

// Verify 验证验证码
func (c *Captcha) Verify(id string, answer string) bool {
	// 从Redis获取验证码
	storedAnswer, err := c.store.Get(id)
	if err != nil {
		return false
	}

	// 验证后删除验证码
	defer c.store.Delete(id)

	return storedAnswer == answer
}
