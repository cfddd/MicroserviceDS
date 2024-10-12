package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"utils/exception"
)

// ErrorMiddleWare 错误处理中间件，捕获panic抛出异常
func ErrorMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			r := recover() // 捕获panic报错
			if r != nil {
				c.JSON(http.StatusOK, gin.H{
					"status_code": exception.ERROR,
					// 打印具体错误
					"status_msg": fmt.Sprintf("%s", r),
				})
				// 中断后续的中间件和处理函数的执行。这意味着当捕获到 panic 并返回错误响应后，Gin 将不再继续处理该请求。
				c.Abort()
			}
		}()
		c.Next()
	}
}
