package getuserinput

import (
	"fmt"
	"os"
	"strings"
	// "time"
	"github.com/fatih/color"
	"github.com/sinalalebakhsh/Gocron/features"
)

func GetFirstArg() {

	if len(os.Args) == 1 {
		color.HiBlue(features.LOGO)
	}else if len(os.Args) > 1 {
		AllArgs := os.Args

		go func(){
			for _, value := range features.TitleOfHelp {
				FirstArg := strings.ToLower(AllArgs[1])
				if value == FirstArg {
					features.HelpMessage()
				}
			}
		}()
		

	
		go func(){
			for _, value := range features.TitleOfEveryWords {
				FirstArg := strings.ToLower(AllArgs[1])
				if value == FirstArg {
						features.HelpMessage()
						color.HiBlue(fmt.Sprintln(features.OriginalFeatures))			
				}
			}
		}()
	}
}

