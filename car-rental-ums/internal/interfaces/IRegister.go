package interfaces

import (
	"car-rental-ums/internal/models"
	"github.com/gin-gonic/gin"
)

type IRegisterService interface {
	Register(request *models.User) (*models.User, error)
}

type IRegisterAPI interface {
	Register(c *gin.Context)
}
