package middleware

import (
	"Gin_todo/pkg/e"
	"Gin_todo/pkg/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := 200
		//var data interface{}
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 30004
		} else {
			claim, err := utils.ParseToken(token)
			if err != nil {
				code = 30001 // 无权限
			} else if time.Now().Unix() > claim.ExpiresAt {
				code = 30002 // Token过期
			}
		}
		if code != 200 {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
