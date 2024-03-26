package repository

import (
	"adx-admin/internal/domian/request"
	"adx-admin/internal/domian/response"
)

type ReportRepo interface {
	QueryAdxData(req *request.QueryReportReq) ([]*response.AdxReportResponse, error)
}
