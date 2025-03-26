package bootstrap

import (
	"codeReleaseSystem/app/http/middlewares"
	"codeReleaseSystem/routes"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetupRoute(router *gin.Engine) {
	// 注册全局中间件
	registerGlobalMiddleware(router)

	// 注册 API 路由
	routes.RegisterAPIRoutes(router)

	// 配置 404 路由
	setup404Handler(router)

}

func registerGlobalMiddleware(router *gin.Engine) {
	router.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
	)
}

func setup404Handler(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		// 获取头信息的 accept 信息
		acceptString := c.Request.Header.Get("Accept")
		fmt.Println("acceptString", acceptString)
		if strings.Contains(acceptString, "text/html") {
			// HTML返回页面错误信息
			c.String(http.StatusNotFound, "<h1>404 Page Not Found</h1>")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}
