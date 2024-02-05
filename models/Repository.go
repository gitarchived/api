package models

import "gorm.io/gorm"

type Repository struct {
	gorm.Model
	ID         uint
	Owner      string
	Name       string
	Host       string // `github`
	Deleted    bool
	LastCommit string
}

type RepositoryResponse struct {
	ID         uint   `json:"id"`
	Owner      string `json:"owner"`
	Name       string `json:"name"`
	Host       string `json:"host"`
	Deleted    bool   `json:"deleted"`
	CreatedAt  string `json:"createdAt"`
	LastCommit string `json:"lastCommit"`
}
