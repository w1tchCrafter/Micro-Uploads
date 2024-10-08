package controllers

import (
	"errors"
	"fmt"
	"micro_uploads/internal/middleware"
	"micro_uploads/internal/models"
	"micro_uploads/internal/services"
	"net/http"
	"path/filepath"
	"strings"

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
		uploads.POST("/", uc.uploadFile)
		uploads.GET("/:filename", uc.getFile)
		uploads.DELETE("/:filename", uc.deleteFile)
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

	ctx.JSON(201, gin.H{
		"name": dbfile.OriginalName,
		"link": "/api/v1/uploads/" + strings.Split(dbfile.Filename, "/")[1],
		"size": dbfile.Size,
	})
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

	tx := uc.DB.Begin()

	if err := tx.Where("Filename = ?", fullFilePath).First(&file).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusNotFound, gin.H{"error": NOT_FOUND})
		return
	}

	if author == "" || author != file.Author {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": ACCESS_DENIED})
		return
	}

	if err := tx.Delete(&file).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": SERVER_ERR})
		return
	}

	tx.Commit()

	ctx.JSON(http.StatusOK, gin.H{"deletd": "file deleted successfully"})
}
