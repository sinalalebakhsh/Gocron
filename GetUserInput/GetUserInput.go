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
		} else if FinalInput == "regEx" ||
			FinalInput == "REGEX" ||
			FinalInput == "RegGex" ||
			FinalInput == "Regular Expression" {
			color.Green(fmt.Sprintln("---------------------------------------------------------------"))
			color.Green(fmt.Sprintln(features.OriginalAllRegex))
			color.Green(fmt.Sprintln("---------------------------------------------------------------"))
			Regulators = true
			break
			} else if FinalInput == "time" || 
			FinalInput == "datatime" ||
			FinalInput == "alldatatime" ||
			FinalInput == "thetime" {
			color.Green(fmt.Sprintln("---------------------------------------------------------------"))
			color.Green(fmt.Sprintln(features.OriginalTimeData))
			color.Green(fmt.Sprintln("---------------------------------------------------------------"))
			Regulators = true
			break
			}
	}
	


	if !Regulators {
		color.Red(fmt.Sprintln("---------------------"))
		color.Red(fmt.Sprintf("Do not add %s yet.", FinalInput))
		color.Red(fmt.Sprintln("---------------------"))
	} 

}
