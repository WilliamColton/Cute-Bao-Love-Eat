package model

import (
	"gorm.io/gorm"
)

type Level struct {
	Name     string `binding:"required"`
	ParentID *uint
	gorm.Model
}

type Item struct {
	Name      string `binding:"required"`
	LevelID   uint   `binding:"required"`
	ImagePath string `binding:"required"`
	gorm.Model
}

type Evaluation struct {
	UserID      uint `binding:"required"`
	ItemID      uint `binding:"required"`
	Recommended bool `binding:"required"`
	gorm.Model
}

type User struct {
	UnionID  string `gorm:"uniqueIndex" binding:"required"`
	UserName string `binding:"required"`
	Role     string
	gorm.Model
}