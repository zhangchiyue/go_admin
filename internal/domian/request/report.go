package request

type QueryReportReq struct {
	GroupBy *ReportGroupBy `json:"group_data"`
}

type ReportGroupBy struct {
	Area bool `json:"area,omitempty"`
	Hour bool `json:"hour,omitempty"`
}
