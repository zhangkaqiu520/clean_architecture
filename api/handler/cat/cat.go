package cat

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"my-clean-rchitecture/domain/entity"
	"my-clean-rchitecture/domain/service/cat"
	"net/http"
)

type HandlerInterface interface {
	GetCat(c *gin.Context)
	PostCat(c *gin.Context)
}

type Cat struct {
	cat.ServiceInterface
}

func NewCatHandler(s cat.ServiceInterface) HandlerInterface {
	return Cat{s}
}

func (cat Cat) GetCat(c *gin.Context) {
	id, _ := uuid.Parse(c.Param(`id`))
	nu, err := cat.GetCatByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, nu)
}

func (cat Cat) PostCat(c *gin.Context) {
	var r entity.CreateCatRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusUnprocessableEntity, nil)
		return
	}

	nu, err := cat.CreateCat(r)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, nil)
		return
	}

	c.JSON(http.StatusOK, nu)
}
