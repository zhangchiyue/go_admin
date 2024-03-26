package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data，omitempty"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "ok",
		Data:    data,
	})
}
func ServerError(c *gin.Context, err error) {
	c.JSON(http.StatusOK, &Response{
		Code:    500,
		Message: err.Error(),
		Data:    nil,
	})
}

func ParamError(c *gin.Context, err error) {
	c.JSON(http.StatusOK, &Response{
		Code:    400,
		Message: err.Error(),
		Data:    nil,
	})
}
