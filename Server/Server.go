package server

import (
	"io"
	"log"
	"net/http"
	"os"
)


type StringHandler struct {
	Message string
}

func MyServer() {
	logFile, err := os.OpenFile("Server/LogFile/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	err = http.ListenAndServe(":5000", StringHandler{Message: "Hello, World"})
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
	io.WriteString(writer, sh.Message)
	log.Printf("============================================◉🧭🧭🧭🧭🧭🧭🧭◉==========================================")
}
