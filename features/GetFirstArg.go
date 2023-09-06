package features

import (
	"fmt"
	"os"
	"strings"
	"github.com/fatih/color"
	"github.com/sinalalebakhsh/Gocron/features/helpMessage"

)

func GetFirstArg() {

	if len(os.Args) > 1 {
		AllArgs := os.Args
		FirstArg := strings.ToLower(AllArgs[1])
		if FirstArg == "-h" || 
		FirstArg == "help" || 
		FirstArg == strings.ToLower("-HELP") ||
		FirstArg == strings.ToLower("--HELP") {
			features.
		} else if FirstArg == strings.ToLower("all") || 
		FirstArg == strings.ToLower("everyexample") || 
		FirstArg == strings.ToLower("examples") {
			color.Cyan(fmt.Sprintln(features.OriginalFeatures))			
		} 
	}
}

