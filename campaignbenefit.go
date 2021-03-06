package main

import (
	"database/sql"
	"database/sql/driver"
	"io"
	"math"
	"strconv"
	"time"

	_ "gopkg.in/goracle.v2"
	goracle "gopkg.in/goracle.v2"

	cm "github.com/armbodyslam/campaignbenefitbackend/common"
	st "github.com/armbodyslam/campaignbenefitbackend/structs"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//GetCampaign for get list campaign
func GetCampaign() []st.Campaign {

	var dbMgSource cm.MongoDBInfo
	dbMgSource = cm.GetDatasourceNameMongo("TVSCAMPAIGN")

	session, err := mgo.Dial(dbMgSource.URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(dbMgSource.Database).C("campaign")
	res := []st.Campaign{}

	err = c.Find(bson.M{"status": "A"}).Sort("-campaignid").All(&res)

	if err != nil {
		return nil
		//log.Fatal(err)
	}

	return res
}

//GetCustProfile for get list Profile
func GetCustProfile() []st.CustProfileMaster {

	var dbMgSource cm.MongoDBInfo
	dbMgSource = cm.GetDatasourceNameMongo("TVSCAMPAIGN")

	session, err := mgo.Dial(dbMgSource.URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(dbMgSource.Database).C("custprofilemaster")
	res := []st.CustProfileMaster{}

	err = c.Find(nil).All(&res)

	if err != nil {
		return nil
		//log.Fatal(err)
	}

	return res
}

//GetPackageMaster for get list campaign
func GetPackageMaster() []st.PackageMaster {

	var dbMgSource cm.MongoDBInfo
	dbMgSource = cm.GetDatasourceNameMongo("TVSCAMPAIGN")

	session, err := mgo.Dial(dbMgSource.URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(dbMgSource.Database).C("packagemaster")
	res := []st.PackageMaster{}

	err = c.Find(nil).Sort("packageid").All(&res)

	if err != nil {
		return nil
		//log.Fatal(err)
	}

	return res
}

//GetPreviewProduct for get list PreviewProduct
func GetPreviewProduct() []st.PreviewProduct {

	var dbMgSource cm.MongoDBInfo
	dbMgSource = cm.GetDatasourceNameMongo("TVSCAMPAIGN")

	session, err := mgo.Dial(dbMgSource.URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(dbMgSource.Database).C("previewproduct")
	res := []st.PreviewProduct{}

	err = c.Find(nil).Sort("previewproductid").All(&res)

	if err != nil {
		return nil
		//log.Fatal(err)
	}

	return res
}

//GetOffer for get list Offer
func GetOffer() []st.Offer {

	var dbMgSource cm.MongoDBInfo
	dbMgSource = cm.GetDatasourceNameMongo("TVSCAMPAIGN")

	session, err := mgo.Dial(dbMgSource.URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(dbMgSource.Database).C("offer")
	res := []st.Offer{}

	err = c.Find(nil).Sort("offerid").All(&res)

	if err != nil {
		return nil
		//log.Fatal(err)
	}

	return res
}

//GetLastCampaignID for get last campaignID
func GetLastCampaignID() int {

	var dbMgSource cm.MongoDBInfo
	dbMgSource = cm.GetDatasourceNameMongo("TVSCAMPAIGN")

	var res int

	session, err := mgo.Dial(dbMgSource.URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(dbMgSource.Database).C("campaign")
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
func CreateCampaign(camp st.Campaign) *st.CreateCampaignResponse {

	var dbMgSource cm.MongoDBInfo
	dbMgSource = cm.GetDatasourceNameMongo("TVSCAMPAIGN")

	oRes := st.NewCreateCampaignResponse()
	schedule := camp.Schedule

	if schedule.Type == `daily` {
		if len(schedule.Execute) == 1 {
			schedule.Execute = `0` + schedule.Execute
		}
	}

	camp.Schedule = schedule

	session, err := mgo.Dial(dbMgSource.URL)
	if err != nil {
		oRes.ErrorCode = 1
		oRes.ErrorDesc = err.Error()

		return oRes
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(dbMgSource.Database).C("campaign")

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
func GetCampaignByID(ID string) *st.GetCampaignResponse {

	var dbMgSource cm.MongoDBInfo
	dbMgSource = cm.GetDatasourceNameMongo("TVSCAMPAIGN")

	res := st.NewGetCampaignResponse()

	nID, err := strconv.Atoi(ID)
	if err != nil {
		res.ErrorCode = 2
		res.ErrorDesc = err.Error()
		return res
	}

	session, err := mgo.Dial(dbMgSource.URL)
	if err != nil {
		res.ErrorCode = 3
		res.ErrorDesc = err.Error()
		return res
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(dbMgSource.Database).C("campaign")

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

//GetReportCampaignByID for get list report campaign
func GetReportCampaignByID(reportSearch st.GetListReportCampignRequest) *st.GetListReportCampignResponse {

	res := st.NewGetListReportCampignResponse()
	//
	iCampaignID := reportSearch.CampaignID
	iFullName := reportSearch.Fullname
	iCustomernr := reportSearch.CustomerNR
	pageNo := reportSearch.Page
	size := reportSearch.PageSize
	csort := "offerdate"
	desc := "desc"

	if len(reportSearch.Sorted) > 0 {
		csort = reportSearch.Sorted[0].ID
		if !reportSearch.Sorted[0].Desc {
			desc = "asc"
		}
	}

	if pageNo <= 0 {
		res.ErrorCode = 2
		res.ErrorDesc = "invalid page number, should start with 1"
		return res
	}

	skip := size * (pageNo - 1)

	var dbsource string
	dbsource = cm.GetDatasourceName("SIT62")
	db, err := sql.Open("goracle", dbsource)
	if err != nil {

		res.ErrorCode = 3
		res.ErrorDesc = err.Error()
		return res
	}
	defer db.Close()

	var statement string
	var resultC driver.Rows
	var resultCount driver.Rows

	var totalCount int64

	statement = "begin TVS_CAMPAIGN.GetCampaignReportById(:0,:1,:2,:3,:4,:5,:6,:7,:8); end;"
	//var resultC driver.Rows
	if _, err := db.Exec(statement, iCampaignID, iFullName, iCustomernr, skip, size, csort, desc, sql.Out{Dest: &resultC},
		sql.Out{Dest: &resultCount}); err != nil {

		res.ErrorCode = 5
		res.ErrorDesc = err.Error()
		return res
	}
	defer resultC.Close()
	defer resultCount.Close()

	values := make([]driver.Value, len(resultC.Columns()))
	var oLReportCampaign []st.Reportcampign
	var i int64
	for {

		colmap := cm.Createmapcol(resultC.Columns())
		err = resultC.Next(values)
		if err != nil {

			if err == io.EOF {

				break
			}

			res.ErrorCode = 6
			res.ErrorDesc = err.Error()
			return res
		}

		var oReportCampaign st.Reportcampign
		i++
		oReportCampaign.Runnumber = i
		/* if values[cm.Getcolindex(colmap, "RUNNUMBER")] != nil {
			g, _ := values[cm.Getcolindex(colmap, "RUNNUMBER")].(goracle.Number)
			i := string(g)
			oReportCampaign.Runnumber, _ = strconv.ParseInt(i, 10, 64)
		} */
		if values[cm.Getcolindex(colmap, "CUSTOMERNR")] != nil {
			oReportCampaign.CustomerNR = values[cm.Getcolindex(colmap, "CUSTOMERNR")].(int64)
		}

		oReportCampaign.Fullname = values[cm.Getcolindex(colmap, "FULLNAME")].(string)
		oReportCampaign.StatusResult = values[cm.Getcolindex(colmap, "STATUSRESULT")].(string)

		if values[cm.Getcolindex(colmap, "OFFERDATE")] != nil {
			oReportCampaign.OfferDate = values[cm.Getcolindex(colmap, "OFFERDATE")].(time.Time)
		}

		oLReportCampaign = append(oLReportCampaign, oReportCampaign)
	}

	values = make([]driver.Value, len(resultCount.Columns()))
	for {

		colmap := cm.Createmapcol(resultCount.Columns())
		err = resultCount.Next(values)
		if err != nil {

			if err == io.EOF {

				break
			}

			res.ErrorCode = 6
			res.ErrorDesc = err.Error()
			return res
		}

		if values[cm.Getcolindex(colmap, "COUNTREPORT")] != nil {
			g, _ := values[cm.Getcolindex(colmap, "COUNTREPORT")].(goracle.Number)
			i := string(g)
			totalCount, _ = strconv.ParseInt(i, 10, 64)
		}
	}

	d := float64(totalCount) / float64(size)
	totalPages := int(math.Ceil(d))
	res.Pages = int64(totalPages)

	//oListReportCampign.Reportcampigns = oLReportCampaign
	res.Reportcampigns = oLReportCampaign

	if res.Reportcampigns == nil {

		res.ErrorCode = 6
		res.ErrorDesc = "not found"
	}

	if res.ErrorCode == 1 {
		res.ErrorCode = 0
		res.ErrorDesc = "Success"
	}

	return res
}

//CancelCampaign for cancel campaign
func CancelCampaign(campID int) *st.CancelCampaignResponse {

	var dbMgSource cm.MongoDBInfo
	dbMgSource = cm.GetDatasourceNameMongo("TVSCAMPAIGN")

	res := st.NewCancelCampaignResponse()

	session, err := mgo.Dial(dbMgSource.URL)
	if err != nil {
		res.ErrorCode = 2
		res.ErrorDesc = err.Error()

		return res
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB(dbMgSource.Database).C("campaign")
	err = c.Update(bson.M{"campaignid": campID}, bson.M{"$set": bson.M{"status": "C"}})
	if err != nil {
		res.ErrorCode = 3
		res.ErrorDesc = err.Error()

		return res
		//log.Fatal(err)
	}

	if res.ErrorCode == 1 {
		res.ErrorCode = 0
		res.ErrorDesc = ""
		res.ResultValue = "Success"
		return res
	}
	return res
}

//SearchCampaign for Search Campaign
func SearchCampaign(campSearch st.SearchCampaignRequest) *st.SearchCampaignResponse {

	var dbMgSource cm.MongoDBInfo
	dbMgSource = cm.GetDatasourceNameMongo("TVSCAMPAIGN")

	res := st.NewSearchCampaignResponse()

	var sorted string

	//var checkArraySort []st.SortCampaign

	if len(campSearch.Sorted) > 0 {
		if campSearch.Sorted[0].Desc {

			sorted = "-" + campSearch.Sorted[0].ID
		} else {
			sorted = campSearch.Sorted[0].ID
		}
	} else {
		sorted = "-campaignid"
	}

	pageNo := campSearch.Page
	size := campSearch.PageSize

	if pageNo <= 0 {
		res.ErrorCode = 2
		res.ErrorDesc = "invalid page number, should start with 1"
		return res
	}

	skip := size * (pageNo - 1)

	session, err := mgo.Dial(dbMgSource.URL)
	if err != nil {

		res.ErrorCode = 3
		res.ErrorDesc = err.Error()
		return res
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(dbMgSource.Database).C("campaign")

	var camID bson.M
	if campSearch.CampaignID == 0 {
		camID = bson.M{"campaignid": bson.M{"$gt": 0}}
	} else {
		camID = bson.M{"campaignid": campSearch.CampaignID}
	}

	camName := bson.M{"campaignname": bson.M{"$regex": campSearch.CampaignName, "$options": "i"}}

	var camStartDate bson.M
	if campSearch.StartDate.IsZero() {
		camStartDate = bson.M{"startdate": bson.M{"$gt": campSearch.StartDate}}
	} else {
		camStartDate = bson.M{"startdate": bson.M{"$eq": campSearch.StartDate}}
	}

	var camEndDate bson.M
	if campSearch.EndDate.IsZero() {
		camEndDate = bson.M{"enddate": bson.M{"$gt": campSearch.EndDate}}
	} else {
		camEndDate = bson.M{"enddate": bson.M{"$eq": campSearch.EndDate}}
	}

	camScheduleType := bson.M{"schedule.type": bson.M{"$regex": campSearch.Schedule.Type,
		"$options": "i"}}
	camScheduleExecute := bson.M{"schedule.execute": bson.M{"$regex": campSearch.Schedule.Execute,
		"$options": "i"}}

	camStatus := bson.M{"status": "A"}

	totalCount, err := c.Find(bson.M{"$and": []bson.M{camID, camName, camStartDate, camEndDate, camScheduleType,
		camScheduleExecute, camStatus}}).Sort("-campaignid").Count()
	if err != nil {

		res.ErrorCode = 4
		res.ErrorDesc = err.Error()
		return res
		//log.Fatal(err)
	}

	d := float64(totalCount) / float64(size)
	totalPages := int(math.Ceil(d))

	err = c.Find(bson.M{"$and": []bson.M{camID, camName, camStartDate, camEndDate, camScheduleType,
		camScheduleExecute, camStatus}}).Sort(sorted).Skip(skip).Limit(size).All(&res.Campaigns)
	if err != nil {

		res.ErrorCode = 5
		res.ErrorDesc = err.Error()
		return res
		//log.Fatal(err)
	}

	res.Pages = totalPages
	if res.ErrorCode == 1 {
		res.ErrorCode = 0
		res.ErrorDesc = ""

		return res
	}

	return res
}
