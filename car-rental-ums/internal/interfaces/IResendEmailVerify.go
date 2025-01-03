package interfaces

import (
	"car-rental-ums/internal/models"
	"context"
	"github.com/gin-gonic/gin"
)

type IResendEmailVerifyService interface {
	ResendEmailVerify(ctx context.Context, email *models.ResendEmailVerifyRequest) error
}

type IResendEmailVerifyAPI interface {
	ResendEmailVerify(c *gin.Context)
}
