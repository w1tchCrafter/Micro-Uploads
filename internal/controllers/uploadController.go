package controllers

import (
	"errors"
	"fmt"
	"micro_uploads/internal/middleware"
	"micro_uploads/internal/models"
	"micro_uploads/internal/services"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		uploads.POST("/", middleware.GetUsername(), uc.uploadFile)
		uploads.GET("/:filename", uc.getFile)
	}
}

func (uc UploadControllers) uploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": BAD_REQUEST})
		return
	}

	filename, err := uc.filesystem.Save(file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": SERVER_ERR})
		fmt.Println(err)
		return
	}

	username := ctx.GetString("username")
	ismidia := uc.filesystem.IsMidia(file.Header.Get("Content-Type"))
	dbfile := models.FileModel{Filename: filename, Size: file.Size, Author: username, OriginalName: file.Filename, IsMidia: ismidia}
	if err := uc.DB.Create(&dbfile).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": SERVER_ERR})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"created": SUCCESS_UPDATE})
}

func (uc UploadControllers) getFile(ctx *gin.Context) {
	filename := ctx.Param("filename")
	fullFilePath := filepath.Join(uc.filesystem.UploadPath, filename)
	dbFileData := models.FileModel{}
	err := uc.DB.Where("Filename = ?", fullFilePath).First(&dbFileData).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": NOT_FOUND})
		return
	}

	fileData, err := uc.filesystem.Open(dbFileData.Filename, dbFileData.IsMidia)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": SERVER_ERR})
		return
	}

	defer fileData.Close()

	if dbFileData.IsMidia {
		ctx.File(dbFileData.Filename)
		return
	}

	ctx.DataFromReader(http.StatusOK, dbFileData.Size, "application/octet-stream", fileData, nil)
}
