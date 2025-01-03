package cmd

import (
	"car-rental-framework/helpers"
	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	r := gin.Default()

	err := r.Run(":" + helpers.GetEnv("APP_PORT"))
	if err != nil {
		helpers.Logger.Fatal("failed to run server: ", err)
	}
}
