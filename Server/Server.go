package server

import (
	"fmt"
	"net/http"

	"github.com/sinalalebakhsh/Gocron/features"
	"github.com/gorilla/mux"

	"html/template"
)

var templates *template.Template

func HandlerAllFeatures(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func HandlerSingleDefinitions(w http.ResponseWriter, r *http.Request) {
	Gold := fmt.Sprint(features.OriginSingleDef)
	fmt.Fprint(w, Gold)
}



func Servers() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
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
