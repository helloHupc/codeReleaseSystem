package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// defaultMessage 获取默认消息，如果未传参则返回默认值
func defaultMessage(defaultMsg string, msg ...string) string {
	if len(msg) == 0 {
		return defaultMsg
	}
	return msg[0]
}

// JSON 响应 200 和 JSON 数据
func JSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

// Success 响应 200 和预设『操作成功！』的 JSON 数据
func Success(c *gin.Context) {
	JSON(c, gin.H{
		"message": "操作成功！",
		"result":  true,
	})
}

// SuccessData 响应 200 和带 data 键的 JSON 数据
func SuccessData(c *gin.Context, data interface{}) {
	JSON(c, gin.H{
		"message": "操作成功！",
		"result":  true,
		"data":    data,
	})
}

func Result(c *gin.Context, result interface{}, msg ...string) {
	JSON(c, gin.H{
		"message": defaultMessage("message", msg...),
		"result":  result,
	})

}

// Abort404 响应 404，未传参 msg 时使用默认消息
func Abort404(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"message": defaultMessage("数据不存在，请确定请求正确", msg...),
	})
}

// Abort403 响应 403，未传参 msg 时使用默认消息
func Abort403(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"message": defaultMessage("权限不足，请确定您有对应的权限", msg...),
	})
}

// Abort500 响应 500，未传参 msg 时使用默认消息
func Abort500(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"message": defaultMessage("服务器发生错误，请稍后再试", msg...),
	})
}

// BadRequest 响应 400，传参 err 对象，未传参 msg 时使用默认消息
func BadRequest(c *gin.Context, err error, msg ...string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": defaultMessage("请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。", msg...),
		"error":   err.Error(),
	})
}
