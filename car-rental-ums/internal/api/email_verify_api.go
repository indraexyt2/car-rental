package api

import (
	"car-rental-ums/constants"
	"car-rental-ums/helpers"
	"car-rental-ums/internal/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EmailVerifyAPI struct {
	EmailVerifySVC interfaces.IEmailVerifyService
}

func (api *EmailVerifyAPI) EmailVerify(c *gin.Context) {
	var (
		log = helpers.Logger
	)
	tokenEmailVerify := c.Param("token")

	err := api.EmailVerifySVC.EmailVerify(tokenEmailVerify)
	if err != nil {
		log.Error("failed to verify email: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, constants.StatusBadRequest, nil)
		return
	}

	helpers.SendResponse(c, http.StatusOK, constants.StatusSuccess, nil)
}
