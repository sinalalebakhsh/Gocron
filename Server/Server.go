package server

import (
	"fmt"
	"net/http"

	"github.com/sinalalebakhsh/Gocron/features"
	"github.com/gorilla/mux"
)


func HandlerAllFeatures(w http.ResponseWriter, r *http.Request) {
	Gold := fmt.Sprint(features.OriginalFeatures)
	fmt.Fprint(w, Gold)
}

func HandlerSingleDefinitions(w http.ResponseWriter, r *http.Request) {
	Gold := fmt.Sprint(features.OriginSingleDef)
	fmt.Fprint(w, Gold)
}



func Servers() {
	router := mux.NewRouter()
	http.Handle("/", router)
	router.HandleFunc("/", HandlerAllFeatures).Methods("GET")
	router.HandleFunc("/SingleDefinitions", HandlerSingleDefinitions).Methods("GET")
	http.ListenAndServe(":8080", nil)
}


// Future possibilities 
	// For flexible URL:
		// "github.com/sinalalebakhsh/Gocron/GetUserInput"
		// getuserinput.GetUserInput()
