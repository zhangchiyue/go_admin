package server

import (
	"adx-admin/api"
	"github.com/gin-gonic/gin"
)

func NewEngine(healthApi *api.HealthAPI, reportApi *api.ReportApi) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	setMiddlewares(engine)
	setRoutes(engine, healthApi, reportApi)
	return engine
}
