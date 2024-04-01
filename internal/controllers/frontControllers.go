package controllers

import (
	"micro_uploads/internal/middleware"
	"micro_uploads/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewFrontControllers(r *gin.RouterGroup, db *gorm.DB) FrontControllers {
	fc := FrontControllers{}
	fc.R = r
	fc.DB = db
	return fc
}

func (fc *FrontControllers) StartRoutes() {
	fc.R.GET("/", middleware.GetUsername(), fc.index)
	fc.R.GET("/user", middleware.GetUsername(), fc.user)
}

func (fc FrontControllers) index(ctx *gin.Context) {
	username := ctx.GetString("username")
	logged := fc.front.IsLogged(username)

	ctx.HTML(http.StatusOK, "index", gin.H{
		"title":  "micro uploads - login",
		"logged": logged,
	})
}

func (fc FrontControllers) user(ctx *gin.Context) {
	username := ctx.GetString("username")
	logged := fc.front.IsLogged(username)
	files := make([]models.FileModel, 0)

	if logged {
		err := fc.DB.Where("author = ?", username).Find(&files).Error

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}
	} else {
		ctx.Status(http.StatusForbidden)
		return
	}

	ctx.HTML(http.StatusOK, "user", gin.H{
		"title":  "micro uploads - " + username,
		"logged": logged,
		"files":  fc.front.SetData(files...),
	})
}
