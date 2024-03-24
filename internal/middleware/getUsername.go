package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetUsername() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		username := session.Get("username").(string)

		ctx.Set("username", username)
		ctx.Next()
	}
}
