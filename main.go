package main

import (
	"fmt"
	"github.com/sinalalebakhsh/Gocron/GetUserInput"
	"github.com/fatih/color"
)

func main()  {


	getuserinput.GetFirstArg()




	UserInput , err := getuserinput.GetUserInput()
	if err != nil {
		fmt.Println(err)
	}
	color.Cyan(fmt.Sprintln(UserInput))
		

	
}