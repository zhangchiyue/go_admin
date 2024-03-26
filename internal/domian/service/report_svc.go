package service

import (
	"adx-admin/internal/domian/repository"
	"adx-admin/internal/domian/request"
	"adx-admin/internal/domian/response"
)

type ReportService struct {
	reportRepo repository.ReportRepo
}

func NewReportService(reportRepo repository.ReportRepo) *ReportService {
	return &ReportService{reportRepo: reportRepo}
}

func (p *ReportService) QueryAdxReport(req *request.QueryReportReq) ([]*response.AdxReportResponse, error) {
	return p.reportRepo.QueryAdxData(req)
}
