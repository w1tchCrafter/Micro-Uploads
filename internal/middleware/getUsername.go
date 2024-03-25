package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetUsername() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)

		if username, ok := session.Get("username").(string); ok {
			ctx.Set("username", username)
		} else {
			ctx.Set("username", "")
		}

		ctx.Next()
	}
}
