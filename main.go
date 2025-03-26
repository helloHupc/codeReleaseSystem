package main

import (
	"codeReleaseSystem/bootstrap"
	btsConfig "codeReleaseSystem/config"
	"codeReleaseSystem/pkg/config"
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	btsConfig.Initialize()
}

func main() {

	// 配置初始化 依赖命令行 --env 参数
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	// 初始化 Logger
	bootstrap.SetupLogger()

	// 初始化Gin实例
	router := gin.New()

	// 初始化 DB
	bootstrap.SetupDB()

	// 初始化路由绑定
	bootstrap.SetupRoute(router)

	// 运行服务
	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		fmt.Println(err.Error())
	}
}
