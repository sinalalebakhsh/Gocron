package server

import (
	"fmt"
	"github.com/sinalalebakhsh/Gocron/features"
	"github.com/sinalalebakhsh/Gocron/GetUserInput"
	"net/http"
	"os"
	"strings"
)

func HandlerAllFeatures(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, features.OriginalFeatures)
}

func HandletSingleDefinition(w http.ResponseWriter, r *http.Request){
	firstArg := fmt.Sprint(getuserinput.GetUserInput())

	fmt.Fprint(w, firstArg)
}

func Servers() {
	if len(os.Args) == 1 {
		HOME := "/"
		http.HandleFunc(HOME, HandlerAllFeatures)
	} else if len(os.Args) > 1 {
		HOME := "/"
		go http.HandleFunc(HOME, HandlerAllFeatures)
		AllArgs := os.Args
		FirstArg := strings.ToLower(AllArgs[1])
		HELP := FirstArg
		go http.HandleFunc(HELP, HandletSingleDefinition)
	}
	http.ListenAndServe(":8080", nil)
}

// Future possibilities
// For flexible URL:
// "github.com/sinalalebakhsh/Gocron/GetUserInput"
// getuserinput.GetUserInput()
