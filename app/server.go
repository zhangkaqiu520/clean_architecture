package app

import (
	"github.com/gin-gonic/gin"
	"my-clean-rchitecture/api/handler"
	"my-clean-rchitecture/api/router"
)

type Server struct {
	*gin.Engine
}

func NewServer(hs handler.Handlers) Server {
	r := router.NewRouter(hs)

	return Server{r}
}
