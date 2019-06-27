package main

import (
	"database/sql"
	"database/sql/driver"
	"io"
	"strconv"

	_ "gopkg.in/goracle.v2"

	cm "github.com/armbodyslam/campaignbenefitbackend/common"
	st "github.com/armbodyslam/campaignbenefitbackend/structs"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//GetCampaign for get list campaign
func (db *MongoDBInfo) GetCampaign() []st.Campaign {

	session, err := mgo.Dial(db.URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(db.database).C(db.collection)
	res := []st.Campaign{}

	err = c.Find(bson.M{"status": "A"}).All(&res)

	if err != nil {
		return nil
		//log.Fatal(err)
	}

	return res
}

//GetCustProfile for get list Profile
func (db *MongoDBInfo) GetCustProfile() []st.CustProfileMaster {

	session, err := mgo.Dial(db.URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(db.database).C(db.collection)
	res := []st.CustProfileMaster{}

	err = c.Find(nil).All(&res)

	if err != nil {
		return nil
		//log.Fatal(err)
	}

	return res
}

//GetPackageMaster for get list campaign
func (db *MongoDBInfo) GetPackageMaster() []st.PackageMaster {

	session, err := mgo.Dial(db.URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(db.database).C(db.collection)
	res := []st.PackageMaster{}

	err = c.Find(nil).Sort("packageid").All(&res)

	if err != nil {
		return nil
		//log.Fatal(err)
	}

	return res
}

//GetPreviewProduct for get list PreviewProduct
func (db *MongoDBInfo) GetPreviewProduct() []st.PreviewProduct {

	session, err := mgo.Dial(db.URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(db.database).C(db.collection)
	res := []st.PreviewProduct{}

	err = c.Find(nil).Sort("previewproductid").All(&res)

	if err != nil {
		return nil
		//log.Fatal(err)
	}

	return res
}

//GetOffer for get list Offer
func (db *MongoDBInfo) GetOffer() []st.Offer {

	session, err := mgo.Dial(db.URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(db.database).C(db.collection)
	res := []st.Offer{}

	err = c.Find(nil).Sort("offerid").All(&res)

	if err != nil {
		return nil
		//log.Fatal(err)
	}

	return res
}

//GetLastCampaignID for get last campaignID
func (db *MongoDBInfo) GetLastCampaignID() int {

	var res int

	session, err := mgo.Dial(db.URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(db.database).C(db.collection)
	campaign := st.Campaign{}

	err = c.Find(nil).Sort("-campaignid").One(&campaign)
	if err != nil {
		return 0
		//log.Fatal(err)
	}

	res = campaign.CampaignID

	return res
}

// CreateCampaign for insert new document
func (db *MongoDBInfo) CreateCampaign(camp st.Campaign) *st.CreateCampaignResponse {

	oRes := st.NewCreateCampaignResponse()
	schedule := camp.Schedule

	if schedule.Type == `daily` {
		if len(schedule.Execute) == 1 {
			schedule.Execute = `0` + schedule.Execute
		}
	}

	camp.Schedule = schedule

	session, err := mgo.Dial(db.URL)
	if err != nil {
		oRes.ErrorCode = 1
		oRes.ErrorDesc = err.Error()

		return oRes
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(db.database).C(db.collection)

	lastcampaign := st.Campaign{}
	err = c.Find(nil).Sort("-campaignid").One(&lastcampaign)
	if err != nil {
		oRes.ErrorCode = 2
		oRes.ErrorDesc = err.Error()

		return oRes
		//log.Fatal(err)
	}

	lastID := lastcampaign.CampaignID
	campaignID := lastID + 1

	camp.ID = bson.NewObjectId()
	camp.CampaignID = campaignID
	camp.UpdateSQL = "Y"
	err = c.Insert(camp)
	if err != nil {
		oRes.ErrorCode = 3
		oRes.ErrorDesc = err.Error()

		return oRes
		//log.Fatal(err)
	}

	oRes.ErrorCode = 0
	oRes.ResultValue = "Success"
	oRes.ErrorDesc = ""
	return oRes
}

//GetKeyword for get list Keyword
func GetKeyword() *st.GetListKeywordResult {

	res := st.NewGetListKeywordResult()
	var oListKeyword st.ListKeyword
	var dbsource string
	dbsource = cm.GetDatasourceName("ICC")
	dbicc, err := sql.Open("goracle", dbsource)
	if err != nil {

		res.ErrorCode = 2
		res.ErrorDesc = err.Error()
		return res
	}

	defer dbicc.Close()
	var statement string
	statement = "begin TVS_CAMPAIGN.GetKeyword(:0); end;"
	var resultC driver.Rows
	if _, err := dbicc.Exec(statement, sql.Out{Dest: &resultC}); err != nil {

		res.ErrorCode = 3
		res.ErrorDesc = err.Error()
		return res
	}
	defer resultC.Close()
	values := make([]driver.Value, len(resultC.Columns()))
	var oLKeyword []st.Keyword

	for {

		colmap := cm.Createmapcol(resultC.Columns())
		err = resultC.Next(values)
		if err != nil {
			if err == io.EOF {
				break
			}

			res.ErrorCode = 4
			res.ErrorDesc = err.Error()
			return res
		}

		var oKeyword st.Keyword
		if values[cm.Getcolindex(colmap, "ID")] != nil {
			oKeyword.KeywordID = values[cm.Getcolindex(colmap, "ID")].(int64)
		}

		oKeyword.KaKeyword = values[cm.Getcolindex(colmap, "KAKEYWORD")].(string)
		oKeyword.KaAttribute = values[cm.Getcolindex(colmap, "KAATTRIBUTE")].(string)
		oKeyword.KaLongDescr = values[cm.Getcolindex(colmap, "KALONGDESCR")].(string)

		if values[cm.Getcolindex(colmap, "KAHIDE")] != nil {
			oKeyword.KaHide = values[cm.Getcolindex(colmap, "KAHIDE")].(int64)
		}

		if values[cm.Getcolindex(colmap, "KEYTYPES_ID")] != nil {
			oKeyword.KeyTypesID = values[cm.Getcolindex(colmap, "KEYTYPES_ID")].(int64)
		}

		oLKeyword = append(oLKeyword, oKeyword)
	}

	oListKeyword.Keywords = oLKeyword
	res.MyListKeyword = oListKeyword

	if res.ErrorCode == 1 {

		res.ErrorCode = 0
		res.ErrorDesc = "Success"
	}

	return res
}

//GetCampaignByID for get list campaign
func (db *MongoDBInfo) GetCampaignByID(ID string) *st.GetCampaignResponse {

	res := st.NewGetCampaignResponse()

	nID, err := strconv.Atoi(ID)
	if err != nil {
		res.ErrorCode = 2
		res.ErrorDesc = err.Error()
		return res
	}

	session, err := mgo.Dial(db.URL)
	if err != nil {
		res.ErrorCode = 3
		res.ErrorDesc = err.Error()
		return res
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(db.database).C(db.collection)

	var campaign st.Campaign

	err = c.Find(bson.M{"campaignid": nID, "status": "A"}).One(&campaign)

	if err != nil {
		res.ErrorCode = 4
		res.ErrorDesc = err.Error()
		return res
		//log.Fatal(err)
	}

	res.MyCampaign = campaign
	res.ErrorCode = 0
	res.ErrorDesc = ""

	return res
}
