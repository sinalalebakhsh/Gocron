package getuserinput

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/sinalalebakhsh/Gocron/features"
)

func GetFirstArg() bool {

	if len(os.Args) == 1 {
		color.HiBlue(features.LOGO)
		return true
	} else if len(os.Args) > 1 {
		AllArgs := os.Args

		for _, value := range features.TitleOfHelp {
			FirstArg := strings.ToLower(AllArgs[1])
			if value == FirstArg {
				features.HelpMessage()
				return false
			}
		}

		for _, value := range features.TitleOfEveryWords {
			FirstArg := strings.ToLower(AllArgs[1])
			if value == FirstArg {
				features.HelpMessage()
				color.HiBlue(fmt.Sprintln(features.OriginalFeatures))
			}
		}

	}
	return true
}
