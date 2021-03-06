package goods

import (
	"my-clean-rchitecture/domain/entity"
	"my-clean-rchitecture/domain/service/goods"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HandlerInterface interface {
	GetGoods(c *gin.Context)
	GetPagination(c *gin.Context)
	PostGoods(c *gin.Context)
}

type Goods struct {
	goods.ServiceInterface
}

func NewGoodsHandler(s goods.ServiceInterface) HandlerInterface {
	return Goods{s}
}

func (g Goods) GetGoods(c *gin.Context) {
	id, _ := uuid.Parse(c.Param(`id`))
	nu, err := g.GetGoodsByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, nu)
}

func (g Goods) GetPagination(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query(`page`))
	perPage, _ := strconv.Atoi(c.Query(`perPage`))
	nu, err := g.GetGoodsPagination(page, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nu)
}

func (g Goods) PostGoods(c *gin.Context) {
	var r entity.CreateGoodsRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	goods, err := g.CreateGoods(r)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, goods)
}
