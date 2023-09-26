package getuserinput

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/sinalalebakhsh/Gocron/features"
	"os"
	// "sort"
	"strings"
)

// after ensure user dont wrote arg more than one like -h or --help
// get user input in this condition
func GetUserInput() {

	// I use For loop for getting user input for Repeatedly.
	// if you want just get one more time delete this loop
	for {

		// use bufio Package and os Package for reading and get clear, good and clean,
		// user input. maybe you want use fmt.Scan or another ways
		// You can do it. :)
		UserInput := bufio.NewReader(os.Stdin)
		FinalInput, _ := UserInput.ReadString('\n')
		FinalInput = strings.TrimSuffix(FinalInput, "\n")

		FinalInput = strings.ToLower(FinalInput)
		FinalInput = strings.TrimSpace(FinalInput)
		FinalInputTrimSpaced := strings.TrimSpace(FinalInput)

		
		SliceOfWords := strings.Split(FinalInput, " ")

		// Checks for each entry into the loop
		// if user input was "exit" so break loop and log out
		// from Gocron program.
		Exit := "exit"
		Help := "help"
		if FinalInput == strings.ToLower(Exit) {
			features.GoodByePrint()
			os.Exit(0)
		} else if FinalInput == strings.ToLower(Help) {
			features.HelpMessage()
		} else if len(SliceOfWords) == 2 {
			FirstInput, SecondInput := SliceOfWords[0], SliceOfWords[1]
			
			for Index := range features.OriginalSingleDefExamples.MapSingleDefEx {
				if FirstInput == Index && SecondInput == "example" {
					words := features.SplitIntoWords(features.OriginSingleDef.SingleDefinition[FirstInput])
					features.PrintWordByWord(words)
					fmt.Println()
					color.HiMagenta(fmt.Sprintln("============================================◉🧭🧭🧭🧭🧭🧭🧭◉=========================================="))
					color.HiMagenta(fmt.Sprintln(features.OriginalSingleDefExamples.MapSingleDefEx[Index]))
					color.HiMagenta(fmt.Sprintln("============================================◉⭐⭐⭐⭐⭐⭐⭐◉=========================================="))
				}
			}
		} 
	
	
		for _, Value := range features.TitleOfAllIndexSlices {
			FinalInputTrimSpaced = strings.ToUpper(FinalInputTrimSpaced)
			if FinalInputTrimSpaced == Value {
				MyChann1 := make(chan features.DataBase)
				go GetUserInputSendChannel(FinalInputTrimSpaced, MyChann1)

				result := <-MyChann1
				if result.Alldatafield != "" {
					color.HiBlue(fmt.Sprintln("============================================◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉=========================================="))
					color.HiBlue(fmt.Sprintln(result))
					color.HiBlue(fmt.Sprintln("============================================◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉=========================================="))
				} else  {
					PrintNotAddYet(FinalInputTrimSpaced)
				}
			}
		}

	
	}

}

func PrintNotAddYet(FinalInput string) {
	color.HiRed(fmt.Sprintln("================================="))
	color.HiRed(fmt.Sprintf("Not add %v yet.", FinalInput))
	color.HiRed(fmt.Sprintln("================================="))
}

func GetUserInputSendChannel(FinalInput string, MyChann1 chan<- features.DataBase) {

	// ===============================================
	// RegEx
	// for searching on slice = "ALL REGEX", "ALLREGEX",
	if result := searchSlice(FinalInput, features.TitleOfRegEx, features.OriginalAllRegex); result.Alldatafield != "" {
		MyChann1 <- result
		return
	}

	// // ===============================================
	// // Time
	// for searching on slice = "ALL TIME", "ALLTIME",
	if result := searchSlice(FinalInput, features.TitleOfTimeData, features.OriginalTimeData); result.Alldatafield != "" {
		MyChann1 <- result
		return
	}
	
	
	// // ===============================================
	// // Read & Writing
	// go func() {
	// 	for _, value := range features.TitleOfReadingWriting {
	// 		FinalInput = strings.ToUpper(FinalInput)
	// 		if FinalInput == value {
	// 			color.HiBlue(fmt.Sprintln("---------------------------------------------------------------"))
	// 			color.HiBlue(fmt.Sprintln(features.OriginalReadingandWriting))
	// 			color.HiBlue(fmt.Sprintln("---------------------------------------------------------------"))
	// 			break
	// 		}
	// 	}
	// }()
	// // ===============================================
	// // Working with JSON Files
	// go func() {
	// 	for _, Value := range features.TitleOfJSON {
	// 		FinalInput = strings.ToUpper(FinalInput)
	// 		if FinalInput == Value {
	// 			color.HiBlue(fmt.Sprintln("---------------------------------------------------------------"))
	// 			color.HiBlue(fmt.Sprintln(features.OriginalJSONData))
	// 			color.HiBlue(fmt.Sprintln("---------------------------------------------------------------"))
	// 			break
	// 		}
	// 	}
	// }()
	// // ===============================================
	// // Working with Files
	// go func() {
	// 	for _, value := range features.TitleOfWorkingFiles {
	// 		FinalInput = strings.ToUpper(FinalInput)
	// 		if FinalInput == value {
	// 			color.HiBlue(fmt.Sprintln("---------------------------------------------------------------"))
	// 			color.HiBlue(fmt.Sprintln(features.OriginalWorkWithFiles))
	// 			color.HiBlue(fmt.Sprintln("---------------------------------------------------------------"))
	// 			break
	// 		}
	// 	}
	// }()
	// // ===============================================
	// // HTML & Text Templates
	// go func() {
	// 	for _, value := range features.TitleOfUsingHTMLAndTextTemplates {
	// 		FinalInput = strings.ToUpper(FinalInput)
	// 		if FinalInput == value {
	// 			color.HiBlue(fmt.Sprintln("---------------------------------------------------------------"))
	// 			color.HiBlue(fmt.Sprintln(features.OriginalHTMLAndTemplates))
	// 			color.HiBlue(fmt.Sprintln("---------------------------------------------------------------"))
	// 			break
	// 		}
	// 	}
	// }()

	MyChann1 <- features.DataBase{} // Return an empty struct if not found

}

func searchSlice(input string, slice []string, obj features.DataBase) features.DataBase {
	for _, item := range slice {
		if strings.Contains(item, input) {
			return obj
		}
	}
	return features.DataBase{}
}
