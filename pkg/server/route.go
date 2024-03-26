package server

import (
	"adx-admin/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func addRoute[T func(c *gin.Context)](engine *gin.Engine, httpMethod, relativePath string, handler gin.HandlerFunc) {
	engine.Handle(httpMethod, relativePath, handler)
}

func setRoutes(e *gin.Engine, healthApi *api.HealthAPI, reportApi *api.ReportApi) {
	addRoute(e, http.MethodGet, "health", healthApi.OnHealth)
	addRoute(e, http.MethodPost, "report", reportApi.ReportList)
}
