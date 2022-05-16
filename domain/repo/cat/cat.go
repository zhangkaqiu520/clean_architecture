package cat

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"my-clean-rchitecture/domain/entity"
)

type RepoInterface interface {
	Get(uuid.UUID) (entity.Cat, error)
	Create(entity.CreateCatRequest) (entity.Cat, error)
}

type Repo struct {
	db *gorm.DB
}

func NewCatRepo(db *gorm.DB) RepoInterface {
	return Repo{db}
}

func (c Repo) Get(uid uuid.UUID) (cat entity.Cat, err error) {
	err = c.db.Where(`id`, uid).Find(&cat).Error

	return
}

func (c Repo) Create(req entity.CreateCatRequest) (entity.Cat, error) {
	cat := entity.Cat{
		ID:    uuid.New(),
		Name:  req.Name,
		Breed: req.Breed,
	}
	err := c.db.Create(&cat).Error

	return cat, err
}
