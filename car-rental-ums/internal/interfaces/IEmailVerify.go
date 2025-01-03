package interfaces

import "github.com/gin-gonic/gin"

type IEmailVerifyService interface {
	EmailVerify(emailVerifyToken string) error
}

type IEmailVerifyAPI interface {
	EmailVerify(c *gin.Context)
}
