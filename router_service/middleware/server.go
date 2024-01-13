package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

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
