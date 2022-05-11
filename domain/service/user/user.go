package user

import (
	"github.com/google/uuid"
	"my-clean-rchitecture/domain/entity"
	"my-clean-rchitecture/domain/repo/user"
)

type ServiceInterface interface {
	GetByUID(uuid uuid.UUID) (entity.User, error)
}

type Service struct {
	user.RepoInterface
}

func NewUserService(repo user.RepoInterface) ServiceInterface {
	return Service{repo}
}

func (s Service) GetByUID(uid uuid.UUID) (entity.User, error) {
	return s.Get(uid)
}
