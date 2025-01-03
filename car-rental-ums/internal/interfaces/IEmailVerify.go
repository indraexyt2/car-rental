package interfaces

import (
	"context"
	"github.com/gin-gonic/gin"
)

type IEmailVerifyService interface {
	EmailVerify(ctx context.Context, emailVerifyToken string) error
}

type IEmailVerifyAPI interface {
	EmailVerify(c *gin.Context)
}
