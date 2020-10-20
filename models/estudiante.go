package models

import (
	"gorm.io/gorm"
)

type Estudiante struct {
	gorm.Model
	Name     string
	Paternal string
	Maternal string
	Age      string
	State    string
}
