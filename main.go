package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/sinalalebakhsh/Gocron/GetUserInput"
	"github.com/sinalalebakhsh/Gocron/features"

	"github.com/fatih/color"
)

func main()  {


	if os.Args[1] == strings.ToLower("-h") || os.Args[1] == strings.ToLower("help") {
		fmt.Println(`
Gocron Project:
is a API for learning GO language with example.
my name is Sina LalehBakhsh, I hope this API is useful for you
after running program, write your single word about any of GO language.
if your perpuse is more than one word, for convenience searching, just write keywords.
like this:
	map slice `)
	} else if os.Args[1] == strings.ToLower("all") || 
	os.Args[1] == strings.ToLower("everyexample") || 
	os.Args[1] == strings.ToLower("examples") {
		color.Cyan(fmt.Sprintln(features.OriginalFeatures))
	}





	UserInput , err := getuserinput.GetUserInput()
	if err != nil {
		fmt.Println(err)
	}
	color.Blue(fmt.Sprintln(UserInput))
		

	
}