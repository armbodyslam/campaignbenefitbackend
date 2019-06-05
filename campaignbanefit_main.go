package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	st "github.com/armbodyslam/campaignbenefitbackend/structs"
	"github.com/gorilla/mux"
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

func main() {

	fmt.Println("Service Start...")
	mainRouter := mux.NewRouter().StrictSlash(true)
	mainRouter.HandleFunc("/campaign", index)
	mainRouter.HandleFunc("/campaign/getcampaign", getcampaign)
	mainRouter.HandleFunc("/custprofilemaster/getcustprofile", getcustprofile)
	mainRouter.HandleFunc("/packagemaster/getpackage", getpackage)
	log.Fatal(http.ListenAndServe(":8000", mainRouter))
}

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to campaign Restful")
}

func getcampaign(w http.ResponseWriter, r *http.Request) {

	// Create db connection to mongo
	db := Create("", "", "172.19.218.104", "27017", "tvscampaigndb", "campaign")

	var res st.ListCampaign
	res.Campaigns = db.GetCampaign()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(res)
}

func getcustprofile(w http.ResponseWriter, r *http.Request) {
	// Create db connection to mongo
	db := Create("", "", "172.19.218.104", "27017", "tvscampaigndb", "custprofilemaster")
	var res st.ListCustProfileMaster
	res.CustProfileMasters = db.GetCustProfile()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(res)

}

func getpackage(w http.ResponseWriter, r *http.Request) {

	// Create db connection to mongo
	db := Create("", "", "172.19.218.104", "27017", "tvscampaigndb", "packagemaster")
	var res st.ListPackageMaster
	res.PackageMasters = db.GetPackageMaster()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(res)
}
