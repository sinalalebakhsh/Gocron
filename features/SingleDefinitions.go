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
	
	"flag": `5.flag package:
Command-line flags are a common way to specify options for command-line programs. 
For example, in wc -l the -l is a command-line flag.
Go provides a flag package supporting basic command-line flag parsing.
We'll use this package to implement our example command-line program.`,
	
	"flag package": `5.flag package:
Command-line flags are a common way to specify options for command-line programs. 
For example, in wc -l the -l is a command-line flag.
Go provides a flag package supporting basic command-line flag parsing.
We'll use this package to implement our example command-line program.`,
    
    "clean":`2.clean
The go clean command removes the output produced by the go build command, 
including the executable and any temporary files that were created during the build.`,

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

"Create Environment GO":`0.Create Environment GO
in Command Line Interface Go:
1- go mod init YOURNAME
2- go work init YOURWORKDIRECTORY
3- go run main.go  OR  go run projectName.go
`,


},

}
