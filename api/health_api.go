package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthAPI struct {
}

func NewHealthApi() *HealthAPI {
	return &HealthAPI{}
}

func (p *HealthAPI) OnHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Server is healthy",
	})
}
