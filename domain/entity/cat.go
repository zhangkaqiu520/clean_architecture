package entity

import "github.com/google/uuid"

// Cat 猫
type Cat struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"string"`
	Breed Breed     `json:"breed"`
}

// CreateCatRequest 新增猫
type CreateCatRequest struct {
	Name  string `json:"name"`
	Breed Breed  `json:"breed"`
}
