package getuserinput

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/fatih/color"
	"github.com/sinalalebakhsh/Gocron/features"
	"os/exec"
	"runtime"
	// "github.com/sinalalebakhsh/Gocron/Server"
)

// after ensure user dont wrote arg more than one like -h or --help
// get user input in this condition
func GetUserInput() {

	fmt.Println("Would you like to join the gocron project on GitHub? (yes/no)")
	FirstCount := true

	// I use For loop for getting user input for Repeatedly.
	// if you want just get one more time delete this loop
	for {
		// use bufio Package and os Package for reading and get clear, good and clean,
		// user input. maybe you want use fmt.Scan or another ways
		// You can do it. :)
		UserInput := bufio.NewReader(os.Stdin)
		FinalInput, _ := UserInput.ReadString('\n')
		FinalInput = strings.TrimSuffix(FinalInput, "\n")
		FinalInput = strings.TrimSpace(FinalInput)
		SliceOfWords := strings.Split(FinalInput, " ")

		// Checks for each entry into the loop
		// if user input was "exit" so break loop and log out
		// from Gocron program.
		Regular := true

		Regular = IfUsrisEXIT(FinalInput) // If FinalInput is "exit" return os.Exit(0)

		if Regular {
			Regular = IfUserisHELP(FinalInput) // If FinalInput is or does't "help" return true
		}

		if Regular {
			Regular = IfUserisCLEAR(FinalInput)
		}


		if Regular {
			if FirstCount {
				Regular = GetUserForBrowser(FinalInput)	
				FirstCount = false
			}
		}

		if Regular {
			Regular = IfUserWantSingleDefinition(FinalInput)
		}

		if Regular {
			if len(SliceOfWords) >= 2 {
				FirstInput, SecondInput := SliceOfWords[0], SliceOfWords[1]
				Regular = IfUsris2orMoreWords(FirstInput, SecondInput)
			}
		}

		if Regular {
			Regular = IfUsrisALL(FinalInput)
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
		features.ClearTerminal()
		os.Exit(0)
	}
	return true
}

func IfUserisHELP(FinalInput string) bool {
	Help := "help"
	if FinalInput == strings.ToLower(Help) {
		features.HelpMessage()
		return false
	}
	return true
}

func IfUserisCLEAR(FinalInput string) bool {
	if FinalInput == strings.ToLower("clear") {
		features.ClearTerminal()
		return false
	}
	return true
}

func GetUserForBrowser(userInput string) bool {
	if strings.ToLower(userInput) == "yes" {
		openBrowser()
		// server.MyServer()
		return false
	} else {
		fmt.Println("Ok. Please search for the topic you are looking for with keywords:")
		return false
	}
}

func openBrowser() {
	var command string

	switch runtime.GOOS {
	case "darwin":
		command = "open"
	case "windows":
		command = "start"
	case "linux":
		command = "xdg-open"
	default:
		fmt.Println("Unsupported platform.")
		return
	}

	exec.Command(command, "https://github.com/sinalalebakhsh/Gocron").Start()
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
			return false
		}
	}
	return true
}

func IfUserWantSingleDefinition(FinalInput string) bool {
	keys := features.GetMapReturnSlice()
	for _, Value := range keys {
		if FinalInput == Value {
			words := features.SplitIntoWords(features.OriginSingleDef.SingleDef[Value])
			features.PrintWordByWord(words)
			// fmt.Println(features.OriginSingleDef.SingleDef[Value])
			fmt.Println()
			return false
		}
	}
	return true
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
				return false
			}
		}
	}
	return true
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
	// ===============================================
	// Creating HTTP Servers
	// searching for on slice = "ALL HTTP SERVERS", "ALLHTTPSERVERS", "ALL CREATING HTTP SERVERS", "ALLCREATINGHTTPSERVERS",
	if result := searchSlice(FinalInput, features.TitleHTTPServers, features.OriginalHTTPServers); result.Alldatafield != "" {
		MyChann1 <- result
		return
	}

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

func PrintNotAddYet(FinalInput string) {
	color.HiRed(fmt.Sprintf("=================================\nNot add %q yet.\n=================================", FinalInput))
}
