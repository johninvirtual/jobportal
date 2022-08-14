package models

import (
	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
}
