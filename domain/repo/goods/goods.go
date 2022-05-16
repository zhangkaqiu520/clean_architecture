package goods

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"my-clean-rchitecture/domain/entity"
	"my-clean-rchitecture/domain/response"
)

type RepoInterface interface {
	Get(uuid.UUID) (entity.Goods, error)
	GetAll() ([]entity.Goods, error)
	GetPagination(page, perPage int) (response.Pagination, error)
}

type Repo struct {
	db *gorm.DB
}

func NewGoodsRepo(db *gorm.DB) RepoInterface {
	return Repo{db}
}

func (c Repo) Get(uid uuid.UUID) (goods entity.Goods, err error) {
	err = c.db.Where(`id`, uid).Find(&goods).Error

	return
}

func (c Repo) GetAll() ([]entity.Goods, error) {
	var goods []entity.Goods
	err := c.db.Create(&goods).Error

	return goods, err
}

func (c Repo) GetPagination(page, perPage int) (response.Pagination, error) {
	var goods []entity.Goods
	return response.NewPagination(page, perPage, &goods, c.db.Model(&entity.Goods{}))
}
