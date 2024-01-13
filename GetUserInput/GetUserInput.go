// This is first
package getuserinput

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"github.com/fatih/color"
	"github.com/sinalalebakhsh/Gocron/Crawl"
	"github.com/sinalalebakhsh/Gocron/features"

	// "github.com/sinalalebakhsh/Gocron/Server"
	"math/rand"
)

// After ensure user dont wrote arg more than one like -h or --help
// get user input in this condition
func GetUserInput(GetBoleanFromGetFirstArgFunction bool) {

	// If user wrote another argument after ./Gocron for running program,
	// Like -help, so this here below will not print. 
	if GetBoleanFromGetFirstArgFunction {
		fmt.Println("💠💠💠 Would you like to join the gocron project on GitHub? (yes/no)")
	}

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
		
		SaveUserInputInFileTXT(FinalInput)


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
			Regular = IfUserInputIsJustNumber(FinalInput)
		}

		if Regular {
			Regular = IfUserWantSingleDefinition(FinalInput)
		}

		if len(SliceOfWords) >= 2 {
			// FirstInput, SecondInput := SliceOfWords[0], SliceOfWords[1]
			if Regular {
				Regular = IfUsris2orMoreWords(SliceOfWords)
			}

			if Regular {
				Regular = IfUserInputIsCrawlURL(SliceOfWords)
			}
		}
		
		
		if Regular {
			Regular = IfUsrisALL(FinalInput)
		}

		if Regular {
			Regular = IfUserInputIsEmpty(FinalInput)
		}

		if Regular {
			PrintNotAddYet(FinalInput)
		}

	}
}

