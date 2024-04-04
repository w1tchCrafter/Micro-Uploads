package controllers

import (
	"errors"
	"micro_uploads/internal/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func NewAuthController(r *gin.RouterGroup, db *gorm.DB) AuthControllers {
	ac := AuthControllers{}
	ac.R = r
	ac.DB = db
	return ac
}

func (ac *AuthControllers) StartRoutes() {
	auth := ac.R.Group("auth")
	{
		auth.POST("/register", ac.register)
		auth.POST("/login", ac.login)
		auth.GET("/logout", ac.logout)
	}
}

func (ac AuthControllers) register(ctx *gin.Context) {
	newUser := models.UserForm{}

	if err := ctx.Bind(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": BAD_REQUEST})
		return
	}

	hashed, err := ac.authentication.HashPassword(newUser.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": SERVER_ERR})
		return
	}

	userModel := models.NewUserModel(newUser)
	userModel.Password = hashed
	err = ac.DB.Create(&userModel).Error

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": USER_EXISTS})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": SERVER_ERR})
		return
	}

	session := sessions.Default(ctx)
	session.Set("id", userModel.ID)
	session.Set("username", userModel.Username)
	if err := session.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": SERVER_ERR})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"created": CREATED_USER})
}

func (ac AuthControllers) login(ctx *gin.Context) {
	loginUser := models.UserForm{}

	if err := ctx.Bind(&loginUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": BAD_REQUEST})
		return
	}

	userModel := models.UserModel{}
	err := ac.DB.Where("username = ?", loginUser.Username).First(&userModel).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": ACCESS_DENIED})
		return
	}

	if err := ac.authentication.ValidatePassword(
		loginUser.Password, userModel.Password,
	); err == bcrypt.ErrMismatchedHashAndPassword {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": ACCESS_DENIED})
		return
	}

	session := sessions.Default(ctx)
	session.Set("id", userModel.ID)
	session.Set("username", userModel.Username)
	if err := session.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": SERVER_ERR})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": LOGIN_MSG})
}

func (ac AuthControllers) logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()
	ctx.JSON(http.StatusOK, gin.H{"success": LOGOUT_MSG})
}
