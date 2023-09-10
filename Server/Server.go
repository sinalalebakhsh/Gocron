package server

import (
	"fmt"
	"github.com/sinalalebakhsh/Gocron/features"
	"net/http"
	"os"
)

func HandlerAllFeatures(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, features.OriginalFeatures)
}

func Servers() {
	if len(os.Args) == 1 {
		http.HandleFunc("/", HandlerAllFeatures)
	} else if len(os.Args) > 1 {
		
	}
	http.ListenAndServe(":8080", nil)
}

// Future possibilities
// For flexible URL:
// "github.com/sinalalebakhsh/Gocron/GetUserInput"
// getuserinput.GetUserInput()
