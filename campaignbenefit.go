package main

import (
	"strings"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	st "github.com\armbodyslam\campaignbenefitbackend\structs"
)

//MongoDBInfo for ..
type MongoDBInfo struct {
	user       string
	password   string
	host       string
	port       string
	database   string
	collection string
	URL        string
}



// Create for create dbinfo
func Create(user string, pass string, host string, port string, database string, collection string) *MongoDBInfo {
	db := &MongoDBInfo{user: user, password: pass, host: host, port: port, database: database, collection: collection}

	var url string
	if user == "" || pass == "" {
		url = "mongodb://$host:$port/$db"
	} else {
		url = "mongodb://$user:$pass@$host:$port/$db"
	}
	url = strings.Replace(url, "$user", user, -1)
	url = strings.Replace(url, "$pass", pass, -1)
	url = strings.Replace(url, "$host", host, -1)
	url = strings.Replace(url, "$port", port, -1)
	url = strings.Replace(url, "$db", database, -1)

	db.URL = url

	return db
}

//GetCampaign for get list campaign
func (db *MongoDBInfo) GetCampaign() []Campaign {

	session, err := mgo.Dial(db.URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(db.database).C(db.collection)
	res := []Campaign{}
	err = c.Find(nil).All(&res)
	//fmt.Println("GetCampaign...")
	if err != nil {
		return nil
		//log.Fatal(err)
	}

	return res
}
