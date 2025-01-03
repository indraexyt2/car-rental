package api

import (
	"car-rental-ums/constants"
	"car-rental-ums/helpers"
	"car-rental-ums/internal/interfaces"
	"car-rental-ums/internal/models"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

type ResendEmailVerifyAPI struct {
	ResendEmailVerifySVC interfaces.IResendEmailVerifyService
}

func (api *ResendEmailVerifyAPI) ResendEmailVerify(c *gin.Context) {
	var (
		log = helpers.Logger
		req = &models.ResendEmailVerifyRequest{}
	)

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*5)
	defer cancel()

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
	err = api.ResendEmailVerifySVC.ResendEmailVerify(ctx, req)
	if err != nil {
		log.Error("failed to resend email verify: ", err)
		if errors.Is(err, context.DeadlineExceeded) {
			helpers.SendResponse(c, http.StatusRequestTimeout, constants.StatusTimeout, nil)
		} else {
			helpers.SendResponse(c, http.StatusBadGateway, constants.StatusServerError, nil)
		}
		return
	}

	helpers.SendResponse(c, http.StatusOK, constants.StatusSuccess, nil)
}
