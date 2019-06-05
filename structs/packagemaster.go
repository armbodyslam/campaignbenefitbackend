package structs

import (
	"gopkg.in/mgo.v2/bson"
)

// PackageMaster ...
type PackageMaster struct {
	ID        bson.ObjectId `bson:"_id"`
	PackageID int           `json:"packageid"`
	Name      string        `json:"name"`
}

//ListPackageMaster for list
type ListPackageMaster struct {
	PackageMasters []PackageMaster `json:"packagemaster"`
}
