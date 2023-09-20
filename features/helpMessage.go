package features

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
	"time"
)

var Logo string = fmt.Sprintln(`┌────────────────────────────────────────────────────────────────────┐ 
│                       GoCron v2.49.1                               │
│ █████████  █████████  ████████  █████████  █████████  ███      ██  │
│ █          █       █  █         █       █  █       █  ████     ██  │
│ █          █       █  █         █       █  █       █  ██ ██    ██  │
│ █     ███  █       █  █         █████████  █       █  ██  ██   ██  │
│ █       █  █       █  █         █    ██    █       █  ██   ██  ██  │
| █       █  █       █  █         █     ██   █       █  ██    ██ ██  │
| █████████  █████████  ████████  █      ██  █████████  ██     ████  │
│       is a API for learning GO language with example               │
└────────────────────────────────────────────────────────────────────┘ `)

var Message string = fmt.Sprintln(`
my name is Sina LalehBakhsh, I hope this API is useful for you
after running program, write your single word about any of GO language.
if your perpuse is more than one word, for convenience searching, just write keywords.
like this:
	map slice`)

func HelpMessage(){
	color.HiBlue(Logo)
	words := SplitIntoWords(Message)
	PrintWordByWord(words)

}


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
		counter ++
		if counter == 12 {
			fmt.Println()
			counter = 0
		}
		fmt.Print(word + " ")
		time.Sleep(85 * time.Millisecond) // Adjust the delay as needed
	}
}

