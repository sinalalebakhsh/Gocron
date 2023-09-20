package features

import (
	"fmt"
	"github.com/fatih/color"
)

var LOGO string = fmt.Sprintln(`┌────────────────────────────────────────────────────────────────────┐ 
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
	color.HiBlue(LOGO)
	words := SplitIntoWords(Message)
	PrintWordByWord(words)

}




