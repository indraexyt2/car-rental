package interfaces

import (
	"car-rental-ums/internal/models"
	"context"
	"github.com/gin-gonic/gin"
)

type ILoginService interface {
	Login(ctx context.Context, request *models.LoginRequest) (*models.LoginResponse, error)
}

type ILoginAPI interface {
	Login(c *gin.Context)
}
