package features

type WordsInGolangNumbered struct {
	MapOfNumbered map[string]string
}

var OriginalWordsGoNum = WordsInGolangNumbered{
	MapOfNumbered: map[string]string{
		"0":`0.Create Environment GO
in Command Line Interface Go:
1- go mod init YOURNAME
2- go work init YOURWORKDIRECTORY
3- go run main.go  OR  go run projectName.go

package main = first executable file main.go

func main() {} = every execute files in "package main"

import = for importing another package

function
A function is an independent section of code 
that mapszero or more input parameters to zero or more outputparameters.   
Functions (also known as procedures orsubroutines) 
are often represented as a black box: (theblack box represents the function)
func FunctionName(ParameterName ParameterType) ReturnType { body of function}
example:
	func getUserName(UserInput string) string { return UserInput }

returning multiple values
	Go is also capable of returning multiple values from afunction

variadic functions
func add(args ...int) int {}`,
// ====================================================================================
		"1":`1.build
Using the Go Command
The go build command compiles the source code in the current directory 
and generates an executable file.`,
// ====================================================================================
		"2":`2.clean
go clean command removes the output produced by the go build command, 
the executable and any temporary files that were created during the build.`,
// ====================================================================================
		"3":`3.doc
The go doc command generates documentation from source code.`,
// ====================================================================================
		"4":`4.fmt
The go fmt command ensures consistent indentation and alignment in source code files.
fmt = for printing in CLI
fmt.Printf("%v %s %f", variable, string, float32)`,
// ====================================================================================
		"5":`5.get
The go get command downloads and installs external packages.
flag
flag package:
Command-line flags are a common way to specify options for command-line programs. 
For example, in wc -l the -l is a command-line flag.
Go provides a flag package supporting basic command-line flag parsing.
We'll use this package to implement our example command-line program.`,
// ====================================================================================
		"6":`6.install
The go install command downloads packages and is usually used to install tool packages.`,
// ====================================================================================
		"7":`7.help
The go help command displays help information for other Go features. The command go
help build, for example, displays information about the build argument.`,
// ====================================================================================
		"8":`8.mod
The go mod command is used to create and manage a Go module.`,
// ====================================================================================
		"9":`9.run
The go run command builds and executes the source code in a specified folder without
creating an executable output`,
// ====================================================================================
		"10":`10.test
The go test command executes unit tests`,
// ====================================================================================
		"11":`11.version
The go version command writes out the Go version number.`,
// ====================================================================================
		"12":`12.vet
The go vet command detects common problems in Go code`,
// ====================================================================================
		"13":`13.print <expr>
Useful Debugger State Commands
This command evaluates an expression and displays the result. It can
be used to display a value (print i) or perform a more complex test
(print i > 0).`,
// ====================================================================================
		"14":`14.set <variable> = <value>
This command changes the value of the specified variable.
این دستور مقدار متغیر مشخص شده را تغییر می دهد.`,
// ====================================================================================
		"15":`15.locals
This command prints the value of all local variables.`,
// ====================================================================================
		"16":`16.whatis <expr>
This command prints the type of the specified expression such as whatis`,
// ====================================================================================
		"17":``,
// ====================================================================================
		"18":``,
// ====================================================================================
		"19":``,
// ====================================================================================
		"20":``,
// ====================================================================================
		"21":``,
// ====================================================================================
		"22":``,
// ====================================================================================
		"23":``,
// ====================================================================================
		"24":``,
// ====================================================================================
		"25":``,
// ====================================================================================
		"26":``,
// ====================================================================================
// ====================================================================================
// ====================================================================================
// ====================================================================================
// ====================================================================================
// ====================================================================================
// ====================================================================================


	},
}