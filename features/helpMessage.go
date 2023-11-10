package features

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

var LOGO string = fmt.Sprintln(` ┌─────────────────────────────────────────────────────────────────────────────────────────────────────────────┐ 
 •                                      GOCRON-VERSION 💠 v1.0.32 💠                                           •
 •                                       Go learning with CLI tool                                             • 
 • ██████████████████████████████████████████████████████████████████████████████████████████████████████████  •
 • ██              ███             ███              ██               █████             ███       ██████     █  •
 • ██              ██               ██              ██               ████               ██         ████     █  • 
 • ██    ████████████     █████     ██    ███████   ██    ███████    ████     █████     ██    █     ███     █  •  
 • ██    ████████████    ███████    ██    ████████████    ███████    ████    ███████    ██    ██     ██     █  •  
 • ██    █████     ██    ███████    ██    ████████████    █████      ████    ███████    ██    ███     █     █  •  
 • ██    █████     ██    ███████    ██    ████████████    ██      ███████    ███████    ██    ████          █  •  
 • ██    ███████   ██    ███████    ██    ████████████    ████      █████    ███████    ██    █████         █  •  
 • ██    ███████   ██     █████     ██    ███████   ██    ██████     ████     █████     ██    ██████        █  •  
 • ██              ██               ██              ██    ███████     ███               ██    ███████       █  •   
 • ██              ███             ███              ██    ███████      ███             ███    ████████      █  •  
 • ██████████████████████████████████████████████████████████████████████████████████████████████████████████  •
 └─────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
  ─────────────────────────────────────────────────────────────────────────────────────────────────────────────`)
var TitleOfHelp = []string{
	"-h",
	"help",
	"-help",
	"--help",
}

var Message string = fmt.Sprintln(`
my name is Sina LalehBakhsh, I hope this API is useful for you
after running program, write your single word about any of GO language.
if your perpuse is more than one word, for convenience searching, just write keywords.
like this:
	map slice
	`)

func HelpMessage() {
	color.HiBlue(`
┌───────────────────┐
|   GoCron v1.0.16  |
└───────────────────┘`)
	words := SplitIntoWords(Message)
	PrintWordByWord(words)
	fmt.Println()
	time.Sleep(time.Second * 2)
	color.HiBlue(fmt.Sprintln("You can use this command before start:"))
	time.Sleep(time.Second * 2)
	color.HiBlue(fmt.Sprintln(`Run
  go run .

Build
  go build .
  ./Gocron

Help
  ./Gocron -h 
  ./Gocron help
  ./Gocron -help
  ./Gocron --help

Show All
	./Gocron all
	./Gocron -all
	./Gocron --all`))
	time.Sleep(time.Second * 2)
	color.HiCyan(fmt.Sprintln("example command after run program:"))
	time.Sleep(time.Second * 2)
	color.HiCyan(fmt.Sprintln(`	help
	for example
	all regex
	goroutines and channels`))
	time.Sleep(time.Second * 2)
	color.HiMagenta(fmt.Sprintln("if just write one input you get all about that:"))
	time.Sleep(time.Second * 2)
	color.HiMagenta(fmt.Sprintln(`	READING AND WRITING DATA
	READINGANDWRITINGDATA
	ALL READING AND WRITING DATA
	ALLREADINGANDWRITINGDATA
	ALL REGEX 
	ALLREGEX
	ALL TIME
	ALLTIME
	ALL HTML AND TEMPLATE
	ALLHTMLANDTEMPLATE
	ALL WORKING WITH FILES
	ALLWORKINGWITHFILES
	ALL JSON
	ALLJSON
	ALL JSON DATA
	ALLJSONDATA
	ALL WORK WITH JSON DATA
	ALLWORKWITHJSONDATA
	ALL WORKING WITH JSON DATA
	ALLWORKINGWITHJSONDATA
	ALL HTTP SERVERS
	ALLHTTPSERVERS
	ALL CREATING HTTP SERVERS
	ALLCREATINGHTTPSERVERS
	ALL HTTP CLIENTS
	ALLHTTPCLIENTS
	ALL CREATING HTTP CLIENTS
	ALLCREATINGHTTPCLIENTS`))
	time.Sleep(time.Second * 2)
	color.HiBlue(fmt.Sprintln("💠💠💠 Would you like to join the gocron project on GitHub? (yes/no)"))


}
