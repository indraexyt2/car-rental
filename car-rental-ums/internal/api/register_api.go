package api

import (
	"car-rental-ums/constants"
	"car-rental-ums/helpers"
	"car-rental-ums/internal/interfaces"
	"car-rental-ums/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterAPI struct {
	RegisterSVC interfaces.IRegisterService
}

func (api *RegisterAPI) Register(c *gin.Context) {
	var (
		log = helpers.Logger
		req = &models.User{}
	)

	err := c.ShouldBindJSON(req)
	if err != nil {
		log.Error("failed to bind json: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, constants.StatusBadRequest, nil)
		return
	}

	if err := req.Validate(); err != nil {
		log.Error("failed to validate request: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, constants.StatusBadRequest, nil)
		return
	}

	resp, err := api.RegisterSVC.Register(req)
	if err != nil {
		log.Error("failed to register user: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, constants.StatusServerError, nil)
		return
	}

	helpers.SendResponse(c, http.StatusOK, constants.StatusSuccess, resp)
}
