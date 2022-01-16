package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Token() gin.HandlerFunc {
	return func(c *gin.Context) {
		apitoken := c.Request.Header.Get("ApiToken")
		method := c.Request.Method
		if apitoken == "" && method == "POST" {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "接口参数错误",
			})
			c.Abort()
			return
		} else {
			// 这里进行验证

		}
		c.Next()
	}
}
