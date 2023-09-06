package getuserinput

import (
	"bufio"
	"os"
	"strings"
)

func GetUserInput() string {
	UserInput := bufio.NewReader(os.Stdin)
	Input, _ := UserInput.ReadString('\n') 
	Input = strings.TrimSuffix(Input, "\n")
	return Input
}