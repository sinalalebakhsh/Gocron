package getuserinput

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/sinalalebakhsh/Gocron/features"
	"net/http"
)

func GetFirstArg() {

	if len(os.Args) > 1 {
		AllArgs := os.Args
		FirstArg := strings.ToLower(AllArgs[1])
		if FirstArg == "-h" || 
		FirstArg == "help" || 
		FirstArg == "-help" ||
		FirstArg == "--help" {
		features.HelpMessage()
		
		} else if FirstArg == "all" || 
		FirstArg == "-all" || 
		FirstArg == "--all" || 
		FirstArg == "everything" || 
		FirstArg == "fullexplaination" || 
		FirstArg == "explain" || 
		FirstArg == "example" || 
		FirstArg == "allexample" || 
		FirstArg == "allcodes" || 
		FirstArg == "allcode" || 
		FirstArg == "everyexample" || 
		FirstArg == "examples" {
			time.Sleep(time.Second * 2)		
			fmt.Sprint(features.HelpMessage(),features.OriginalFeatures)
		} 
	} else {
		errors.New("this Argument not add yet")
	} 


}

