package api

import (
	"car-rental-ums/constants"
	"car-rental-ums/helpers"
	"car-rental-ums/internal/interfaces"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type EmailVerifyAPI struct {
	EmailVerifySVC interfaces.IEmailVerifyService
}

func (api *EmailVerifyAPI) EmailVerify(c *gin.Context) {
	var (
		log = helpers.Logger
	)
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*5)
	defer cancel()

	tokenEmailVerify := c.Param("token")

	err := api.EmailVerifySVC.EmailVerify(ctx, tokenEmailVerify)
	if err != nil {
		log.Error("failed to verify email: ", err)
		if errors.Is(err, context.DeadlineExceeded) {
			helpers.SendResponse(c, http.StatusRequestTimeout, constants.StatusTimeout, nil)
		} else {
			helpers.SendResponse(c, http.StatusInternalServerError, constants.StatusServerError, nil)
		}
		return
	}

	helpers.SendResponse(c, http.StatusOK, constants.StatusSuccess, nil)
}
