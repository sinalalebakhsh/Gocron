package getuserinput

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"github.com/fatih/color"
	"github.com/sinalalebakhsh/Gocron/features"
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

		// get length of map from Single Definition in side file here
		SliceOfMap := make([]string, 0, len(features.OriginSingleDef.SingleDefinition))

		// in here processing on the map for getting keys and ignore values 
		for key, error1 := range features.OriginSingleDef.SingleDefinition {
			if error1 == "nil" {
				fmt.Println(error1, "!!!!!!")
			}
			// converting map keys to slice for iteration
			SliceOfMap = append(SliceOfMap, key)
		}
		// sort for iteration convenience 
		sort.Strings(SliceOfMap)
	
		
		go func(){
			for _, value := range SliceOfMap {
				FinalInput = strings.ToLower(FinalInput)
				if FinalInput == value {
					words := features.SplitIntoWords(features.OriginSingleDef.SingleDefinition[value])
					features.PrintWordByWord(words)
					fmt.Println()
					// Regulators = false
					break
				}
			}
			FinalInput = strings.ToLower(FinalInput)
			FinalInput = strings.TrimSpace(FinalInput)
			SliceOfWords := strings.Split(FinalInput, " ")
			if len(SliceOfWords) == 2  {
				FirstInput , SecondInput := SliceOfWords[0], SliceOfWords[1]
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
		}()

		go func(){
			for _, Value := range features.TitleOfRegEx {
				FinalInput = strings.ToUpper(FinalInput)
				if FinalInput == Value {
					color.HiBlue(fmt.Sprintln("---------------------------------------------------------------"))
					color.HiBlue(fmt.Sprintln(features.OriginalAllRegex))
					color.HiBlue(fmt.Sprintln("---------------------------------------------------------------"))
					break
				} 
			}
		}()
		go func() {
			for _, value := range features.TitleOfTimeData {
				FinalInput = strings.ToUpper(FinalInput)
				if FinalInput == value {
					color.HiBlue(fmt.Sprintln("---------------------------------------------------------------"))
					color.HiBlue(fmt.Sprintln(features.OriginalTimeData))
					color.HiBlue(fmt.Sprintln("---------------------------------------------------------------"))
					break
				} 
			}
		}()
		go func() {
			for _, value := range features.TitleOfReadingWriting {
				FinalInput = strings.ToUpper(FinalInput)
				if FinalInput == value {
					color.HiBlue(fmt.Sprintln("---------------------------------------------------------------"))
					color.HiBlue(fmt.Sprintln(features.OriginalReadingandWriting))
					color.HiBlue(fmt.Sprintln("---------------------------------------------------------------"))
					break
				} 
			}
		}()
		go func() {
			for _, Value := range features.TitleOfSlice {
				FinalInput = strings.ToUpper(FinalInput)
				if FinalInput == Value {
					color.HiBlue(fmt.Sprintln("---------------------------------------------------------------"))
					color.HiBlue(fmt.Sprintln(features.OriginalJSONData))
					color.HiBlue(fmt.Sprintln("---------------------------------------------------------------"))
					break
				} 
			}
		}()
		go func() {
			for _, value := range features.TitleOfWorkingFiles {
				FinalInput = strings.ToUpper(FinalInput)
				if FinalInput == value {
					color.HiBlue(fmt.Sprintln("---------------------------------------------------------------"))
					color.HiBlue(fmt.Sprintln(features.OriginalWorkWithFiles))
					color.HiBlue(fmt.Sprintln("---------------------------------------------------------------"))
					break
				} 
			}	
		}()
		go func() {
			for _, value := range features.TitleOfUsingHTMLAndTextTemplates {
				FinalInput = strings.ToUpper(FinalInput)
				if FinalInput == value {
					color.HiBlue(fmt.Sprintln("---------------------------------------------------------------"))
					color.HiBlue(fmt.Sprintln(features.OriginalHTMLAndTemplates))
					color.HiBlue(fmt.Sprintln("---------------------------------------------------------------"))
					break
				} 
			}
		}()
	
		// Checks for each entry into the loop
		// if user input was "exit" so break loop and log out
		// from Gocron program. 
		Exit := "exit"
		if FinalInput == strings.ToLower(Exit) {
			// Regulators = false
			break
		}
		// time.Sleep(time.Second * 3)
		FinalInput = strings.ToLower(FinalInput)
		if FinalInput == "help" {
			features.HelpMessage()
		}
	}
}


func PrintNotAddYet(FinalInput string){
	color.HiRed(fmt.Sprintln("================================="))
	color.HiRed(fmt.Sprintf("Not add %v yet.",FinalInput))
	color.HiRed(fmt.Sprintln("================================="))
}