func IfUsrisEXIT(FinalInput string) bool {
	Exit := "exit"
	if Exit == strings.ToLower(FinalInput) {
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
		fmt.Println("Ok. Please search for the topic you are looking for with keywords:")
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

func IfUserInputIsJustNumber(userInput string) bool {
	for Index := range features.OriginalWordsGoNum.MapOfNumbered {
		if userInput == Index {
			color.HiMagenta(fmt.Sprintln("============================================◉💠💠💠💠💠💠💠◉=========================================="))
			color.HiMagenta(fmt.Sprintln(features.OriginalWordsGoNum.MapOfNumbered[Index]))
			color.HiMagenta(fmt.Sprintln("============================================◉💠💠💠💠💠💠💠◉=========================================="))
			return false
		} 
	}
	return true
}

func IfUsris2orMoreWords(SliceOfWords []string) bool {
	// Initialize an empty string
	result := ""

	// Iterate through the slice excluding the last element
	for i := 0; i < len(SliceOfWords)-1; i++ {
		result += SliceOfWords[i]

		// Add a separator (comma and space) if it's not the last element
		if i < len(SliceOfWords)-2 {
			result += " "
		}
	}

	// Get the last index
	lastIndex := len(SliceOfWords) - 1
	// Access the last element
	lastElement := SliceOfWords[lastIndex]

	for Index := range features.OriginalSingleDefExamples.MapSingleDefEx {
		if result == Index && lastElement == "example" {
			words := features.SplitIntoWords(features.OriginSingleDef.SingleDef[result])
			features.PrintWordByWord(words)
			fmt.Println()
			color.HiMagenta(fmt.Sprintln("============================================◉🧭🧭🧭🧭🧭🧭🧭◉=========================================="))
			color.HiMagenta(fmt.Sprintln(features.OriginalSingleDefExamples.MapSingleDefEx[Index]))
			color.HiMagenta(fmt.Sprintln("============================================◉⭐⭐⭐⭐⭐⭐⭐◉=========================================="))
			return false
		}
	}

	for Index := range features.QuestionsSample.MapQuestionsSample {
		if result == Index && lastElement == "question" {
			words := features.SplitIntoWords(features.Questions.SingleQues[result])
			features.PrintWordByWord(words)
			fmt.Println()
			color.HiYellow(fmt.Sprintln("============================================◉❓❓❓❓❓❓❓◉=========================================="))
			color.HiRed(fmt.Sprintln(features.QuestionsSample.MapQuestionsSample[Index]))
			color.HiYellow(fmt.Sprintln("============================================◉🔴🔴🔴🔴🔴🔴🔴◉=========================================="))
			return false
		}
	}

	for Index := range features.AnswersSample.MapAnswersSample {
		if result == Index && lastElement == "answer" {
			words := features.SplitIntoWords(features.Answers.SingleAnws[result])
			features.PrintWordByWord(words)
			fmt.Println()
			color.HiGreen(fmt.Sprintln("============================================◉🔰🔰🔰🔰🔰🔰🔰◉=========================================="))
			color.HiCyan(fmt.Sprintln(features.AnswersSample.MapAnswersSample[Index]))
			color.HiGreen(fmt.Sprintln("============================================◉✅✅✅✅✅✅✅◉=========================================="))
			return false
		}
	}

	for Index := range features.MapUsage.SingleDef {
		if result == Index && lastElement == "usage" {
			color.HiCyan(fmt.Sprintln("============================================◉🔰🔰🔰🔰🔰🔰🔰◉=========================================="))
			color.HiCyan(fmt.Sprintln(features.MapUsage.SingleDef[Index]))
			color.HiCyan(fmt.Sprintln("============================================◉✅✅✅✅✅✅✅◉=========================================="))
			return false
		}
	}

	for Index := range features.OriginalSingleDefFunctions.MapSingleDefFuncs {
		result = strings.ToLower(result)
		if result == Index && strings.ToLower(lastElement) == "function" {
			words := features.SplitIntoWords(features.OriginalSingleDefFunctions.MapSingleDefFuncs[result])
			features.PrintWordByWord(words)
			fmt.Println()
			return false
		}
	}

	return true
}

func IfUserWantSingleDefinition(FinalInput string) bool {
	keys := features.GetMapReturnSlice()
	for _, Value := range keys {
		FinalInput = strings.ToLower(FinalInput)
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
	// ===============================================
	// Creating HTTP Clients
	// searching for on slice = "ALL HTTP CLIENTS", "ALLHTTPCLIENTS", "ALL CREATING HTTP CLIENTS", "ALLCREATINGHTTPCLIENTS",
	if result := searchSlice(FinalInput, features.TitleHTTPClients, features.OriginalHTTPClients); result.Alldatafield != "" {
		MyChann1 <- result
		return
	}
	// ===============================================
	// WORKING WITH DATABASES
	// searching for on slice =	"ALL WORKING WITH DATABASES","ALLWORKINGWITHDATABASES","ALL DATABASES","ALLDATABASES",
	if result := searchSlice(FinalInput, features.TitleDataBases, features.OriginalDataBases); result.Alldatafield != "" {
		MyChann1 <- result
		return
	}
	// ===============================================
	// ALL USING REFLECTION
	// 	searching for on slice = "ALLUSINGREFLECTION", "ALL USING REFLECTION",
	if result := searchSlice(FinalInput, features.TitleUsingReflection, features.OriginalUsingReflection); result.Alldatafield != "" {
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

func IfUserInputIsEmpty(FinalInput string) bool {
	if FinalInput == "" {
		// Example slice of strings
		mySlice := []string{
			"I am Gocron!! and I don't have feature for empty!!! say something!!",
			"I am Gocron!! and I don't have feature for empty!!! say something!!",
			"I tell to Sina Lalehbakhsh!!!",
			"I am Gocron!! and I don't have feature for empty!!! say something!!",
			"I tell to Sina Lalehbakhsh!!!",
			"I tell to Sina Lalehbakhsh!!!",
			"I am Gocron!! and I don't have feature for empty!!! say something!!",
			"I tell to Sina Lalehbakhsh!!!",
			"I tell to Sina Lalehbakhsh!!!",
			"I tell to Sina Lalehbakhsh!!!",
			"I am Gocron!! and I don't have feature for empty!!! say something!!",
			"What are you doing!? with empty!",
			"empty search !",
			"empty search !",
			"I am Gocron!! and I don't have feature for empty!!! say something!!",
			"empty search !",
			"empty search !",
			"empty search !",
			"are you kidding me؟!", 
			"say something.", 
			"you think that faunny!?", 
			"I am Gocron!! and I don't have feature for empty!!! say something!!",
			"You entered blank!", 
			"Go language has no definition for empty space!",
			"Do you think the empty space has meaning?!",
			"Seriously, are you joking?!",
			"I am Gocron!! and I don't have feature for empty!!! say something!!",
			"Speak up, please.",
			"Do you find that amusing!?",
			"You just submitted nothing!",
			"empty search !",
			"empty search !",
			"I am Gocron!! and I don't have feature for empty!!! say something!!",
			"empty search !",
			"In Go language, an empty space is undefined!",
			"Contemplate: does an empty space convey meaning to you?!",
			"Are you pulling my leg right now?!",
			"Express yourself, don't stay silent.",
			"Humor me, do you really think that's funny!?",
			"I am Gocron!! and I don't have feature for empty!!! say something!!",
			"A void? You just presented me with emptiness!",
			"In the realm of Go language, emptiness lacks definition!",
			"Does the absence of content hold significance to you?",
			"empty search !",
			"empty search !",
			"I am Gocron!! and I don't have feature for empty!!! say something!!",
			"empty search !",
			"You jest, right? There's nothing here!",
			"Your input is as vacant as an open field in Go!",
			"Is the concept of nothingness profound to you?",
			"I am Gocron!! and I don't have feature for empty!!! say something!!",
			"Seriously, an empty space? That's your contribution?",
			"In Go, silence speaks louder than an undefined space!",
			"Time will tell.",
			"Dream big dreams.",
			"Chase your passion.",
			"I am Gocron!! and I don't have feature for empty!!! say something!!",
			"Love conquers all.",
			"Silence speaks volumes.",
			"Explore new horizons.",
			"Learn from mistakes.",
			"Embrace the unexpected.",
			"Create lasting memories.",
			"I am Gocron!! and I don't have feature for empty!!! say something!!",
			"Find inner strength.",
			"empty search !",
			"empty search !",
			"empty search !",
			"I am Gocron!! and I don't have feature for empty!!! say something!!",
			"empty search !",
			"empty search !",
			"I am Gocron!! and I don't have feature for empty!!! say something!!",
			"empty search !",
		}

		// Get a random element from the slice
		randomElement, err := getRandomElement(mySlice)
		if err != nil {
			fmt.Println(err)
		} else {
			color.HiRed(fmt.Sprintf("=================================\n%s\n=================================", randomElement))
		}
		return false
	}
	return true
}
func getRandomElement(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", fmt.Errorf("slice is empty")
	}
	randomIndex := rand.Intn(len(slice))
	return slice[randomIndex], nil
}

const SearchLogFile = "SearchLogs/user_input.txt"
func SaveUserInputInFileTXT(FinalInput string) {
	// Check if the file exists
	if _, err := os.Stat(SearchLogFile); os.IsNotExist(err) {
		// File doesn't exist, create an empty one
		if err := os.WriteFile(SearchLogFile, []byte(""), 0644); err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
	}

	// Read existing content from the file
	content, err := os.ReadFile(SearchLogFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert file content to lowercase for case-insensitive comparison
	fileContent := strings.ToLower(string(content))


	// Convert user input to lowercase for case-insensitive comparison
	FinalInput = strings.ToLower(FinalInput)

	// Check if the user input exists in the file
	if strings.Contains(fileContent, FinalInput) {
		// fmt.Println("This search was done before.")
	} else {
		// Save the user input to the file
		if err := SaveToFile(FinalInput); err != nil {
			// fmt.Println("Error saving to file:", err)
			return
		}
		// fmt.Println("Sentence saved to the file.")
	}
}


func SaveToFile(sentence string) error {
	// Open the file in append mode
	file, err := os.OpenFile(SearchLogFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the sentence to the file
	if _, err := file.WriteString(sentence + "\n"); err != nil {
		return err
	}

	return nil
}


func IfUserInputIsCrawlURL(SliceOfWords[]string) bool {
	const CRAWL = "crawl"
	var URL string
	if CRAWL == strings.ToLower(SliceOfWords[0]) {
		if SliceOfWords[1] != "" {
			URL = SliceOfWords[1]
			depthStr := SliceOfWords[2]
			depth, err := strconv.Atoi(depthStr)
			if err != nil {
				fmt.Println("Error converting depth to integer:", err) 
			} else {
				crawl.Crawl(URL, depth)
				return false
			}
			return false
		}
		return false
	}
	return true
}