package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	st "github.com/armbodyslam/campaignbenefitbackend/structs"
	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Service Start...")
	mainRouter := mux.NewRouter().StrictSlash(true)
	mainRouter.HandleFunc("/campaign", index)
	mainRouter.HandleFunc("/campaign/getcampaign", getcampaign)
	mainRouter.HandleFunc("/campaign/getlastcampaignid", getlastcampaignid)
	mainRouter.HandleFunc("/campaign/getcampaign/{campaignid}", getcampaignbyid)
	mainRouter.HandleFunc("/campaign/createcampaign", createcampaign).Methods("POST")
	mainRouter.HandleFunc("/campaign/cancelcampaign", cancelcampaign).Methods("POST")
	mainRouter.HandleFunc("/campaign/searchcampaign", searchcampaign).Methods("POST")
	mainRouter.HandleFunc("/custprofilemaster/getcustprofile", getcustprofile)
	mainRouter.HandleFunc("/packagemaster/getpackage", getpackage)
	mainRouter.HandleFunc("/previewproduct/getpreview", getpreview)
	mainRouter.HandleFunc("/offer/getoffer", getoffer)
	mainRouter.HandleFunc("/keyword/getkeyword", getkeyword)
	mainRouter.HandleFunc("/report/getreportcampaign", getreportcampaignbyid).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", mainRouter))
}

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to campaign Restful")
}

func getcampaign(w http.ResponseWriter, r *http.Request) {

	// Create db connection to mongo
	//db := Create("", "", "172.19.218.104", "27017", "tvscampaigndb", "campaign")

	var res st.ListCampaign
	res.Campaigns = GetCampaign()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(res)
}

func getcampaignbyid(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var res *st.GetCampaignResponse

	// Create db connection to mongo
	//db := Create("", "", "172.19.218.104", "27017", "tvscampaigndb", "campaign")

	res = GetCampaignByID(params["campaignid"])
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(res)

}

func createcampaign(w http.ResponseWriter, r *http.Request) {

	// Create db connection to mongo
	//db := Create("", "", "172.19.218.104", "27017", "tvscampaigndb", "campaign")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var req st.Campaign
	err = json.Unmarshal(body, &req)
	if err != nil {
		panic(err)
	}

	var res *st.CreateCampaignResponse

	res = CreateCampaign(req)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(res)
	//log.Println(req)
}

func cancelcampaign(w http.ResponseWriter, r *http.Request) {

	// Create db connection to mongo
	//db := Create("", "", "172.19.218.104", "27017", "tvscampaigndb", "campaign")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var res *st.CancelCampaignResponse

	var req st.CancelCampaignRequest

	err = json.Unmarshal(body, &req)
	if err != nil {

		panic(err)
	}

	res = CancelCampaign(req.CampaignID)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(res)
}

func searchcampaign(w http.ResponseWriter, r *http.Request) {

	// Create db connection to mongo
	//db := Create("", "", "172.19.218.104", "27017", "tvscampaigndb", "campaign")

	var res *st.SearchCampaignResponse

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var req st.SearchCampaignRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		panic(err)
	}

	res = SearchCampaign(req)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(res)
}

func getcustprofile(w http.ResponseWriter, r *http.Request) {

	// Create db connection to mongo
	//db := Create("", "", "172.19.218.104", "27017", "tvscampaigndb", "custprofilemaster")
	var res st.ListCustProfileMaster
	res.CustProfileMasters = GetCustProfile()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(res)

}

func getpackage(w http.ResponseWriter, r *http.Request) {

	// Create db connection to mongo
	//db := Create("", "", "172.19.218.104", "27017", "tvscampaigndb", "packagemaster")
	var res st.ListPackageMaster
	res.PackageMasters = GetPackageMaster()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(res)
}

func getpreview(w http.ResponseWriter, r *http.Request) {

	// Create db connection to mongo
	//db := Create("", "", "172.19.218.104", "27017", "tvscampaigndb", "previewproduct")
	var res st.ListPreviewProduct
	res.PreviewProducts = GetPreviewProduct()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(res)
}

func getoffer(w http.ResponseWriter, r *http.Request) {

	// Create db connection to mongo
	//db := Create("", "", "172.19.218.104", "27017", "tvscampaigndb", "offer")
	var res st.ListOffer
	res.Offers = GetOffer()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(res)
}

func getlastcampaignid(w http.ResponseWriter, r *http.Request) {

	// Create db connection to mongo
	//db := Create("", "", "172.19.218.104", "27017", "tvscampaigndb", "campaign")

	var res int
	res = GetLastCampaignID()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(res)
}

func getkeyword(w http.ResponseWriter, r *http.Request) {

	var res *st.GetListKeywordResult

	res = GetKeyword()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(res)
}

func getreportcampaignbyid(w http.ResponseWriter, r *http.Request) {

	//params := mux.Vars(r)

	var res *st.GetListReportCampignResponse

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var req st.GetListReportCampignRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		panic(err)
	}

	res = GetReportCampaignByID(req)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(res)
}
