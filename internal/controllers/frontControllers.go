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
	fc.R.GET("/", fc.index)
	fc.R.GET("/user", middleware.GetUsername(), fc.user)
}

func (fc FrontControllers) index(ctx *gin.Context) {
	// _, ok := ctx.Get("username")
	// var logged bool

	// if ok {
	// 	logged = true
	// } else {
	// 	logged = false
	// 	ctx.String(http.StatusInternalServerError, "error")
	// 	return
	// }

	// fix the above shit later
	ctx.HTML(http.StatusOK, "index", gin.H{
		"title":  "micro uploads",
		"logged": false,
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
