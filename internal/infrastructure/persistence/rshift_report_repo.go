package persistence

import (
	"adx-admin/internal/domian/repository"
	"adx-admin/internal/domian/request"
	"adx-admin/internal/domian/response"
	"adx-admin/pkg/database/redshift"
	"fmt"
)

type RShiftReportRepo struct {
	client *redshift.Redshift
}

func NewRShiftReportRepo(client *redshift.Redshift) repository.ReportRepo {
	return &RShiftReportRepo{
		client: client,
	}
}

func (p *RShiftReportRepo) QueryAdxData(req *request.QueryReportReq) ([]*response.AdxReportResponse, error) {
	var results []*response.AdxReportResponse
	fmt.Printf("%+v\n", req)
	var count int64
	p.client.DB.Table("adx_report_hour").Count(&count)
	results = append(results, &response.AdxReportResponse{
		Count: count,
	})
	return results, nil
}
