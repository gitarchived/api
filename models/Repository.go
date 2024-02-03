package models

import "gorm.io/gorm"

type Repository struct {
	gorm.Model
	ID   uint
	Name string // `user/repo`
	Host string // `github`
}

type RepositoryResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Host      string `json:"host"`
	CreatedAt string `json:"createdAt"`
}
