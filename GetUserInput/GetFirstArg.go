package getuserinput

import (
	"fmt"
	"os"
	"strings"
	"github.com/fatih/color"
	"github.com/sinalalebakhsh/Gocron/features"
)

func GetFirstArg() {

	if len(os.Args) > 1 {
		AllArgs := os.Args
		FirstArg := strings.ToLower(AllArgs[1])
		if FirstArg == "-h" || 
		FirstArg == "help" || 
		FirstArg == "-help" ||
		FirstArg == strings.ToLower("--help") {
			features.HelpMessage()
		} else if FirstArg == strings.ToLower("all") || 
		FirstArg == strings.ToLower("everyexample") || 
		FirstArg == strings.ToLower("examples") {

			color.Yellow(fmt.Sprintln(features.OriginalFeatures))			
		} 
	}
}

