package main

import (
	"fmt"

	"github.com/sinalalebakhsh/Gocron/features"

	"github.com/fatih/color"
)

func main()  {

	color.Cyan(fmt.Sprintln(features.OriginalFeatures))

	for _, Value := range features.OriginSingleDef.SingleDefinition {
		fmt.Println(Value)
	} 
	// color.Blue(fmt.Sprintln(features.OriginSingleDef))


	
}