package controllers

import (
	"micro_uploads/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewFrontControllers(r *gin.Engine) FrontControllers {
	return FrontControllers{R: r}
}

func (fc *FrontControllers) StartRoutes() {
	fc.R.GET("/", middleware.GetUsername(), fc.index)
	fc.R.GET("/user", middleware.GetUsername(), fc.user)
}

func (fc FrontControllers) isLogged(username string) bool {
	return username != ""
}

func (fc FrontControllers) index(ctx *gin.Context) {
	username := ctx.GetString("username")
	logged := fc.isLogged(username)

	ctx.HTML(http.StatusOK, "index", gin.H{
		"title":  "micro uploads",
		"logged": logged,
	})
}

func (fc FrontControllers) user(ctx *gin.Context) {
	//username := ctx.GetString("username")

	ctx.HTML(http.StatusOK, "user", gin.H{
		"title":  "user",
		"logged": false,
		"files": []struct{ Name, Link string }{
			struct {
				Name string
				Link string
			}{Name: "bruh", Link: "link1"},
			struct {
				Name string
				Link string
			}{Name: "dah", Link: "link2"},
		},
	})
}
