package cat

import (
	"github.com/google/uuid"
	"my-clean-rchitecture/domain/entity"
	"my-clean-rchitecture/domain/repo/cat"
)

type ServiceInterface interface {
	GetByUID(uuid uuid.UUID) (entity.Cat, error)
}

type Service struct {
	cat.RepoInterface
}

func NewCatService(repo cat.RepoInterface) ServiceInterface {
	return Service{repo}
}

func (s Service) GetByUID(uid uuid.UUID) (entity.Cat, error) {
	return s.Get(uid)
}
