package router

import (
	"my-clean-rchitecture/api/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(hs handler.Handlers) *gin.Engine {
	r := gin.Default()

	api := r.Group(`/api`)

	// user路由
	api.GET(`/user`, hs.User.GetUser)

	// cat路由
	api.GET(`/cat`, hs.Cat.GetCat)

	// goods路由
	api.GET(`/goods`, hs.Goods.GetPagination)
	api.GET(`/goods/:id`, hs.Goods.GetGoods)
	api.POST(`/goods`, hs.Goods.PostGoods)

	return r
}
