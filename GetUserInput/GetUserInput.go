package getuserinput

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/sinalalebakhsh/Gocron/features"
)

func GetUserInput() (string, error) {
	UserInput := bufio.NewReader(os.Stdin)
	FinalInput, _ := UserInput.ReadString('\n')
	FinalInput = strings.TrimSuffix(FinalInput, "\n")

	SliceOfMap := make([]string, 0, len(features.OriginSingleDef.SingleDefinition))

	for key, _ := range features.OriginSingleDef.SingleDefinition {
		SliceOfMap = append(SliceOfMap, key)
	}

	sort.Strings(SliceOfMap)

	// Regulators := false

	for _, value := range SliceOfMap {
		if FinalInput == value {
			SingleDefinition := features.OriginSingleDef.SingleDefinition[value]
			
			return SingleDefinition, nil

			// Regulators = true
			// break
		}
	}

	// if Regulators == false {
	// }
	return "", fmt.Errorf("do not add %s yet", FinalInput)

}
