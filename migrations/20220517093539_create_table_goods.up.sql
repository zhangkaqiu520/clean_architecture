CREATE TABLE `goods`
(
    id          varchar(64) NOT NULL,
    name        varchar(32) NOT NULL,
    type_id     varchar(64) NOT NULL,
    price       int         NOT NULL,
    cover_image text,
    images      longtext,
    created_at  datetime    NOT NULL,
    updated_at  datetime    NOT NULL,
    deleted_at  datetime,
    CONSTRAINT Goods_PK PRIMARY KEY (id)
)
    ENGINE = InnoDB


-- type Goods struct {
--	Base
--	Name       string        `json:"name"`
--	ID         uuid.UUID     `json:"id"`
--	TypeID     uuid.UUID     `json:"typeID"`
--	Price      int           `json:"price"`
--	CoverImage string        `json:"coverImage"`
--	Images     types.Strings `json:"images"`
-- }