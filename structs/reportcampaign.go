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

//GetListReportCampignResponse obj
type GetListReportCampignResponse struct {
	ErrorCode      int             `json:"errorcode"`
	ErrorDesc      string          `json:"errordesc"`
	Reportcampigns []Reportcampign `json:"reportcampign"`
}

// NewGetListReportCampignResponse Obj
func NewGetListReportCampignResponse() *GetListReportCampignResponse {
	return &GetListReportCampignResponse{
		ErrorCode: 1,
		ErrorDesc: "Unexpected Error",
	}
}
