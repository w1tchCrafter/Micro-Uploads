package middleware

import (
	"encoding/json"
	"micro_uploads/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUpdateStorage(t *testing.T) {
	// opening db in memory for testing
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{TranslateError: true})
	db.AutoMigrate(&models.UserModel{})

	assert := assert.New(t)
	r := gin.New()

	userdata := models.UserModel{
		ID:       uuid.New().String(),
		Username: "testUser",
		Storage:  0,
	}

	db.Create(&userdata)

	r.Use(UpdateStorage(db))
	r.GET("/", func(ctx *gin.Context) {
		ctx.Set("username", userdata.Username)
		ctx.Set("datasize", 1000)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resp := struct {
		Error string `json:"error"`
	}{}

	json.Unmarshal(w.Body.Bytes(), &resp)
	db.Where("username = ?", userdata.Username).First(&userdata)

	assert.Equal(userdata.Storage, uint(1000))
	assert.Empty(resp.Error)
}

func TestFullStorage(t *testing.T) {
	// opening db in memory for testing
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{TranslateError: true})
	db.AutoMigrate(&models.UserModel{})

	assert := assert.New(t)
	r := gin.New()

	userdata := models.UserModel{
		ID:       uuid.New().String(),
		Username: "testUser2",
		Storage:  GB,
	}

	db.Create(&userdata)

	r.Use(UpdateStorage(db))
	r.GET("/", func(ctx *gin.Context) {
		ctx.Set("username", userdata.Username)
		ctx.Set("datasize", 1000)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resp := struct {
		Error string `json:"error"`
	}{}

	json.Unmarshal(w.Body.Bytes(), &resp)
	db.Where("username = ?", userdata.Username).First(&userdata)

	t.Log(userdata.Storage == GB, userdata.Storage == GB+1000)
	assert.Equal(FULL_STORAGE, resp.Error)
}

func TestSubStorage(t *testing.T) {
	// opening db in memory for testing
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{TranslateError: true})
	db.AutoMigrate(&models.UserModel{})

	assert := assert.New(t)
	r := gin.New()

	userdata := models.UserModel{
		ID:       uuid.New().String(),
		Username: "testUser3",
		Storage:  10000,
	}

	db.Create(&userdata)

	r.Use(UpdateStorage(db))
	r.GET("/", func(ctx *gin.Context) {
		ctx.Set("username", userdata.Username)
		ctx.Set("datasize", -5000)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resp := struct {
		Error string `json:"error"`
	}{}

	json.Unmarshal(w.Body.Bytes(), &resp)
	db.Where("username = ?", userdata.Username).First(&userdata)

	assert.Equal(userdata.Storage, uint(5000))
	assert.Empty(resp.Error)
}
