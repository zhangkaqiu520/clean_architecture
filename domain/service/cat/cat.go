package cat

import (
	"github.com/google/uuid"
	"my-clean-rchitecture/domain/entity"
	"my-clean-rchitecture/domain/repo/cat"
)

type ServiceInterface interface {
	GetCatByID(uuid.UUID) (entity.Cat, error)
	CreateCat(request entity.CreateCatRequest) (entity.Cat, error)
}

type Service struct {
	cat.RepoInterface
}

func NewCatService(repo cat.RepoInterface) ServiceInterface {
	return Service{repo}
}

func (s Service) GetCatByID(uid uuid.UUID) (entity.Cat, error) {
	return s.Get(uid)
}

func (s Service) CreateCat(req entity.CreateCatRequest) (entity.Cat, error) {
	return s.Create(req)
}
