package entity

import (
	"my-clean-rchitecture/domain/types"

	"github.com/google/uuid"
)

// Goods 商品
type Goods struct {
	Base
	Name       string         `json:"name"`
	ID         uuid.UUID      `json:"id"`
	TypeID     uuid.UUID      `json:"typeID"`
	Price      int            `json:"price"`
	CoverImage *string        `json:"coverImage"`
	Images     *types.Strings `json:"images"`
}

func (g Goods) TableName() string {
	return `goods`
}

// CreateGoodsRequest 创建商品
type CreateGoodsRequest struct {
	Name   string    `json:"name"`
	TypeID uuid.UUID `json:"typeID"`
	Price  int       `json:"price"`
}
