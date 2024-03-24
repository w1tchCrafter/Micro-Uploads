package models

import "gorm.io/gorm"

type FileModel struct {
	gorm.Model
	Filename     string `gorm:"filename"`
	OriginalName string `gorm:"originalfilename"`
	Author       string `gorm:"author"`
	Size         int64  `gorm:"size"`
}
