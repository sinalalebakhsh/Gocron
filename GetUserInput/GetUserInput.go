package getuserinput

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/sinalalebakhsh/Gocron/features"
)

func GetUserInput() (string, error) {
	UserInput := bufio.NewReader(os.Stdin)
	FinalInput, _ := UserInput.ReadString('\n') 
	FinalInput = strings.TrimSuffix(FinalInput, "\n")


	Final := features.OriginSingleDef.SingleDefinition[FinalInput]

	for _, Value := range features.OriginSingleDef.SingleDefinition {
		if features.OriginSingleDef.SingleDefinition[Value] == Final {
			return Final , nil
		} 
	}

	return "", errors.New("it's not yet add")

}