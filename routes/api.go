package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterAPIRoutes(r *gin.Engine) {
	// v1路由组
	v1 := r.Group("/v1")
	{
		// 注册一个首页目录
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "hello world",
			})
		})
	}

}
