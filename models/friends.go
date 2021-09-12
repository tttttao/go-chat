package models

import "gorm.io/gorm"

type Friend struct {
	gorm.Model
	Name string
}
