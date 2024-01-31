package models

import "gorm.io/gorm"

type Repository struct {
	gorm.Model
	ID   uint
	Name string // `user/repo`
	Host string // `github`
}
