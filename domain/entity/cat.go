package entity

import "github.com/google/uuid"

type Cat struct {
	ID uuid.UUID `json:"id" bson:"id"`
}
