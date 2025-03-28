package config

import "codeReleaseSystem/pkg/config"

func init() {
	config.Add("redis", func() map[string]interface{} {
		return map[string]interface{}{
			"host":     config.Env("REDIS_HOST", "127.0.0.1"),
			"port":     config.Env("REDIS_PORT", "6379"),
			"username": config.Env("REDIS_USERNAME", ""),
			"password": config.Env("REDIS_PASSWORD", ""),

			// 业务类存储使用 1 (图片验证码、短信验证码、会话)
			"db": config.Env("REDIS_DB", 1),
		}
	})

}
