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

type LoginAPI struct {
	LoginSVC interfaces.ILoginService
}

func (api *LoginAPI) Login(c *gin.Context) {
	var (
		req = &models.LoginRequest{}
		log = helpers.Logger
	)

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*5)
	defer cancel()

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

	resp, err := api.LoginSVC.Login(ctx, req)
	if err != nil {
		log.Error("failed to login: ", err)
		if errors.Is(err, context.DeadlineExceeded) {
			helpers.SendResponse(c, http.StatusRequestTimeout, constants.StatusTimeout, nil)
		} else {
			helpers.SendResponse(c, http.StatusInternalServerError, constants.StatusServerError, nil)
		}
		return
	}

	helpers.SendResponse(c, http.StatusOK, constants.StatusSuccess, resp)
}
