package controllers

import (
	"errors"
	"micro_uploads/internal/middleware"
	"micro_uploads/internal/models"
	"micro_uploads/internal/services"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	BAD_REQ_MSG    = "error uploading file"
	SERVER_ERR_MSG = "error, try again later"
	SUCCESS_MSG    = "file uploaded successfully"
)

func NewUploadController(r *gin.RouterGroup, db *gorm.DB, fsPath string) UploadControllers {
	uc := UploadControllers{}
	uc.R = r
	uc.DB = db
	uc.filesystem = services.NewFileSystem(fsPath)
	return uc
}

func (uc *UploadControllers) StartRoutes() {
	uploads := uc.R.Group("/uploads")
	{
		uploads.POST("/create", middleware.GetUsername(), uc.uploadFile)
		uploads.GET("/retrieve/:filename", uc.getFile)
	}
}

func (uc UploadControllers) uploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": BAD_REQ_MSG})
		return
	}

	filename, err := uc.filesystem.Save(file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": SERVER_ERR_MSG})
		return
	}

	username := ctx.GetString("username")
	dbfile := models.FileModel{Filename: filename, Size: file.Size, Author: username, OriginalName: file.Filename}
	if err := uc.DB.Create(&dbfile).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": SERVER_ERR_MSG})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"created": SUCCESS_MSG})
}

func (uc UploadControllers) getFile(ctx *gin.Context) {
	filename := ctx.Param("filename")
	fullFilePath := filepath.Join(uc.filesystem.UploadPath, filename)
	dbFileData := models.FileModel{}
	err := uc.DB.Where("Filename = ?", fullFilePath).First(&dbFileData).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.Status(http.StatusNotFound)
		return
	}

	fileData, err := uc.filesystem.Open(dbFileData.Filename)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": SERVER_ERR_MSG})
		return
	}

	defer fileData.Close()
	ctx.DataFromReader(http.StatusOK, dbFileData.Size, "application/octet-stream", fileData, nil)
}
