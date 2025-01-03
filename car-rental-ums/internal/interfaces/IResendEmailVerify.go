package interfaces

import (
	"car-rental-ums/internal/models"
	"github.com/gin-gonic/gin"
)

type IResendEmailVerifyService interface {
	ResendEmailVerify(email *models.ResendEmailVerifyRequest) error
}

type IResendEmailVerifyAPI interface {
	ResendEmailVerify(c *gin.Context)
}
