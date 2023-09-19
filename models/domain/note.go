package domain

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Title string
	Tags  string
	Body  string
}
