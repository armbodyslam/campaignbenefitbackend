package structs

import (
	"time"
)

//Reportcampign object
type Reportcampign struct {
	Runnumber    int64     `json:"runnumber"`
	CustomerNR   int64     `json:"customernr"`
	Fullname     string    `json:"fullname"`
	StatusResult string    `json:"statusresult"`
	OfferDate    time.Time `json:"offerdate"`
}

//GetListReportCampignRequest obj
type GetListReportCampignRequest struct {
	CampaignID int64        `json:"campaignid"`
	Fullname   string       `json:"fullname"`
	CustomerNR int64        `json:"customernr"`
	PageSize   int64        `json:"pagesize"`
	Page       int64        `json:"page"`
	Sorted     []SortReport `json:"sorted"`
}

// SortReport Obj
type SortReport struct {
	ID   string `json:"id"`
	Desc bool   `json:"desc"`
}

//GetListReportCampignResponse obj
type GetListReportCampignResponse struct {
	ErrorCode      int             `json:"errorcode"`
	ErrorDesc      string          `json:"errordesc"`
	Reportcampigns []Reportcampign `json:"reportcampign"`
	Pages          int64           `json:"pages"`
}

// NewGetListReportCampignResponse Obj
func NewGetListReportCampignResponse() *GetListReportCampignResponse {
	return &GetListReportCampignResponse{
		ErrorCode: 1,
		ErrorDesc: "Unexpected Error",
	}
}
