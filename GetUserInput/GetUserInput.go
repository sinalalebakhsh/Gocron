package getuserinput

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
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


		// Checks for each entry into the loop
		// if user input was "exit" so break loop and log out
		// from Gocron program. 
		if FinalInput == "exit" {
			break
		}

		SliceOfMap := make([]string, 0, len(features.OriginSingleDef.SingleDefinition))
		for key, NotUsed := range features.OriginSingleDef.SingleDefinition {
			if NotUsed == "nil" {
				fmt.Println(NotUsed, "!!!!!!")
			}
			SliceOfMap = append(SliceOfMap, key)
		}
		sort.Strings(SliceOfMap)
	
		Regulators := false
		for _, value := range SliceOfMap {
			FinalInput = strings.ToLower(FinalInput)
			if FinalInput == value {
				words := features.SplitIntoWords(features.OriginSingleDef.SingleDefinition[value])
				features.PrintWordByWord(words)
				fmt.Println()
				Regulators = true
				break
			}
		}
	
		go func() {
			for _, Value := range features.TitleOfRegEx {
				FinalInput = strings.ToUpper(FinalInput)
				if FinalInput == Value {
					color.HiBlue(fmt.Sprintln("---------------------------------------------------------------"))
					color.HiBlue(fmt.Sprintln(features.OriginalAllRegex))
					color.HiBlue(fmt.Sprintln("---------------------------------------------------------------"))
					Regulators = true
					break
				}
			}
		}()
	
		go func(){
			for _, value := range features.TitleOfTimeData {
				FinalInput = strings.ToUpper(FinalInput)
				if FinalInput == value {
					color.HiBlue(fmt.Sprintln("---------------------------------------------------------------"))
					color.HiBlue(fmt.Sprintln(features.OriginalTimeData))
					color.HiBlue(fmt.Sprintln("---------------------------------------------------------------"))
					Regulators = true
					break
				}
			}
		}()
		
		go func(){
			for _, value := range features.TitleOfReadingWriting {
				FinalInput = strings.ToUpper(FinalInput)
				if FinalInput == value {
					color.HiBlue(fmt.Sprintln("---------------------------------------------------------------"))
					color.HiBlue(fmt.Sprintln(features.OriginalReadingandWriting))
					color.HiBlue(fmt.Sprintln("---------------------------------------------------------------"))
					Regulators = true
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
					Regulators = true
					break
				}
			}
		}()
	
		go func()  {
			for _, value := range features.TitleOfWorkingFiles {
				FinalInput = strings.ToUpper(FinalInput)
				if FinalInput == value {
					color.HiBlue(fmt.Sprintln("---------------------------------------------------------------"))
					color.HiBlue(fmt.Sprintln(features.OriginalWorkWithFiles))
					color.HiBlue(fmt.Sprintln("---------------------------------------------------------------"))
					Regulators = true
					break
				}
			}	
		}()
	
		time.Sleep(time.Second)
		FinalInput = strings.ToLower(FinalInput)
		if FinalInput == "help" {
			features.HelpMessage()
		}else if !Regulators {
			color.Red(fmt.Sprintln("---------------------"))
			color.Red(fmt.Sprintf("Do not add %s yet.", FinalInput))
			color.Red(fmt.Sprintln("---------------------"))
		}
	}
}
