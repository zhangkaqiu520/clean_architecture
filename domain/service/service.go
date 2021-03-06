package service

import (
	"my-clean-rchitecture/domain/repo"
	"my-clean-rchitecture/domain/service/cat"
	"my-clean-rchitecture/domain/service/goods"
	"my-clean-rchitecture/domain/service/user"
)

type Services struct {
	User  user.ServiceInterface
	Cat   cat.ServiceInterface
	Goods goods.ServiceInterface
}

func NewServices(rs repo.Repos) Services {
	var s Services
	s.User = user.NewUserService(rs.User)
	s.Cat = cat.NewCatService(rs.Cat)
	s.Goods = goods.NewGoodsService(rs.Goods)

	return s
}
