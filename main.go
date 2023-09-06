package main

import (
	"fmt"

	"github.com/sinalalebakhsh/Gocron/features"
	"github.com/sinalalebakhsh/Gocron/GetUserInput"

	"github.com/fatih/color"
)

func main()  {

	color.Cyan(fmt.Sprintln(features.OriginalFeatures))

	

	UserInput , err := getuserinput.GetUserInput()
	if err != nil {
		fmt.Println(err)
	}
	color.Blue(fmt.Sprintln(UserInput))
		

	
}