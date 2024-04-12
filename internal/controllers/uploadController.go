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
	uc.R.Use(middleware.GetUsername())

	uploads := uc.R.Group("/uploads")
	{
		uploads.POST("/", middleware.UpdateStorage(uc.DB), uc.uploadFile)
		uploads.GET("/:filename", uc.getFile)
		uploads.DELETE("/:filename", middleware.UpdateStorage(uc.DB), uc.deleteFile)
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
	dbfile := models.FileModel{
		Filename:     filename,
		Size:         file.Size,
		Author:       username,
		OriginalName: file.Filename,
		IsMidia:      ismidia,
	}

	if err := uc.DB.Create(&dbfile).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": SERVER_ERR})
		return
	}

	ctx.Set("username", username)
	ctx.Set("datasize", file.Size)
	ctx.Set("status", http.StatusCreated)
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

func (uc UploadControllers) deleteFile(ctx *gin.Context) {
	filename := ctx.Param("filename")
	author := ctx.GetString("username")
	fullFilePath := filepath.Join(uc.filesystem.UploadPath, filename)
	file := models.FileModel{}

	if author == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": ACCESS_DENIED})
		return
	}

	if err := uc.DB.Where("Filename = ?", fullFilePath).First(&file).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": NOT_FOUND})
		return
	}

	if err := uc.DB.Delete(&file).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": SERVER_ERR})
		return
	}

	ctx.Set("username", author)
	ctx.Set("datasize", -file.Size)
	ctx.Set("status", http.StatusOK)
}
