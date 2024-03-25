package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUserNotLogged(t *testing.T) {
	assert := assert.New(t)
	r := gin.New()
	store := cookie.NewStore([]byte("secret"))

	r.Use(sessions.Sessions("mysession", store))
	r.Use(GetUsername())

	r.GET("/", func(ctx *gin.Context) {
		username := ctx.GetString("username")
		ctx.String(http.StatusOK, username)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	resp := w.Body.String()
	assert.Empty(resp)
}

func TestUserLogged(t *testing.T) {
	assert := assert.New(t)
	r := gin.New()
	store := cookie.NewStore([]byte("secret"))

	r.Use(sessions.Sessions("mysession", store))
	r.Use(func(ctx *gin.Context) {
		s := sessions.Default(ctx)
		s.Set("username", "testuser")
		s.Save()
	})
	r.Use(GetUsername())

	r.GET("/", func(ctx *gin.Context) {
		username := ctx.GetString("username")
		ctx.String(http.StatusOK, username)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	resp := w.Body.String()
	assert.NotEmpty(resp)
}
