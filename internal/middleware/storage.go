package middleware

import (
	"fmt"
	"micro_uploads/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	// defining here to avoid circular imports
	GB uint = 1 << 30

	SERVER_ERR     string = "an error ocurred, try again later"
	FULL_STORAGE   string = "you reached your max storage space"
	SUCCESS_UPLOAD string = "file was uploaded successfully"
	FILE_DELETED   string = "file was deleted successfully"
)

func UpdateStorage(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		// Executes after handler func is done and has set the required data
		username := ctx.GetString("username")
		datasize := ctx.GetInt64("datasize")
		status := ctx.GetInt("status")
		user := models.UserModel{}

		if err := db.Where("username = ?", username).First(&user).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": SERVER_ERR})
			return
		}

		stored := int64(user.Storage)
		current := stored + datasize
		fmt.Println("from storage,", username, datasize, status)

		if uint(current) > GB {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": FULL_STORAGE})
			return
		}

		if err := db.Model(&user).Update("storage", current).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": SERVER_ERR})
			return
		}

		switch status {
		case http.StatusCreated:
			ctx.JSON(status, gin.H{"created": SUCCESS_UPLOAD})
		case http.StatusOK:
			ctx.JSON(status, gin.H{"deleted": FILE_DELETED})
		}
	}
}
