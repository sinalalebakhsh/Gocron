package getuserinput

import (
	// "fmt"
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
		
	
		// } else if FirstArg == "all" || 
		// FirstArg == "-all" || 
		// FirstArg == "--all" || 
		// FirstArg == "everything" || 
		// FirstArg == "fullexplaination" || 
		// FirstArg == "explain" || 
		// FirstArg == "example" || 
		// FirstArg == "allexample" || 
		// FirstArg == "allcodes" || 
		// FirstArg == "allcode" || 
		// FirstArg == "everyexample" || 
		// FirstArg == "examples" {
		// 	features.HelpMessage()
		// 	time.Sleep(time.Second * 2)
		// 	color.Yellow(fmt.Sprintln(features.OriginalFeatures))			
		// } 
	}
}

