package getuserinput

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/sinalalebakhsh/Gocron/features"
	"os"
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
		SliceOfWords := strings.Split(FinalInput, " ")

		// Checks for each entry into the loop
		// if user input was "exit" so break loop and log out
		// from Gocron program.
		Regular := true

		Regular = IfUsrisEXIT(FinalInput) // If FinalInput is "exit" return os.Exit(0)

		IfUserisHELP(FinalInput) // If FinalInput is or does't "help" return true

		if len(SliceOfWords) == 2 {
			FirstInput, SecondInput := SliceOfWords[0], SliceOfWords[1]
			Regular = IfUsris2orMoreWords(FirstInput, SecondInput)
		}

		if Regular {
			Regular = IfUserWantSingleDefinition(FinalInput)
		}

		if Regular {
			PrintNotAddYet(FinalInput)
		}

	}
}

func IfUsrisEXIT(FinalInput string) bool {
	Exit := "exit"
	if FinalInput == strings.ToLower(Exit) {
		features.GoodByePrint()
		os.Exit(0)
	}
	return true
}

func IfUserisHELP(FinalInput string) {
	Help := "help"
	if FinalInput == strings.ToLower(Help) {
		features.HelpMessage()
	}
}

func IfUsris2orMoreWords(FirstInput, SecondInput string) bool {
	for Index := range features.OriginalSingleDefExamples.MapSingleDefEx {
		if FirstInput == Index && SecondInput == "example" {
			words := features.SplitIntoWords(features.OriginSingleDef.SingleDef[FirstInput])
			features.PrintWordByWord(words)
			fmt.Println()
			color.HiMagenta(fmt.Sprintln("============================================◉🧭🧭🧭🧭🧭🧭🧭◉=========================================="))
			color.HiMagenta(fmt.Sprintln(features.OriginalSingleDefExamples.MapSingleDefEx[Index]))
			color.HiMagenta(fmt.Sprintln("============================================◉⭐⭐⭐⭐⭐⭐⭐◉=========================================="))
			return true
		}
	}
	return false
}

func IfUserWantSingleDefinition(FinalInput string) bool {
	for index := range features.OriginSingleDef.SingleDef {
		if FinalInput == index {
			words := features.SplitIntoWords(features.OriginSingleDef.SingleDef[FinalInput])
			features.PrintWordByWord(words)
			fmt.Println()
			return true
		}
	}
	return false
}

func IfUsrisALL(FinalInput string) bool {
	for _, Value := range features.TitleOfAllIndexSlices {
		FinalInput = strings.ToUpper(FinalInput)
		if FinalInput == Value {
			MyChann1 := make(chan features.DataBase)
			go GetUserInputSendChannel(FinalInput, MyChann1)

			result := <-MyChann1
			if result.Alldatafield != "" {
				color.HiBlue(fmt.Sprintln("============================================◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉=========================================="))
				color.HiBlue(fmt.Sprintln(result))
				color.HiBlue(fmt.Sprintln("============================================◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉=========================================="))

			}
		}
	}
	return false
}

func GetUserInputSendChannel(FinalInput string, MyChann1 chan<- features.DataBase) {
	// ===============================================
	// RegEx
	// for searching on slice = "ALL REGEX", "ALLREGEX",
	if result := searchSlice(FinalInput, features.TitleOfRegEx, features.OriginalAllRegex); result.Alldatafield != "" {
		MyChann1 <- result
		return
	}
	// ===============================================
	// Time
	// searching for on slice = "ALL TIME", "ALLTIME",
	if result := searchSlice(FinalInput, features.TitleOfTimeData, features.OriginalTimeData); result.Alldatafield != "" {
		MyChann1 <- result
		return
	}
	// ===============================================
	// Read & Writing
	// searching for on slice = "READING AND WRITING DATA", "READINGANDWRITINGDATA", "ALL READING AND WRITING DATA", "ALLREADINGANDWRITINGDATA",
	if result := searchSlice(FinalInput, features.TitleOfReadingWriting, features.OriginalReadingandWriting); result.Alldatafield != "" {
		MyChann1 <- result
		return
	}
	// ===============================================
	// Working with JSON Files
	// searching for on slice = "ALL JSON", "ALLJSON", "ALL JSON DATA", "ALLJSONDATA", "ALL WORK WITH JSON DATA", "ALLWORKWITHJSONDATA", "ALL WORKING WITH JSON DATA", "ALLWORKINGWITHJSONDATA",
	if result := searchSlice(FinalInput, features.TitleOfJSON, features.OriginalJSONData); result.Alldatafield != "" {
		MyChann1 <- result
		return
	}
	// ===============================================
	// HTML & Text Templates
	// searching for on slice = "ALL HTML AND TEMPLATE", "ALLHTMLANDTEMPLATE",
	if result := searchSlice(FinalInput, features.TitleOfUsingHTMLAndTextTemplates, features.OriginalHTMLAndTemplates); result.Alldatafield != "" {
		MyChann1 <- result
		return
	}
	// ===============================================
	//  Working with Files
	// searching for on slice = "ALL WORKING WITH FILES", "ALLWORKINGWITHFILES",
	if result := searchSlice(FinalInput, features.TitleOfWorkingFiles, features.OriginalWorkWithFiles); result.Alldatafield != "" {
		MyChann1 <- result
		return
	}

	MyChann1 <- features.DataBase{} // Return an empty struct if not found

}

func PrintNotAddYet(FinalInput string) {
	color.HiRed(fmt.Sprintf("=================================\nNot add %v yet.\n=================================", FinalInput))
}

func searchSlice(input string, slice []string, obj features.DataBase) features.DataBase {
	for _, item := range slice {
		if strings.Contains(item, input) {
			return obj
		}
	}
	return features.DataBase{}
}
