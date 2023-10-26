package features

type SingleDefinitions struct {
	SingleDef map[string]string
}

var OriginSingleDef = SingleDefinitions{

	SingleDef: map[string]string{
		"go":                         "Fans of Go (called gophers) describe Go as having the expressiveness of dynamic languages like Python or Ruby, with the performance of compiled languages like C or C++. The language is open source, and was started by engineers at Google. It's written using a C-style syntax, has statically typed variables, manages memory using garbage collection, and is compiled into stand-alone executables. Go is noted for the concurrent programming features built into the language core, the networking packages in the standard library (such as a web server), fast compilation and execution speed. Its simple, minimalistic and consistent language design make for a delightful experience, while the abundant and thoughtful tooling addresses traditional problems such as consistent formatting and documentation. The home page for Go is go.dev, and there is an excellent interactive tutorial at tour.go.dev.",
		"about go":                   "Fans of Go (called gophers) describe Go as having the expressiveness of dynamic languages like Python or Ruby, with the performance of compiled languages like C or C++. The language is open source, and was started by engineers at Google. It's written using a C-style syntax, has statically typed variables, manages memory using garbage collection, and is compiled into stand-alone executables. Go is noted for the concurrent programming features built into the language core, the networking packages in the standard library (such as a web server), fast compilation and execution speed. Its simple, minimalistic and consistent language design make for a delightful experience, while the abundant and thoughtful tooling addresses traditional problems such as consistent formatting and documentation. The home page for Go is go.dev, and there is an excellent interactive tutorial at tour.go.dev.",
		"create environment go":      "0.Create Environment GO in Command Line Interface Go: 1- go mod init YOURNAME 2- go work init YOURWORKDIRECTORY 3- go run main.go  OR  go run projectName.go",
		"create go project":          "0.Create Environment GO in Command Line Interface Go: 1- go mod init YOURNAME 2- go work init YOURWORKDIRECTORY 3- go run main.go  OR  go run projectName.go",
		"go project":                 "0.Create Environment GO in Command Line Interface Go: 1- go mod init YOURNAME 2- go work init YOURWORKDIRECTORY 3- go run main.go  OR  go run projectName.go",
		"environment go":             "0.Create Environment GO in Command Line Interface Go: 1- go mod init YOURNAME 2- go work init YOURWORKDIRECTORY 3- go run main.go  OR  go run projectName.go",
		"go environment":             "0.Create Environment GO in Command Line Interface Go: 1- go mod init YOURNAME 2- go work init YOURWORKDIRECTORY 3- go run main.go  OR  go run projectName.go",
		"create go environment":      "0.Create Environment GO in Command Line Interface Go: 1- go mod init YOURNAME 2- go work init YOURWORKDIRECTORY 3- go run main.go  OR  go run projectName.go",
		"run go environment":         "0.Create Environment GO in Command Line Interface Go: 1- go mod init YOURNAME 2- go work init YOURWORKDIRECTORY 3- go run main.go  OR  go run projectName.go",
		"run go environment project": "0.Create Environment GO in Command Line Interface Go: 1- go mod init YOURNAME 2- go work init YOURWORKDIRECTORY 3- go run main.go  OR  go run projectName.go",
		"build":                      "1.build: The go build command compiles the source code in the current directory and generates an executable file.",
		"clean":                      "2.clean: The go clean command removes the output produced by the go build command,  including the executable and any temporary files that were created during the build.",
		"flag":                       "5.flag package: Command-line flags are a common way to specify options for command-line programs. For example, in wc -l the -l is a command-line flag. Go provides a flag package supporting basic command-line flag parsing. We'll use this package to implement our example command-line program.",
		"flag package":               "5.flag package: Command-line flags are a common way to specify options for command-line programs. For example, in wc -l the -l is a command-line flag. Go provides a flag package supporting basic command-line flag parsing. We'll use this package to implement our example command-line program.",
		"goroutine":                  "138.Goroutines and Channels: What are they? Goroutines are lightweight threads created and managed by the Go runtime. Channels are pipes that carry values of a specific type. Why are they useful? Goroutines allow functions to be executed concurrently, without needing to deal with the complications of operating system threads. Channels allow goroutines to produce results asynchronously. How are they used? Goroutines are created using the go keyword. Channels are defined as data types. Are there any or limitations? pitfalls Care must be taken to manage the direction of channels. Goroutines that share data require additional features. Are there any alternatives? Goroutines and channels are the built-in Go concurrency features, but some applications can rely on a single thread of execution, which is created by default to execute the main function.",
		"go routine":                 "138.Goroutines and Channels: What are they? Goroutines are lightweight threads created and managed by the Go runtime. Channels are pipes that carry values of a specific type. Why are they useful? Goroutines allow functions to be executed concurrently, without needing to deal with the complications of operating system threads. Channels allow goroutines to produce results asynchronously. How are they used? Goroutines are created using the go keyword. Channels are defined as data types. Are there any or limitations? pitfalls Care must be taken to manage the direction of channels. Goroutines that share data require additional features. Are there any alternatives? Goroutines and channels are the built-in Go concurrency features, but some applications can rely on a single thread of execution, which is created by default to execute the main function.",
		"goroutines":                 "138.Goroutines and Channels: What are they? Goroutines are lightweight threads created and managed by the Go runtime. Channels are pipes that carry values of a specific type. Why are they useful? Goroutines allow functions to be executed concurrently, without needing to deal with the complications of operating system threads. Channels allow goroutines to produce results asynchronously. How are they used? Goroutines are created using the go keyword. Channels are defined as data types. Are there any or limitations? pitfalls Care must be taken to manage the direction of channels. Goroutines that share data require additional features. Are there any alternatives? Goroutines and channels are the built-in Go concurrency features, but some applications can rely on a single thread of execution, which is created by default to execute the main function.",
		"channels":                   "138.Goroutines and Channels: What are they? Goroutines are lightweight threads created and managed by the Go runtime. Channels are pipes that carry values of a specific type. Why are they useful? Goroutines allow functions to be executed concurrently, without needing to deal with the complications of operating system threads. Channels allow goroutines to produce results asynchronously. How are they used? Goroutines are created using the go keyword. Channels are defined as data types. Are there any or limitations? pitfalls Care must be taken to manage the direction of channels. Goroutines that share data require additional features. Are there any alternatives? Goroutines and channels are the built-in Go concurrency features, but some applications can rely on a single thread of execution, which is created by default to execute the main function.",
		"goroutines and channels":    "138.Goroutines and Channels: What are they? Goroutines are lightweight threads created and managed by the Go runtime. Channels are pipes that carry values of a specific type. Why are they useful? Goroutines allow functions to be executed concurrently, without needing to deal with the complications of operating system threads. Channels allow goroutines to produce results asynchronously. How are they used? Goroutines are created using the go keyword. Channels are defined as data types. Are there any or limitations? pitfalls Care must be taken to manage the direction of channels. Goroutines that share data require additional features. Are there any alternatives? Goroutines and channels are the built-in Go concurrency features, but some applications can rely on a single thread of execution, which is created by default to execute the main function.",
		"channels and goroutines":    "138.Goroutines and Channels: What are they? Goroutines are lightweight threads created and managed by the Go runtime. Channels are pipes that carry values of a specific type. Why are they useful? Goroutines allow functions to be executed concurrently, without needing to deal with the complications of operating system threads. Channels allow goroutines to produce results asynchronously. How are they used? Goroutines are created using the go keyword. Channels are defined as data types. Are there any or limitations? pitfalls Care must be taken to manage the direction of channels. Goroutines that share data require additional features. Are there any alternatives? Goroutines and channels are the built-in Go concurrency features, but some applications can rely on a single thread of execution, which is created by default to execute the main function.",
		"regular expressions":        "189.Regular Expressions: The regular expressions used in this section perform basic matches, but the regexp package supports an extensive pattern syntax, which is described at https://pkg.go.dev/regexp/syntax@go1.17.1.",
		"regularexpressions":         "189.Regular Expressions: The regular expressions used in this section perform basic matches, but the regexp package supports an extensive pattern syntax, which is described at https://pkg.go.dev/regexp/syntax@go1.17.1.",
		"regex":                      "189.Regular Expressions: The regular expressions used in this section perform basic matches, but the regexp package supports an extensive pattern syntax, which is described at https://pkg.go.dev/regexp/syntax@go1.17.1.",
		"currying":                   "Currying: Function currying is the practice of writing a function that takes a function (or functions) as input, and returns a new function.",
		"replace(s)":                 "Replace(s): This method returns a string for which all the replacements specified with the constructor have been performed on the string s.",
		"replace()":                  "Replace(s): This method returns a string for which all the replacements specified with the constructor have been performed on the string s.",
		"writestring(writer, s)":     "WriteString(writer, s): This method is used to perform the replacements specified with the constructor and write the results to an io.Writer",
		"trimspace(s)":               "TrimSpace(s): This function returns the string s without leading or trailing whitespace characters.",
		"trim(s, set)":               "Trim(s, set): This function returns a string from which any leading or trailing characters contained in the string set are removed from the string s.",
		"trimLeft(s, set)":           "This function returns the string s without any leading character contained in the string set. This function matches any of the specified characters—use the TrimPrefix function to remove a complete substring.",
		"trimright(s, set)":          "This function returns the string s without any trailing character contained in the string set. This function matches any of the specified characters—use the TrimSuffix function to remove a complete substring.",
		"trimprefix(s, prefix)":      "function returns the string s after removing the specified prefix string. This function removes the complete prefix string—use the TrimLeft function to remove characters from a set.",
		"trimsuffix(s, suffix)":      "function returns the string s after removing the specified suffix string. This function removes the complete suffix string—use the TrimRight function to remove characters from a set.",
		"trimfunc(s, func)":          "TrimFunc(s, func): This function returns the string s from which any leading or trailing character for which a custom function returns true are removed.",
		"trimleftFunc(s, func)":      "TrimLeftFunc(s, func): function returns the string s from which any leading character for which a custom function returns true are removed.",
		"trimrightFunc(s,    func)":  "TrimRightFunc(s, func): function returns the string s from which any trailing character for which a custom function returns true are removed.",
		"join(slice, sep)":           "Join(slice, sep): function combines the elements in the specified string slice, with the specified separator string placed between elements.",
		"repeat(s, count)":           "Repeat(s, count): function generates a string by repeating the string s for a specified number of times.",
		"for":                        "56.for: Go allows loops only inside of functions. The for keyword is used to create loops that repeatedly execute statements. The most basic for loops will repeat indefinitely unless interrupted by the break keyword Incorporating the Condition into the Loop",
		"package":                    "Go applications are organized in packages. A package is a collection of source files located in the same directory. All source files in a directory must share the same package name. When a package is imported, only entities (functions, types, variables, constants) whose names start with a capital letter can be used / accessed. The recommended style of naming in Go is that identifiers will be named using camelCase, except for those meant to be accessible across packages which should be PascalCase. EXAMPLE: || package lasagna || ",
		"packages":                   "Go applications are organized in packages. A package is a collection of source files located in the same directory. All source files in a directory must share the same package name. When a package is imported, only entities (functions, types, variables, constants) whose names start with a capital letter can be used / accessed. The recommended style of naming in Go is that identifiers will be named using camelCase, except for those meant to be accessible across packages which should be PascalCase. EXAMPLE: || package lasagna || ",
		"variable":                   "Variables Go is statically-typed, which means all variables must have a defined type at compile-time.Variables can be defined by explicitly specifying a type, EXAMPLE: var explicit int || Explicitly typed||",
		"variables":                  "Variables Go is statically-typed, which means all variables must have a defined type at compile-time.Variables can be defined by explicitly specifying a type, EXAMPLE: var explicit int || Explicitly typed||",
		"constant":                   "Constants hold a piece of data just like variables, but their value cannot change during the execution of the program.Constants are defined using the const keyword and can be numbers, characters, strings or booleans. EXAMPLE:	const Age = 21",
		"constants":                  "Constants hold a piece of data just like variables, but their value cannot change during the execution of the program.Constants are defined using the const keyword and can be numbers, characters, strings or booleans. EXAMPLE:	const Age = 21",
		"function":                   "Go functions accept zero or more parameters. Parameters must be explicitly typed, there is no type inference. Values are returned from functions using the return keyword. A function is invoked by specifying the function name and passing arguments for each of the function's parameters.",
		"functions":                  "Go functions accept zero or more parameters. Parameters must be explicitly typed, there is no type inference. Values are returned from functions using the return keyword. A function is invoked by specifying the function name and passing arguments for each of the function's parameters.",
		"comment":                    "Note that Go supports two types of comments. Single line comments are preceded by // and multiline comments are inserted between /* and */.",
		"comments":                   "Note that Go supports two types of comments. Single line comments are preceded by // and multiline comments are inserted between /* and */.",
		"bool":                       "Booleans in Go are represented by the predeclared boolean type bool, which values can be either true or false. It's a defined type.",
		"boolean":                    "Booleans in Go are represented by the predeclared boolean type bool, which values can be either true or false. It's a defined type.",
		"booleans":                   "Booleans in Go are represented by the predeclared boolean type bool, which values can be either true or false. It's a defined type.",
<<<<<<< HEAD
		"package comments":           "Package comments should be written directly before a package clause (package x) and begin with Package x ...",
		"package comment":            "Package comments should be written directly before a package clause (package x) and begin with Package x ...",
=======
		"function comment":           "A function comment should be written directly before the function declaration. It should be a full sentence that starts with the function name. For example, an exported comment for the function Calculate should take the form Calculate .... It should also explain what arguments the function takes, what it does with them, and what its return values mean, ending in a period):",
		"function comments":          "A function comment should be written directly before the function declaration. It should be a full sentence that starts with the function name. For example, an exported comment for the function Calculate should take the form Calculate .... It should also explain what arguments the function takes, what it does with them, and what its return values mean, ending in a period):",
<<<<<<< HEAD
<<<<<<< HEAD
>>>>>>> beforeMergeToMain
=======
		"number":                     "Go contains basic numeric types that can represent sets of either integer or floating-point values.",
		"numbers":                    "Go contains basic numeric types that can represent sets of either integer or floating-point values.",
>>>>>>> beforeMergeToMain
=======
		"number":                     "Go contains basic numeric types that can represent sets of either integer or floating-point values. Numbers can be converted to other numeric types through Type Conversion.",
		"numbers":                    "Go contains basic numeric types that can represent sets of either integer or floating-point values. Numbers can be converted to other numeric types through Type Conversion.",
		"int":                        "e.g. 0, 255, 2147483647. A signed integer that is at least 32 bits in size (value range of: -2147483648 through 2147483647). But this will depend on the systems architecture. Most modern computers are 64 bit, therefore int will be 64 bits in size (value rate of: -9223372036854775808 through 9223372036854775807).",
		"float64":                    "e.g. 0.0, 3.14. Contains the set of all 64-bit floating-point numbers.",
<<<<<<< HEAD
		"uint":                       "e.g. 0, 255. An unsigned integer that is the same size as int (value range of: 0 through 4294967295 for 32 bits and 0 through 18446744073709551615 for 64 bits)",
		"arithmetic operators":       "Go supports many standard arithmetic operators EXAMPLE: + , - , / , % ",
		"operator":                   "Go supports many standard arithmetic operators EXAMPLE: + , - , / , % ",
		"operators":                  "Go supports many standard arithmetic operators EXAMPLE: + , - , / , % ",
		"basic operator":             "Go supports many standard arithmetic operators EXAMPLE: + , - , / , % ",
		"basic operators":            "Go supports many standard arithmetic operators EXAMPLE: + , - , / , % ",
>>>>>>> beforeMergeToMain
=======
		"uint":                       "e.g. 0, 255. An unsigned integer that is the same size as int (value range of: 0 through 4294967295 for 32 bits and 0 through 18446744073709551615 for 64 bits). Go has shorthand assignment for the operators above (e.g. a += 5 is short for a = a + 5). Go also supports the increment and decrement statements ++ and -- (e.g. a++). For integer division, the remainder is dropped (e.g. 5 / 2 == 2).",
		"arithmetic operators":       "Go supports many standard arithmetic operators. EXAMPLE: + , - , / , % . For integer division, the remainder is dropped (e.g. 5 / 2 == 2).",
		"operator":                   "Go supports many standard arithmetic operators. EXAMPLE: + , - , / , % . For integer division, the remainder is dropped (e.g. 5 / 2 == 2).",
		"operators":                  "Go supports many standard arithmetic operators. EXAMPLE: + , - , / , % . For integer division, the remainder is dropped (e.g. 5 / 2 == 2).",
		"basic operator":             "Go supports many standard arithmetic operators. EXAMPLE: + , - , / , % . For integer division, the remainder is dropped (e.g. 5 / 2 == 2).",
		"basic operators":            "Go supports many standard arithmetic operators. EXAMPLE: + , - , / , % . For integer division, the remainder is dropped (e.g. 5 / 2 == 2).",
		"converting between types":   "Converting between types is done via a function with the name of the type to convert to. ",
		"";"",
>>>>>>> beforeMergeToMain
	},
}

func GetMapReturnSlice() []string {
	// Step 2: Initialize a slice to hold the keys
	var Keys []string

	// Step 3: Iterate over the map and append Keys to the slice
	for key := range OriginSingleDef.SingleDef {
		Keys = append(Keys, key)
	}

	return Keys
}

// ██████████████████████████████████████████████████████████████████████
