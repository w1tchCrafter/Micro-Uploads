package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserForm struct {
	Username string `form:"username" binding:"required,min=4,max=36"`
	Password string `form:"password" binding:"required,min=8,max=72"`
}

type UserModel struct {
	gorm.Model
	ID       string `gorm:"id;type:uuid;primarykey"`
	Username string `gorm:"username;unique;not null"`
	Password string `gorm:"password;not null"`
}

func NewUserModel(userform UserForm) UserModel {
	return UserModel{
		ID:       uuid.New().String(),
		Username: userform.Username,
		Password: "",
	}
}
