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
		"4":``,
// ====================================================================================
		"5":``,
// ====================================================================================
		"6":``,
// ====================================================================================
		"7":``,
// ====================================================================================
		"8":``,
// ====================================================================================
		"9":``,
// ====================================================================================
		"10":``,

	},
}