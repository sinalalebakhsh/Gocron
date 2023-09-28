package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sinalalebakhsh/Gocron/features"
	"io"
	"log"
	"net/http"
	"os"
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


type StringHandler struct {
	message string
}

func MyServer() {
	logFile, err := os.OpenFile("LogFile/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	err = http.ListenAndServe(":5000", StringHandler{message: "Hello, World"})
	if err != nil {
		log.Printf("Error: %v", err.Error())
	}
}

func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Printf("Method: %v", request.Method)
	log.Printf("URL: %v", request.URL)
	log.Printf("HTTP Version: %v", request.Proto)
	log.Printf("Host: %v", request.Host)
	for name, val := range request.Header {
		log.Printf("Header: %v, Value: %v", name, val)
	}
	io.WriteString(writer, sh.message)
	log.Printf("============================================◉🧭🧭🧭🧭🧭🧭🧭◉==========================================")
}
