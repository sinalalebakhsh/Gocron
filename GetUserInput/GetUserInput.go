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


	Final := features.OriginSingleDef.SingleDefinition[FinalInput]

	fmt.Println(features.OriginSingleDef.SingleDefinition[Final])
	
	return "", errors.New("it's not yet add")

}