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

	SliceOfMap := make([]string, 0, len(features.OriginSingleDef.SingleDefinition))

	for key, _ := range features.OriginSingleDef.SingleDefinition {
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
		}
	}

	if Regulators == false {
		color.Red(fmt.Sprintln("---------------------"))
		color.Red(fmt.Sprintf("Do not add %s yet.", FinalInput))
		color.Red(fmt.Sprintln("---------------------"))
	}

}
