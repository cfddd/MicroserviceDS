package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// ServeMiddleware 服务发现的中间件，将服务放到gin的上下文中
// 这个中间件的作用是在每个 HTTP 请求处理过程中，将指定的服务实例 serveInstance 存入到 gin.Context 中的 Keys 字段，
// 以便后续的中间件或处理函数可以从 gin.Context 中获取这些服务实例。
// 在每个请求中都会创建一个独立的 Keys，避免了 map 的并发冲突
func ServeMiddleware(serveInstance map[string]interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 如果直接复制，浅拷贝会导致map冲突
		c.Keys = make(map[string]interface{})
		for key, value := range serveInstance {
			c.Keys[key] = value
		}
		fmt.Println("ServeMiddleware")
		c.Next()
	}
}
