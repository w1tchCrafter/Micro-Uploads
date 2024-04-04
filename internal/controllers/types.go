package controllers

import (
	"micro_uploads/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// defining api messages here
const (
	SUCCESS_UPDATE string = "file was uploaded successfully"
	CREATED_USER   string = "user was created successfully"
	LOGIN_MSG      string = "user logged successfully"
	LOGOUT_MSG     string = "user logged out successfully"
	USER_EXISTS    string = "user already exists"
	ACCESS_DENIED  string = "access denied, wrong credentials or user do not exists"
	NOT_FOUND      string = "page not found"
	SERVER_ERR     string = "an error ocurred, try again later"
	BAD_REQUEST    string = "error, bad request"
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
