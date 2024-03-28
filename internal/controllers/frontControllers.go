package controllers

import (
	"fmt"
	"micro_uploads/internal/middleware"
	"micro_uploads/internal/models"
	"net/http"
	"path/filepath"

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

func (fc FrontControllers) isLogged(username string) bool {
	return username != ""
}

func (fc FrontControllers) setData(filesData ...models.FileModel) []FileResponseData {
	newData := make([]FileResponseData, 0)
	var filename, size string
	const KB = 1 << 10
	const MB = 1 << 20

	for _, v := range filesData {
		if len(v.OriginalName) >= 12 {
			ext := filepath.Ext(v.OriginalName)
			filename = v.OriginalName[0:12] + "... " + ext
		}

		switch {
		case v.Size >= MB:
			size = fmt.Sprint(int(float64(v.Size))/MB, "MB")
		case v.Size >= KB:
			size = fmt.Sprint(int(float64(v.Size))/KB, "KB")
		default:
			size = fmt.Sprint(v.Size, "bytes")
		}

		newData = append(newData, FileResponseData{
			Filename: filename,
			StrSize:  size,
		})
	}

	return newData
}

func (fc FrontControllers) index(ctx *gin.Context) {
	username := ctx.GetString("username")
	logged := fc.isLogged(username)

	ctx.HTML(http.StatusOK, "index", gin.H{
		"title":  "micro uploads - login",
		"logged": logged,
	})
}

func (fc FrontControllers) user(ctx *gin.Context) {
	username := ctx.GetString("username")
	logged := fc.isLogged(username)
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
		"logged": logged,
		"files":  fc.setData(files...),
	})
}
