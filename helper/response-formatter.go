package helper

import (
	"my-app/types"

	"github.com/gin-gonic/gin"
)

func JSONResponse(c *gin.Context, code int, status, message string, data interface{}) {
	c.JSON(code, types.Response{
		Status:  status,
		Message: message,
		Data:    data,
	})
}
