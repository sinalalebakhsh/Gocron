package features

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

func SplitIntoWords(text string) []string {
	// Split the text into words
	return splitAndTrim(text, " ")
}

func splitAndTrim(s, sep string) []string {
	words := strings.Fields(s)
	for i, word := range words {
		words[i] = strings.TrimSpace(word)
	}
	return words
}

func PrintWordByWord(words []string) {
	counter := 0
	for _, word := range words {
		counter++
		if counter == 12 {
			fmt.Println()
			counter = 0
		}
		fmt.Print(word + " ")
		time.Sleep(85 * time.Millisecond) // Adjust the delay as needed
	}
}

func init() {
	color.HiBlue(fmt.Sprintln("---------------------------------------------------------------"))
	color.HiBlue(LOGO)
	fmt.Println()
	color.HiBlue(fmt.Sprintln("---------------------------------------------------------------"))
}
