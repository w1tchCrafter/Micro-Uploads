package controllers

import (
	"micro_uploads/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	// defining api messages here

	SUCCESS_UPLOAD string = "file was uploaded successfully"
	CREATED_USER   string = "user was created successfully"
	LOGIN_MSG      string = "user logged successfully"
	LOGOUT_MSG     string = "user logged out successfully"
	USER_EXISTS    string = "user already exists"
	ACCESS_DENIED  string = "access denied, wrong credentials or user do not exists"
	NOT_FOUND      string = "page not found"
	SERVER_ERR     string = "an error ocurred, try again later"
	BAD_REQUEST    string = "error, bad request"

	MAX_UPLOAD_CAP uint = 1 << 30
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
	front services.Front
}
