package middleware

import (
	"micro_uploads/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	// defining here to avoid circular imports
	GB uint = 1 << 30
)

func UpdateStorage(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		// Executes after handler func is done and has set the required data
		username := ctx.GetString("username")
		datasize := ctx.GetInt("datasize")
		user := models.UserModel{}

		if err := db.Where("username = ?", username).First(&user).Error; err != nil {
			//handle error
		}

		stored := int(user.Storage)
		current := stored + datasize

		if uint(stored) > GB {
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}

		if err := db.Model(&user).Update("storage", current).Error; err != nil {
			//
		}
	}
}
