package features


type SingleDefinitions struct {
    SliceSingleDefinitions []string
	SingleDefinition map[string]string
}

var OriginSingleDef SingleDefinitions = SingleDefinitions{
    SliceSingleDefinitions: []string{
        "build",
        "flag",
        "flag package",

    },

	SingleDefinition: map[string]string{
		"build": `1.build:
The go build command compiles the source code in the current directory 
and generates an executable file.`,
// ██████████████████████████████████████████████████████████████████████
	"flag": `5.flag package:
Command-line flags are a common way to specify options for command-line programs. 
For example, in wc -l the -l is a command-line flag.
Go provides a flag package supporting basic command-line flag parsing.
We'll use this package to implement our example command-line program.`,
// ██████████████████████████████████████████████████████████████████████
	"flag package": `5.flag package:
Command-line flags are a common way to specify options for command-line programs. 
For example, in wc -l the -l is a command-line flag.
Go provides a flag package supporting basic command-line flag parsing.
We'll use this package to implement our example command-line program.`,
// ██████████████████████████████████████████████████████████████████████    
    "clean":`2.clean
The go clean command removes the output produced by the go build command, 
including the executable and any temporary files that were created during the build.`,
// ██████████████████████████████████████████████████████████████████████
"Goroutines":`138.Goroutines and Channels
What are they?
Goroutines are lightweight threads created and managed by the Go runtime.
Channels are pipes that carry values of a specific type.

Why are they useful?
Goroutines allow functions to be executed concurrently, without needing to deal
with the complications of operating system threads. Channels allow goroutines to
produce results asynchronously.

How are they used?
Goroutines are created using the go keyword. Channels are defined as data types.

Are there any or limitations?
pitfalls Care must be taken to manage the direction of channels. 
Goroutines that share data require additional features.

Are there any alternatives?
Goroutines and channels are the built-in Go concurrency features, but some
applications can rely on a single thread of execution, which is created by default to
execute the main function.

Problem                                 Solution
---------------------------------       -------------------------
Execute a function asynchronously       Create a goroutine
Produce a result from a function        Use a channel 
executed asynchronously
Send and receive values                 Use arrow expressions 
using a channel
Indicate that no further values         Use the close function
will be sent over a channel
Enumerate the values received           Use a for loop with the range keyword 
from a channel
`,
// ██████████████████████████████████████████████████████████████████████
    "Channels":`138.Goroutines and Channels
What are they?
Goroutines are lightweight threads created and managed by the Go runtime.
Channels are pipes that carry values of a specific type.

Why are they useful?
Goroutines allow functions to be executed concurrently, without needing to deal
with the complications of operating system threads. Channels allow goroutines to
produce results asynchronously.

How are they used?
Goroutines are created using the go keyword. Channels are defined as data types.

Are there any or limitations?
pitfalls Care must be taken to manage the direction of channels. 
Goroutines that share data require additional features.

Are there any alternatives?
Goroutines and channels are the built-in Go concurrency features, but some
applications can rely on a single thread of execution, which is created by default to
execute the main function.

Problem                                 Solution
---------------------------------       -------------------------
Execute a function asynchronously       Create a goroutine
Produce a result from a function        Use a channel 
executed asynchronously
Send and receive values                 Use arrow expressions 
using a channel
Indicate that no further values         Use the close function
will be sent over a channel
Enumerate the values received           Use a for loop with the range keyword 
from a channel
`,
// ██████████████████████████████████████████████████████████████████████
    "Goroutines and Channels":`138.Goroutines and Channels
What are they?
Goroutines are lightweight threads created and managed by the Go runtime.
Channels are pipes that carry values of a specific type.

Why are they useful?
Goroutines allow functions to be executed concurrently, without needing to deal
with the complications of operating system threads. Channels allow goroutines to
produce results asynchronously.

How are they used?
Goroutines are created using the go keyword. Channels are defined as data types.

Are there any or limitations?
pitfalls Care must be taken to manage the direction of channels. 
Goroutines that share data require additional features.

Are there any alternatives?
Goroutines and channels are the built-in Go concurrency features, but some
applications can rely on a single thread of execution, which is created by default to
execute the main function.

Problem                                 Solution
---------------------------------       -------------------------
Execute a function asynchronously       Create a goroutine
Produce a result from a function        Use a channel 
executed asynchronously
Send and receive values                 Use arrow expressions 
using a channel
Indicate that no further values         Use the close function
will be sent over a channel
Enumerate the values received           Use a for loop with the range keyword 
from a channel
`,
// ██████████████████████████████████████████████████████████████████████
    "Create Environment GO":`0.Create Environment GO
in Command Line Interface Go:
1- go mod init YOURNAME
2- go work init YOURWORKDIRECTORY
3- go run main.go  OR  go run projectName.go
`,
// ██████████████████████████████████████████████████████████████████████
    "Environment GO":`0.Create Environment GO
in Command Line Interface Go:
1- go mod init YOURNAME
2- go work init YOURWORKDIRECTORY
3- go run main.go  OR  go run projectName.go`,
// ██████████████████████████████████████████████████████████████████████
    "GO Environment":`0.Create Environment GO
in Command Line Interface Go:
1- go mod init YOURNAME
2- go work init YOURWORKDIRECTORY
3- go run main.go  OR  go run projectName.go`,
// ██████████████████████████████████████████████████████████████████████
    "Create GO Environment":`0.Create Environment GO
in Command Line Interface Go:
1- go mod init YOURNAME
2- go work init YOURWORKDIRECTORY
3- go run main.go  OR  go run projectName.go`,
// ██████████████████████████████████████████████████████████████████████
    "Run GO Environment Project ":`0.Create Environment GO
in Command Line Interface Go:
1- go mod init YOURNAME
2- go work init YOURWORKDIRECTORY
3- go run main.go  OR  go run projectName.go`,
// ██████████████████████████████████████████████████████████████████████
    "Create GO Project":`0.Create Environment GO
in Command Line Interface Go:
1- go mod init YOURNAME
2- go work init YOURWORKDIRECTORY
3- go run main.go  OR  go run projectName.go`,
// ██████████████████████████████████████████████████████████████████████
    "new() function":`188.new() function
example:
    package main
    import "fmt"
    func main() {
        // 1.
        var ChannelName1 string = "AcronProject"
        fmt.Println("Value=", ChannelName1, "memory address=", &ChannelName1)
        // 2.
        var ChannelName2 *string = new(string)
        fmt.Println("Value ??? = ", ChannelName2, "memory address=", &ChannelName2)
        var ChannelName3 *string = new(string)
        fmt.Println("Value ??? = ", ChannelName3, "memory address=", &ChannelName3)
        fmt.Printf("%q \n",*ChannelName2 + *ChannelName3)
    }
Output:
Value= AcronProject memory address= 0xc000014070
Value ??? =  0xc000014090 memory address= 0xc00004e028
Value ??? =  0xc0000140a0 memory address= 0xc00004e030
""
`,
// ██████████████████████████████████████████████████████████████████████
"Regular Expressions":`189.Regular Expressions
The regular expressions used in this section perform basic matches, but the regexp package
supports an extensive pattern syntax, which is described at https://pkg.go.dev/regexp/syntax@go1.17.1.

The Basic Functions Provided by the regexp Package
Function                Description
-----------------       --------------------------------------------------------------------------
Match(pattern, b)       This function returns a bool that indicates whether a pattern is matched by
                        the byte slice b.

MatchString(patten, s)  This function returns a bool that indicates whether a pattern is matched by
                        the string s.

Compile(pattern)        This function returns a RegExp that can be used to perform repeated pattern
                        matching with the specified pattern.

MustCompile(pattern)    This function provides the same feature as Compile but panics, 
                        if the specified pattern cannot be compiled.

example:
    package main
    import (
        "fmt"
        //"strings"
        "regexp"
    )
    func main() {
        description := "A boat for one person"
        match, err := regexp.MatchString("[A-z]oat", description)
        if (err == nil) {
            fmt.Println("Match:", match)
        } else {
            fmt.Println("Error:", err)
        }
    }
Output:
    Match: true`,
// ██████████████████████████████████████████████████████████████████████
"Currying":`Function currying is the practice of writing a function 
that takes a function (or functions) as input, and returns a new function.`,
// ██████████████████████████████████████████████████████████████████████
"Replace(s)":`This method returns a string for which all the replacements specified with the
constructor have been performed on the string s.`,
// ██████████████████████████████████████████████████████████████████████
"WriteString(writer, s)":`This method is used to perform the replacements specified with the constructor
and write the results to an io.Writer`,
// ██████████████████████████████████████████████████████████████████████
"TrimSpace(s)":`This function returns the string s without leading or trailing whitespace
characters.`,
// ██████████████████████████████████████████████████████████████████████
"Trim(s, set)":`This function returns a string from which any leading or trailing characters
contained in the string set are removed from the string s.`,
// ██████████████████████████████████████████████████████████████████████
"TrimLeft(s, set)":`    This function returns the string s without any leading character contained
in the string set. This function matches any of the specified characters—use
the TrimPrefix function to remove a complete substring.`,
// ██████████████████████████████████████████████████████████████████████
"TrimRight(s, set)":`This function returns the string s without any trailing character contained
in the string set. This function matches any of the specified characters—use
the TrimSuffix function to remove a complete substring.`,
// ██████████████████████████████████████████████████████████████████████
"TrimPrefix(s, prefix)":`function returns the string s after removing the specified prefix string.
This function removes the complete prefix string—use the TrimLeft
function to remove characters from a set.`,
// ██████████████████████████████████████████████████████████████████████
"TrimSuffix(s, suffix)":`function returns the string s after removing the specified suffix string.
This function removes the complete suffix string—use the TrimRight
function to remove characters from a set.`,
// ██████████████████████████████████████████████████████████████████████

// ██████████████████████████████████████████████████████████████████████
// ██████████████████████████████████████████████████████████████████████
// ██████████████████████████████████████████████████████████████████████
// ██████████████████████████████████████████████████████████████████████
// ██████████████████████████████████████████████████████████████████████
// ██████████████████████████████████████████████████████████████████████
// ██████████████████████████████████████████████████████████████████████
// ██████████████████████████████████████████████████████████████████████
// ██████████████████████████████████████████████████████████████████████
// ██████████████████████████████████████████████████████████████████████
// ██████████████████████████████████████████████████████████████████████
// ██████████████████████████████████████████████████████████████████████
// ██████████████████████████████████████████████████████████████████████
// ██████████████████████████████████████████████████████████████████████
// ██████████████████████████████████████████████████████████████████████
// ██████████████████████████████████████████████████████████████████████
// ██████████████████████████████████████████████████████████████████████
// ██████████████████████████████████████████████████████████████████████
// ██████████████████████████████████████████████████████████████████████
// ██████████████████████████████████████████████████████████████████████
// ██████████████████████████████████████████████████████████████████████
// ██████████████████████████████████████████████████████████████████████
// ██████████████████████████████████████████████████████████████████████
// ██████████████████████████████████████████████████████████████████████
// ██████████████████████████████████████████████████████████████████████
// ██████████████████████████████████████████████████████████████████████
// ██████████████████████████████████████████████████████████████████████
// ██████████████████████████████████████████████████████████████████████
// ██████████████████████████████████████████████████████████████████████

},

}
