package entity

import "github.com/google/uuid"

// Breed 品种
type Breed struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"string"`
}
