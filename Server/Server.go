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

func Servers() {
	router := mux.NewRouter()

	router.HandleFunc("/", HandlerAllFeatures).Methods("GET")
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}


