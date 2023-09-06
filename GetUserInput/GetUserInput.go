package getuserinput

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/sinalalebakhsh/Gocron/features"
)

func GetUserInput() (string, error) {
	UserInput := bufio.NewReader(os.Stdin)
	FinalInput, _ := UserInput.ReadString('\n') 
	FinalInput = strings.TrimSuffix(FinalInput, "\n")

	for _ , Value := range features.OriginSingleDef.SingleDefinition {
		if Value == features.OriginSingleDef.SingleDefinition[FinalInput] {
			Final := fmt.Sprint(features.OriginSingleDef.SingleDefinition[FinalInput])			
			return Final, nil
		} else {
			goto target
		}
	} 

	target: return "", errors.New("it's not yet add")

}