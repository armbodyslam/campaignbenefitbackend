package structs

import (
	"gopkg.in/mgo.v2/bson"
)

// CustProfileMaster ...
type CustProfileMaster struct {
	ID         bson.ObjectId `bson:"_id"`
	Profile    string        `json:"profile"`
	SubProfile string        `json:"subprofile"`
}

//ListCustProfileMaster for list
type ListCustProfileMaster struct {
	CustProfileMasters []CustProfileMaster `json:"custprofilemaster"`
}
