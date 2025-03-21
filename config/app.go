package config

import "codeReleaseSystem/pkg/config"

func init() {
	config.Add("app", func() map[string]interface{} {
		return map[string]interface{}{
			// 应用名称
			"name": config.Env("APP_NAME", "codeReleaseSystem"),

			// 当前环境，用以区分多环境，一般为 local, stage, production, test
			"env": config.Env("APP_ENV", "production"),
		}
	})
}
