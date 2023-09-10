package server

import (
	"fmt"
	"net/http"
	"html/template"

	"github.com/sinalalebakhsh/Gocron/features"
	"github.com/gorilla/mux"
	"github.com/go-redis/redis"
)

var client *redis.Client
var templates *template.Template

func HandlerAllFeatures(w http.ResponseWriter, r *http.Request) {
	comments, err := client.LRange("comments", 0, 10).Result()
	if err != nil {
		return
	}
	templates.ExecuteTemplate(w, "index.html", comments)
}

func HandlerSingleDefinitions(w http.ResponseWriter, r *http.Request) {
	Gold := fmt.Sprint(features.OriginSingleDef)
	fmt.Fprint(w, Gold)
}



func Servers() {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
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
