package main

import (
	"codeReleaseSystem/bootstrap"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化Gin实例
	router := gin.New()

	// 初始化路由绑定
	bootstrap.SetupRoute(router)

	// 运行服务
	err := router.Run(":8080")
	if err != nil {
		fmt.Println(err.Error())
	}
}
