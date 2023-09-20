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

func GetUserInput() {
	UserInput := bufio.NewReader(os.Stdin)
	FinalInput, _ := UserInput.ReadString('\n')
	FinalInput = strings.TrimSuffix(FinalInput, "\n")

	FinalInput = strings.ToLower(FinalInput)

	SliceOfMap := make([]string, 0, len(features.OriginSingleDef.SingleDefinition))
	for key, NotUsed:= range features.OriginSingleDef.SingleDefinition {
		if NotUsed == "nil" {
			fmt.Println(NotUsed, "!!!!!!")
		}
		SliceOfMap = append(SliceOfMap, key)
	}
	sort.Strings(SliceOfMap)

	Regulators := false
	
	for _, value := range SliceOfMap {
		if FinalInput == value {
			color.Green(fmt.Sprintln("---------------------------------------------------------------"))
			color.Green(fmt.Sprintln(features.OriginSingleDef.SingleDefinition[value]))
			color.Green(fmt.Sprintln("---------------------------------------------------------------"))
			Regulators = true
			break
		// } else if FinalInput == "regex" ||
		// 	FinalInput == "allregex" ||
		// 	FinalInput == "all regex" ||
		// 	FinalInput == "regexs" ||
		// 	FinalInput == "allregexs" ||
		// 	FinalInput == "all regexs" ||
		// 	FinalInput == "regula rexpression" ||
		// 	FinalInput == "regularexpression" {
		// 		color.Green(fmt.Sprintln("---------------------------------------------------------------"))
		// 		color.Green(fmt.Sprintln(features.OriginalAllRegex))
		// 		color.Green(fmt.Sprintln("---------------------------------------------------------------"))
		// 		Regulators = true
		// 		break
		// 	} else if FinalInput == "time" || 
		// 	FinalInput == "datatime" ||
		// 	FinalInput == "data time" ||
		// 	FinalInput == "alldatatime" ||
		// 	FinalInput == "all data time" ||
		// 	FinalInput == "all data times" ||
		// 	FinalInput == "thetime" ||
		// 	FinalInput == "the time" {
		// 		color.Green(fmt.Sprintln("---------------------------------------------------------------"))
		// 		color.Green(fmt.Sprintln(features.OriginalTimeData))
		// 		color.Green(fmt.Sprintln("---------------------------------------------------------------"))
		// 		Regulators = true
		// 		break
		// 	} else if FinalInput == "reading and writing data" || 
		// 	FinalInput == "readingandwritingdata" ||
		// 	FinalInput == "readingwritingdata" ||
		// 	FinalInput == "readingwriting" ||
		// 	FinalInput == "reading writing" ||
		// 	FinalInput == "reading writing datas" ||
		// 	FinalInput == "reading & writing data" ||
		// 	FinalInput == "reading & writing"{
		// 		color.Green(fmt.Sprintln("---------------------------------------------------------------"))
		// 		color.Green(fmt.Sprintln(features.OriginalReadingandWriting))	
		// 		color.Green(fmt.Sprintln("---------------------------------------------------------------"))
		// 		Regulators = true
		// 		break

			}
	}
	
	go func ()  {
		for _, Value := range features.TitleOfSlice {
			if FinalInput == Value {
				color.Green(fmt.Sprintln("---------------------------------------------------------------"))
				color.Green(fmt.Sprintln(features.OriginalJSONData))	
				color.Green(fmt.Sprintln("---------------------------------------------------------------"))
				Regulators = true
			}
		}	
	}()


	time.Sleep(time.Second)
	if !Regulators {
		color.Red(fmt.Sprintln("---------------------"))
		color.Red(fmt.Sprintf("Do not add %s yet.", FinalInput))
		color.Red(fmt.Sprintln("---------------------"))
	} 

}
