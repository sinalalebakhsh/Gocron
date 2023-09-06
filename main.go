package main

import (
	"fmt"

	"github.com/sinalalebakhsh/Gocron/features"
	"github.com/sinalalebakhsh/Gocron/GetUserInput"

	"github.com/fatih/color"
)

func main()  {

	color.Cyan(fmt.Sprintln(features.OriginalFeatures))

	UserInput := getuserinput.GetUserInput()
	
	

	for _, Value := range features.OriginSingleDef.SingleDefinition {
		if UserInput == Value {
			fmt.Println(Value)
		}
	} 
	// color.Blue(fmt.Sprintln(features.OriginSingleDef))


	
}