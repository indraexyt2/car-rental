package helpers

import "github.com/gin-gonic/gin"

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, Response{
		Message: message,
		Data:    data,
	})
}
