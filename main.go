package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/sinalalebakhsh/Gocron/GetUserInput"
	"github.com/sinalalebakhsh/Gocron/features"

	"github.com/fatih/color"
)

func main()  {


	HelpFlag := flag.String("-h", "help", "--help")
	HelpFlagPlus := flag.String("-help", "-HELP", "--HELP")
	HelpFlagPlusPlus := flag.String("HELP", "--h", "helpme!")

	var SliceOfFlags []string 
	SliceOfFlags = append(SliceOfFlags, *HelpFlag)
	SliceOfFlags = append(SliceOfFlags, *HelpFlagPlus)
	SliceOfFlags = append(SliceOfFlags, *HelpFlagPlusPlus)

	for _, Value := range SliceOfFlags {
		fmt.Println(Value)
	}

	HelpMessage := `
Gocron Project:
-------------------------------------------------------
is a API for learning GO language with example.
-------------------------------------------------------
my name is Sina LalehBakhsh, I hope this API is useful for you
after running program, write your single word about any of GO language.
if your perpuse is more than one word, for convenience searching, just write keywords.
like this:
	map slice
-------------------------------------------------------`

	// if AllArgs[1] == strings.ToLower("-h") || AllArgs[1] == strings.ToLower("help") {
	// 	fmt.Println()
	// } else if AllArgs[1] == strings.ToLower("all") || 
	// AllArgs[1] == strings.ToLower("everyexample") || 
	// AllArgs[1] == strings.ToLower("examples") {
	// 	} 
		
	color.Cyan(fmt.Sprintln(features.OriginalFeatures))




	UserInput , err := getuserinput.GetUserInput()
	if err != nil {
		fmt.Println(err)
	}
	color.Blue(fmt.Sprintln(UserInput))
		

	
}