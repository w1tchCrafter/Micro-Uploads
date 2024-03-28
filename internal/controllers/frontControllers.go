package controllers

import (
	"fmt"
	"micro_uploads/internal/middleware"
	"micro_uploads/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewFrontControllers(r *gin.Engine, db *gorm.DB) FrontControllers {
	return FrontControllers{
		R:  r,
		DB: db,
	}
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
	username := ctx.GetString("username")
	files := []models.FileModel{}

	if username != "" {
		err := fc.DB.Where("author = ?", username).Find(&files).Error

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
	}

	ctx.HTML(http.StatusOK, "user", gin.H{
		"title":  "user",
		"logged": false,
		"files":  files,
	})
}
