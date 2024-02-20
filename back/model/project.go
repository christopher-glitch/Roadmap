package model

import (
	"gorm.io/gorm"
)


type Project struct {
    gorm.Model
    ProjectName string
    UserID     int
		User  User `gorm:"foreignKey:UserID"` 
}