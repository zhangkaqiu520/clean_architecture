package cat

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"my-clean-rchitecture/domain/entity"
)

type RepoInterface interface {
	Get(uuid uuid.UUID) (entity.Cat, error)
}

type Repo struct {
	db *gorm.DB
}

func NewCatRepo(db *gorm.DB) RepoInterface {
	return Repo{db}
}

func (c Repo) Get(uid uuid.UUID) (entity.Cat, error) {
	return entity.Cat{}, nil
}
