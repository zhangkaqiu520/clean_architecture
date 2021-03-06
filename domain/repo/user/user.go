package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"my-clean-rchitecture/domain/entity"
)

type RepoInterface interface {
	Get(uuid uuid.UUID) (entity.User, error)
}

type Repo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) RepoInterface {
	return Repo{db}
}

func (u Repo) Get(uid uuid.UUID) (entity.User, error) {
	var us entity.User
	err := u.db.Where(`id`, uid).Find(&us).Error
	return us, err
}
