package cmd

import (
	"car-rental-ums/constants"
	"car-rental-ums/helpers"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

func (d *Dependency) MiddlewareValidateAuthToken(c *gin.Context) {
	var (
		log = helpers.Logger
	)

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*10)
	defer cancel()

	token, err := c.Cookie("token")
	if err != nil {
		log.Error("failed to get token: ", err)
		helpers.SendResponse(c, http.StatusUnauthorized, constants.StatusUnauthorized, nil)
		c.Abort()
		return
	}

	if token == "" {
		log.Error("token not found")
		helpers.SendResponse(c, http.StatusUnauthorized, constants.StatusUnauthorized, nil)
		c.Abort()
		return
	}

	_, err = d.UserRepository.GetUserSessionByToken(ctx, token)
	if err != nil {
		log.Error("failed to get user session by token: ", err)
		if errors.Is(err, context.DeadlineExceeded) {
			helpers.SendResponse(c, http.StatusRequestTimeout, constants.StatusTimeout, nil)
		} else {
			helpers.SendResponse(c, http.StatusInternalServerError, constants.StatusServerError, nil)
		}
		c.Abort()
		return
	}

	_, err = helpers.ValidateToken(ctx, token)
	if err != nil {
		log.Error("failed to validate token: ", err)
		if errors.Is(err, context.DeadlineExceeded) {
			helpers.SendResponse(c, http.StatusRequestTimeout, constants.StatusTimeout, nil)
		} else {
			helpers.SendResponse(c, http.StatusUnauthorized, constants.StatusUnauthorized, nil)
		}
		c.Abort()
		return
	}
	c.Next()
}
