package controllers

import (
	"micro_uploads/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IControllers interface {
	StartRoutes()
}

type BaseControllers struct {
	R  *gin.RouterGroup
	DB *gorm.DB
}

type UploadControllers struct {
	BaseControllers
	filesystem services.FS
}

type AuthControllers struct {
	BaseControllers
	authentication services.Auth
}

type FrontControllers struct {
	BaseControllers
}

type FileResponseData struct {
	Filename string
	StrSize  string
	Link     string
}
