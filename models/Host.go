package models

import "gorm.io/gorm"

type Host struct {
	gorm.Model
	ID     uint
	Name   string
	URL    string
	Prefix string
}
