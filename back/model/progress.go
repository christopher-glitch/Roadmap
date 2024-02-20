package model

import (
	"gorm.io/gorm"
)

type Progress struct {
		gorm.Model
    ProgressName  string
    Note      string
    ProjectID int
		Project Project `gorm:"foreignKey:ProjectID"` 
}