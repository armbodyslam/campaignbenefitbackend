package structs

import (
	"gopkg.in/mgo.v2/bson"
)

// PreviewProduct ...
type PreviewProduct struct {
	ID               bson.ObjectId `bson:"_id"`
	PreviewProductID int           `json:"previewproductid"`
	Name             string        `json:"name"`
}

//ListPreviewProduct for list
type ListPreviewProduct struct {
	PreviewProducts []PreviewProduct `json:"previewproduct"`
}
