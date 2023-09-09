package server

import (
	"fmt"
	"net/http"

	"github.com/sinalalebakhsh/Gocron/features"
)

func HandlerAllFeatures(w http.ResponseWriter, r *http.Request) {
	Gold := fmt.Sprint(features.OriginalFeatures)

	fmt.Fprint(w, Gold)

}

func Servers() {
	http.HandleFunc("/", HandlerAllFeatures)
	http.ListenAndServe(":8080", nil)
}


