package cat

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"my-clean-rchitecture/domain/service/cat"
	"net/http"
)

type HandlerInterface interface {
	GetCat(c *gin.Context)
}

type Cat struct {
	cat.ServiceInterface
}

func NewCatHandler(s cat.ServiceInterface) HandlerInterface {
	return Cat{s}
}

func (cat Cat) GetCat(c *gin.Context) {
	nu, _ := cat.GetByUID(uuid.New())
	c.JSON(http.StatusOK, nu)
}
