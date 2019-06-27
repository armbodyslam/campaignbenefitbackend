package structs

import (
	"gopkg.in/mgo.v2/bson"
)

// Offer ...
type Offer struct {
	ID          bson.ObjectId `bson:"_id"`
	OfferID     int           `json:"offerid"`
	Description string        `json:"description"`
}

//ListOffer for list
type ListOffer struct {
	Offers []Offer `json:"offer"`
}
