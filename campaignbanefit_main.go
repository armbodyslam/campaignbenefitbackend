package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Service Start...")
	mainRouter := mux.NewRouter().StrictSlash(true)
	mainRouter.HandleFunc("/campaign", index)
	mainRouter.HandleFunc("/campaign/getcampaign", getcampaign)
	log.Fatal(http.ListenAndServe(":8000", mainRouter))
}

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to campaign Restful")
}

func getcampaign(w http.ResponseWriter, r *http.Request) {

	res := GetCampaign()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
