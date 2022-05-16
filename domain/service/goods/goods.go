package goods

import (
	"github.com/google/uuid"
	"my-clean-rchitecture/domain/entity"
	"my-clean-rchitecture/domain/repo/goods"
	"my-clean-rchitecture/domain/response"
)

type ServiceInterface interface {
	GetGoodsByID(uuid.UUID) (entity.Goods, error)
	GetGoodsPagination(page, perPage int) (response.Pagination, error)
}

type Service struct {
	goods.RepoInterface
}

func NewGoodsService(repo goods.RepoInterface) ServiceInterface {
	return Service{repo}
}

func (s Service) GetGoodsByID(uid uuid.UUID) (entity.Goods, error) {
	return s.Get(uid)
}

func (s Service) GetGoodsPagination(page, perPage int) (response.Pagination, error) {
	return s.GetPagination(page, perPage)
}
