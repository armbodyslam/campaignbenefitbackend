package main

import (
	st "github.com/armbodyslam/campaignbenefitbackend/structs"
	mgo "gopkg.in/mgo.v2"
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

	err = c.Find(nil).All(&res)

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

	err = c.Find(nil).All(&res)

	if err != nil {
		return nil
		//log.Fatal(err)
	}

	return res
}
