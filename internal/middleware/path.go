package middleware

import (
	"webconsole/pkg/respcode"

	"github.com/gin-gonic/gin"
)

// 获取查询参数的中间件
func QueryParse(c *gin.Context) {
	column := c.Query("column")
	value := c.Query("value")

	if column == "" || value == "" {
		respcode.ResponseError(c, respcode.CodeInvalidPath)
		c.Abort()
	} else {
		c.Set("column", column)
		c.Set("value", value)
		c.Next()
	}
}
