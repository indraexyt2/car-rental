package api

import (
	"car-rental-ums/constants"
	"car-rental-ums/helpers"
	"car-rental-ums/internal/interfaces"
	"car-rental-ums/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResendEmailVerifyAPI struct {
	ResendEmailVerifySVC interfaces.IResendEmailVerifyService
}

func (api *ResendEmailVerifyAPI) ResendEmailVerify(c *gin.Context) {
	var (
		log = helpers.Logger
		req = &models.ResendEmailVerifyRequest{}
	)

	// validate
	err := c.ShouldBindJSON(req)
	if err != nil {
		log.Error("failed to bind json: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, constants.StatusBadRequest, nil)
		return
	}

	err = req.Validate()
	if err != nil {
		log.Error("failed to validate request: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, constants.StatusBadRequest, nil)
		return
	}

	// service
	err = api.ResendEmailVerifySVC.ResendEmailVerify(req)
	if err != nil {
		log.Error("failed to resend email verify: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.SendResponse(c, http.StatusOK, constants.StatusSuccess, nil)
}
