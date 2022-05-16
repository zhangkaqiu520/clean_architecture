package entity

// Goods 商品
type Goods struct {
	Name       string        `json:"name"`
	//ID         uuid.UUID     `json:"id"`
	//Type       GoodsType     `json:"type"`
	//TypeID     uuid.UUID     `json:"typeID"`
	//Price      int           `json:"price"`
	//CoverImage string        `json:"coverImage"`
	//Images     types.Strings `json:"images"`
}

func (g Goods) TableName() string {
	return `goods`
}
