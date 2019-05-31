package structs

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Campaign ...
type Campaign struct {
	ID           bson.ObjectId `bson:"_id"`
	CampaignID   int           `json:"campaignid"`
	CampaignName string        `json:"campaignName"`
	Description  string        `json:"description"`
	StartDate    time.Time     `bson:"startdate"`
	EndDate      time.Time     `bson:"enddate"`
	PeriodBy     string        `json:"periodby"`
	Schedule     struct {
		Type    string `json:"type"`
		Execute string `json:"execute"`
	} `json:"schedule"`
	Status       string   `json:"status"`
	ProfileAllow []string `json:"profileallow,omitempty"`
	PackageAllow []int    `json:"packageallow,omitempty"`
	KeywordAllow []int    `json:"keywordallow,omitempty"`
	ProductAdd   []struct {
		ProductNr int    `json:"productnr"`
		DayAdd    int    `json:"dayadd"`
		EndDate   string `json:"enddate"`
	} `json:"productadd,omitempty"`
	OfferAdd []struct {
		OfferNr     int    `json:"offernr"`
		DayAdd      int    `json:"dayadd"`
		DayFormular int    `json:"dayformular"`
		MonthAdd    int    `json:"monthadd"`
		EndDate     string `json:"enddate"`
	} `json:"offeradd,omitempty"`
	SQL       string `json:"sql"`
	UpdateSQL string `json:"updatesql"`
}
