package user

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"my-clean-rchitecture/domain/service/user"
	"net/http"
)

type HandlerInterface interface {
	GetUser(c *gin.Context)
}

type User struct {
	user.ServiceInterface
}

func NewUserHandler(s user.ServiceInterface) HandlerInterface {
	return User{s}
}

func (u User) GetUser(c *gin.Context) {
	nu, _ := u.GetByUID(uuid.New())
	c.JSON(http.StatusOK, nu)
}
