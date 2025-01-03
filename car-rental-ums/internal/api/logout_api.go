package api

import (
	"car-rental-ums/constants"
	"car-rental-ums/helpers"
	"car-rental-ums/internal/interfaces"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

type LogoutAPI struct {
	LogoutSVC interfaces.ILogoutService
}

func (api *LogoutAPI) Logout(c *gin.Context) {
	var (
		log = helpers.Logger
	)

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*10)
	defer cancel()

	cookie, err := c.Cookie("token")
	if err != nil {
		log.Error("failed to get token: ", err)
		helpers.SendResponse(c, http.StatusUnauthorized, constants.StatusUnauthorized, nil)
		c.Abort()
		return
	}

	err = api.LogoutSVC.Logout(ctx, cookie)
	if err != nil {
		log.Error("failed to delete user session: ", err)
		if errors.Is(err, context.DeadlineExceeded) {
			helpers.SendResponse(c, http.StatusRequestTimeout, constants.StatusTimeout, nil)
		} else {
			helpers.SendResponse(c, http.StatusInternalServerError, constants.StatusServerError, nil)
		}
		return
	}
	helpers.SendResponse(c, http.StatusOK, constants.StatusSuccess, nil)
}
