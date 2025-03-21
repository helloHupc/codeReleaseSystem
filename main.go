package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	// 初始化Gin实例
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	// 处理404
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

	r.Run(":8080")

}
