package entity

import "github.com/google/uuid"

type GoodsType struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
