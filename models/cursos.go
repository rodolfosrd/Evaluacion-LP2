package models

import (
	"gorm.io/gorm"
)

type Cursos struct {
	gorm.Model
	Name   string
	Period string
	State  string
}
