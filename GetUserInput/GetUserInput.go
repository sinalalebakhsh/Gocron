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
			color.Green(fmt.Sprintln(features.OriginSingleDef.SingleDefinition[value]))
			Regulators = true
			break
		}
	}

	if Regulators == false {
		fmt.Println("---------------------")
		color.Red(fmt.Sprintf("Do not add %s yet.", FinalInput))
		fmt.Println("---------------------")

	}

}
