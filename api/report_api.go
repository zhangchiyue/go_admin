package api

import (
	"adx-admin/internal/domian/request"
	"adx-admin/internal/domian/service"
	"adx-admin/pkg/microlog"
	"github.com/gin-gonic/gin"
)

type ReportApi struct {
	log           microlog.Logger
	reportService *service.ReportService
}

func NewReportApi(log microlog.Logger, reportService *service.ReportService) *ReportApi {
	return &ReportApi{
		log:           log,
		reportService: reportService,
	}
}

func (p *ReportApi) ReportList(c *gin.Context) {
	var req request.QueryReportReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ParamError(c, err)
		return
	}
	report, err := p.reportService.QueryAdxReport(&req)
	if err != nil {
		p.log.Errorf("failed to query adx report, err:%s\n", err.Error())
		ServerError(c, err)
		return
	}
	Success(c, report)
}
