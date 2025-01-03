package interfaces

import (
	"car-rental-ums/internal/models"
	"context"
	"github.com/gin-gonic/gin"
)

type IRegisterService interface {
	Register(ctx context.Context, request *models.User) (*models.User, error)
}

type IRegisterAPI interface {
	Register(c *gin.Context)
}
