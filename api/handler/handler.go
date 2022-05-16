package handler

import (
	"my-clean-rchitecture/api/handler/cat"
	"my-clean-rchitecture/api/handler/goods"
	"my-clean-rchitecture/api/handler/user"
	"my-clean-rchitecture/domain/service"
)

type Handlers struct {
	User  user.HandlerInterface
	Cat   cat.HandlerInterface
	Goods goods.HandlerInterface
}

func NewHandlers(ss service.Services) Handlers {
	var hs Handlers
	hs.User = user.NewUserHandler(ss.User)
	hs.Cat = cat.NewCatHandler(ss.Cat)
	hs.Goods = goods.NewGoodsHandler(ss.Goods)

	return hs
}
