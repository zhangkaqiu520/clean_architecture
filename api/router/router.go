package router

import (
	"github.com/gin-gonic/gin"
	"my-clean-rchitecture/api/handler"
)

func NewRouter(hs handler.Handlers) *gin.Engine {
	r := gin.Default()

	// user路由
	r.GET(`/user`, hs.User.GetUser)

	// cat路由
	r.GET(`/cat`, hs.Cat.GetCat)

	return r
}
