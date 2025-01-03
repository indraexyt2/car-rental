package interfaces

import (
	"car-rental-ums/internal/models"
	"github.com/gin-gonic/gin"
)

type ILoginService interface {
	Login(request *models.LoginRequest) (*models.LoginResponse, error)
}

type ILoginAPI interface {
	Login(c *gin.Context)
}
