package api

import (
	"car-rental-ums/constants"
	"car-rental-ums/helpers"
	"car-rental-ums/internal/interfaces"
	"car-rental-ums/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginAPI struct {
	LoginSVC interfaces.ILoginService
}

func (api *LoginAPI) Login(c *gin.Context) {
	var (
		req = &models.LoginRequest{}
		log = helpers.Logger
	)

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

	resp, err := api.LoginSVC.Login(req)
	if err != nil {
		log.Error("failed to login: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.SendResponse(c, http.StatusOK, constants.StatusSuccess, resp)
}
