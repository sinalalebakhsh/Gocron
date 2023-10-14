package features

type Features struct {
	EveryWordsInGolang string
}

var TitleOfEveryWords = []string{
	"all",
	"-all",
	"--all",
}

var TitleOfAllIndexSlices = map[string]string{
	"READING AND WRITING DATA":     "READING AND WRITING DATA",
	"READINGANDWRITINGDATA":        "READINGANDWRITINGDATA",
	"ALL READING AND WRITING DATA": "ALL READING AND WRITING DATA",
	"ALLREADINGANDWRITINGDATA":     "ALLREADINGANDWRITINGDATA",
	"ALL REGEX":                    "ALL REGEX",
	"ALLREGEX":                     "ALLREGEX",
	"ALL TIME":                     "ALL TIME",
	"ALLTIME":                      "ALLTIME",
	"ALL HTML AND TEMPLATE":        "ALL HTML AND TEMPLATE",
	"ALLHTMLANDTEMPLATE":           "ALLHTMLANDTEMPLATE",
	"ALL WORKING WITH FILES":       "ALL WORKING WITH FILES",
	"ALLWORKINGWITHFILES":          "ALLWORKINGWITHFILES",
	"ALL JSON":                     "ALL JSON",
	"ALLJSON":                      "ALLJSON",
	"ALL JSON DATA":                "ALL JSON DATA",
	"ALLJSONDATA":                  "ALLJSONDATA",
	"ALL WORK WITH JSON DATA":      "ALL WORK WITH JSON DATA",
	"ALLWORKWITHJSONDATA":          "ALLWORKWITHJSONDATA",
	"ALL WORKING WITH JSON DATA":   "ALL WORKING WITH JSON DATA",
	"ALLWORKINGWITHJSONDATA":       "ALLWORKINGWITHJSONDATA",
	"ALL HTTP SERVERS":             "ALL HTTP SERVERS",
	"ALLHTTPSERVERS":               "ALLHTTPSERVERS",
	"ALL CREATING HTTP SERVERS":    "ALL CREATING HTTP SERVERS",
	"ALLCREATINGHTTPSERVERS":       "ALLCREATINGHTTPSERVERS",
	"ALL HTTP CLIENTS":             "ALL HTTP CLIENTS",
	"ALLHTTPCLIENTS":               "ALLHTTPCLIENTS",
	"ALL CREATING HTTP CLIENTS":    "ALL CREATING HTTP CLIENTS",
	"ALLCREATINGHTTPCLIENTS":       "ALLCREATINGHTTPCLIENTS",
}

var OriginalFeatures Features = Features{
	EveryWordsInGolang: `
████████████████████████████████████████████████████████████████████████
0.Create Environment GO
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
    func add(args ...int) int {}

████████████████████████████████████████████████████████████████████████
1.build
    Using the Go Command
    The go build command compiles the source code in the current directory 
    and generates an executable file.
████████████████████████████████████████████████████████████████████████
2.clean
    The go clean command removes the output produced by the go build command, 
    including the executable and any temporary files that were created during the build.
████████████████████████████████████████████████████████████████████████
3.doc
    The go doc command generates documentation from source code.
████████████████████████████████████████████████████████████████████████
4.fmt
    The go fmt command ensures consistent indentation and alignment in source code files.
    fmt = for printing in CLI
    fmt.Printf("%v %s %f", variable, string, float32)
████████████████████████████████████████████████████████████████████████
5.get
    The go get command downloads and installs external packages.
    flag
    flag package:
    Command-line flags are a common way to specify options for command-line programs. 
    For example, in wc -l the -l is a command-line flag.
    Go provides a flag package supporting basic command-line flag parsing.
    We'll use this package to implement our example command-line program.
████████████████████████████████████████████████████████████████████████
6.install
    The go install command downloads packages and is usually used to install tool packages.
████████████████████████████████████████████████████████████████████████
7.help
    The go help command displays help information for other Go features. The command go
    help build, for example, displays information about the build argument.
████████████████████████████████████████████████████████████████████████
8.mod
    The go mod command is used to create and manage a Go module.
████████████████████████████████████████████████████████████████████████
9.run
    The go run command builds and executes the source code in a specified folder without
    creating an executable output
████████████████████████████████████████████████████████████████████████
10.test
    The go test command executes unit tests
████████████████████████████████████████████████████████████████████████
11.version
    The go version command writes out the Go version number.
████████████████████████████████████████████████████████████████████████
12.vet
    The go vet command detects common problems in Go code
████████████████████████████████████████████████████████████████████████
13.print <expr>
    Useful Debugger State Commands
    This command evaluates an expression and displays the result. It can
    be used to display a value (print i) or perform a more complex test
    (print i > 0).
████████████████████████████████████████████████████████████████████████
14.set <variable> = <value>
    This command changes the value of the specified variable.
    این دستور مقدار متغیر مشخص شده را تغییر می دهد.
████████████████████████████████████████████████████████████████████████
15.locals
    This command prints the value of all local variables.
████████████████████████████████████████████████████████████████████████
16.whatis <expr>
    This command prints the type of the specified expression such as whatis
████████████████████████████████████████████████████████████████████████
17.continue
    Useful Debugger Commands for Controlling Execution
    This command resumes execution of the application.
████████████████████████████████████████████████████████████████████████
18.next
    This command moves to the next statement.
████████████████████████████████████████████████████████████████████████
19.step
    This command steps into the current statement.
████████████████████████████████████████████████████████████████████████
20.stepout
    This command steps out of the current statement.
████████████████████████████████████████████████████████████████████████
21.restart
    This command restarts the process. Use the continue command to begin execution.
████████████████████████████████████████████████████████████████████████
22.exit
    This command exits the debugger.
████████████████████████████████████████████████████████████████████████
23.int = integers
    Understanding the Basic Data Types
    This type represents a whole number, which can be positive or negative. The
    int type size is platform-dependent and will be either 32 or 64 bits. There are
    also integer types that have a specific size, such as int8, int16, int32, and
    int64, but the int type should be used unless you need a specific size.
    
    Literal Value Examples:
        20, -20. Values can also be expressed in hex (0x14), octal (0o24), and binary notation
        (0b0010100).
████████████████████████████████████████████████████████████████████████
24.unit
    uint8,  uint16,  uint32,  uint64,int8,  int16,  int32
    This type represents a positive whole number. The uint type size is platform-
    dependent and will be either 32 or 64 bits. There are also unsigned integer
    types that have a specific size, such as uint8, uint16, uint32, and uint64, but
    the uint type should be used unless you need a specific size.
    Literal Value Examples:
        There are no uint literals. All literal whole numbers are treated as int values.
████████████████████████████████████████████████████████████████████████
25.byte = uint8
    This type is an alias for uint8 and is typically used to represent a byte of data.
    byte = 8 bits, 
    1024bytes = 1 kilobyte
    1024 kilobytes = 1 megabyte
    Literal Value Examples:
        There are no byte literals. Bytes are typically expressed as integer literals (such as 101) or run
        literals ('e') since the byte type is an alias for the uint8 type.
████████████████████████████████████████████████████████████████████████
26.float32, float64
    Floating Point Numbers. e.g 3.14, 123.541, 100.4020401
    These types represent numbers with a fraction. These types allocate 32 or 64
    bits to store the value.
    Literal Value Examples:
        20.2, -20.2, 1.2e10, 1.2e-10. Values can also be expressed in hex notation (0x2p10), although
        the exponent is expressed in decimal digits.
████████████████████████████████████████████████████████████████████████
27.complex64, complex128
    These types represent numbers that have real and imaginary components.
    These types allocate 64 or 128 bits to store the value.
████████████████████████████████████████████████████████████████████████
28.bool 
    bolean
    This type represents a Boolean truth with the values true and false.

    Literal Value Examples:
        true, false.
████████████████████████████████████████████████████████████████████████
29.string
    This type represents a sequence of characters.
    Literal Value Examples:
        "Hello". Character sequences escaped with a backslash are interpreted if the value is enclosed
        in double quotes ("Hello\n"). Escape sequences are not interpreted if the value is enclosed in
        backquotes ('Hello\n').
████████████████████████████████████████████████████████████████████████
30.rune = int32
    This type represents a single Unicode code point. Unicode is complicated,
    but—loosely—this is the representation of a single character. The rune type is
    an alias for int32.
   
    Literal Value Examples:
        'A', '\n', '\u00A5', '¥'. Characters, glyphs, and escape sequences are enclosed in single
        quotes (the ' character).
████████████████████████████████████████████████████████████████████████
31.var = variable 
    Variables are defined using the var keyword, and, unlike constants, the value assigned to a variable can be
    changed
████████████████████████████████████████████████████████████████████████
32.const = constant
    Constants are basically variables whose values cannot be changed later.

    The Zero Values for the Basic Data Types
    Type    Zero Value
    ------  -----------
    int         0
    unit        0
    byte        0
    float64     0
    bool        false
    string      “” (the empty string)
    rune        0
████████████████████████████████████████████████████████████████████████
33.+, -, *, /, %
    Operations and Conversions
    These operators are used to perform arithmetic using numeric values.
████████████████████████████████████████████████████████████████████████
34.==, !=, <, <=, >, >=
    These operators compare two values.
████████████████████████████████████████████████████████████████████████
35., &&, !
    These are the logical operators, which are applied to bool values and return a bool value.
████████████████████████████████████████████████████████████████████████
36.=, :=
    These are the assignment operators. The standard assignment operator (=) is used to set
    the initial value when a constant or variable is defined, or to change the value assigned to
    a previously defined variable. The shorthand operator (:=) is used to define a variable and
    assign a value.
████████████████████████████████████████████████████████████████████████
37.-=, +=, ++, --
    These operators increment and decrement numeric values.
████████████████████████████████████████████████████████████████████████
38.&, , ^, &^, <<, >>
    These are the bitwise operators, which can be applied to integer values. These operators are
    not often required in mainstream development
    where the  operator is used to configure the Go logging features.

    The Arithmetic Operators
    Operator    Description
    -------     ---------------------
    +        This operator returns the sum of two operands.
    -        This operator returns the difference between two operands.
    *        This operator returns the product of two operands.
    /        This product returns the quotient of two operators.
    %        This product returns the remainder, which is similar to the modulo operator provided by
                other programming languages but can return negative values, as described in the “Using the
                Remainder Operator” section.

    The Comparison Operators
    Operator    Description
    -------     ---------------------
    ==       This operator returns true if the operands are equal.
    !=       This operator returns true if the operands are not equal.
    <        This operator returns true if the first operand is less than the second operand.
    >        This operator returns true if the first operand is greater than the second operand.
    <=       This operator returns true if the first operand is less than or equal to the second operand.
    >=       This operator returns true if the first operand is greater than or equal to the second
                operand.

    The Logical Operators
    Operator    Description
    -------     ---------------------
            This operator returns true if either operand is true. 
                If the first operand is true, then the second
                operand will not be evaluated.

    &&       This operator returns true if both operands are true. 
                If the first operand is false, then the
                second operand will not be evaluated.

    !        This operator is used with a single operand. 
                It returns true if the operand is false and false if
                the operand is true.
████████████████████████████████████████████████████████████████████████
39.Ceil(value)
    Converting Floating-Point Values to Integers
    Functions in the math Package for Converting Numeric Types
    This function returns the smallest integer that is greater than the specified floating-
    point value. The smallest integer that is greater than 27.1, for example, is 28.
████████████████████████████████████████████████████████████████████████
40.Floor(value)
    This function returns the largest integer that is less than the specified floating-point
    value. The largest integer that is less than 27.1, for example, is 28.
████████████████████████████████████████████████████████████████████████
41.Round(value)
    This function rounds the specified floating-point value to the nearest integer.
████████████████████████████████████████████████████████████████████████
42.RoundToEven(value)
    This function rounds the specified floating-point value to the nearest even integer.
████████████████████████████████████████████████████████████████████████
Parsing from Strings
Functions for Parsing Strings into Other Data Types
43.ParseBool(str)
    This function parses a string into a bool value. Recognized string values are "true",
    "false", "TRUE", "FALSE", "True", "False", "T", "F", "0", and "1".
████████████████████████████████████████████████████████████████████████
44.ParseFloat(str,size)
    This function parses a string into a floating-point value with the specified size, as
    described in the “Parsing Floating-Point Numbers” section.
████████████████████████████████████████████████████████████████████████
45.ParseInt(str,base, size)
    This function parses a string into an int64 with the specified base and size. Acceptable
    base values are 2 for binary, 8 for octal, 16 for hex, and 10.
████████████████████████████████████████████████████████████████████████
46.ParseUint(str,base, size)
    This function parses a string into an unsigned integer value with the specified base and size.
████████████████████████████████████████████████████████████████████████
47.Atoi(str)
    This function parses a string into a base 10 int and is equivalent to calling
    ParseInt(str, 10, 0)
████████████████████████████████████████████████████████████████████████
48.FormatBool(val)
    Formatting Values as Strings
    The strconv Functions for Converting Values into Strings
    This function returns the string true or false based on the value of the
    specified bool.
████████████████████████████████████████████████████████████████████████
49.FormatInt(val, base)
    This function returns a string representation of the specified int64 value,
    expressed in the specified base.
████████████████████████████████████████████████████████████████████████
50.FormatUint(val, base)
    This function returns a string representation of the specified uint64 value,
    expressed in the specified base.
████████████████████████████████████████████████████████████████████████
51.FormatFloat(val, format, precision, size) 
    This function returns a string representation of the specified float64 value,
    expressed using the specified format, precision, and size.
████████████████████████████████████████████████████████████████████████
52.Itoa(val)
    This function returns a string representation of the specified int value,
    expressed using base 10.

    Commonly Used Format Options for Floating-Point String Formatting
    Function
    f
        The floating-point value will be expressed in the form ±ddd.ddd without an exponent, such as
        49.95.

    e, E
        The floating-point value will be expressed in the form ±ddd.ddde±dd, such as 4.995e+01 or
        4.995E+01. The case of the letter denoting the exponent is determined by the case of the rune
        used as the formatting argument.

    g, G
        The floating-point value will be expressed using format e/E for large exponents or format f for
        smaller values.
████████████████████████████████████████████████████████████████████████
53.if
    Understanding Flow Control
    The if keyword is followed by the expression and then the group of statements to be executed,
    surrounded by braces
████████████████████████████████████████████████████████████████████████
54.else
    The else keyword can be used to create additional clauses in an if statement
████████████████████████████████████████████████████████████████████████
55.else if
    The else/if combination can be repeated to create a sequence of clauses
████████████████████████████████████████████████████████████████████████
56.for
    Go allows loops only inside of functions.
    The for keyword is used to create loops that repeatedly execute statements. The most basic for loops will
    repeat indefinitely unless interrupted by the break keyword
    Incorporating the Condition into the Loop
    
    example: 
        for (counter <= 3) {}
    
    Enumerating Sequences:
        for index, character := range product {
                    fmt.Println("Index:", index, "Character:", string(character))
                }
    
    range : is for variable like slice or array like this :
    for name := range sliceOfNames
████████████████████████████████████████████████████████████████████████
57.switch
    A switch statement provides an alternative way to control execution flow, based on matching the result of an
    expression to a specific value, as opposed to evaluating a true or false result
    example:
        switch (character) {
            case 'K':
                fmt.Println("K at position", index)
            case 'y':
                fmt.Println("y at position", index)
        }
    
    Matching Multiple Values:
    example:
        switch (character) {
            case 'K', 'k':
                fmt.Println("K or k at position", index)
            case 'y':
                fmt.Println("y at position", index)
        }
    
    Terminate case Statement Execution:
    example:
        switch (character) {
            case 'K', 'k':
        if (character == 'k') {
            fmt.Println("Lowercase k at position", index)
            break
        }
        fmt.Println("Uppercase K at position", index)
            case 'y':
        fmt.Println("y at position", index)
        }
████████████████████████████████████████████████████████████████████████
58.Falling Through
    Go switch statements don't automatically fall through, but this behavior can be enabled using the
    fallthrough keyword
    example:
        switch (character) {
            case 'K':
                fmt.Println("Uppercase character")
                fallthrough
            case 'k':
                fmt.Println("k at position", index)
            case 'y':
                fmt.Println("y at position", index)
            }
████████████████████████████████████████████████████████████████████████
59.default
    The default keyword is used to define a clause that will be executed when none of the case statements
    matches the switch statement's value
    
    example:
        switch (character) {
                case 'K', 'k':
                if (character == 'k') {
                    fmt.Println("Lowercase k at position", index)
                    break
                }
                fmt.Println("Uppercase K at position", index)
                case 'y':
                fmt.Println("y at position", index)
                default:
                fmt.Println("Character", string(character), "at position", index)
            }
████████████████████████████████████████████████████████████████████████
60.goto label 
    Label statements allow execution to jump to a different point, 
    giving greater flexibility than other flow control features
    example:
        counter := 0
        target: fmt.Println("Counter", counter)
        counter++
        if (counter < 5) {
                goto target
        }
████████████████████████████████████████████████████████████████████████
61.Array
    Go arrays are a fixed length and contain elements of a single type, which are accessed by index
    var names [3]string

    An array is a numbered sequence of elements of a sin-gle type with a fixed length.


    Array Literal Syntax:
        names := [3]string { "Kayak", "Lifejacket", "Paddle" }

    The number of elements specified with the literal syntax can be less than the capacity of the array. 
    Any position in the array for which a value is not provided will be assigned the zero value for the array type.

    Enumerating Arrays:
    example:
        names := [3]string { "Kayak", "Lifejacket", "Paddle" }
        for index, value := range names {
            fmt.Println("Index:", index, "Value:", value)
        }
████████████████████████████████████████████████████████████████████████
62.Slices
    The slice type in this example is []string
    The best way to think of slices is as a variable-length array because they are useful when you don't know how
    many values you need to store or when the number changes over time. One way to define a slice is to use the
    built-in make function.

    A slice is a segment of an array. Like arrays slices areindexable and have a length. 
    Unlike arrays this lengthis allowed to change.

    if you want to create a slice you should use the built-inmake function

    example:
        names := make([]string, 3)
        names[0] = "Kayak"
        names[1] = "Lifejacket"
        names[2] = "Paddle"
    
    Literal Syntax:
    example:
        names := []string {"Kayak", "Lifejacket", "Paddle"}
████████████████████████████████████████████████████████████████████████
63.append
    If you define a slice variable but don't initialize it, then the result is a slice that has a length of zero
    and a capacity of zero, and this will cause an error when an element is appended to it.
    One of the key advantages of slices is that they can be expanded to accommodate additional elements
    
    Appending Elements to a Slice:
    example:
        names := []string {"Kayak", "Lifejacket", "Paddle"}
        names = append(names, "Hat", "Gloves")

    Appending Items to a Slice:
    example:
        names := []string {"Kayak", "Lifejacket", "Paddle"}
        appendedNames := append(names, "Hat", "Gloves")
        names[0] = "Canoe"

    Allocating Additional Slice Capacity:
    example:
        names := make([]string, 3, 6)
        names[0] = "Kayak"
        names[1] = "Lifejacket"
        names[2] = "Paddle"

    Adding Elements to a Slice:
    example:
        names := make([]string, 3, 6)
        names[0] = "Kayak"
        names[1] = "Lifejacket"
        names[2] = "Paddle"
        appendedNames := append(names, "Hat", "Gloves")
        names[0] = "Canoe"
        fmt.Println("names:",names)
        fmt.Println("appendedNames:", appendedNames)
    

    Appending One Slice to Another:
    example:
        names := make([]string, 3, 6)
        names[0] = "Kayak"
        names[1] = "Lifejacket"
        names[2] = "Paddle"
        moreNames := []string { "Hat Gloves"}
        appendedNames := append(names, moreNames...)
        fmt.Println("appendedNames:", appendedNames)
    

    Creating Slices from Existing Arrays:
    example:
        products := [4]string { "Kayak", "Lifejacket", "Paddle", "Hat"}
        someNames := products[1:3]
        allNames := products[:]

    Specifying Capacity When Creating a Slice from an Array:
    example:
        products := [4]string { "Kayak", "Lifejacket", "Paddle", "Hat"}
        someNames := products[1:3:3]
        allNames := products[:]
        someNames = append(someNames, "Gloves")
    
    Creating Slices from Other Slices:
    example:
        products := [4]string { "Kayak", "Lifejacket", "Paddle", "Hat"}
        allNames := products[1:]
        someNames := allNames[1:3]
        allNames = append(allNames, "Gloves")
        allNames[1] = "Canoe"

    Comparing Slices:
    example:
        Go restricts the use of the comparison operator so that slices can be compared only to the nil value.
            p1 := []string { "Kayak", "Lifejacket", "Paddle", "Hat"}
            p2 := p1
            fmt.Println("Equal:", p1 == p2)
    Output:
        .\main.go:13:30: invalid operation: p1 == p2 (slice can only be compared to nil)
████████████████████████████████████████████████████████████████████████
64.copy
    The copy function is used to copy elements between slices. This function can be used to ensure that slices
    have separate arrays and to create slices that combine elements from different sources.

    Duplicating a Slice:
    The copy function accepts two arguments:
    example:
        products := [4]string { "Kayak", "Lifejacket", "Paddle", "Hat"}
        allNames := products[1:]
        someNames := make([]string, 2)
        copy(someNames, allNames)
    Output:
        allNames [Lifejacket Paddle Hat]
        someNames: [Lifejacket Paddle]

    the Uninitialized Slice Pitfall:
    the copy function doesn't resize the destination slice. A common
    pitfall is to try to copy elements into a slice that has not been initialized
    example:
        products := [4]string { "Kayak", "Lifejacket", "Paddle", "Hat"}
        allNames := products[1:]
        var someNames []string
        copy(someNames, allNames)
    Output:
        someNames: []
        allNames [Lifejacket Paddle Hat]

    Specifying Ranges When Copying Slices:
    Fine-grained control over the elements that are copied can be achieved using ranges
    example:
        products := [4]string { "Kayak", "Lifejacket", "Paddle", "Hat"}
        allNames := products[1:]
        someNames := []string { "Boots", "Canoe"}
        copy(someNames[1:], allNames[2:3])
    Output:
        allNames [Lifejacket Paddle Hat]
        someNames: [Boots Hat]
    
    Copying Slices with Different Sizes:
        products := []string { "Kayak", "Lifejacket", "Paddle", "Hat"}
        replacementProducts := []string { "Canoe", "Boots"}
        copy(products, replacementProducts)
    Output:
        products: [Canoe Boots Paddle Hat]

    Copying a Larger Source Slice
    example:
        products := []string { "Kayak", "Lifejacket", "Paddle", "Hat"}
            replacementProducts := []string { "Canoe", "Boots"}
            copy(products[0:1], replacementProducts)
            fmt.Println("products:", products)
    Output:
        products: [Canoe Lifejacket Paddle Hat]
    Deleting Slice Elements:
        products := [4]string { "Kayak", "Lifejacket", "Paddle", "Hat"}
        deleted := append(products[:2], products[3:]...)
        fmt.Println("Deleted:", deleted)
    Enumerating Slices:
    example:
        products := []string { "Kayak", "Lifejacket", "Paddle", "Hat"}
        for index, value := range products[2:] {
            fmt.Println("Index:", index, "Value:", value)
        }
    Output:
        Index: 0 Value: Paddle
        Index: 1 Value: Hat
████████████████████████████████████████████████████████████████████████
65.sort
    There is no built-in support for sorting slices, but the standard library includes the sort package, which
    defines functions for sorting different types of slice.

    example:
        products := []string { "Kayak", "Lifejacket", "Paddle", "Hat"}
        sort.Strings(products)
        for index, value := range products {
            fmt.Println("Index:", index, "Value:", value)
        }
    Output:
        Index: 0 Value: Hat
        Index: 1 Value: Kayak
        Index: 2 Value: Lifejacket
        Index: 3 Value: Paddle
████████████████████████████████████████████████████████████████████████
66.DeepEqual
    The DeepEqual function can be used to compare a wider range of data types than the
    equality operator, including slices.

    example:
        package main
        import (
            "fmt"
            "reflect"
        )
        func main() {
            p1 := []string { "Kayak", "Lifejacket", "Paddle", "Hat"}
            p2 := p1
            fmt.Println("Equal:", reflect.DeepEqual(p1, p2))
        }    
    Output:
        Equal: true
████████████████████████████████████████████████████████████████████████
67.map
    Maps are a built-in data structure that associates data values with keys. 
    Unlike arrays, where values are associated with sequential integer locations, 
    maps can use other data types as keys.

    A map is an unordered collection of key-value pairs.
    Also known as an associative array, a hash table or dictionary, 
    maps are used to look up a value by its associated key.

    example:
        products := make(map[string]float64, 10)
        products["Kayak"] = 279
        products["Lifejacket"] = 48.95
    
    Map Literal Syntax:
        products := map[string]float64 {
            "Kayak" : 279,
            "Lifejacket": 48.95,
        }
    Checking for Items in a Map:
    example:
        products := map[string]float64 {
                    "Kayak" : 279,
                    "Lifejacket": 48.95,
                    "Hat": 0,
                }
                value, ok := products["Hat"]
                if (ok) {
                    fmt.Println("Stored value:", value)
                } else {
                    fmt.Println("No stored value")
                }

    Removing Items from a Map:
    example:
        products := map[string]float64 {
            "Kayak" : 279,
            "Lifejacket": 48.95,
            "Hat": 0,
        }
        delete(products, "Hat")

    Enumerating the Contents of a Map:
    example:
        products := map[string]float64 {
                "Kayak" : 279,
                "Lifejacket": 48.95,
                "Hat": 0,
        }
        for key, value := range products {
            fmt.Println("Key:", key, "Value:", value)
        }

    Enumerating a Map in Order:
    example:
        import (
                "fmt"
                "sort"
            )
        func main() {
            products := map[string]float64 {
                "Kayak" : 279,
                "Lifejacket": 48.95,
                "Hat": 0,
            }
            keys := make([]string, 0, len(products))
            for key,  := range products {
                keys = append(keys, key)
            }
            sort.Strings(keys)
            for , key := range keys {
                fmt.Println("Key:", key, "Value:", products[key])
            }
        }
████████████████████████████████████████████████████████████████████████
68.strconv
    Understanding the Dual Nature of Strings
    Indexing and Slicing a String:
        var price string = "$48.95"
        var currency byte = price[0]
        var amountString string = price[1:]
        amount, parseErr := strconv.ParseFloat(amountString, 64)
        fmt.Println("Currency:", currency)
        if (parseErr == nil) {
                fmt.Println("Amount:", amount)
        } else {
                fmt.Println("Parse Error:", parseErr)
        }
    Output:
        Currency: 36
        Amount: 48.95

    Converting the Result:
        var price string = "$48.95"
        var currency string = string(price[0])
        var amountString string = price[1:]
        amount, parseErr  := strconv.ParseFloat(amountString, 64)
        fmt.Println("Currency:", currency)
        if (parseErr == nil) {
            fmt.Println("Amount:", amount)
        } else {
            fmt.Println("Parse Error:", parseErr)
        }   
    Output:
        Currency: $
        Amount: 48.95
    Changing the Currency Symbol:
    example:
       var price string = "€48.95"
       var currency string = string(price[0])
       var amountString string = price[1:]
       amount, parseErr  := strconv.ParseFloat(amountString, 64)
       fmt.Println("Currency:", currency)
       if (parseErr == nil) {
           fmt.Println("Amount:", amount)
       } else {
           fmt.Println("Parse Error:", parseErr)
       }
    Output:
        Currency: â
        Parse Error: strconv.ParseFloat: parsing "\x82\xac48.95": invalid syntax
    
    Obtaining the Length:
        fmt.Println("Length:", len(price))
    Output:
        Length: 8
████████████████████████████████████████████████████████████████████████
69.rune
    Converting a String to Runes
    The rune type represents a Unicode code point, 
    which is essentially a single character. 
    To avoid slicing strings in the middle of characters, 
    an explicit conversion to a rune slice can be performed
    Unicode is incredibly complex, as you would expect 
    from any standard that aims to describe multiple
    writing systems evolved over thousands of years.
    Converting to Runes:
        var price []rune = []rune("€48.95")
        var amountString string = string(price[1:])

    Enumerating Strings        
    example:
        var price = "€48.95"
        for index, char := range price {
            fmt.Println(index, char, string(char))
        }
    Output:
        0 8364 €
        3 52 4
        4 56 8
        5 46 .
        6 57 9
        7 53 5
    
    Enumerating the Bytes
    example:
        var price = "€48.95"
        for index, char := range []byte(price) {
            fmt.Println(index, char)
        }
    Output:
        0 226
        1 130
        2 172
        3 52
        4 56
        5 46
        6 57
        7 53
████████████████████████████████████████████████████████████████████████
70.func
    Functions are groups of code statements that are executed only when the function is
    invoked during the flow of execution.
████████████████████████████████████████████████████████████████████████
71.Function Parameters
    Parameters allow a function to receive data values when it is called, 
    allowing its behavior to be altered.
    Values for parameters are supplied as arguments when invoking the function, meaning that different
    values can be provided each time the function is called. Arguments are provided between the parentheses
    that follow the function name, separated by commas and in the same order in which the parameters have
    been defined
    example:
        func printPrice(product string, price float64, taxRate float64) {
            taxAmount := price * taxRate
            fmt.Println(product, "price:", price, "Tax:", taxAmount)
        }
████████████████████████████████████████████████████████████████████████
72.Defining Variadic Parameters
    example:
        func printSuppliers(product string, suppliers []string ) {
            for , supplier := range suppliers {
                fmt.Println("Product:", product, "Supplier:", supplier)
            }
        }
    example:
        func printSuppliers(product string, suppliers ...string ) {
            for , supplier := range suppliers {
                fmt.Println("Product:", product, "Supplier:", supplier)
            }
        }
████████████████████████████████████████████████████████████████████████
73.Dealing with No Arguments for a Variadic Parameter
    example:
        func printSuppliers(product string, suppliers ...string ) {
            for , supplier := range suppliers {
                fmt.Println("Product:", product, "Supplier:", supplier)
            }
        }
        func main() {
            printSuppliers("Kayak", "Acme Kayaks", "Bob's Boats", "Crazy Canoes")
            printSuppliers("Lifejacket", "Sail Safe Co")
            printSuppliers("Soccer Ball")
        }
    Output:
        Product: Kayak Supplier: Acme Kayaks
        Product: Kayak Supplier: Bob's Boats
        Product: Kayak Supplier: Crazy Canoes
        Product: Lifejacket Supplier: Sail Safe Co
████████████████████████████████████████████████████████████████████████
74.return Function Results
    example:
        func calcTax(price float64) float64 {
            return price + (price * 0.2)
        }
████████████████████████████████████████████████████████████████████████
75.Returning Multiple Function Results
    example:
        func swapValues(first, second int) (int, int) {
            return second, first
        }
        func main() {
            val1, val2 := 10, 20
            fmt.Println("Before calling function", val1, val2)
            val1, val2 = swapValues(val1, val2)
            fmt.Println("After calling function", val1, val2)
        }
████████████████████████████████████████████████████████████████████████
76.Using Named Results
    example:
        func calcTax(price float64) (float64, bool) {
            if (price > 100) {
                return price * 0.2, true
            }
            return 0, false
        }
        func calcTotalPrice(products map[string]float64,
                minSpend float64) (total, tax float64)  {
            total = minSpend
            for , price := range products {
                if taxAmount, due := calcTax(price); due {
                    total += taxAmount;
                    tax += taxAmount
                } else {
                    total += price
                }
            }
            return
        }
████████████████████████████████████████████████████████████████████████
77.defer
    The defer keyword is used to schedule a function call that will be performed immediately before the current
    function returns
    The defer keyword lets you group the statements that create, use, and
    release the resource together.
    The defer keyword can be used with any function call
    a single function can use the defer keyword multiple times.
    Immediately before the function returns, Go will perform the
    calls scheduled with the defer keyword in the order in which they were defined.
████████████████████████████████████████████████████████████████████████
78.Function Types
    Functions in Go have a data type, which describes the combination of parameters the
    function consumes and the results the function produces. This type can be specified
    explicitly or inferred from a function defined using a literal syntax.
    Function types are defined using the func keyword, followed by a signature that
    describes the parameters and results. No function body is specified.

    Go does not support arrow functions, where functions are expressed more concisely using the =>
    operator, without the func keyword and a code block surrounded by braces. In Go, functions must always be
    defined with the keyword and a body.
████████████████████████████████████████████████████████████████████████
79.Function Comparisons and the Zero Type
    example:
        func calcWithTax(price float64) float64 {
            return price + (price * 0.2)
        }
        func calcWithoutTax(price float64) float64 {
            return price
        }
        func main() {
            products := map[string]float64 {
                "Kayak" : 275,
                "Lifejacket": 48.95,
            }
            for product, price := range products {
                var calcFunc func(float64) float64
                fmt.Println("Function assigned:", calcFunc == nil)
                if (price > 100) {
                    calcFunc = calcWithTax
                } else {
                    calcFunc = calcWithoutTax
                }
                fmt.Println("Function assigned:", calcFunc == nil)
                totalPrice := calcFunc(price)
                fmt.Println("Product:", product, "Price:", totalPrice)
            }
        }
████████████████████████████████████████████████████████████████████████
80.Functions as Arguments
    example:
        func calcWithTax(price float64) float64 {
            return price + (price * 0.2)
        }
        func calcWithoutTax(price float64) float64 {
            return price
        }
        func printPrice(product string, price float64, calculator func(float64) float64 ) {
            fmt.Println("Product:", product, "Price:", calculator(price))
        }
        func main() {
            products := map[string]float64 {
                "Kayak" : 275,
                "Lifejacket": 48.95,
            }
            for product, price := range products {
                if (price > 100) {
                    printPrice(product, price, calcWithTax)
                } else {
                    printPrice(product, price, calcWithoutTax)
                }
            }
        }
████████████████████████████████████████████████████████████████████████
81.Functions as Results
    example:
        func calcWithTax(price float64) float64 {
            return price + (price * 0.2)
        }
        func calcWithoutTax(price float64) float64 {
            return price
        }
        func printPrice(product string, price float64, calculator func(float64) float64 ) {
            fmt.Println("Product:", product, "Price:", calculator(price))
        }
        func selectCalculator(price float64) func(float64) float64 {
            if (price > 100) {
                return calcWithTax
            }
            return calcWithoutTax
        }
        func main() {
            products := map[string]float64 {
                "Kayak" : 275,
                "Lifejacket": 48.95,
            }
            for product, price := range products {
                printPrice(product, price, selectCalculator(price))
            }
        }
████████████████████████████████████████████████████████████████████████
82.Function Type Aliases
    Go supports type aliases, which can be used to assign a name to
    a function signature so that the parameter and result types are not specified every time the function type is
    used.
    example:
        type calcFunc func(float64) float64
        func calcWithTax(price float64) float64 {
            return price + (price * 0.2)
        }
        func calcWithoutTax(price float64) float64 {
            return price
        }
        func printPrice(product string, price float64, calculator calcFunc) {
            fmt.Println("Product:", product, "Price:", calculator(price))
        }
        func selectCalculator(price float64) calcFunc {
            if (price > 100) {
                return calcWithTax
            }
            return calcWithoutTax
        }
        func main() {
            products := map[string]float64 {
                "Kayak" : 275,
                "Lifejacket": 48.95,
            }
            for product, price := range products {
                printPrice(product, price, selectCalculator(price))
            }
        }
████████████████████████████████████████████████████████████████████████
83.the Literal Function Syntax
    example:
        func selectCalculator(price float64) calcFunc {
            if (price > 100) {
                var withTax calcFunc = func (price float64) float64 {
                    return price + (price * 0.2)
                }
                return withTax
            }
████████████████████████████████████████████████████████████████████████
84.Function Variable Scope
    example:
        func selectCalculator(price float64) calcFunc {
            if (price > 100) {
                var withTax calcFunc = func (price float64) float64 {
                    return price + (price * 0.2)
                }
                return withTax
            } else if (price < 10) {
                return withTax
            }
             withoutTax := func (price float64) float64 {
                    return price
                }
            return withoutTax
        }
████████████████████████████████████████████████████████████████████████
85.Functions Values Directly
    example:
        func selectCalculator(price float64) calcFunc {
            if (price > 100) {
                return func (price float64) float64 {
                    return price + (price * 0.2)
                }
            }
             return func (price float64) float64 {
                return price
            }
        }
████████████████████████████████████████████████████████████████████████
86.Literal Function Argument
    example:
        func main() {
            products := map[string]float64 {
                "Kayak" : 275,
                "Lifejacket": 48.95,
            }
            for product, price := range products {
                printPrice(product, price, func (price float64) float64 {
                    return price + (price * 0.2)
                })
            }
        }
████████████████████████████████████████████████████████████████████████
87.Function Closure
    Functions defined using the literal syntax can reference variables from the surrounding code, a feature
    known as closure.
    example:
        type calcFunc func(float64) float64
        func printPrice(product string, price float64, calculator calcFunc) {
            fmt.Println("Product:", product, "Price:", calculator(price))
        }
        func main() {
            watersportsProducts := map[string]float64 {
                "Kayak" : 275,
                "Lifejacket": 48.95,
            }
            soccerProducts := map[string] float64 {
                "Soccer Ball": 19.50,
                "Stadium": 79500,
            }
            calc := func(price float64) float64 {
                if (price > 100) {
                    return price + (price * 0.2)
                }
                return price;
            }
            for product, price := range watersportsProducts {
                printPrice(product, price, calc)
            }
            calc = func(price float64) float64 {
                if (price > 50) {
                    return price + (price * 0.1)
                }
                return price
            }
            for product, price := range soccerProducts {
                printPrice(product, price, calc)
            }
        }
████████████████████████████████████████████████████████████████████████    
88.struct
    What are they?
    Structs are data types, comprised of fields.
    
    Why are they useful?
    Structs allow custom data types to be defined.

    How are they used?
    The type and struct keywords are used to define a type, 
    allowing field names and types to be specified.
    A struct can mix regular and embedded field types.
    example:
        func main() {
            type Product struct {
                name, category string
                price float64
            }
            kayak := Product {
                name: "Kayak",
                category: "Watersports",
                price: 275,
            }
            fmt.Println(kayak.name, kayak.category, kayak.price)
            kayak.price = 300
            fmt.Println("Changed price:", kayak.price)
        }
    Go doesn't differentiate between structs and classes, in the way that other languages do. All custom
    data types are defined as structs, and the decision to pass them by reference or by value is made
    depending on whether a pointer is used.

    Go doesn't allow structs to be used with the const keyword, and the compiler will report an error if
    you try to define a constant struct.
████████████████████████████████████████████████████████████████████████
89.struct tag
    The struct type can be defined with tags, which provide additional information about how a field should
    be processed. Struct tags are just strings that are interpreted by the code that processes struct values,
    using the features provided by the reflect package.
████████████████████████████████████████████████████████████████████████
90.struct values
    Values do not have to be provided for all fields when creating a struct value
    
    example:
        func main() {
            type Product struct {
                name, category string
                price float64
            }
            kayak := Product {
                    name: "Kayak",
                    category: "Watersports",
            }
        }
████████████████████████████████████████████████████████████████████████
91.new 
    the new function to create struct values
        var lifejacket = new(Product)

    The result is a pointer to a struct value whose fields are 
    initialized with their type's zero value. 
    This is equivalent to this statement:
        var lifejacket = &Product{}

    These approaches are interchangeable, and choosing between them is a matter of preference.
████████████████████████████████████████████████████████████████████████
92.Field Positions to Create Struct Values
    example:
        func main() {
            type Product struct {
                name, category string
                price float64
            }
            var kayak = Product { "Kayak", "Watersports", 275.00 }
            fmt.Println("Name:", kayak.name)
            fmt.Println("Category:", kayak.category)
            fmt.Println("Price:", kayak.price)
        }
████████████████████████████████████████████████████████████████████████
93.Defining Embedded Fields
    If a field is defined without a name, 
    it is known as an embedded field, and it is accessed using the name of its type.
    example:
        func main() {
                type Product struct {
                    name, category string
                    price float64
                }
                type StockLevel struct {
                    Product
                    count int
                }
                stockItem := StockLevel {
                    Product: Product { "Kayak", "Watersports", 275.00 },
                    count: 100,
                }
                fmt.Println("Name:", stockItem.Product.name)
                fmt.Println("Count:", stockItem.count)
        }
████████████████████████████████████████████████████████████████████████
94.Defining an Additional Field
    example:
        func main() {
            type Product struct {
                name, category string
                price float64
            }
            type StockLevel struct {
                Product
                Alternate Product
                count int
            }
            stockItem := StockLevel {
                Product: Product { "Kayak", "Watersports", 275.00 },
                Alternate: Product{"Lifejacket", "Watersports", 48.95 },
                count: 100,
            }
            fmt.Println("Name:", stockItem.Product.name)
            fmt.Println("Alt Name:", stockItem.Alternate.name)
        }
████████████████████████████████████████████████████████████████████████
95.Comparing Struct Values
    example:
        func main() {
            type Product struct {
                name, category string
                price float64
            }
            p1 := Product { name: "Kayak", category: "Watersports", price: 275.00 }
            p2 := Product { name: "Kayak", category: "Watersports", price: 275.00 }
            p3 := Product { name: "Kayak", category: "Boats", price: 275.00 }
            fmt.Println("p1 == p2:", p1 == p2)
            fmt.Println("p1 == p3:", p1 == p3)
        }
████████████████████████████████████████████████████████████████████████
96.Anonymous Struct Types
    Anonymous struct types are defined without using a name
    example:
        package main
        import "fmt"
        func writeName(val struct {
                name, category string
                price float64}) {
            fmt.Println("Name:", val.name)
        }
        func main() {
            type Product struct {
                name, category string
                price float64
                //otherNames []string
            }
            type Item struct {
                name string
                category string
                price float64
            }
            prod := Product { name: "Kayak", category: "Watersports", price: 275.00 }
            item := Item { name: "Stadium", category: "Soccer", price: 75000 }
            writeName(prod)
            writeName(item)
        }
████████████████████████████████████████████████████████████████████████
97.Creating Arrays, Slices, and Maps Containing Struct Values
    Omitting the Struct Type
    example:
        package main
        import "fmt"
        func main() {
            type Product struct {
                name, category string
                price float64
                //otherNames []string
            }
            type StockLevel struct {
                Product
                Alternate Product
                count int
            }
            array := [1]StockLevel {
                {
                    Product: Product { "Kayak", "Watersports", 275.00 },
                    Alternate: Product{"Lifejacket", "Watersports", 48.95 },
                    count: 100,
                },
            }
            fmt.Println("Array:", array[0].Product.name)
            slice := []StockLevel {
                {
                    Product: Product { "Kayak", "Watersports", 275.00 },
                    Alternate: Product{"Lifejacket", "Watersports", 48.95 },
                    count: 100,
                },
            }
            fmt.Println("Slice:", slice[0].Product.name)
            kvp := map[string]StockLevel {
                "kayak": {
                    Product: Product { "Kayak", "Watersports", 275.00 },
                    Alternate: Product{"Lifejacket", "Watersports", 48.95 },
                    count: 100,
                },
            }
            fmt.Println("Map:", kvp["kayak"].Product.name)
        }
████████████████████████████████████████████████████████████████████████
98.Structs and Pointers
    Assigning a struct to a new variable or using a struct as a function parameter 
    creates a new value that copies the field values.
████████████████████████████████████████████████████████████████████████
99.Copying a Struct Value
    example:
        package main
        import "fmt"
        func main() {
            type Product struct {
                name, category string
                price float64
            }
            p1 := Product {
                name: "Kayak",
                category: "Watersports",
                price: 275,
            }
            p2 := p1
            p1.name = "Original Kayak"
            fmt.Println("P1:", p1.name)
            fmt.Println("P2:", p2.name)
        }
████████████████████████████████████████████████████████████████████████
100.Using a Pointer to a Struct
    example:
        package main
        import "fmt"
        func main() {
            type Product struct {
                name, category string
                price float64
            }
            p1 := Product {
                name: "Kayak",
                category: "Watersports",
                price: 275,
            }
            p2 := &p1
            p1.name = "Original Kayak"
            fmt.Println("P1:", p1.name)
            fmt.Println("P2:", (*p2).name)
        }
████████████████████████████████████████████████████████████████████████
101.the Struct Pointer Convenience Syntax
    example:
        package main
        import "fmt"
        
        type Product struct {
            name, category string
            price float64
        }
        
        func calcTax(product *Product) {
            if ((*product).price > 100) {
                (*product).price += (*product).price * 0.2
            }
        }
        
        func main() {
            kayak := Product {
                name: "Kayak",
                category: "Watersports",
                price: 275,
            }
            calcTax(&kayak)
            fmt.Println("Name:", kayak.name, "Category:",
            kayak.category, "Price", kayak.price)
        }
████████████████████████████████████████████████████████████████████████
102.Struct Constructor Functions
    A constructor function is responsible for creating struct values using values received through parameters
    Constructor functions are used to create struct values consistently. Constructor functions are usually
    named new or New followed by the struct type so that the constructor function for creating Product values
    is named newProduct.
    example:
        package main
        import "fmt"
        type Product struct {
            name, category string
            price float64
        }
        
        func newProduct(name, category string, price float64) *Product {
            return &Product{name, category, price}
        }
        
        func main() {
            products := [2]*Product {
                newProduct("Kayak", "Watersports", 275),
                newProduct("Hat", "Skiing", 42.50),
            }
            for , p := range products {
                fmt.Println("Name:", p.name, "Category:",  p.category, "Price", p.price)
            }
        }
████████████████████████████████████████████████████████████████████████
103.Modifying a Constructor
    The benefit of using constructor functions is consistency, 
    ensuring that changes to the construction
    process are reflected in all the struct values created by the function.

    example:
        func newProduct(name, category string, price float64) *Product {
            return &Product{name, category, price - 10}
        }
████████████████████████████████████████████████████████████████████████  
104.Pointer Types for Struct Fields
    example:
        package main
        import "fmt"
        type Product struct {
            name, category string
            price float64
            *Supplier
        }
        type Supplier struct {
            name, city string
        }
        func newProduct(name, category string, price float64, supplier *Supplier) *Product {
            return &Product{name, category, price -10, supplier}
        }
        func main() {
            acme := &Supplier { "Acme Co", "New York"}
            products := [2]*Product {
                newProduct("Kayak", "Watersports", 275, acme),
                newProduct("Hat", "Skiing", 42.50, acme),
            }
            for , p := range products {
                fmt.Println("Name:", p.name, "Supplier:",
                    p.Supplier.name, p.Supplier.city)
            }
        }
████████████████████████████████████████████████████████████████████████
105.Pointer Field Copying
    Care must be taken when copying structs to consider the effect on pointer fields    
    example:
        package main
        import "fmt"
        type Product struct {
            name, category string
            price float64
            *Supplier
        }
        type Supplier struct {
            name, city string
        }
        func newProduct(name, category string, price float64, supplier *Supplier) *Product {
            return &Product{name, category, price -10, supplier}
        }
        func main() {
            acme := &Supplier { "Acme Co", "New York"}
            p1 := newProduct("Kayak", "Watersports", 275, acme)
            p2 := *p1
            p1.name = "Original Kayak"
            p1.Supplier.name = "BoatCo"
            for , p := range []Product { *p1, p2 } {
                fmt.Println("Name:", p.name, "Supplier:",
                    p.Supplier.name, p.Supplier.city)
            }
        }
████████████████████████████████████████████████████████████████████████
106.Method
    What are they?
    Methods are functions that are invoked on a struct and have access to all of the
    fields defined by the value's type. Interfaces define sets of methods, which can be
    implemented by struct types.

    Why are they useful?
    These features allow types to be mixed and used through their common
    characteristics.

    How are they used?
    Methods are defined using the func keyword, but with the addition of a receiver.
    Interfaces are defined using the type and interface keywords.

    Are there any pitfalls or limitations?
    Careful use of pointers is important when creating methods, and care must be taken
    when using interfaces to avoid problems with the underlying dynamic types.
████████████████████████████████████████████████████████████████████████
107.Defining and Using Method
    example:
        package main
        import "fmt"
        type Product struct {
            name, category string
            price float64
        }
        func newProduct(name, category string, price float64) *Product {
                return &Product{ name, category, price }
            }
            func (product *Product) printDetails() {
                fmt.Println("Name:", product.name, "Category:", product.category,
                    "Price", product.price)
            }
        func main() {
            products := []*Product {
                newProduct("Kayak", "Watersports", 275),
                newProduct("Lifejacket", "Watersports", 48.95),
                newProduct("Soccer Ball", "Soccer", 19.50),
            }
            for , p := range products {
                p.printDetails()
            }
        }
████████████████████████████████████████████████████████████████████████
108.Defining Method Parameters and Results
    example:
        package main
        import "fmt"
        type Product struct {
            name, category string
            price float64
        }
        func newProduct(name, category string, price float64) *Product {
                return &Product{ name, category, price }
            }
            func (product *Product) printDetails() {
                fmt.Println("Name:", product.name, "Category:", product.category,
                    "Price", product.price)
            }
        func main() {
            products := []*Product {
                newProduct("Kayak", "Watersports", 275),
                newProduct("Lifejacket", "Watersports", 48.95),
                newProduct("Soccer Ball", "Soccer", 19.50),
            }
            for , p := range products {
                p.calcTax(0.2, 100) //<-------------------------------here
            }
        }
████████████████████████████████████████████████████████████████████████
109.Defining and Using Interfaces
    One interface can enclose another, with the effect that types must implement all the methods defined
    by the enclosing and enclosed interfaces. Interfaces are simpler than structs, and there are no fields or
    method to promote. The result of composing interfaces is a union of the method defined by the enclosing
    and enclosed types.
    
    example:
        package main
        import "fmt"
        type Expense interface {
            getName() string
            getCost(annual bool) float64
        }
        func main() {
            expenses := []Expense {
                Product { "Kayak", "Watersports", 275 },
                Service {"Boat Cover", 12, 89.50 },
            }
            for , expense := range expenses {
                fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
            }
        }    
████████████████████████████████████████████████████████████████████████
110.an Interface in a Function
    example:
        package main
        import "fmt"
        type Expense interface {
            getName() string
            getCost(annual bool) float64
        }
        func calcTotal(expenses []Expense) (total float64) {
            for , item := range expenses {
                total += item.getCost(true)
            }
            return
        }
        func main() {
            expenses := []Expense {
                Product { "Kayak", "Watersports", 275 },
                Service {"Boat Cover", 12, 89.50 },
            }
            for , expense := range expenses {
                fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
            }
            fmt.Println("Total:", calcTotal(expenses))
        }
████████████████████████████████████████████████████████████████████████
111.an Interface for Struct Fields
    example:
        package main
        import "fmt"
        type Expense interface {
            getName() string
            getCost(annual bool) float64
        }
        func calcTotal(expenses []Expense) (total float64) {
            for , item := range expenses {
                total += item.getCost(true)
            }
            return
        }
        type Account struct {
            accountNumber int
            expenses []Expense
        }
        func main() {
            account := Account {
                accountNumber: 12345,
                expenses: []Expense {
                    Product { "Kayak", "Watersports", 275 },
                    Service {"Boat Cover", 12, 89.50 },
                },
            }
            for , expense := range account.expenses {
                fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
            }
            fmt.Println("Total:", calcTotal(account.expenses))
        }    
████████████████████████████████████████████████████████████████████████
112.Comparing Interface Values
    Care must be taken when comparing interface values, and inevitably, some knowledge of the dynamic
    types is required.
    The first two Expense values are not equal. 
    That's because the dynamic type for these values is a pointer
    type, and pointers are equal only if they point to the same memory location
    
    example:
        package main
        import "fmt"
        type Expense interface {
            getName() string
            getCost(annual bool) float64
        }
        func main() {
            var e1 Expense = &Product { name: "Kayak" }
            var e2 Expense = &Product { name: "Kayak" }
            var e3 Expense = Service { description: "Boat Cover" }
            var e4 Expense = Service { description: "Boat Cover" }
            fmt.Println("e1 == e2", e1 == e2)
            fmt.Println("e3 == e4", e3 == e4)
        }
    Output:
        e1 == e2 false
        e3 == e4 true
████████████████████████████████████████████████████████████████████████
113.Empty Interface
    Go allows the user of the empty interface—which means an interface that defines no methods—to represent
    any type, which can be a useful way to group disparate types that share no common features
    The empty interface represents all types, including the built-in types and any structs and interfaces
    that have been defined.
████████████████████████████████████████████████████████████████████████
114.Empty Interface for Function Parameters
    The empty interface can be used as the type for a function parameter, allowing a function to be called with
    any value
    
    example:
        package main
        import "fmt"
        type Expense interface {
            getName() string
            getCost(annual bool) float64
        }
        type Person struct {
            name, city string
        }
        func processItem(item interface{}) {
                switch value := item.(type) {
                    case Product:
                        fmt.Println("Product:", value.name, "Price:", value.price)
                    case *Product:
                        fmt.Println("Product Pointer:", value.name, "Price:", value.price)
                    case Service:
                        fmt.Println("Service:", value.description, "Price:",
                            value.monthlyFee * float64(value.durationMonths))
                    case Person:
                        fmt.Println("Person:", value.name, "City:", value.city)
                    case *Person:
                        fmt.Println("Person Pointer:", value.name, "City:", value.city)
                    case string, bool, int:
                        fmt.Println("Built-in type:", value)
                    default:
                        fmt.Println("Default:", value)
                }
            }
        func main() {
            var expense Expense = &Product { "Kayak", "Watersports", 275 }
            data := []interface{} {
                expense,
                Product { "Lifejacket", "Watersports", 48.95 },
                Service {"Boat Cover", 12, 89.50, []string{} },
                Person { "Alice", "London"},
                &Person { "Bob", "New York"},
                "This is a string",
                100,
                true,
            }
            for , item := range data {
                processItem(item)
            }
        }
████████████████████████████████████████████████████████████████████████
115.Package
    Packages are the Go feature that allows projects to be structured so that related functionality can be grouped
    together, without the need to put all the code into a single file or folder.
    
    What are they?
    Packages allow projects to be structured so that related features can be developed together.

    Why are they useful?
    Packages are how Go implements access controls so that the implementation of
    a feature can be hidden from the code that consumes it.
    
    How are they used?
    Packages are defined by creating code files in folders and using the package
    keyword to denote which package they belong to.
████████████████████████████████████████████████████████████████████████
116.the Module File
    This name is important because it is used to import features from other packages created within
    the same project and third-party packages
    The go statement specifies the version of Go that is used.
████████████████████████████████████████████████████████████████████████
117.Package Access Control
    Go has an unusual approach to access control. Instead of relying on dedicated keywords, like public
    and private, Go examines the first letter of the names given to the features in a code file, such as types,
    functions, and methods. If the first letter is lowercase, then the feature can be used only within the package
    that defines it. Features are exported for use outside of the package by giving them an uppercase first letter.

    The access control rules do not apply to individual function or method parameters, which means that
    the NewProduct function has to have an uppercase first character to be exported, but the parameter names
    can be lowercase.
████████████████████████████████████████████████████████████████████████
118.Adding Code Files to Packages
    Packages can contain multiple code files, and to simplify development, access control rules and package
    prefixes do not apply when accessing features defined in the same package.
████████████████████████████████████████████████████████████████████████
119.func init()
    Each code file can contain an initialization function that is executed only when all packages have been
    loaded and all other initialization—such as defining constants and variables—has been done. The most
    common use for initialization functions is to perform calculations that are difficult to perform or that require
    duplication to perform

    The initialization function is called init, and it is defined without parameters and a result. The init
    function is called automatically and provides an opportunity to prepare the package for use.

    The init function is not a regular Go function and cannot be invoked directly. And, unlike regular
    functions, a single file can define multiple init functions, all of which will be executed.

    AVOIDING THE MULTIPLE INITIALIZATION FUNCTION PITFALL
    Each code file can have its own initialization function. When using the standard Go compiler, the
    initialization functions are executed based on the alphabetic order of the filenames, so the function in
    the a.go file will be executed before the function in the b.go file, and so on.
    But this order is not part of the Go language specification and should not be relied on. Your initialization
    functions should be self-contained and not rely on other init functions having been invoked previously.
████████████████████████████████████████████████████████████████████████
120.Creating Nested Packages
    Packages can be defined within other packages, making it easy to break up complex features into as many
    units as possible.
    The package statement is used just as with any other package, without the need to include the name
    of the parent or enclosing package. And dependency on custom packages must include the full package
    path. 

    example:
        Create the packages/store/cart folder 
        and add to it a file named cart.go with the contents:
        #1
        contents:
            package cart
            import "packages/store"
            type Cart struct {
                CustomerName string
                Products []store.Product
            }
            func (cart *Cart) GetTotal() (total float64) {
                for , p := range cart.Products {
                    total += p.Price()
                }
                return
            }
    The features defined by the nested package are accessed using the package name, just like any other
    package. When importing a nested package, the package path starts with the module name and lists the
    sequence of packages.
    example:
        package main
        import (
            "fmt"
            "packages/store"
            . "packages/fmt"
            "packages/store/cart"
        )
        func main() {
            product := store.NewProduct("Kayak", "Watersports", 279)
            cart := cart.Cart {
                CustomerName: "Alice",
                Products: []store.Product{ *product },
            }
            fmt.Println("Name:", cart.CustomerName)
            fmt.Println("Total:",  ToCurrency(cart.GetTotal()))
        }
████████████████████████████████████████████████████████████████████████
121.Initialization Function
    example:
        func init() {
                for category, price := range categoryMaxPrices {
                    categoryMaxPrices[category] = price + (price * defaultTaxRate)
                }
        }
████████████████████████████████████████████████████████████████████████
122.Importing a Package Only for Initialization Effects
    Go prevents packages from being imported but not used, 
    which can be a problem if you rely on the effect of
    an initialization function but don't need to use any of the features the package exports.
    If I need the effect of the initialization function, 
    but I don't need to use the GetData function the package exports.

    example:
        package main
        import (
            "fmt"
            "packages/store"
            . "packages/fmt"
            "packages/store/cart"
             "packages/data"
        )
████████████████████████████████████████████████████████████████████████
123.Finding Go Packages
    #1 https://pkg.go.dev
    #2 https://github.com/golang/go/wiki/Projects

    Many Go modules are written by individual developers
    to solve a problem and then published for anyone else to use. 
    This creates a rich module ecosystem,
    but it does mean that maintenance and support can be inconsistent.
████████████████████████████████████████████████████████████████████████
124.indirect
    The indirect comment at the end of the statements is added automatically because
    the packages are not used by the code in the project. A file named go.sum is created when the module is
    obtained and contains checksums used to validate the packages.

    example:
        module packages
        go 1.17
        require (
            github.com/fatih/color v1.10.0 // indirect
            github.com/mattn/go-colorable v0.1.8 // indirect
            github.com/mattn/go-isatty v0.0.12 // indirect
            golang.org/x/sys v0.0.0-20200223170610-d5e6a3e2c0ae // indirect
        )

    Note:
        You can also use the go.mod file to create dependencies on projects you have created locally

    The first time you run the go get command, you will see a list of the modules that are
    downloaded, which illustrated that modules have their own dependencies and that these are resolved
    automatically:
        go: downloading github.com/fatih/color v1.10.0
        go: downloading github.com/mattn/go-isatty v0.0.12
        go: downloading github.com/mattn/go-colorable v0.1.8
        go: downloading golang.org/x/sys v0.0.0-20200223170610-d5e6a3e2c0ae
████████████████████████████████████████████████████████████████████████
125.Managing External Packages
    Removing a Package
    To update the go.mod file to reflect the change, run the command:
████████████████████████████████████████████████████████████████████████
go mod tidy126.Putting Type and Interface Composition in Context
    What is it?
    Composition is the process by which new types are created by combining
    structs and interfaces.

    Why is it useful?
    Composition allows types to be defined based on existing types.

    How is it used?
    Existing types are embedded in new types.
    
    Are there any pitfalls or limitations?
    Composition doesn't work in the same way as inheritance, and care must be
    taken to achieve the desired outcome.

    Are there any alternatives? 
    Composition is optional, and you can create entirely independent types.


    Go doesn't support classes or inheritance and focuses on composition instead. But, despite the
    differences, composition can be used to create hierarchies of types, just in a different way.

    Steps:
        1-Defining the Base Type -> struct - method 
        2-Defining a Constructor -> for create struct
        example:
        // struct:
        type Product struct {
                Name, Category string
                price float64
        }
        // method:
        func (p *Product) Price(taxRate float64) float64 {
            return p.price + (p.price * taxRate)
        }
        // Constructor:
        func NewProduct(name, category string, price float64) *Product {
            return &Product{ name, category, price }
        }
████████████████████████████████████████████████████████████████████████
127.Creating Struct Values in Packages
    example:
    package main
    import (
        "fmt"
        "composition/store"
    )
    func main() {
        
        // use Constructor:
        kayak := store.NewProduct("Kayak", "Watersports", 275)

        // use the literal syntax:
        lifejacket := &store.Product{ Name: "Lifejacket", Category:  "Watersports"}
        for , p := range []*store.Product { kayak, lifejacket} {
            fmt.Println("Name:", p.Name, "Category:", p.Category, "Price:", p.Price(0.2))
        }
    }
████████████████████████████████████████████████████████████████████████
128.Steps of composition
    1-Defining the Base Type -> struct - method 
    2-Defining a Constructor -> for create struct
    3-Composing Types
    4-Using the Boat Struct in the main.go
    A struct can mix regular and embedded field types, but the embedded fields are an important part of
    the composition feature

    Go gives special treatment to struct types that have fields whose type is another struct type, in the way
    that the Boat type has a *Product field in the example project. You can see this special treatment in the
    statement in the for loop, which is responsible for writing out details of each Boat.
    Go allows the fields of the nested type to be accessed in two ways. The first is the conventional approach
    of navigating the hierarchy of types to reach the value that is required. The *Product field is embedded, which
    means that its name its its type.

    The Boat type doesn't define a Name field, but it can be treated as though it did because of the direct
    access feature. This is known as field promotion, and Go essentially flattens the types so that the Boat
    type behaves as though it defines the fields that are provided by the nested Product type

    example:
    product.go:
        // struct:
        type Product struct {
                Name, Category string
                price float64
        }
        // method:
        func (p *Product) Price(taxRate float64) float64 {
            return p.price + (p.price * taxRate)
        }
        // Constructor:
        func NewProduct(name, category string, price float64) *Product {
            return &Product{ name, category, price }
        }
    

    boat.go:
       package store
       The Boat struct type defines an embedded *Product field
       type Boat struct {
           *Product       -------------------> Embedded Type
           Capacity int   -----------------> Reguler Fields
           Motorized bool 
       }
       func NewBoat(name string, price float64, capacity int, motorized bool) *Boat {
               return &Boat {
                   NewProduct(name, "Watersports", price), capacity, motorized,
               }
           }

    main.go:
        package main
        import (
            "fmt"
            "composition/store"
        )
    
        func main() {
           boats := []*store.Boat {
               store.NewBoat("Kayak", 275, 1, false),
               store.NewBoat("Canoe", 400, 3, false),
               store.NewBoat("Tender", 650.25, 2, true),
           }
           for , b := range boats {
                fmt.Println("Conventional:", b.Product.Name, "Direct:", b.Name)
            }
        }    
    Output:
        Conventional: Kayak Direct: Kayak
        Conventional: Canoe Direct: Canoe
        Conventional: Tender Direct: Tender

    If the field type is a value, such as Product, then any methods defined with Product or *Product
    receivers will be promoted. If the field type is a pointer, such as *Product, then only methods with *Product
    receivers will be prompted.
    There is no Price method defined for the *Boat type, but Go promotes the method defined with a
    *Product receiver.

    Calling a Method in the main.go:
        package main
        import (
            "fmt"
            "composition/store"
        )
        func main() {
            boats := []*store.Boat {
                store.NewBoat("Kayak", 275, 1, false),
                store.NewBoat("Canoe", 400, 3, false),
                store.NewBoat("Tender", 650.25, 2, true),
            }
            for , b := range boats {
                fmt.Println("Boat:", b.Name, "Price:", b.Price(0.2))
            }
        }
    Output:
        Boat: Kayak Price: 330
        Boat: Canoe Price: 480
        Boat: Tender Price: 780.3
████████████████████████████████████████████████████████████████████████
129.Creating a Chain of Nested Types
    The composition feature can be used to create complex chains of nested types, 
    whose fields and methods are promoted to the top-level enclosing type.

    Go performs promotion so that the fields defined by all three types in the chain can be
    accessed directly.

    Go promotes fields from the nested Boat and 
    Product types so they can be accessed through the top-
    level RentalBoat type, which allows the Name field to be read.
    Methods are also promoted to the top-level type, 
    which is why I can use the Price method, even though it is defined on the *Product type,
    which is at the end of the chain.

    example:
    rentalboats.go:
        package store
    
        type RentalBoat struct {
           *Boat
           IncludeCrew bool
        }
    
        func NewRentalBoat(name string, price float64, capacity int,
              motorized, crewed bool) *RentalBoat {
           return &RentalBoat{NewBoat(name, price, capacity, motorized), crewed}
        }
    

    
    Accessing Nested Fields Directly in the main.go:
        package main
        import (
            "fmt"
            "composition/store"
        )
        func main() {
            rentals := []*store.RentalBoat {
                store.NewRentalBoat("Rubber Ring", 10, 1, false, false),
                store.NewRentalBoat("Yacht", 50000, 5, true, true),
                store.NewRentalBoat("Super Yacht", 100000, 15, true, true),
            }
            for , r := range rentals {
                fmt.Println("Rental Boat:", r.Name, "Rental Price:", r.Price(0.2))
            }
        }        
    Output:
        Rental Boat: Rubber Ring Rental Price: 12
        Rental Boat: Yacht Rental Price: 60000
        Rental Boat: Super Yacht Rental Price: 120000
████████████████████████████████████████████████████████████████████████
130.Multiple Nested Types in the Same Struct
    example:
    rentalboats.go:
        package store
        type Crew struct {
            Captain, FirstOfficer string
        }
        type RentalBoat struct {
            *Boat
            IncludeCrew bool
            *Crew
        }
        func NewRentalBoat(name string, price float64, capacity int,
                motorized, crewed bool, captain, firstOfficer string) *RentalBoat {
            return &RentalBoat{NewBoat(name, price, capacity, motorized), crewed,
                &Crew{captain, firstOfficer}}
        }
    

    Using Promoted Fields in the main.go:
        package main
        import (
           "fmt"
           "composition/store"
        )
        func main() {
           rentals := []*store.RentalBoat {
              store.NewRentalBoat("Rubber Ring", 10, 1, false, false, "N/A", "N/A"),
              store.NewRentalBoat("Yacht", 50000, 5, true, true, "Bob", "Alice"),
              store.NewRentalBoat("Super Yacht", 100000, 15, true, true,
                 "Dora", "Charlie"),
           }
           for , r := range rentals {
              fmt.Println("Rental Boat:", r.Name, "Rental Price:", r.Price(0.2),
                 "Captain:", r.Captain)
           }
        }
    Output:
        Rental Boat: Rubber Ring Rental Price: 12 Captain: N/A
        Rental Boat: Yacht Rental Price: 60000 Captain: Bob
        Rental Boat: Super Yacht Rental Price: 120000 Captain: Dora
████████████████████████████████████████████████████████████████████████
131.When Promotion Cannot Be Performed
    example:
    specialdeal.go File in the store Folder:
        package store
        type SpecialDeal struct {
           Name string
           *Product
           price float64
        }
        func NewSpecialDeal(name string, p *Product, discount float64) *SpecialDeal {
           return &SpecialDeal{ name, p, p.price - discount }
        }
        func (deal *SpecialDeal ) GetDetails() (string, float64, float64) {
           return deal.Name, deal.price, deal.Price(0)
        }
    

    Using a New Type in the main.go:
        package main
        import (
            "fmt"
            "composition/store"
        )
        func main() {
            product := store.NewProduct("Kayak", "Watersports", 279)
            deal := store.NewSpecialDeal("Weekend Special", product, 50)
            Name, price, Price := deal.GetDetails()
            fmt.Println("Name:", Name)
            fmt.Println("Price field:", price)
            fmt.Println("Price method:", Price)
        }
    Output:
        Name: Weekend Special
        Price field: 229
        Price method: 279
    

    is a more concise way of expressing this statement:
        return deal.Name, deal.price, deal.Product.Price(0)

    The new Price method stops Go from promoting the Product method 
    and produces the following result
    Defining a Method in the specialdeal.go:
        package store
        type SpecialDeal struct {
           Name string
           *Product
           price float64
        }
        func NewSpecialDeal(name string, p *Product, discount float64) *SpecialDeal {
           return &SpecialDeal{ name, p, p.price - discount }
        }
        func (deal *SpecialDeal ) GetDetails() (string, float64, float64) {
           return deal.Name, deal.price, deal.Price(0)
        }
        func (deal *SpecialDeal) Price(taxRate float64) float64 {
           return deal.price
        }    
    Output:
        Name: Weekend Special
        Price field: 229
        Price method: 229
████████████████████████████████████████████████████████████████████████
132.Promotion Ambiguity
    A related issue arises when two embedded fields use the same field or method names
    example:
    An Ambiguous Method in the main.go:
        package main
        import (
           "fmt"
           "composition/store"
        )
        func main() {
           kayak := store.NewProduct("Kayak", "Watersports", 279)
           type OfferBundle struct {
              *store.SpecialDeal
              *store.Product
           }
           bundle := OfferBundle {
              store.NewSpecialDeal("Weekend Special", kayak, 50),
              store.NewProduct("Lifrejacket", "Watersports", 48.95),
           }
           fmt.Println("Price:", bundle.Price(0))
        }
    Output:
        .\main.go:22:33: ambiguous selector bundle.Price    
████████████████████████████████████████████████████████████████████████
133.Composition and Interfaces
    example:
    Mixing Types in the main.go:
        package main
        import (
            "fmt"
            "composition/store"
        )
        func main() {
            products := map[string]*store.Product {
                "Kayak": store.NewBoat("Kayak", 279, 1, false),
                "Ball": store.NewProduct("Soccer Ball", "Soccer", 19.50),
            }
            for , p := range products {
                fmt.Println("Name:", p.Name, "Category:", p.Category, "Price:", p.Price(0.2))
            }
        }
    Output:
        .\main.go:11:9: cannot use store.NewBoat("Kayak", 279, 1, false) (type *store.Boat) as
        type *store.Product in map value
████████████████████████████████████████████████████████████████████████
134.Composition to Implement Interfaces
    Go takes promoted methods into account when determining whether a type conforms to an interface,
    which avoids the need to duplicate methods that are already present through an embedded field.

    add a file named forsale.go to the store folder:
    The Contents of the forsale.go:
        package store
        type ItemForSale interface {
            Price(taxRate float64) float64
        }
            

    Using an Interface in the main.go:
        package main
        import (
           "fmt"
           "composition/store"
        )
        func main() {
           products := map[string]store.ItemForSale {
              "Kayak": store.NewBoat("Kayak", 279, 1, false),
              "Ball": store.NewProduct("Soccer Ball", "Soccer", 19.50),
           }
           for key, p := range products {
              fmt.Println("Key:", key, "Price:", p.Price(0.2))
           }
        }
    Output:
        Key: Kayak Price: 334.8
        Key: Ball Price: 23.4
████████████████████████████████████████████████████████████████████████
135.the Type Switch Limitation
    
    example:
    main.go:
        package main
        import (
            "fmt"
            "composition/store"
        )
        func main() {
            products := map[string]store.ItemForSale {
                "Kayak": store.NewBoat("Kayak", 279, 1, false),
                "Ball": store.NewProduct("Soccer Ball", "Soccer", 19.50),
            }
            for key, p := range products {
                switch item := p.(type) {
                    case *store.Product, *store.Boat:
                        fmt.Println("Name:", item.Name, "Category:", item.Category,
                            "Price:", item.Price(0.2))
                    default:
                        fmt.Println("Key:", key, "Price:", p.Price(0.2))
                }
            }
        }
    Output:
        .\main.go:21:42: item.Name undefined (type store.ItemForSale has no field or method Name)
        .\main.go:21:66: item.Category undefined (type store.ItemForSale has no field or method
        Category)
    

    Using Separate case Statements in the main.go:
        package main
        import (
           "fmt"
           "composition/store"
        )
        func main() {
           products := map[string]store.ItemForSale {
              "Kayak": store.NewBoat("Kayak", 279, 1, false),
              "Ball": store.NewProduct("Soccer Ball", "Soccer", 19.50),
           }
           for key, p := range products {
              switch item := p.(type) {
                 case *store.Product:
                    fmt.Println("Name:", item.Name, "Category:", item.Category,
                       "Price:", item.Price(0.2))
                 case *store.Boat:
                    fmt.Println("Name:", item.Name, "Category:", item.Category,
                       "Price:", item.Price(0.2))
                 default:
                    fmt.Println("Key:", key, "Price:", p.Price(0.2))
              }
           }
        }
    Output:
        Name: Kayak Category: Watersports Price: 334.8
        Name: Soccer Ball Category: Soccer Price: 23.4
████████████████████████████████████████████████████████████████████████
136.the Type Switch Limitation An alternative solution
    example:
    Defining an Interface in the product.go:
        package store
        type Product struct {
            Name, Category string
            price float64
        }
        func NewProduct(name, category string, price float64) *Product {
            return &Product{ name, category, price }
        }
        func (p *Product) Price(taxRate float64) float64 {
            return p.price + (p.price * taxRate)
        }
        type Describable interface  {
            GetName() string
            GetCategory() string
        }
        func (p *Product) GetName() string {
            return p.Name
        }
        func (p *Product) GetCategory() string {
            return p.Category
        }
    

    Using Interfaces in the main.go:
        package main
        import (
           "fmt"
           "composition/store"
        )
        func main() {
           products := map[string]store.ItemForSale {
              "Kayak": store.NewBoat("Kayak", 279, 1, false),
              "Ball": store.NewProduct("Soccer Ball", "Soccer", 19.50),
           }
           for key, p := range products {
              switch item := p.(type) {
                 case store.Describable:
                    fmt.Println("Name:", item.GetName(), "Category:", item.GetCategory(),
                       "Price:", item.(store.ItemForSale).Price(0.2))
                 default:
                    fmt.Println("Key:", key, "Price:", p.Price(0.2))
              }
           }
        }
    

    Output:
        Name: Kayak Category: Watersports Price: 334.8
        Name: Soccer Ball Category: Soccer Price: 23.4
████████████████████████████████████████████████████████████████████████
137.Composing Interfaces
    Go allows interfaces to be composed from other interfaces
    
    example:
    Composing an Interface in the product.go:
        package store
        type Product struct {
            Name, Category string
            price float64
        }
        func NewProduct(name, category string, price float64) *Product {
            return &Product{ name, category, price }
        }
        func (p *Product) Price(taxRate float64) float64 {
            return p.price + (p.price * taxRate)
        }
        type Describable interface  {
            GetName() string
            GetCategory() string
            ItemForSale
        }
        func (p *Product) GetName() string {
            return p.Name
        }
        func (p *Product) GetCategory() string {
            return p.Category
        }
    
████████████████████████████████████████████████████████████████████████
138.Goroutines and Channels
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
    
    Receiving from a channel is a blocking operation, 
    meaning that execution will not continue until a value
    has been received
    
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
████████████████████████████████████████████████████████████████████████
139.How Go Executes Code
    All Go programs use at least one goroutine because this is how Go executes the code in the
    main function.

    When compiled Go code is executed, the runtime creates a goroutine that starts executing
    the statements in the entry point, which is the main function in the main package. 
    Each statement in the main function is executed in the order in which they are defined. 
    The goroutine keeps executing statements until
    it reaches the end of the main function, at which point the application terminates.

    The goroutine executes each statement in the main function synchronously, 
    which means that it waits
    for the statement to complete before moving on to the next statement. 
    The statements in the main function
████████████████████████████████████████████████████████████████████████
140.Creating Additional Goroutines
    Go allows the developer to create additional goroutines, 
    which execute code at the same time as the main
    goroutine. Go makes it easy to create new goroutines

    example:
        package main
        import (
            "fmt"
            "time
        )

        func main(){
            fmt.Println("first statement")
            fmt.Println("second statement")
            time.Sleep(time.Second * 2)

            go fmt.Println("first statement")
            
            fmt.Println("first statement")

        }
████████████████████████████████████████████████████████████████████████
141.Returning Results from Goroutines
    Receiving from a channel is a blocking operation, 
    meaning that execution will not continue until a value
    has been received

    Getting a result from a function that is being executed asynchronously 
    can be complicated because it requires coordination between the goroutine 
    that produces the result and the goroutine that consumes the result.
    To address this issue, Go provides channels, which are 
    conduits through which data can be sent and received.

    Defining a Channel:
    example:
        var channel chan float64 = make(chan float64)
████████████████████████████████████████████████████████████████████████
142.Sending a Result Using a Channel
    Receiving from a channel is a blocking operation, 
    meaning that execution will not continue until a value
    has been received
    example:
        resultChannel <- total
████████████████████████████████████████████████████████████████████████
143.Receiving a Result Using a Channel
    Receiving from a channel is a blocking operation, 
    meaning that execution will not continue until a value
    has been received
    
    example:
        storeTotal += <- channel
████████████████████████████████████████████████████████████████████████
144.using adapters to execute functions asynchronously
    It isn't always possible to rewrite existing functions or methods to use channels, but it is a simple matter
    to execute synchronous functions asynchronously in a wrapper, like this:
    The syntax is a little awkward because the arguments used to invoke the function are expressed
    immediately following the function definition. But the result is the same, which is that a synchronous
    function can be executed by a goroutine with the result being sent through a channel.

    example:
        ...
        calcTax := func(price float64) float64 {
            return price + (price * 0.2)
        }
        wrapper := func (price float64, c chan float64)  {
            c <- calcTax(price)
        }
        resultChannel := make(chan float64)
        go wrapper(275, resultChannel)
        result := <- resultChannel
        fmt.Println("Result:", result)
        ...
        The wrapper function receives a channel, which it uses to send the value received from executing the
        calcTax function synchronously. This can be expressed more concisely by defining a function without
        assigning it to a variable, like this:
        ...
        go func (price float64, c chan float64) {
            c <- calcTax(price)
        }(275, resultChannel)
        ...
████████████████████████████████████████████████████████████████████████
145.Coordinating Channels
    By default, sending and receiving through a channel are blocking operations. This means a goroutine that
    sends a value will not execute any further statements until another goroutine receives the value from the
    channel. If a second goroutine sends a value, it will be blocked until the channel is cleared, causing a queue
    of goroutines waiting for values to be received. This happens in the other direction, too, so that goroutines
    that receive values will block until another goroutine sends one.

    example:
        If Bob has a message for Alice, the default channel behavior requires 
        Alice and Bob to agree on a meeting place, and
        whoever gets there first will wait for the other to arrive. 
        Bob will only give the message to Alice when they are
        both present. When Charlie also has a message for Alice, he will form a queue behind Bob. 
        Everyone waits patiently, messages are transferred only when the sender 
        and receiver are both available, and messages are
        processed sequentially.
████████████████████████████████████████████████████████████████████████
146.Buffered Channel
    The default channel behavior can lead to bursts of activity as goroutines do their work, followed by a long
    idle period waiting for messages to be received. This doesn't have an impact on the example application
    because the goroutines finish once their messages are received, but in a real project goroutines often have
    repetitive tasks to perform, and waiting for a receiver can cause a performance bottleneck.

    An alternative approach is to create a channel with a buffer, 
    which is used to accept values from a
    sender and store them until a receiver becomes available. 
    This makes sending a message a nonblocking
    operation, allowing a sender to pass its value to the channel and continue working without having to wait
    for a receiver. 
    
    example:
        This is similar to Alice having an inbox on her desk. Senders come to Alice's office and put
        their message into the inbox, leaving it for Alice to read when she is ready. But, if the inbox is full, then they
        will have to wait until she has processed some of her backlog before sending a new message.
    Creating a Buffered Channel
    example:
        var channel chan float64 = make(chan float64, 2)
    Result explain:
        For this example, I have set the size of the buffer to 2, meaning that two senders will be able to send
        values through the channel without having to wait for them to be received.
████████████████████████████████████████████████████████████████████████
147.Inspecting a Channel Buffer
    You can determine the size of a channel's buffer using 
    the built-in cap function and determine how many
    values are in the buffer using the len function
    The modified statement uses the len and cap functions to report 
    the number of values in the channel's
    buffer and the overall size of the buffer.

    example:
        len(channel), "items in buffer, size", cap(channel))
████████████████████████████████████████████████████████████████████████
148.Closing a Channel
    The solution for this problem is for the sender to indicate when no further values are coming through the
    channel, which is done by closing the channel
    The built-in close function accepts a channel as its argument and is used to indicate that there will be
    no further values sent through the channel. Receivers can check if a channel is closed when requesting a
    value.
    You need to close channels only when it is helpful to do so to coordinate your goroutines. 
    Go doesn't require channels to be closed to free up resources or perform any kind of housekeeping task.

    example:
        close(channel)
    


    The receive operator can be used to obtain two values. 
    The first value is assigned the value received
    from the channel, and the second value indicates whether the channel is closed
    It is illegal to send values to a channel once it has been closed.

    example:
        if details, open := <- dispatchChannel; open {
                fmt.Println("Dispatch to", details.Customer, ":", details.Quantity,
                "x", details.Product.Name)
            } else {
                fmt.Println("Channel has been closed")
            break
        }    
████████████████████████████████████████████████████████████████████████
149.Enumerating Channel Values
    A for loop can be used with the range keyword to enumerate the values sent through a channel, allowing
    the values to be received more easily and terminating the loop when the channel is closed

    The range expression produces one value per iteration, which is the value received from the channel.
    The for loop will continue to receive values until the channel is closed. (You can use a for...range loop on a
    channel that isn't closed, in which case the loop will never exit.)

    example:
        go DispatchOrders(dispatchChannel)
        for details := range dispatchChannel {
            fmt.Println("Dispatch to", details.Customer, ":", details.Quantity,
                "x", details.Product.Name)
        }
        fmt.Println("Channel has been closed")
████████████████████████████████████████████████████████████████████████
150.Restricting Channel Direction
    By default, channels can be used to send and receive data, but this can be restricted when using channels
    as arguments, such that only send or receive operations can be performed.
    
    example:
        func DispatchOrders(channel chan<- DispatchNotification)

    The location of the arrow specifies the direction of the channel. 
    When the arrow follows the chan keyword,
    then the channel can be used only to send.
    The channel can be used to receive only if the arrow precedes the chan keyword (<-chan, for example).

    Attempting to receive from a send-only (and vice versa) channel is a compile-time error,

    example:
        # concurrency
        .\orderdispatch.go:29:29: invalid operation: <-channel (receive from send-only type chan<-
        DispatchNotification)
████████████████████████████████████████████████████████████████████████
151.Restricting Channel Argument Direction
    Go allows bidirectional channels to be assigned to unidirectional channel variables, 
    allowing restrictions to be applied
    example:
        func receiveDispatches(channel <-chan DispatchNotification) {
            for details := range channel {
                fmt.Println("Dispatch to", details.Customer, ":", details.Quantity,
                    "x", details.Product.Name)
            }
            fmt.Println("Channel has been closed")
        }
    
    Restrictions on channel direction can also be created through explicit conversion
    The explicit conversion for the receive-only channel requires parentheses around the channel type
    to prevent the compiler from interpreting a conversion to the DispatchNotification type.

    example:
        func main() {
            dispatchChannel := make(chan DispatchNotification, 100)
            go DispatchOrders(chan<- DispatchNotification(dispatchChannel))
            receiveDispatches((<-chan DispatchNotification)(dispatchChannel))
        }
████████████████████████████████████████████████████████████████████████
152.Select Statements
    The select keyword is used to group operations that will send or receive from channels, 
    which allows for complex arrangements of goroutines and channels to be created. 
    There are several uses for select statements.

    The simplest use for select statements is to receive from a channel without blocking, ensuring that a
    goroutine won't have to wait when the channel is empty.
    A select statement has a similar structure to a switch statement, except that the case statements are
    channel operations.
    When the select statement is executed, each channel operation is evaluated until one
    that can be performed without blocking is reached. The channel operation is performed, and the statements
    enclosed in the case statement are executed. If none of the channel operations can be performed, the
    statements in the default clause are executed.

    example:
      for {
            select {
                case details, ok := <- dispatchChannel:
                if ok {
                    fmt.Println("Dispatch to", details.Customer, ":",
                        details.Quantity, "x", details.Product.Name)
                } else {
                    fmt.Println("Channel has been closed")
                    goto alldone
                }
            default:
                fmt.Println("-- No message ready to be received")
                time.Sleep(time.Millisecond * 500)
            }
        }
        alldone: fmt.Println("All values received")
████████████████████████████████████████████████████████████████████████
153.Receiving from Multiple Channels
    A select statement can be used to receive without blocking,
    when there are multiple channels, through which values are sent at different
    rates. A select statement will allow the receiver to obtain values from whichever channel has them, without
    blocking on any single channel


    In this example, the select statement is used to receive values from two channels, one that carries
    DispatchNofitication values and one that carries Product values. Each time the select statement is
    executed, it works its way through the case statements, building up a list of the ones from which a value can
    be read without blocking. One of the case statements is selected from the list at random and executed. If
    none of the case statements can be performed, the default clause is executed.

    example:
        package main
        import (
            "fmt"
            "time"
        )
        func enumerateProducts(channel chan<- *Product) {
            for , p := range ProductList[:3] {
                channel <- p
                time.Sleep(time.Millisecond * 800)
            }
            close(channel)
        }
        func main() {
            dispatchChannel := make(chan DispatchNotification, 100)
            go DispatchOrders(chan<- DispatchNotification(dispatchChannel))
            productChannel := make(chan *Product)
            go enumerateProducts(productChannel)
            openChannels := 2
            for  {
                select {
                    case details, ok := <- dispatchChannel:
                        if ok {
                            fmt.Println("Dispatch to", details.Customer, ":",
                                details.Quantity, "x", details.Product.Name)
                        } else {
                            fmt.Println("Dispatch channel has been closed")
                            dispatchChannel = nil
                            openChannels--
                        }
                    case product, ok := <- productChannel:
                        if ok {
                            fmt.Println("Product:", product.Name)
                        } else {
                            fmt.Println("Product channel has been closed")
                            productChannel = nil
                            openChannels--
                        }
                    default:
                        if (openChannels == 0) {
                            goto alldone
                        }
                        fmt.Println("-- No message ready to be received")
                        time.Sleep(time.Millisecond * 500)
                }
            }
            alldone: fmt.Println("All values received")
        }
████████████████████████████████████████████████████████████████████████
154.Sending Without Blocking
    A select statement can also be used to send to a channel without blocking

    example:
        func enumerateProducts(channel chan<- *Product) {
            for , p := range ProductList {
                select {
                    case channel <- p:
                        fmt.Println("Sent product:", p.Name)
                    default:
                        fmt.Println("Discarding product:", p.Name)
                        time.Sleep(time.Second)
                }
            }
            close(channel)
        }
████████████████████████████████████████████████████████████████████████
155.Sending to Multiple Channels
    If there are multiple channels available, a select statement can be used to find a channel for which sending
    will not block
    You can combine case statements with send and receive operations in the same select statement.
    When the select statement is executed, the Go runtime builds a combined list of case statements that can be
    executed without blocking and picks one at random, which can be either a send or a receive statement.

    This example has two channels with small buffers. As with receiving, the select statement builds a list
    of the channels through which a value can be sent without blocking and then picks one at random from that
    list. If none of the channels can be used, then the default clause is executed. There is no default clause
    in this example, which means that the select statement will block until one of the channels can receive
    a value.

    example:
        func enumerateProducts(channel1, channel2 chan<- *Product) {
            for , p := range ProductList {
                select {
                    case channel1 <- p:
                        fmt.Println("Send via channel 1")
                    case channel2 <- p:
                        fmt.Println("Send via channel 2")
                }
            }
            close(channel1)
            close(channel2)
        }
████████████████████████████████████████████████████████████████████████
156.Error Handling
    What is it?
    Go's error handling allows exceptional conditions and failures to be represented and dealt with.

    Why is it useful?
    Applications will often encounter unexpected situations, 
    and the error handling features provide a way to respond to those situations when they arise.

    How is it used?
    The error interface is used to define error conditions, 
    which are typically returned as function results. 
    The panic function is called when an unrecoverable error occurs.

    Are there any pitfalls or limitations?
    Care must be taken to ensure that errors are communicated to the part of the
    application that can best decide how serious the situation is.

    Are there any alternatives?
    You don't have to use the error interface in your code, 
    but it is employed throughout the Go standard library and is difficult to avoid.

    Go provides a predefined interface named error that provides one way to resolve this issue. Here is the
    definition of the interface:

    type error interface {
        Error() string
    }

    Problem                                 Solution
    ----------------------------------      ---------------------------------------------------
    Indicate that an error has occurred     Create a struct that implements the error interface
                                            and return it as a function result
    Report an error over a channel          Add an error field to the struct type used for
                                            channel messages
    Indicate that an unrecoverable          Call the panic function 
    error has occurred
    Recover from a panic                    Use the defer keyword to register a function that
                                            calls the recover function
████████████████████████████████████████████████████████████████████████
157.Generating Errors
    Functions and methods can express exceptional or unexpected outcomes by producing error responses
    Defining an Error
    example:
        package main
        import "fmt"
        func main() {
            categories := []string { "Watersports", "Chess", "Running" }
            for , cat := range categories {
                total, err := Products.TotalPrice(cat)
                if (err == nil) {
                    fmt.Println(cat, "Total:", ToCurrency(total))
                } else {
                    fmt.Println(cat, "(no such category)")
                }
            }
        }
    Output:
        Watersports Total: $328.95
        Chess Total: $1291.00
        Running (no such category)
████████████████████████████████████████████████████████████████████████
158.Ignoring Error Results
    I don't recommend ignoring error results because it means you will lose important information.

    but if you don't need to know when something goes wrong, 
    then you can use the blank identifier instead of a
    name for the error result, like this:
    
    example:
        package main
        import "fmt"
        func main() {
            categories := []string { "Watersports", "Chess", "Running" }
            for , cat := range categories {
                total,  := Products.TotalPrice(cat)
                fmt.Println(cat, "Total:", ToCurrency(total))
            }
        }
████████████████████████████████████████████████████████████████████████
159.Reporting Errors via Channels
    example:
    operations.go:
        package main
        type CategoryError struct {
           requestedCategory string
        }
        func (e *CategoryError) Error() string {
           return "Category " + e.requestedCategory + " does not exist"
        }
        type ChannelMessage struct {
           Category string
           Total float64
           *CategoryError
        }
        func (slice ProductSlice) TotalPrice(category string) (total float64,
              err *CategoryError) {
           productCount := 0
           for , p := range slice {
              if (p.Category == category) {
                 total += p.Price
                 productCount++
              }
           }
           if (productCount == 0) {
              err = &CategoryError{ requestedCategory: category}
           }
           return
        }
        func (slice ProductSlice) TotalPriceAsync (categories []string,
              channel chan<- ChannelMessage) {
           for , c := range categories {
              total, err := slice.TotalPrice(c)
              channel <- ChannelMessage{
                 Category: c,
                 Total: total,
                 CategoryError: err,
              }
           }
           close(channel)
        }
    

    main.go:
        package main
        import "fmt"
        func main() {
           categories := []string { "Watersports", "Chess", "Running" }
           channel := make(chan ChannelMessage, 10)
           go Products.TotalPriceAsync(categories, channel)
           for message := range channel {
              if message.CategoryError == nil {
                 fmt.Println(message.Category, "Total:", ToCurrency(message.Total))
              } else {
                 fmt.Println(message.Category, "(no such category)")
              }
           }
        }
    Output:
        Watersports Total: $328.95
        Chess Total: $1291.00
        Running (no such category)
████████████████████████████████████████████████████████████████████████
160.String Processing and Regular Expressions
    What are they?
    String processing includes a wide range of operations, from trimming
    whitespace to splitting a string into components. Regular expressions
    are patterns that allow string matching rules to be concisely defined.

    Why are they useful?
    These operations are useful when an application needs to process
    string values. A common example is processing HTTP requests.

    How are they used?
    These features are contained in the strings and regexp packages,
    which are part of the standard library.

    Are there any pitfalls or limitations?
    There are some quirks in the way that some of these operations are
    performed, but they mostly behave as you would expect.

    Are there any alternatives?
    The use of these packages is optional, and they do not have
    to be used. That said, there is little point in creating your own
    implementations of these features since the standard library is well-
    written and thoroughly tested.

    Problem                     Solution
    -------------------         -------------------------------------------------------------------
    Compare strings             Use the Contains, EqualFold, or Has* function in the strings package
    Convert string case         Use the ToLower, ToUpper, Title, or ToTitle function in the
                                strings package
    Check or change             Use the functions provided by the unicode package
    character case
    Find content in strings     Use the functions provided by the strings or regexp package
    Split a string              Use the Fields or Split* function in the strings and regexp packages
    Join strings                Use the Join or Repeat function in the strings package
    Trim characters from        Use the Trim* functions in the strings package
    a string
    Perform a substitution      Use the Replace* or Map function in the strings package,
    تعویض انجام دهید            use a Replacer, or use the Replace* functions in the regexp package
    Efficiently build a         Use the Builder type in the strings package
    string

████████████████████████████████████████████████████████████████████████
161.Comparing Strings
    The strings Functions for Comparing Strings

    Function                    Description
    -------------               ------------------------
    Contains(s, substr)         This function returns true if the string s contains substr and false if it does not.
    ContainsAny(s, substr)      This function returns true if the string s contains any of the characters
                                contained in the string substr.
    ContainsRune(s, rune)       This function returns true if the string s contains a specific rune.
    EqualFold(s1, s2)           This function performs a case-insensitive comparison and returns true of
                                strings s1 and s2 are the same.
    HasPrefix(s, prefix)        This function returns true if the string s begins with the string prefix.
    HasSuffix(s, suffix)        This function returns true if the string ends with the string suffix.

    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
            product := "Kayak"
            fmt.Println("Contains:", strings.Contains(product, "yak"))
            fmt.Println("ContainsAny:", strings.ContainsAny(product, "abc"))
            fmt.Println("ContainsRune:", strings.ContainsRune(product, 'K'))
            fmt.Println("EqualFold:", strings.EqualFold(product, "KAYAK"))
            fmt.Println("HasPrefix:", strings.HasPrefix(product, "Ka"))
            fmt.Println("HasSuffix:", strings.HasSuffix(product, "yak"))
        }
    Output:
        Contains: true
        ContainsAny: true
        ContainsRune: true
        HasPrefix: true
        HasSuffix: true
        EqualFold: true
████████████████████████████████████████████████████████████████████████
162.Using The Byte-Oriented Functions
    example:
        package main
        import (
            "fmt"
            "strings"
            "bytes"
        )
        func main() {
            price := "€100"
            fmt.Println("Strings Prefix:", strings.HasPrefix(price, "€"))
            fmt.Println("Bytes Prefix:", bytes.HasPrefix([]byte(price),
                []byte { 226, 130 }))
        }
    Output:
        Strings Prefix: true
        Bytes Prefix: true
████████████████████████████████████████████████████████████████████████
163.Converting String Case
    The Case Functions in the strings Package
    Function        Description
    --------------  ------------------------------------------------------
    ToLower(str)    This function returns a new string containing the characters in the specified string
                    mapped to lowercase.
    ToUpper(str)    This function returns a new string containing the characters in the specified string
                    mapped to lowercase.
    Title(str)      This function converts the specific string so that the first character of each word is
                    uppercase and the remaining characters are lowercase.
    ToTitle(str)    This function returns a new string containing the characters in the specified string
                    mapped to title case.

    Care must be taken with the Title and ToTitle functions, which don't work the way you might expect.
    The Title function returns a string that is suitable for use as a title, but it treats all words the same

    Creating a Title:
    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
            description := "A boat for sailing"
            fmt.Println("Original:", description)
            fmt.Println("Title:", strings.Title(description))
    	fmt.Println("Title:", strings.ToTitle(description))
        }
    Output:
        Original: A boat for sailing
        Title: A Boat For Sailing
        Title: A BOAT FOR SAILING
████████████████████████████████████████████████████████████████████████
164.Title Case
    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
            specialChar := "\u01c9"
            fmt.Println("Original:", specialChar, []byte(specialChar))
            upperChar := strings.ToUpper(specialChar)
            fmt.Println("Upper:", upperChar, []byte(upperChar))
            titleChar := strings.ToTitle(specialChar)
            fmt.Println("Title:", titleChar, []byte(titleChar))
        }
    Output:
        Original: ǉ [199 137]
        Upper: Ǉ [199 135]
        Title: ǈ [199 136]
████████████████████████████████████████████████████████████████████████
165.Working with Character Case
    The unicode package provides functions that can be used to determine or change the case of individual
    characters
    Functions in the unicode Package for Character Case
    Function        Description
    -------------   -----------------------------------------------------------------
    IsLower(rune)   This function returns true if the specified rune is lowercase.
    ToLower(rune)   This function returns the lowercase rune associated with the specified rune.
    IsUpper(rune)   This function returns true if the specified rune is uppercase.
    ToUpper(rune)   This function returns the upper rune associated with the specified rune.
    IsTitle(rune)   This function returns true if the specified rune is title case.
    ToTitle(rune)   This function returns the title case rune associated with the specified rune.
████████████████████████████████████████████████████████████████████████
166.the Rune Case Functions
    example:
        package main
        import (
            "fmt"
            "unicode"
        )
        func main() {
            product := "Kayak"
            for _, char := range product {
                fmt.Println(string(char), "Upper case:", unicode.IsUpper(char))
            }
        }
    Output:
        K Upper case: true
        a Upper case: false
        y Upper case: false
        a Upper case: false
        k Upper case: false        
████████████████████████████████████████████████████████████████████████
167.Inspecting Strings
    The strings Functions for Inspecting Strings
    Function                Description
    ------------            --------------------------------------------------
    Count(s, sub)           This function returns an int that reports how many times the specified
                            substring is found in the string s.
    Index(s, sub)           These functions return the index of the first or last occurrence of a specified
    LastIndex(s, sub)       substring string within the string s, or -1 if there is no occurrence.
    IndexAny(s, chars)      These functions return the first or last occurrence of any character in the
    LastIndexAny(s, chars)  specified string within the string s, or -1 if there is no occurrence.
    IndexByte(s, b)         These functions return the index of the first or last occurrence of a specified
    LastIndexByte(s, b)     byte within the string s, or -1 if there is no occurrence.
    IndexFunc(s, func)      These functions return the index of the first or last occurrence of the
    LastIndexFunc(s, func)  character in the string s for which the specified function returns true, as
                            described in the “Inspecting Strings with Custom Functions” section.
    
    example:
        package main
        import (
            "fmt"
            "strings"
            //"unicode"
        )
        func main() {
            description := "A boat for one person"
            fmt.Println("Count:", strings.Count(description, "o"))
            fmt.Println("Index:", strings.Index(description, "o"))
            fmt.Println("LastIndex:", strings.LastIndex(description, "o"))
            fmt.Println("IndexAny:", strings.IndexAny(description, "abcd"))
            fmt.Println("LastIndex:", strings.LastIndex(description, "o"))
            fmt.Println("LastIndexAny:", strings.LastIndexAny(description, "abcd"))
        }
    Output:
        Count: 4
        Index: 3
        LastIndex: 19
        IndexAny: 2
        LastIndex: 19
        LastIndexAny: 4
████████████████████████████████████████████████████████████████████████
168.IndexFunc and LastIndexFunc functions 
    Inspecting Strings with Custom Functions
    The IndexFunc and LastIndexFunc functions use a custom function to inspect strings, using custom functions

    Custom functions receive a rune and return a bool result that indicates if the character meets the
    desired condition. The IndexFunc function invokes the custom function for each character in the string until
    a true result is obtained, at which point the index is returned.
    The isLetterB variable is assigned a custom function that receives a rune and returns true if the rune
    is a uppercase or lowercase B. The custom function is passed to the strings.IndexFunc function
    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
            description := "A boat for one person"
            isLetterB := func (r rune) bool {
                return r == 'B' || r == 'b'
            }
            fmt.Println("IndexFunc:", strings.IndexFunc(description, isLetterB))
        }
    Output:
        IndexFunc: 2
████████████████████████████████████████████████████████████████████████
169.Splitting Strings
    The Functions for Splitting Strings in the strings Package
    Function                    Description
    ---------------             --------------------------------
    Fields(s)                   This function splits a string on whitespace characters and returns a slice
                                containing the nonwhitespace sections of the string s.
    FieldsFunc(s, func)         This function splits the string s on the characters for which a custom function
                                returns true and returns a slice containing the remaining sections of the string.
    Split(s, sub)               This function splits the string s on every occurrence of the specified substring,
                                returning a string slice. If the separator is the empty string, then the slice will
                                contain strings for each character.
    SplitN(s, sub, max)         This function is similar to Split, but accepts an additional int argument that
                                specifies the maximum number of substrings to return. The last substring in the
                                result slice will contain the unsplit portion of the source string.
    SplitAfter(s, sub)          This function is similar to Split but includes the substring used in the results.
    SplitAfterN(s, sub, max)    This function is similar to SplitAfter, but accepts an additional int argument
                                that specifies the maximum number of substrings to return.

    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
            description := "A boat for one person"
            splits := strings.Split(description, " ")
            for _, x := range splits {
                fmt.Println("Split >>" + x + "<<")
            }
            splitsAfter := strings.SplitAfter(description, " ")
            for _, x := range splitsAfter {
                fmt.Println("SplitAfter >>" + x + "<<")
            }
        }
    Output:
        Split >>A<<
        Split >>boat<<
        Split >>for<<
        Split >>one<<
        Split >>person<<
        SplitAfter >>A <<
        SplitAfter >>boat <<
        SplitAfter >>for <<
        SplitAfter >>one <<
        SplitAfter >>person<<
████████████████████████████████████████████████████████████████████████
170.SplitN and SplitAfterN functions 
    Restricting the Number of Results
    The SplitN and SplitAfterN functions accept an int argument that specifies the maximum number of
    results that should be included in the results
    Restricting the Results
    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
            description := "A boat for one person"
            splits := strings.SplitN(description, " ", 3)
            for _, x := range splits {
                fmt.Println("Split >>" + x + "<<")
            }
        }
    Output:
        Split >>A<<
        Split >>boat<<
        Split >>for one person<<
████████████████████████████████████████████████████████████████████████
171.strings.SplitN function
    Splitting on Whitespace Characters
    One limitation of the Split, SplitN, SplitAfter, and SplitAfterN functions is they do not deal with
    repeated sequences of characters, which can be a problem when splitting a string on whitespace characters

    The words in the source string are double-spaced, but the SplitN function splits only on the first space
    character, which produces odd results.
    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
            description := "This  is  double  spaced"
            splits := strings.SplitN(description, " ", 3)
            for _, x := range splits {
                fmt.Println("Split >>" + x + "<<")
            }
        }
    Output:
        Split >>This<<
        Split >><<
        Split >>is  double  spaced<<    
████████████████████████████████████████████████████████████████████████
172.Fields Function
    The Fields function doesn't support a limit on the number of results 
    but does deal with the double spaces properly.
    The Fields function has a better approach, which is to split on any character for
    which the IsSpace function in the unicode package returns true.
    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
            description := "This  is  double  spaced"
            splits := strings.Fields(description)
            for _, x := range splits {
                fmt.Println("Field >>" + x + "<<")
            }
        }
    Output:
        Field >>This<<
        Field >>is<<
        Field >>double<<
        Field >>spaced<<
████████████████████████████████████████████████████████████████████████
173.FieldsFunc function
    Splitting Using a Custom Function to Split Strings
    The FieldsFunc function splits a string by passing each character to a custom function and splitting when
    that function returns true
    The custom function receives a rune and returns true if that rune should cause the string to split.
    The FieldsFunc function is smart enough to deal with repeated characters
    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
            description := "This  is  double  spaced"
            splitter := func(r rune) bool {
                return r == ' '
            }
            splits := strings.FieldsFunc(description, splitter)
            for _, x := range splits {
                fmt.Println("Field >>" + x + "<<")
            }
        }
    Output:
        Field >>This<<
        Field >>is<<
        Field >>double<<
        Field >>spaced<<
████████████████████████████████████████████████████████████████████████
174.Trimming Strings
    The Functions for Trimming Strings in the strings Package
    
    Function                Description
    ------------            ---------------------------------------
    TrimSpace(s)            This function returns the string s without leading or trailing whitespace characters.
    Trim(s, set)            This function returns a string from which any leading or trailing characters
                            contained in the string set are removed from the string s.
    TrimLeft(s, set)        This function returns the string s without any leading character contained
                            in the string set. This function matches any of the specified characters—use
                            the TrimPrefix function to remove a complete substring.
    TrimRight(s, set)       This function returns the string s without any trailing character contained
                            in the string set. This function matches any of the specified characters—use
                            the TrimSuffix function to remove a complete substring.
    TrimPrefix(s, prefix)   This function returns the string s after removing the specified prefix string.
                            This function removes the complete prefix string—use the TrimLeft
                            function to remove characters from a set.
    TrimSuffix(s, suffix)   This function returns the string s after removing the specified suffix string.
                            This function removes the complete suffix string—use the TrimRight
                            function to remove characters from a set.
    TrimFunc(s, func)       This function returns the string s from which any leading or trailing
                            character for which a custom function returns true are removed.
    TrimLeftFunc(s, func)   This function returns the string s from which any leading character for
                            which a custom function returns true are removed.
    TrimRightFunc(s, func)  This function returns the string s from which any trailing character for
                            which a custom function returns true are removed.
████████████████████████████████████████████████████████████████████████
175.TrimSpace function
    Trimming Whitespace
    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
            username := " Alice"
            trimmed := strings.TrimSpace(username)
            fmt.Println("Trimmed:", ">>" + trimmed + "<<")
        }
    Output:
        Trimmed: >>Alice<<
████████████████████████████████████████████████████████████████████████
176.Trim, TrimLeft, and TrimRight functions
    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
            description := "A boat for one person"
            trimmed := strings.Trim(description, "Anor ")
            fmt.Println("Trimmed:>>"+ trimmed+ "<<")
        }
    Ourput:
        Trimmed:>>boat for one pers<<
████████████████████████████████████████████████████████████████████████
177.TrimPrefix and TrimSuffix functions
    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
            description := "A boat for one person"
            prefixTrimmed := strings.TrimPrefix(description, "A boat ")
            wrongPrefix := strings.TrimPrefix(description, "A hat ")
            fmt.Println("Trimmed:", prefixTrimmed)
            fmt.Println("Not trimmed:", wrongPrefix)
        }
    Output:
        Trimmed: for one person
        Not trimmed: A boat for one person
████████████████████████████████████████████████████████████████████████
178.TrimFunc, TrimLeftFunc, and TrimRightFunc functions
    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
            description := "A boat for one person"
            trimmer := func(r rune) bool {
                return r == 'A' || r == 'n'
            }
            trimmed := strings.TrimFunc(description, trimmer)
            fmt.Println("Trimmed:", trimmed)
        }
    Output:
        Trimmed:  boat for one perso
████████████████████████████████████████████████████████████████████████
179.Altering Strings
    تغییر رشته‌ها 
    
    The Functions for Altering Strings in the strings Package
    Function                    Description
    ----------------            -----------------------------------
    Replace(s, old, new, n)     This function alters the string s by replacing occurrences of the string old with the
                                string new. The maximum number of occurrences that will be replaced is specified by
                                the int argument n.
    ReplaceAll(s, old, new)     This function alters the string s by replacing all occurrences of the string old with
                                the string new. Unlike the Replace function, there is no limit on the number of
                                occurrences that will be replaced.
    Map(func, s)                This function generates a string by invoking the custom function for each character in
                                the string s and concatenating the results. If the function produces a negative value,
                                the current character is dropped without a replacement.
████████████████████████████████████████████████████████████████████████
180.Replace and ReplaceAll functions
    The Replace function allows a maximum number of changes to be specified, 
    while the ReplaceAll function will replace all the
    occurrences of the substring it finds
    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
            text := "It was a boat. A small boat."
            replace := strings.Replace(text, "boat", "canoe", 1)
            replaceAll := strings.ReplaceAll(text, "boat", "truck")
            fmt.Println("Replace:", replace)
            fmt.Println("Replace All:", replaceAll)
        }
    Output:
        Replace: It was a canoe. A small boat.
        Replace All: It was a truck. A small truck.
████████████████████████████████████████████████████████████████████████
181.Map function
    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
        text := "It was a boat. A small boat."
        mapper := func(r rune) rune {
                if r == 'b' {
                    return 'c'
                }
                return r
            }
            mapped := strings.Map(mapper, text)
            fmt.Println("Mapped:", mapped)
        }
    Output:
        Mapped: It was a coat. A small coat.
████████████████████████████████████████████████████████████████████████
182.NewReplacer function
    The strings package exports a struct type named Replacer that is used to replace strings
    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
            text := "It was a boat. A small boat.   111"
            replacer := strings.NewReplacer("boat", "kayak", 
            "small", "huge",
            "111", "222",
        )
            replaced := replacer.Replace(text)
            fmt.Println("Replaced:", replaced)
        }
    Output:
        Replaced: It was a kayak. A huge kayak.   222
████████████████████████████████████████████████████████████████████████
183.The Replacer Methods
    Name                    Description
    -------------------     --------------------------------------------
    Replace(s)              This method returns a string for which all the replacements specified with the
                            constructor have been performed on the string s.

    WriteString(writer, s)  This method is used to perform the replacements specified with the constructor
                            and write the results to an io.Writer
████████████████████████████████████████████████████████████████████████
184.Building and Generating Strings
    The strings Functions for Generating Strings
    Function            Description
    ----------------    ----------------------------------------------------------------------------------------
    Join(slice, sep)    This function combines the elements in the specified string slice, with the specified
                        separator string placed between elements.

    Repeat(s, count)    This function generates a string by repeating the string s for a specified number of times.
████████████████████████████████████████████████████████████████████████
185.Join and Repeat functions
    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
            text := "It was a boat. A small boat."
           elements := strings.Fields(text)
            joined := strings.Join(elements, "--")
            fmt.Println("Joined:", joined)
            esplited := strings.Split(text, " ")
            fmt.Printf("%q\n",esplited)
        }
    Output:
        Joined: It--was--a--boat.--A--small--boat.
        ["It" "was" "a" "boat." "A" "small" "boat."]
████████████████████████████████████████████████████████████████████████
186.Building Strings
    The strings.Builder Methods
    Name                Description
    ---------------     --------------------------------------------
    WriteString(s)      This method appends the string s to the string being built.
    WriteRune(r)        This method appends the character r to the string being built.
    WriteByte(b)        This method appends the byte b to the string being built.
    String()            This method returns the string that has been created by the builder.
    Reset()             This method resets the string created by the builder.
    Len()               This method returns the number of bytes used to store the string created by the builder.
    Cap()               This method returns the number of bytes that have been allocated by the builder.
    Grow(size)          This method increases the number of bytes used allocated by the builder to store the
                        string that is being built.
████████████████████████████████████████████████████████████████████████
187.builder.String()
    Creating the string using the Builder is more efficient than using the concatenation operator on regular
    string values, especially if the Grow method is used to allocate storage in advance.
    Care must be taken to use pointers when passing Builder values to and from functions and
    methods; otherwise, the efficiency gains will be lost when the Builder is copied.
    example:
        package main
        import (
            "fmt"
            "strings"
        )
        func main() {
            text := "It was a boat. A small boat."
            var builder strings.Builder
            for _, sub := range strings.Fields(text) {
                if (sub == "small") {
                    builder.WriteString("very ")
                }
                builder.WriteString(sub)
                builder.WriteRune(' ')
            }
            fmt.Println("String:", builder.String())
        }
████████████████████████████████████████████████████████████████████████
188.new() function
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
████████████████████████████████████████████████████████████████████████
189.Regular Expressions
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
        Match: true
████████████████████████████████████████████████████████████████████████
190.Compiling and Reusing Patterns
    The MatchString function is simple and convenient, 
    but the full power of regular expressions is accessed
    through the Compile function
████████████████████████████████████████████████████████████████████████
191.regexp.Compile() function
    This is more efficient because the pattern has to be compiled only once. 
    The result of the Compile
    function is an instance of the RegExp type, 
    which defines the MatchString function.

    example:
        package main
        import (
            "fmt"
            "regexp"
        )
        func main() {
            pattern, compileErr := regexp.Compile("[A-z]oat")
            description := "A boat for one person"
            question := "Is that a goat?"
            preference := "I like oats"
            if (compileErr == nil) {
                fmt.Println("Description:", pattern.MatchString(description))
                fmt.Println("Question:", pattern.MatchString(question))
                fmt.Println("Preference:", pattern.MatchString(preference))
            } else {
                fmt.Println("Error:", compileErr)
            }
        }
    Output:
        Description: true
        Question: true
        Preference: false
████████████████████████████████████████████████████████████████████████
192.Useful Basic Regexp Methods
    Function                    Description
    ---------------             ----------------------------------------------------------------------
    MatchString(s)              This method returns true if the string s matches the compiled pattern.
    FindStringIndex(s)          This method returns an int slice containing the location for the left-
                                most match made by the compiled pattern in the string s. A nil result
                                indicates that no matches were made.
    FindAllStringIndex(s, max)  This method returns a slice of int slices that contain the location for all
                                the matches made by the compiled pattern in the string s. A nil result
                                indicates that no matches were made.
    FindString(s)               This method returns a string containing the left-most match made by the
                                compiled pattern in the string s. An empty string will be returned if no
                                match is made.
    FindAllString(s, max)       This method returns a string slice containing the matches made by the
                                compiled pattern in the string s. The int argument max specifies the
                                maximum number of matches, with -1 specifying no limit. A nil result is
                                returned if there are no matches.
    Split(s, max)               This method splits the string s using matches from the compiled pattern
                                as separators and returns a slice containing the split substrings.

    example:
        package main
        import (
            "fmt"
            "regexp"
        )
        func getSubstring(s string, indices []int) string {
            return string(s[indices[0]:indices[1]])
        }
        func main() {
            pattern := regexp.MustCompile("K[a-z]{4}|[A-z]oat")
            description := "Kayak. A boat for one person."
            firstIndex := pattern.FindStringIndex(description)
            allIndices := pattern.FindAllStringIndex(description, -1)
            fmt.Println("First index", firstIndex[0], "-", firstIndex[1],
                "=", getSubstring(description, firstIndex))
            for i, idx := range allIndices {
                fmt.Println("Index", i, "=", idx[0], "-",
                    idx[1], "=", getSubstring(description, idx))
            }
        }
    Output:
        First index 0 - 5 = Kayak
        Index 0 = 0 - 5 = Kayak
        Index 1 = 9 - 13 = boat

████████████████████████████████████████████████████████████████████████
193.FindString and FindAllString methods
    If you dont need to know the location of the matches, then the FindString and FindAllString
    methods are more useful because their results are the substrings matched by the regular expression
    Getting Match Substrings
    example:
        package main
        import (
            "fmt"
            "regexp"
        )
        func main() {
            pattern := regexp.MustCompile("K[a-z]{4}|[A-z]oat")
            description := "Kayak. A boat for one person."
            firstMatch := pattern.FindString(description)
            allMatches := pattern.FindAllString(description, -1)
            fmt.Println("First match:", firstMatch)
            for i, m := range allMatches {
                fmt.Println("Match", i, "=", m)
            }
        }
    Output:
        First match: Kayak
        Match 0 = Kayak
        Match 1 = boat
████████████████████████████████████████████████████████████████████████
194.Splitting Strings Using a Regular Expression
    The Split method splits a string using the matches made by a regular expression, which can provide a more
    flexible alternative to the splitting functions
    example:
        package main
        import (
            "fmt"
            "regexp"
        )
        func main() {
            pattern := regexp.MustCompile(" |boat|one")
            description := "Kayak. A boat for one person."
            split := pattern.Split(description, -1)
            for _, s := range split {
                if s != "" {
                    fmt.Println("Substring:", s)
                }
            }
        }
    Output:
        Substring: Kayak.
        Substring: A
        Substring: for
        Substring: person.
    example:
        package main
        import (
            "fmt"
            "regexp"
        )
        func main() {
            pattern := regexp.MustCompile(" |boat|one")
            description := "Kayak. A boat | | | | for one person."
            split := pattern.Split(description, -2)
            for _, s := range split {
                if s != "" {
                    fmt.Println("Substring:", s)
                }
            }
        }
    Output:
        Substring: Kayak.
        Substring: A
        Substring: |
        Substring: |
        Substring: |
        Substring: |
        Substring: for
        Substring: person.
████████████████████████████████████████████████████████████████████████
195.Subexpressions
    Subexpressions allow parts of a regular expression to be accessed, which can make it easier to extract
    substrings from within a matched region.
    The pattern in this example matches a specific sentence structure.
    example:
        package main
        import (
            "fmt"
            "regexp"
        )
        func main() {
            pattern := regexp.MustCompile("A [A-z]* for [A-z]* person")
            description := "Kayak. A boat for one person."
            str := pattern.FindString(description)
            fmt.Println("Match:", str)
        }
    Output:
        Match: A boat for one person
████████████████████████████████████████████████████████████████████████
196.FindStringSubmatch method
    The FindStringSubmatch method performs the same
    task as FindString, but also includes the substrings matched by the expressions in its result.
    example:
        package main
        import (
            "fmt"
            "regexp"
        )
        func main() {
            pattern := regexp.MustCompile("A ([A-z]*) for ([A-z]*) person")
            description := "Kayak. A boat for one person."
            subs := pattern.FindStringSubmatch(description)
            for _, s := range subs {
                fmt.Println("Match:", s)
            }
        }
    Output:
        Match: A boat for one person
        Match: boat
        Match: one
████████████████████████████████████████████████████████████████████████
197.The Regexp Methods for Subexpressions
    Name                            Description
    ---------------------------     ------------------------------------------------------------------
    FindStringSubmatch(s)           This method returns a slice containing the first match made by the
                                    pattern and the text for the subexpressions that the pattern defines.
    FindAllStringSubmatch(s, max)   This method returns a slice containing all the matches and the text
                                    for the subexpressions. The int argument is used to specify the
                                    maximum number of matches. A value of -1 specifies all matches.
    FindStringSubmatchIndex(s)      This method is equivalent to FindStringSubmatch but returns
                                    indices rather than substrings.
                                    FindAllStringSubmatchIndex
    (s, max)                        This method is equivalent to FindAllStringSubmatch but returns
                                    indices rather than substrings.
    NumSubexp()                     This method returns the number of subexpressions.
    SubexpIndex(name)               This method returns the index of the subexpression with the
                                    specified name or -1 if there is no such subexpression.
    SubexpNames()                   This method returns the names of the subexpressions, expressed in
                                    the order in which they are defined.
████████████████████████████████████████████████████████████████████████
198.SubexpIndex(name) method
    he syntax for assigning names to subexpressions is awkward: within the parentheses, a question mark,
    followed by an uppercase P, followed by the name within angle brackets.
    pattern := regexp.MustCompile("A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) person")
    The subexpressions are given the names type and capacity. The SubexpIndex method returns the
    position of a named subexpression in the results, which allows me to get the substrings matched by the type
    and capacity subexpressions.

    example:
        package main
        import (
            "fmt"
            "regexp"
        )
        func main() {
            pattern := regexp.MustCompile(
                "A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) person")
            description := "Kayak. A boat for one person."
            subs := pattern.FindStringSubmatch(description)
            for _, name := range []string { "type", "capacity" } {
                fmt.Println(name, "=", subs[pattern.SubexpIndex(name)])
            }
        }
    Output:
        type = boat
        capacity = one
████████████████████████████████████████████████████████████████████████
199.Replacing Substrings Using a Regular Expression
    Name                                Description
    -----------------------------       ---------------------------------------------------------
    ReplaceAllString(s, template)       This method replaces the matched portion of the string s with the
                                        specified template, which is expanded before it is included in the result to
                                        incorporate subexpressions.
    ReplaceAllLiteralString(s, sub)     This method replaces the matched portion of the string s with the
                                        specified content, which is included in the result without being expanded
                                        for subexpressions.
    ReplaceAllStringFunc(s, func)       This method replaces the matched portion of the string s with the result
                                        produced by the specified function.
████████████████████████████████████████████████████████████████████████
200.ReplaceAllString method
    The result from the ReplaceAllString method is a string with the replaced content.
    Notice that the template is responsible for only part of the result from the ReplaceAllString method,
    The first part of the description string—the word Kayak, followed by a period and a space, is
    not matched by the regular expression and is included in the result without being modified.

    ■ Tip
        Use the ReplaceAllLiteralString method if you want to replace content without the new
        substring being interpreted for subexpressions.

    example:
        package main
        import (
            "fmt"
            "regexp"
        )
        func main() {
            pattern := regexp.MustCompile(
                "A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) person")
            description := "Kayak. A boat for one person."
           template := "(type: ${type}, capacity: ${capacity})"
            replaced := pattern.ReplaceAllString(description, template)
            fmt.Println(replaced)
        }
    Output:
        Kayak. (type: boat, capacity: one).
████████████████████████████████████████████████████████████████████████
201.ReplaceAllStringFunc method
    The ReplaceAllStringFunc method replaces the matched section of a string 
    with content generated by a function
    example:
        package main
        import (
            "fmt"
            "regexp"
        )
        func main() {
            pattern := regexp.MustCompile(
                "A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) person")
            description := "Kayak. A boat for one person."
            replaced := pattern.ReplaceAllStringFunc(description, func(s string) string {
                return "This is the replacement content"
            })
            fmt.Println(replaced)
        }
    Output:
        Kayak. This is the replacement content.
████████████████████████████████████████████████████████████████████████
202.Formatting and Scanning Strings
    Formatting is the process of composing a new string from one or more data values, 
    while scanning is the process of parsing values from a string.
    
    What are they?
    Formatting is the process of composing values into a string. Scanning is the process
    of parsing a string for the values it contains.

    Why are they useful?
    Formatting a string is a common requirement and is used to produce strings for
    everything from logging and debugging to presenting the user with information.
    Scanning is useful for extracting data from strings, such as from HTTP requests or
    user input.

    How are they used?
    Both sets of features are provided through functions defined in the fmt package.

    Are there any pitfalls or limitations?
    The templates used to format strings can be hard to read, and there is no built-in
    function that allows a formatted string to be created to which a newline character
    is appended automatically.

    Are there any alternatives?
    Larger amounts of text and HTML content can be generated using the template features.
████████████████████████████████████████████████████████████████████████
203.fmt package
    The fmt package provides functions for composing and writing strings.
    Some of these functions use writers, which are part of the Go support for input/output.
    The Basic fmt Functions for Composing and Writing Strings.

    Name                        Description
    --------------------------  --------------------------------------------------------------------
    Print(...vals)              This function accepts a variable number of arguments and writes out their
                                values to the standard out. Spaces are added between values that are not
                                strings.
    Println(...vals)            This function accepts a variable number of arguments and writes out
                                their values to the standard out, separated by spaces and followed by a
                                newline character.
    Fprint(writer, ...vals)     This function writes out a variable number of arguments to the specified writer,
                                Spaces are added between values that are not strings.
    Fprintln(writer, ...vals)   This function writes out a variable number of arguments to the specified writer
                                followed by a newline character. Spaces are added between all values.
████████████████████████████████████████████████████████████████████████
204.Println and Fprintln functions
    The Println and Fprintln functions add spaces between all the values, but the
    Print and Fprint functions only add spaces between values that are not strings.

    since the Print function adds spaces only between pairs of nonstring values, the results are different.
    
    example:
        package main
        import "fmt"
        func main() {
            fmt.Println("Product:", 1234 , "Price:", 5555)
            fmt.Print("Product:", 1234 , "Price:", 5555, "\n")
        }
    Output:
        Product: 1234 Price: 5555
        Product:1234Price:5555
████████████████████████████████████████████████████████████████████████
205.fmt.Printf
    The Printf function accepts a template string and a series of values. The template is scanned for verbs,
    which are denoted by the percentage sign (the % character) followed by a format specifier.
    
    The first verb is %v, and it specifies the default representation for a type. For a string value, for example,
    %v simply includes the string in the output. The %4.2f verb specifies the format for a floating-point value,
    with 4 digits before the decimal point and 2 digits after. The values for the template verbs are taken from the
    remaining arguments, used in the order they are specified. For the example, this means the %v verb is used
    to format the Product.Name value, and the %4.2f verb is used to format the Product.Price value. These
    values are formatted, inserted into the template string, and written out to the console.

    example:
        package main
        import "fmt"
        type Product struct {
            Name, Category string
            Price float64
        }
        var Kayak = Product {
            Name: "Kayak",
            Category: "Watersports",
            Price: 275,
        }
        var Products = []Product {
            { "Kayak", "Watersports", 279 },
            { "Lifejacket", "Watersports", 49.95 },
            { "Soccer Ball", "Soccer", 19.50 },
            { "Corner Flags", "Soccer", 34.95 },
            { "Stadium", "Soccer", 79500 },
            { "Thinking Cap", "Chess", 16 },
            { "Unsteady Chair", "Chess", 75 },
            { "Bling-Bling King", "Chess", 1200 },
        }
        func main() {
            fmt.Printf("Product: %v, Price: $%4.2f", Kayak.Name, Kayak.Price)
        }
    Output:
        Product: Kayak, Price: $275.00  
████████████████████████████████████████████████████████████████████████
206.The fmt Functions for Formatting Strings
    Name                            Description
    ---------------------------     -----------------------------------------
    Sprintf(t, ...vals)             This function returns a string, which is created by processing the template t.
                                    The remaining arguments are used as values for the template verbs.
    Printf(t, ...vals)              This function creates a string by processing the template t.
                                    The remaining arguments are used as values for the template verbs.
                                    The string is written to the standard out.
    Fprintf(writer, t, ...vals)     This function creates a string by processing the template t.
                                    The remaining arguments are used as values for the template verbs.
                                    The string is written to a Writer, which is described in Chapter 20.
    Errorf(t, ...values)            This function creates an error by processing the template t.
                                    The remaining arguments are used as values for the template verbs. The
                                    result is an error value whose Error method returns the formatted string.
████████████████████████████████████████████████████████████████████████
207.fmt.Sprintf
    Both of the formatted strings in this example use the %v value, 
    which writes out values in their default form.

    example:
        package main

        import "fmt"
        
        type Product struct {
            Name, Category string
            Price          float64
        }
        
        var Kayak = Product{
            Name:     "Kayak",
            Category: "Watersports",
            Price:    275,
        }
        var Products = []Product{
            {"Kayak", "Watersports", 279},
            {"Lifejacket", "Watersports", 49.95},
            {"Soccer Ball", "Soccer", 19.50},
            {"Corner Flags", "Soccer", 34.95},
            {"Stadium", "Soccer", 79500},
            {"Thinking Cap", "Chess", 16},
            {"Unsteady Chair", "Chess", 75},
            {"Bling-Bling King", "Chess", 1200},
        }
        
        func getProductName(index int) (name string, err error) {
            if len(Products) > index {
                name = fmt.Sprintf("Name of product: %v", Products[index].Name)
            } else {
                err = fmt.Errorf("error for index %v", index)
            }
            return
        }
        func main() {
            name, _ := getProductName(1)
            fmt.Println(name)
            _, err := getProductName(10)
            fmt.Println(err.Error())
        }
    Output:
        Name of product: Lifejacket
        error for index 10
████████████████████████████████████████████████████████████████████████
208.the General-Purpose Formatting Verbs
    The general-purpose verbs can be used to display any value.
    The Formatting Verbs for Any Value

    Verb    Description
    ------  -----------------------------------------------
    %v      This verb displays the default format for the value. Modifying the verb with a plus sign (%+v) includes
            field names when writing out struct values.
    %#v     This verb displays a value in a format that could be used to re-create the value in a Go code file.
    %T      This verb displays the Go type of a value.

    example:
        package main
        import "fmt"
        type Product struct {
            Name, Category string
            Price          float64
        }
        var Kayak = Product{
            Name:     "Kayak",
            Category: "Watersports",
            Price:    275,
        }
        var Products = []Product{
            {"Kayak", "Watersports", 279},
            {"Lifejacket", "Watersports", 49.95},
            {"Soccer Ball", "Soccer", 19.50},
            {"Corner Flags", "Soccer", 34.95},
            {"Stadium", "Soccer", 79500},
            {"Thinking Cap", "Chess", 16},
            {"Unsteady Chair", "Chess", 75},
            {"Bling-Bling King", "Chess", 1200},
        }
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
        func main() {
            Printfln("Value: %v", Kayak)
            Printfln("Go syntax: %#v", Kayak)
            Printfln("Type: %T", Kayak)
        }
    Output:
        Value: {Kayak Watersports 275}
        Go syntax: main.Product{Name:"Kayak", Category:"Watersports", Price:275}
        Type: main.Product
████████████████████████████████████████████████████████████████████████
209.Controlling Struct Formatting
    Go has a default format for all data types that the %v verb relies on. 
    For structs, the default value lists the field values within curly braces. 
    The default verb can be modified with a plus sign to include the field names in the output.
    example:
        package main
        import "fmt"
        type Product struct {
            Name, Category string
            Price          float64
        }
        var Kayak = Product{
            Name:     "Kayak",
            Category: "Watersports",
            Price:    275,
        }
        var Products = []Product{
            {"Kayak", "Watersports", 279},
            {"Lifejacket", "Watersports", 49.95},
            {"Soccer Ball", "Soccer", 19.50},
            {"Corner Flags", "Soccer", 34.95},
            {"Stadium", "Soccer", 79500},
            {"Thinking Cap", "Chess", 16},
            {"Unsteady Chair", "Chess", 75},
            {"Bling-Bling King", "Chess", 1200},
        }
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
        func main() {
            Printfln("Value: %v", Kayak)
            Printfln("Value with fields: %+v", Kayak)
        }    
    Output:
        Value: {Kayak Watersports 275}
        Value with fields: {Name:Kayak Category:Watersports Price:275}
████████████████████████████████████████████████████████████████████████
210.Custom Struct Format
    example:
        package main
        import "fmt"
        type Product struct {
            Name, Category string
            Price          float64
        }
        var Kayak = Product{
            Name:     "Kayak",
            Category: "Watersports",
            Price:    275,
        }
        var Products = []Product{
            {"Kayak", "Watersports", 279},
            {"Lifejacket", "Watersports", 49.95},
            {"Soccer Ball", "Soccer", 19.50},
            {"Corner Flags", "Soccer", 34.95},
            {"Stadium", "Soccer", 79500},
            {"Thinking Cap", "Chess", 16},
            {"Unsteady Chair", "Chess", 75},
            {"Bling-Bling King", "Chess", 1200},
        }
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
        func (p Product) String() string {
            return fmt.Sprintf("Product: %v, Price: $%4.2f", p.Name, p.Price)
        }
        func main() {
            Printfln("Value: %v", Kayak)
            Printfln("Value with fields: %+v", Kayak)
        }
    Output:
        Value: Product: Kayak, Price: $275.00
        Value with fields: Product: Kayak, Price: $275.00
████████████████████████████████████████████████████████████████████████
211.Formating Arrays, Slices, Maps
    When arrays and slices are represented as strings, the output is a set of square brackets, within which
    are the individual elements, like this:
    example:
        [Kayak Lifejacket Paddle]
    
    Notice that no commas are separating the elements. When maps are represented as strings, the key-
    value pairs are displayed within square brackets, preceded by the map keyword, like this:
    example:
        map[1:Kayak 2:Lifejacket 3:Paddle]
████████████████████████████████████████████████████████████████████████
212.Integer Formatting Verbs
    Verb    Description
    ------  ------------------------------------
    %b      This verb displays an integer value as a binary string.
    %d      This verb displays an integer value as a decimal string. This is the default format for integer
            values, applied when the %v verb is used.
    %o, %O  These verbs display an integer value as an octal string. The %O verb adds the 0o prefix.
    %x, %X  These verbs display an integer value as a hexadecimal string. The letters A–F are displayed in
            lowercase by the %x verb and in uppercase by the %X verb.
    
    example:
        package main
        import "fmt"
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
        func main() {
            number := 250
            Printfln("Binary: %b", number)
            Printfln("Decimal: %d", number)
            Printfln("Octal: %o, %O", number, number)
            Printfln("Hexadecimal: %x, %X", number, number)
        }
    Output:
        Binary: 11111010
        Decimal: 250
        Octal: 372, 0o372
        Hexadecimal: fa, FA
████████████████████████████████████████████████████████████████████████
213.Floating-Point Formatting Verbs
    The Formatting Verbs for Floating-Point Values
    Verb    Description
    -----   ------------------------------------
    %b      This verb displays a floating-point value with an exponent and without a decimal place.
    %e, %E  These verbs display a floating-point value with an exponent and a decimal place. The %e uses a
            lowercase exponent indicator, while %E uses an uppercase indicator.
    %f, %F  These verbs display a floating-point value with a decimal place but no exponent. The %f and %F
            verbs produce the same output.
    %g      This verb adapts to the value it displays. The %e format is used for values with large exponents,
            and the %f format is used otherwise. This is the default format, applied when the %v verb is
            used.
    %G      This verb adapts to the value it displays. The %E format is used for values with large exponents,
            and the %f format is used otherwise.
    %x, %X  These verbs display a floating-point value in hexadecimal notation, with lowercase (%x) or
            uppercase (%X) letters.

    example:
        package main
        import "fmt"
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template + "\n", values...)
        }
        func main() {
            number := 279.00
            Printfln("Decimalless with exponent: %b", number)
            Printfln("Decimal with exponent: %e", number)
            Printfln("Decimal without exponent: %f", number)
            Printfln("Hexadecimal: %x, %X", number, number)
        }
    Output:
        Decimalless with exponent: 4908219906392064p-44
        Decimal with exponent: 2.790000e+02
        Decimal without exponent: 279.000000
        Hexadecimal: 0x1.17p+08, 0X1.17P+08
████████████████████████████████████████████████████████████████████████
214.Controlling Formatting
    example:
        package main
        import "fmt"
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template + "\n", values...)
        }
        func main() {
            number := 279.00
            Printfln("Decimal without exponent: >>%8.2f<<", number)
        }
    Output:
        Decimal without exponent: >>  279.00<<
████████████████████████████████████████████████████████████████████████
215.Specifying Precision
    example:
        package main
        import "fmt"
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template + "\n", values...)
        }
        func main() {
            number := 279.00
            Printfln("Decimal without exponent: >>%.2f<<", number)
        }
    Output:
        Decimal without exponent: >>279.00<<
████████████████████████████████████████████████████████████████████████
216.The Formatting Verb Modifiers
    Modifier    Description
    --------    --------------------------
    +           This modifier (the plus sign) always prints a sign, positive or negative, for numeric values.

    0           This modifier uses zeros, rather than spaces, as padding when the width is greater than the
                number of characters required to display the value.

    -           This modifier (the subtracts symbol) adds padding to the right of the number, rather than
                the left.
    
    example:
        package main
        import "fmt"
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template + "\n", values...)
        }
        func main() {
            number := 279.00
            Printfln("Sign: >>%+.2f<<", number)
            Printfln("Zeros for Padding: >>%07.2f<<", number)
            Printfln("Zeros for Padding: >>%08.2f<<", number)
            Printfln("Right Padding: >>%-8.2f<<", number)
        }
    Output:
        Sign: >>+279.00<<
        Zeros for Padding: >>0279.00<<
        Zeros for Padding: >>00279.00<<
        Right Padding: >>279.00  <<
████████████████████████████████████████████████████████████████████████
217.The Formatting Verbs for Strings and Runes
    Verb    Description
    ----    ----------------------------------------------------
    %s      This verb displays a string. This is the default format, applied when the %v verb is used.

    %c      This verb displays a character. Care must be taken to avoid slicing strings into individual bytes, as
            explained in the text after the table.

    %U      This verb displays a character in the Unicode format so that the output begins with U+ followed by
            a hexadecimal character code.
    
    example:
        package main
        import "fmt"
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template + "\n", values...)
        }
        func main() {
            name := "Kayak"
            Printfln("String: %s", name)
            Printfln("Character: %c", []rune(name)[0])
            Printfln("Unicode: %U", []rune(name)[0])
        }
    Output:
        String: Kayak
        Character: K
        Unicode: U+004B
████████████████████████████████████████████████████████████████████████
218.The bool Formatting Verb
    Verb    Description
    ----    -------------
    %t      This verb formats bool values and displays true or false.

    example:
        package main
        import "fmt"
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template + "\n", values...)
        }
        func main() {
            name := "Kayak"
            Printfln("Bool: %t", len(name) > 1)
            Printfln("Bool: %t", len(name) > 100)
        }
    Output:
        Bool: true
        Bool: false
████████████████████████████████████████████████████████████████████████
219.The Pointer Formatting Verb
    Verb    Description
    ----    -----------------
    %p      This verb displays a hexadecimal representation of the pointer's storage location.

    example:
        package main
        import "fmt"
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template + "\n", values...)
        }
        func main() {
            name := "Kayak"
            Printfln("Pointer: %p", &name)
        }
    Output:
        Pointer: 0xc00004a240
████████████████████████████████████████████████████████████████████████
220.The fmt Functions for Scanning Strings
    Name                                Description
    ------------------------            ---------------------------------------------------------------------
    Scan(...vals)                       This function reads text from the standard in and stores the space-
                                        separated values into specified arguments. Newlines are treated as
                                        spaces, and the function reads until it has received values for all of its
                                        arguments. The result is the number of values that have been read and an
                                        error that describes any problems.
    Scanln(...vals)                     This function works in the same way as Scan but stops reading when it
                                        encounters a newline character.
    Scanf(template, ...vals)            This function works in the same way as Scan but uses a template string
                                        to select the values from the input it receives.
    Fscan(reader, ...vals)              This function reads space-separated values from the specified reader,
                                        which is described in Chapter 20. Newlines are treated as spaces, and
                                        the function returns the number of values that have been read and an
                                        error that describes any problems.
    Fscanln(reader, ...vals)            This function works in the same way as Fscan but stops reading when it
                                        encounters a newline character.
    Fscanf(reader, template, ...vals)   This function works in the same way as Fscan but uses a template to
                                        select the values from the input it receives.
    Sscan(str, ...vals)                 This function scans the specified string for space-separated values,
                                        which are assigned to the remaining arguments. The result is the
                                        number of values scanned and an error that describes any problems.
    Sscanf(str, template, ...vals)      This function works in the same way as Sscan but uses a template to
                                        select values from the string.
    Sscanln(str, template, ...vals)     This function works in the same way as Sscanf but stops scanning the
                                        string as soon as a newline character is encountered.
████████████████████████████████████████████████████████████████████████
221.Scanning a String
    example:
        package main
        import "fmt"
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template + "\n", values...)
        }
        func main() {
            var name string
            var category string
            var price float64
            fmt.Print("Enter text to scan: ")
            n, err := fmt.Scan(&name, &category, &price)
            if (err == nil) {
                Printfln("Scanned %v values", n)
                Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
            } else {
                Printfln("Error: %v", err.Error())
            }
        }
    Output:
        Enter text to scan: asd asd 123
        Scanned 3 values
        Name: asd, Category: asd, Price: 123.00
████████████████████████████████████████████████████████████████████████
222.Scanln Function
    example:
        package main
        import "fmt"
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template + "\n", values...)
        }
        func main() {
            var name string
            var category string
            var price float64
            fmt.Print("Enter text to scan: ")
            n, err := fmt.Scanln(&name, &category, &price)
            if (err == nil) {
                Printfln("Scanned %v values", n)
                Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
            } else {
                Printfln("Error: %v", err.Error())
            }
        }
    Output:
      Error: unexpected newline
████████████████████████████████████████████████████████████████████████
223.Sscan Function
    example:
        package main
        import "fmt"
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template + "\n", values...)
        }
        func main() {
            var name string
            var category string
            var price float64
            source := "Lifejacket Watersports 48.95"
            n, err := fmt.Sscan(source, &name, &category, &price)
            if (err == nil) {
                Printfln("Scanned %v values", n)
                Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
            } else {
                Printfln("Error: %v", err.Error())
            }
        }
    Output:
        Scanned 3 values
        Name: Lifejacket, Category: Watersports, Price: 48.95   
████████████████████████████████████████████████████████████████████████
224.Scanning Template
    The template ignores the term Product, skipping that part of the string and
    allowing the scanning to begin with the next term.
    example:
        package main
        import "fmt"
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template + "\n", values...)
        }
        func main() {
            var name string
            var category string
            var price float64
            source := "Product Lifejacket Watersports 48.95"
            template := "Product %s %s %f"
            n, err := fmt.Sscanf(source, template, &name, &category, &price)
            if (err == nil) {
                Printfln("Scanned %v values", n)
                Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
            } else {
                Printfln("Error: %v", err.Error())
            }
        }
    Output:
        Scanned 3 values
        Name: Lifejacket, Category: Watersports, Price: 48.95
████████████████████████████████████████████████████████████████████████
225.Putting Math Functions and Sorting Data in Context
    What are they?
    The math functions allow common calculations to be performed. Random
    numbers are numbers generated in a sequence that is difficult to predict. Sorting
    is the process of placing a sequence of values in a predetermined order.

    Why are they useful?
    These are features that are used throughout development.

    How are they used?
    These features are provided in the math, math/rand, and sort packages.

    Are there any pitfalls or limitations?
    Unless initialized with a seed value, the numbers produced by the math/rand
    package are not random.

    Are there any alternatives?
    You could implement both sets of features from scratch, although these packages
    are provided so that this is not required.
████████████████████████████████████████████████████████████████████████
226.Useful Functions from the math Package
    Name                Description
    --------------      ----------------------------------------------------------
    Abs(val)            This function returns the absolute value of a float64 value, meaning the distance
                        from zero without considering direction.
    Ceil(val)           This function returns the smallest integer that is equal to or greater than the specified
                        float64 value. The result is also a float64 value, even though it represents an integer
                        number.
    Copysign(x, y)      This function returns a float64 value, which is the absolute value of x with the sign of y.
    Floor(val)          This function returns the largest integer that is smaller or equal to the specified
                        float64 value. The result is also a float64 value, even though it represents an integer number.
    Max(x, y)           This function returns whichever of the specified float64 value is the largest.
    Min(x, y)           This function returns whichever of the specified float64 value is smallest.
    Mod(x, y)           This function returns the remainder of x/y.
    Pow(x, y)           This function returns x raised to the exponent y.
    Round(val)          This function rounds the specified value to the nearest integer, rounding half values
                        up. The result is a float64 value, even though it represents an integer.
    RoundToEven(val)    This function rounds the specified value to the nearest integer, rounding half values
                        to the nearest even number. The result is a float64 value, even though it represents
                        an integer.

    example:
        package main
        import "math"
        func main() {
            val1 := 279.00
            val2 := 48.95
            Printfln("Abs: %v", math.Abs(val1))
            Printfln("Ceil: %v", math.Ceil(val2))
            Printfln("Copysign: %v", math.Copysign(val1, -5))
            Printfln("Floor: %v", math.Floor(val2))
            Printfln("Max: %v", math.Max(val1, val2))
            Printfln("Min: %v", math.Min(val1, val2))
            Printfln("Mod: %v", math.Mod(val1, val2))
            Printfln("Pow: %v", math.Pow(val1, 2))
            Printfln("Round: %v", math.Round(val2))
            Printfln("RoundToEven: %v", math.RoundToEven(val2))
        }
    Output:
        Abs: 279
        Ceil: 49
        Copysign: -279
        Floor: 48
        Max: 279
        Min: 48.95
        Mod: 34.249999999999986
        Pow: 77841
        Round: 49
        RoundToEven: 49
████████████████████████████████████████████████████████████████████████
227.The Limit Constants
    Name                        Description
    ----------------            -----------------------------
    MaxInt8                     These constants represent the largest and smallest values that can be stored
    MinInt8                     using an int8.
    MaxInt16                    These constants represent the largest and smallest values that can be stored
    MinInt16                    using an int16.
    MaxInt32                    These constants represent the largest and smallest values that can be stored
    MinInt32                    using an int32.
    MaxInt64                    These constants represent the largest and smallest values that can be stored
    MinInt64                    using an int64.
    MaxUint8                    This constant represents the largest value that can be represented using a
                                uint8.The smallest value is zero.
    MaxUint16                   This constant represents the largest value that can be represented using a
                                uint16. The smallest value is zero.
    MaxUint32                   This constant represents the largest value that can be represented using a
                                uint32. The smallest value is zero.
    MaxUint64                   This constant represents the largest value that can be represented using a
                                uint64. The smallest value is zero.
    MaxFloat32                  These constants represent the largest values that can be represented using
    MaxFloat64                  float32 and float64 values.
    SmallestNonzeroFloat32      These constants represent the smallest nonzero values that can be
    SmallestNonzeroFloat32      represented using float32 and float64 values.
████████████████████████████████████████████████████████████████████████
228.Generating Random Numbers
    Useful math/rand Functions
    Name                    Description
    --------------------    ---------------------------------------
    Seed(s)                 This function sets the seed value using the specified int64 value.
    Float32()               This function generates a random float32 value between 0 and 1.
    Float64()               This function generates a random float64 value between 0 and 1.
    Int()                   This function generates a random int value.
    Intn(max)               This function generates a random int smaller than a specified value, as
                            described after the table.
    UInt32()                This function generates a random uint32 value.
    UInt64()                This function generates a random uint64 value.
    Shuffle(count, func)    This function is used to randomize the order of elements, as described after
                            the table.
████████████████████████████████████████████████████████████████████████
229.rand.Int()
    example:
        package main
        import "math/rand"
        func main() {
            for i := 0; i < 5; i++ {
                Printfln("Value %v : %v", i, rand.Int())
            }
        }
    Output:
        Value 0 : 5577006791947779410
        Value 1 : 8674665223082153551
        Value 2 : 6129484611666145821
        Value 3 : 4037200794235010051
        Value 4 : 3916589616287113937
████████████████████████████████████████████████████████████████████████
230.rand.Seed()
    Example:
        package main
        import (
            "math/rand"
            "time"
        )
        func main() {
            rand.Seed(time.Now().UnixNano())
            for i := 0; i < 5; i++ {
                Printfln("Value %v : %v", i, rand.Int())
            }
        }
    Output:
        Value 0 : 8113726196145714527
        Value 1 : 3479565125812279859
        Value 2 : 8074476402089812953
        Value 3 : 3916870404047362448
        Value 4 : 8226545715271170755
████████████████████████████████████████████████████████████████████████
231.rand.Intn(rangeNumber)
    Generating a Random Number Within a Specific Range

    example:
        package main
        import (
            "math/rand"
            "time"
        )
        func main() {
            rand.Seed(time.Now().UnixNano())
            for i := 0; i < 5; i++ {
                Printfln("Value %v : %v", i, rand.Intn(10))
            }
        }
    Output:
        Value 0 : 7
        Value 1 : 5
        Value 2 : 4
        Value 3 : 0
        Value 4 : 7
████████████████████████████████████████████████████████████████████████
232.Specifying a Lower Bound
    example:
    package main
    import (
        "math/rand"
        "time"
    )
    func IntRange(min, max int) int {
        return rand.Intn(max - min) + min
    }
    func main() {
        rand.Seed(time.Now().UnixNano())
        for i := 0; i < 5; i++ {
            Printfln("Value %v : %v", i, IntRange(10, 20))
        }
    }
    Output:
        Value 0 : 10
        Value 1 : 19
        Value 2 : 11
        Value 3 : 10
        Value 4 : 17
████████████████████████████████████████████████████████████████████████
233.Shuffling Elements
    example:
        package main
        import (
            "math/rand"
            "time"
            "fmt"
        )
        var names = []string{"Alice", "Bob", "Charlie", "Dora", "Edith"}
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
        func main() {
            rand.Seed(time.Now().UnixNano())
            rand.Shuffle(len(names), func(first, second int) {
                names[first], names[second] = names[second], names[first]
            })
            for i, name := range names {
                Printfln("Index %v: Name: %v", i, name)
            }
        }
    Output:
        Index 0: Name: Edith
        Index 1: Name: Dora
        Index 2: Name: Charlie
        Index 3: Name: Alice
        Index 4: Name: Bob
████████████████████████████████████████████████████████████████████████
234.The Basic Functions for Sorting
    Name                        Description
    --------------------        -------------------------------------------
    Float64s(slice)             This function sorts a slice of float64 values. The elements are sorted in place.
    Float64sAreSorted(slice)    This function returns true if the elements in the specified float64 slice are in order.
    Ints(slice)                 This function sorts a slice of int values. The elements are sorted in place.
    IntsAreSorted(slice)        This function returns true if the elements in the specified int slice are in order.
    Strings(slice)              This function sorts a slice of string values. The elements are sorted in place.
    StringsAreSorted(slice)     This function returns true if the elements in the specified string slice are in order.
████████████████████████████████████████████████████████████████████████
235.Sorting Slices
    example:
        package main
        import (
            "fmt"
            "sort"
        )
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
        func main() {
            ints := []int{9, 4, 2, -1, 10}
            Printfln("Ints: %v", ints)
            
            sort.Ints(ints)
            Printfln("Ints Sorted: %v", ints)
            
            floats := []float64{279, 48.95, 19.50}
            Printfln("Floats: %v", floats)
            
            sort.Float64s(floats)
            Printfln("Floats Sorted: %v", floats)
            
            strings := []string{"Kayak", "Lifejacket", "Stadium"}
            Printfln("Strings: %v", strings)
            
            
            if !sort.StringsAreSorted(strings) {
                sort.Strings(strings)
                Printfln("Strings Sorted: %v", strings)
            } else {
                Printfln("Strings Already Sorted: %v", strings)
            }
        }
    Output:
        Ints: [9 4 2 -1 10]
        Ints Sorted: [-1 2 4 9 10]
        Floats: [279 48.95 19.5]
        Floats Sorted: [19.5 48.95 279]
        Strings: [Kayak Lifejacket Stadium]
        Strings Already Sorted: [Kayak Lifejacket Stadium]
████████████████████████████████████████████████████████████████████████
236.Creating a Sorted Copy of a Slice
    example:
        package main
        import (
            "fmt"
            "sort"
        )
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
        func main() {
            ints := []int{9, 4, 2, -1, 10}
            
            sortedInts := make([]int, len(ints))
            
            copy(sortedInts, ints)
            
            sort.Ints(sortedInts)
            
            
            Printfln("Ints: %v", ints)
            Printfln("Ints Sorted: %v", sortedInts)
        }
    Output:
        Ints: [9 4 2 -1 10]
        Ints Sorted: [-1 2 4 9 10]
        
████████████████████████████████████████████████████████████████████████
237.The Functions for Searching Sorted Data
    Name                        Description
    ----------------------      ----------------------------------
    SearchInts(slice, val)      This function searches the sorted slice for the specified int value. The
                                result is the index of the specified value or, if the value is not found, the
                                index at which the value can be inserted while maintaining the sorted order.
    SearchFloat64s(slice, val)  This function searches the sorted slice for the specified float64 value.
                                The result is the index of the specified value or, if the value is not found,
                                the index at which the value can be inserted while maintaining the
                                sorted order.
    SearchStrings(slice, val)   This function searches the sorted slice for the specified string value.
                                The result is the index of the specified value or, if the value is not found,
                                the index at which the value can be inserted while maintaining the
                                sorted order.
    Search(count, testFunc)     This function invokes the test function for the specified number of
                                elements. The result is the index for which the function returns true.
                                If there is no match, then the result is the index at which the specified
                                value can be inserted to maintain the sorted order.
████████████████████████████████████████████████████████████████████████
238.sort.SearchInts()
    When a value is located, the functions return its position in the slice. 
    But unusually, if the value is not found, then the result is the position it can be
    inserted while maintaining the sort order.

    example:
        package main
    import (
        "fmt"
        "sort"
    )
    func Printfln(template string, values ...interface{}) {
        fmt.Printf(template+"\n", values...)
    }
    func main() {
        ints := []int{9, 1, 2, -1, 10,5}
        sortedInts := make([]int, len(ints))
        copy(sortedInts, ints)
        sort.Ints(sortedInts)
        Printfln("Ints: %v", ints)
        Printfln("Ints Sorted: %v", sortedInts)
        
        // ===========================================
        indexOf4 := sort.SearchInts(sortedInts, 4)
        Printfln("Index of 4: %v", indexOf4)
        
        indexOf3 := sort.SearchInts(sortedInts, 3)
        Printfln("Index of 3: %v", indexOf3)
    }
Output:
    Ints: [9 1 2 -1 10 5]
    Ints Sorted: [-1 1 2 5 9 10]
    Index of 4: 3
    Index of 3: 3
████████████████████████████████████████████████████████████████████████
239.sort.SearchInts() bool returned
    example:
        package main
        import (
            "fmt"
            "sort"
        )
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
        func main() {
            ints := []int{9, 1, 2, -1, 10,5}
            sortedInts := make([]int, len(ints))
            copy(sortedInts, ints)
            sort.Ints(sortedInts)
            Printfln("Ints: %v", ints)
            Printfln("Ints Sorted: %v", sortedInts)	
            // =========================
            indexOf4:= sort.SearchInts(sortedInts, 4)
            indexOf3 := sort.SearchInts(sortedInts, 3)
            Printfln("Index of 4: %v (present: %v)", indexOf4, sortedInts[indexOf4] == 4)
            Printfln("Index of 3: %v (present: %v)", indexOf3, sortedInts[indexOf3] == 3)
        }
Output:
    Ints: [9 1 2 -1 10 5]
    Ints Sorted: [-1 1 2 5 9 10]
    Index of 4: 3 (present: false)
    Index of 3: 3 (present: false)

████████████████████████████████████████████████████████████████████████
240.The Methods Defined by the sort.Interface Interface
    Name            Description
    ----------      ----------------
    Len()           This method returns the number of items that will be sorted.
    Less(i, j)      This method returns true if the element at index i should appear in the sorted sequence
                    before the element j. If Less(i,j) and Less(j, i) are both false, then the elements are
                    considered equal.
    Swap(i, j)      This method swaps the elements at the specified indices.
████████████████████████████████████████████████████████████████████████
241.The Functions for Sorting Types That Implement Interface
    Name            Description
    ----------      ---------------------
    Sort(data)      This function uses the methods described in 240 Number to sort the specified data.
    Stable(data)    This function uses the methods described in 240 Number to sort the specified data
                    without changing the order of elements of equal value.
    IsSorted(data)  This function returns true if the data is in sorted order.
    Reverse(data)   This function reverses the order of the data.


    example:
    main.go:
        package main

        import (
            "fmt"
        )

        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }

        func main() {
            products := []Product{
                {"Kayak", 279},
                {"Lifejacket", 49.95},
                {"Soccer Ball", 19.50},
            }
            ProductSlices(products)
            
            for _, p := range products {
                Printfln("Name: %v, Price: %.2f", p.Name, p.Price)
            }

        }

        Name: Soccer Ball, Price: 19.50
        Name: Lifejacket, Price: 49.95
        Name: Kayak, Price: 279.00


████████████████████████████████████████████████████████████████████████
242.Sorting with a Comparison Function
    go mod init
    AND
    main.go File:
        package main
        import (
            "fmt"
        )
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
        func main() {
            products := []Product{
                {"Kayak", 279},
                {"Lifejacket", 49.95},
                {"Soccer Ball", 19.50},
            }
            SortWith(products, func(p1, p2 Product) bool {
                return p1.Name < p2.Name
            })
            for _, p := range products {
                Printfln("Name: %v, Price: %.2f", p.Name, p.Price)
            }    
        }

    productsort.go File:
        package main
        import "sort"
        type Product struct {
            Name  string
            Price float64
        }
        type ProductSlice []Product
        func ProductSlices(p []Product) {
            sort.Sort(ProductSlice(p))
        }
        func ProductSlicesAreSorted(p []Product) {
            sort.IsSorted(ProductSlice(p))
        }
        func (products ProductSlice) Len() int {
            return len(products)
        }
        func (products ProductSlice) Less(i, j int) bool {
            return products[i].Price < products[j].Price
        }
        func (products ProductSlice) Swap(i, j int) {
            products[i], products[j] = products[j], products[i]
        }
        type ProductComparison func(p1, p2 Product) bool
        type ProductSliceFlex struct {
		ProductSlice
		ProductComparison
	}
	func (flex ProductSliceFlex) Less(i, j int) bool {
		return flex.ProductComparison(flex.ProductSlice[i], flex.ProductSlice[j])
	}
	func SortWith(prods []Product, f ProductComparison) {
		sort.Sort(ProductSliceFlex{ prods, f})
	}

████████████████████████████████████████████████████████████████████████
243.Putting Dates, Times, and Durations in Context
    What are they?
    The features provided by the time package are used to represent
    specific moments in time and intervals or durations.

    Why are they useful?
    These features are useful in any application that needs to deal with
    calendaring or alarm and for the development of any feature that
    requires delays or notifications in the future.

    How are they used?
    The time package defines data types for representing dates and
    individual units of time and functions for manipulating them. There
    are also features integrated into the Go channel system.

    Are there any pitfalls or limitations?
    Dates can be complex, and care must be taken to deal with calendar
    and time zone issues.

    Are there any alternatives?
    These are optional features, and their use is not required.

████████████████████████████████████████████████████████████████████████
244.The Functions in the time Package for Creating Time Values
    Name                                    Description
    --------                                -----------------------------
    Now()                                   This function creates a Time representing the current moment in time.
    Date(y, m, d, h, min, sec, nsec, loc)   This function creates a Time representing a specified moment in time, which is
                                            expressed by the year, month, day, hour, minute, second, nanosecond, and Location
                                            arguments. (The Location type is described in the “Parsing Time Values from Strings”
                                            section.)
    Unix(sec, nsec)                         This function creates a Time value from the number of seconds and nanoseconds since
                                            January 1, 1970, UTC, commonly known as Unix time.
                                        
████████████████████████████████████████████████████████████████████████
245.The Methods for Accessing Time Components
    Name            Description
    --------        ----------------------------
    Date()          This method returns the year, month, and day components. The year and day are
                    expressed as int values and the month as a Month value.
    Clock()         This method returns the hour, minutes, and seconds components of the Time.
    Year()          This method returns the year component, expressed as an int.
    YearDay()       This method returns the day of the year, expressed as an int between 1 and 366 (to accommodate leap years).
    Month()         This method returns the month component, expressed using the Month type.
    Day()           This method returns the day of the month, expressed as an int.
    Weekday()       This method returns the day of the week, expressed as a Weekday.
    Hour()          This method returns the hour of the day, expressed as an int between 0 and 23.
    Minute()        This method returns the number of minutes elapsed into the hour of the day, expressed as an int between 0 and 59.
    Second()        This method returns the number of seconds elapsed into the minute of the hour, expressed as an int between 0 and 59.
    Nanosecond()    This method returns the number of nanoseconds elapsed into the second of the minute,
                    expressed as an int between 0 and 999,999,999.
████████████████████████████████████████████████████████████████████████
246.The Types Used to Describe Time Components
    Name        Description
    -------     ------------------------------------------------------------------------
    Month       This type represents a month, and the time package defines constant values for the English-
                language month names: January, February, etc. The Month type defines a String method that
                uses these names when formatting strings.
    Weekday     This type represents a day of the week, and the time package defines constant values for the
                English-language weekday names: Sunday, Monday, etc. The Weekday type defines a String
                method that uses these names when formatting strings.
████████████████████████████████████████████████████████████████████████
247.Creating Time Values
    example:
        package main
        import "time"
        func PrintTime(label string, t *time.Time) {
            Printfln("%s: Day: %v: Month: %v Year: %v",
                label, t.Day(), t.Month(), t.Year())
        }
        func main() {
            current := time.Now()
            specific := time.Date(1995, time.June, 9, 0, 0, 0, 0, time.Local)
            unix := time.Unix(1433228090, 0)
            PrintTime("Current", &current)
            PrintTime("Specific", &specific)
            PrintTime("UNIX", &unix)
        }
    Output:
        Current: Day: 15: Month: September Year: 2023
        Specific: Day: 9: Month: June Year: 1995
        UNIX: Day: 2: Month: June Year: 2015

    example:
        package main
        import "time"
        func PrintTime(label string, t *time.Time) {
            Printfln("%s: Day: %v: Month: %v Year: %v",
                label, t.Day(), t.Month(), t.Year())
        }
        func main() {
            current := time.Now()
            specific := time.Date(1993, time.June, 0, 0, 0, 0, 0, time.Local)
            unix := time.Unix(0, 0)
            PrintTime("Current", &current)
            PrintTime("Specific", &specific)
            PrintTime("UNIX", &unix)
        }
    Output:
        Current: Day: 15: Month: September Year: 2023
        Specific: Day: 31: Month: May Year: 1993
        UNIX: Day: 31: Month: December Year: 1969

████████████████████████████████████████████████████████████████████████
248.The Time Method for Creating Formatted Strings
    Name                Description
    --------------      --------------------
    Format(layout)      This method returns a formatted string, which is created using the specified layout.

    example:
        package main
        import (
            "fmt"
            "time"
        )
        func PrintTime(label string, t *time.Time) {
            layout := "Day: 02 Month: Jan Year: 2006"
            fmt.Println(label, t.Format(layout))
        }
        func main() {
            current := time.Now()
            specific := time.Date(1993, time.June, 0, 0, 0, 0, 0, time.Local)
            unix := time.Unix(0, 0)
            PrintTime("Current", &current)
            PrintTime("Specific", &specific)
            PrintTime("UNIX", &unix)
        }
    Output:
        Current Day: 16 Month: Sep Year: 2023
        Specific Day: 31 Month: May Year: 1993
        UNIX Day: 31 Month: Dec Year: 1969
████████████████████████████████████████████████████████████████████████
249.The Layout Constants Defined by the time Package
    Name            Reference Date Format
    -----------     ----------------------------------
    ANSIC           Mon Jan _2 15:04:05 2006
    UnixDate        Mon Jan _2 15:04:05 MST 2006
    RubyDate        Mon Jan 02 15:04:05 -0700 2006
    RFC822          02 Jan 06 15:04 MST
    RFC822Z         02 Jan 06 15:04 -0700
    RFC850          Monday, 02-Jan-06 15:04:05 MST
    RFC1123         Mon, 02 Jan 2006 15:04:05 MST
    RFC1123Z        Mon, 02 Jan 2006 15:04:05 -0700
    RFC3339         2006-01-02T15:04:05Z07:00
    RFC3339Nano     2006-01-02T15:04:05.999999999Z07:00
    Kitchen         3:04PM
    Stamp           Jan _2 15:04:05
    StampMilli      Jan _2 15:04:05.000
    StampMicro      Jan _2 15:04:05.000000
    StampNano       Jan _2 15:04:05.000000000
████████████████████████████████████████████████████████████████████████
250.The time Package Functions for Parsing Strings into Time Values
    Name                                        Description
    --------------------------------------      -----------------------------------
    Parse(layout, str)                          This function parses a string using the specified layout to create a Time value.
                                                An error is returned to indicate problems parsing the string.
    ParseInLocation(layout, str, location)      This function parses a string, using the specified layout and using the
                                                Location if no time zone is included in the string. An error is returned to
                                                indicate problems parsing the string.

    example:
        package main
        import (
            "fmt"
            "time"
        )
        func PrintTime(label string, t *time.Time) {
            fmt.Println(label, t.Format(time.RFC822Z))
        }
        func main() {
            dates := []string{
                "09 Jun 95 00:00 GMT",
                "02 Jun 15 00:00 GMT",
            }
            
            for _, d := range dates {
                time, err := time.Parse(time.RFC822, d)
                if err == nil {
                    PrintTime("Parsed", &time)
                } else {
                    Printfln("Error: %s", err.Error())
                }
            }
        }
    Output:
        Parsed 09 Jun 95 00:00 +0000
        Parsed 02 Jun 15 00:00 +0000
████████████████████████████████████████████████████████████████████████
251.time.ParseInLocation()
    example:
        package main
        import (
            "fmt"
            "time"
        )
        func PrintTime(label string, t *time.Time) {
            //layout := "Day: 02 Month: Jan Year: 2006"
            fmt.Println(label, t.Format(time.RFC822Z))
        }
        func main() {
            layout := "02 Jan 06 15:04"
            date := "09 Jun 95 19:30"
            london, lonerr := time.LoadLocation("Europe/London")
            newyork, nycerr := time.LoadLocation("America/New_York")
            Tehran, TehranErr := time.LoadLocation("Asia/Tehran")
        
            if lonerr == nil && nycerr == nil  && TehranErr == nil{
        
                TehranTime, _ := time.ParseInLocation(layout, date, Tehran)
                PrintTime("Tehran:",&TehranTime)
        
                nolocation, _ := time.Parse(layout, date)
                PrintTime("No location:", &nolocation)
        
                londonTime, _ := time.ParseInLocation(layout, date, london)
                PrintTime("London:", &londonTime)
        
                newyorkTime, _ := time.ParseInLocation(layout, date, newyork)
                PrintTime("New York:", &newyorkTime)
        
            } else {
                fmt.Println(lonerr.Error(), nycerr.Error())
            }
        }
    Output:
        Tehran: 09 Jun 95 19:30 +0430
        No location: 09 Jun 95 19:30 +0000
        London: 09 Jun 95 19:30 +0100
        New York: 09 Jun 95 19:30 -0400
████████████████████████████████████████████████████████████████████████
252.The Functions for Creating Locations
    Name                                Description
    -------------------------           -----------------------------------------
    LoadLocation(name)                  This function returns a *Location for the specified name and an
                                        error that indicates any problems.
    LoadLocationFromTZData(name,data)   This function returns a *Location from a byte slice that contains a
                                        formatted time zone database.
    FixedZone(name, offset)             This function returns a *Location that always uses the specified
                                        name and offset from UTC.

    The place names are defined in the IANA time zone database, 
    https://www.iana.org/time-zones , 
    and are listed by 
    https://en.wikipedia.org/wiki/List_of_tz_database_time_zones

████████████████████████████████████████████████████████████████████████
253.Embedding The Time Zone Database
    example:
        package main
        import (
            "fmt"
            "time"
        )
        func PrintTime(label string, t *time.Time) {
            //layout := "Day: 02 Month: Jan Year: 2006"
            fmt.Println(label, t.Format(time.RFC822Z))
        }
        func main() {
            layout := "02 Jan 06 15:04"
            date := "09 Jun 95 19:30"
            Tehran, TehranErr := time.LoadLocation("Asia/Tehran")
            local, _ := time.LoadLocation("Local")
            if TehranErr == nil{
                TehranTime, _ := time.ParseInLocation(layout, date, Tehran)
                PrintTime("Tehran:",&TehranTime)
                localTime, _ := time.ParseInLocation(layout, date, local)
                PrintTime("Local:", &localTime)
                nolocation, _ := time.Parse(layout, date)
                PrintTime("No location:", &nolocation)
            } else {
                fmt.Println(TehranErr.Error())
            }
        }
    Output:
        Tehran: 09 Jun 95 19:30 +0430
        Local: 09 Jun 95 19:30 -0400
        No location: 09 Jun 95 19:30 +0000
████████████████████████████████████████████████████████████████████████
254.Specifying Time Zones
    The arguments to the FixedZone function are a name and the number of seconds offset from UTC. This
    example creates three fixed time zones, one of which is an hour ahead of UTC, one of which is four hours
    behind, and one of which has no offset.
    example:
        package main
        import (
            "fmt"
            "time"
        )
        func PrintTime(label string, t *time.Time) {
            //layout := "Day: 02 Month: Jan Year: 2006"
            fmt.Println(label, t.Format(time.RFC822Z))
        }
        func main() {
            layout := "02 Jan 06 15:04"
            date := "09 Jun 95 19:30"
            london := time.FixedZone("BST", 1*60*60)
            newyork := time.FixedZone("EDT", -4*60*60)
            local := time.FixedZone("Local", 0)
            
            nolocation, _ := time.Parse(layout, date)
            londonTime, _ := time.ParseInLocation(layout, date, london)
            newyorkTime, _ := time.ParseInLocation(layout, date, newyork)
            localTime, _ := time.ParseInLocation(layout, date, local)
            
            PrintTime("No location:", &nolocation)
            PrintTime("London:", &londonTime)
            PrintTime("New York:", &newyorkTime)
            PrintTime("Local:", &localTime)
        }
    Output:
        No location: 09 Jun 95 19:30 +0000
        London: 09 Jun 95 19:30 +0100
        New York: 09 Jun 95 19:30 -0400
        Local: 09 Jun 95 19:30 +0000
████████████████████████████████████████████████████████████████████████
255.The Methods for Working with Time Values
    Name                Description
    ----------------    -------------------------------------------------
    Add(duration)       This method adds the specified Duration to the Time and returns the result.
    Sub(time)           This method returns a Duration that expresses the difference between the Time on
                        which the method has been called and the Time provided as the argument.
    AddDate(y, m, d)    This method adds the specified number of years, months, and days to the Time and
                        returns the result.
    After(time)         This method returns true if the Time on which the method has been called occurs
                        after the Time provided as the argument.
    Before(time)        This method returns true if the Time on which the method has been called occurs
                        before the Time provided as the argument.
    Equal(time)         This method returns true if the Time on which the method has been called is equal
                        to the Time provided as the argument.
    IsZero()            This method returns true if the Time on which the method has been called
                        represents the zero-time instant, which is January 1, year 1, 00:00:00 UTC.
    In(loc)             This method returns the Time value, expressed in the specified Location.
    Location()          This method returns the Location that is associated with the Time, effectively
                        allowing a time to be expressed in a different time zone.
    Round(duration)     This method rounds the Time to the nearest interval represented by a Duration
                        value.
    Truncate(duration)  This method rounds the Time down to the nearest interval represented by a
                        Duration value.
████████████████████████████████████████████████████████████████████████
256.time.Parse()
    example:
        package main
        import (
            "fmt"
            "time"
        )
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
        func main() {
            t, err := time.Parse(time.RFC822, "09 Jun 95 04:59 BST")
            if err == nil {
                Printfln("After: %v", t.After(time.Now()))
                Printfln("Round: %v", t.Round(time.Hour))
                Printfln("Truncate: %v", t.Truncate(time.Hour))
            } else {
                fmt.Println(err.Error())
            }
        }
    Output:
        After: false
        Round: 1995-06-09 05:00:00 +0100 BST
        Truncate: 1995-06-09 04:00:00 +0100 BST
████████████████████████████████████████████████████████████████████████
257.Comparing Time Values
    example:
        package main
        import (
            "fmt"
            "time"
        )
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
        func main() {
            t1, _ := time.Parse(time.RFC822Z, "09 Jun 95 04:59 +0100")
            t2, _ := time.Parse(time.RFC822Z, "08 Jun 95 23:59 -0400")
            
            Printfln("Equal Method: %v", t1.Equal(t2))
            Printfln("Equality Operator: %v", t1 == t2)
        }
    Output:
        Equal Method: true
        Equality Operator: false
████████████████████████████████████████████████████████████████████████
258.The Duration Constants in the time Package
    Name            Description
    ------------    ----------------------------------------
    Hour            This constant represents 1 hour.
    Minute          This constant represents 1 minute.
    Second          This constant represents 1 second.
    Millisecond     This constant represents 1 millisecond.
    Microsecond     This constant represents 1 microsecond.
    Nanosecond      This constant represents 1 nanosecond.
████████████████████████████████████████████████████████████████████████
259.The Duration Methods
    Name                Description
    ----------------    ---------------------------------------------
    Hours()             This method returns a float64 that represents the Duration in hours.
    Minutes()           This method returns a float64 that represents the Duration in minutes.
    Seconds()           This method returns a float64 that represents the Duration in seconds.
    Milliseconds()      This method returns an int64 that represents the Duration in milliseconds.
    Microseconds()      This method returns an int64 that represents the Duration in microseconds.
    Nanoseconds()       This method returns an int64 that represents the Duration in nanoseconds.
    Round(duration)     This method returns a Duration, which is rounded to the nearest multiple of the
                        specified Duration.
    Truncate(duration)  This method returns a Duration, which is rounded down to the nearest multiple of
                        the specified Duration.
████████████████████████████████████████████████████████████████████████
260.Hours() - Minutes() - Seconds() - rounded.Hours() - rounded.Minutes()
    example:
        package main
        import (
            "fmt"
            "time"
        )
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
        func main() {
            var d time.Duration = time.Hour + (30 * time.Minute)
            Printfln("Hours: %v", d.Hours())
            Printfln("Mins: %v", d.Minutes())
            Printfln("Seconds: %v", d.Seconds())
            Printfln("Millseconds: %v", d.Milliseconds())
        
        
            rounded := d.Round(time.Hour)
            Printfln("Rounded Hours: %v", rounded.Hours())
            Printfln("Rounded Mins: %v", rounded.Minutes())
        
            trunc := d.Truncate(time.Hour)
            Printfln("Truncated Hours: %v", trunc.Hours())
            Printfln("Rounded Mins: %v", trunc.Minutes())
        }
    Output:
        Hours: 1.5
        Mins: 90
        Seconds: 5400
        Millseconds: 5400000
        Rounded Hours: 2
        Rounded Mins: 120
        Truncated Hours: 1
        Rounded Mins: 60
████████████████████████████████████████████████████████████████████████
261.The time Functions for Creating Duration Values relative to a Time
    Name            Description
    -----------     ----------------------------------------
    Since(time)     This function returns a Duration expressing the elapsed time since the specified Time value.
    Until(time)     This function returns a Duration expressing the elapsed time until the specified Time value.
████████████████████████████████████████████████████████████████████████
262.time.Until(time) - Since(time)
    example:
        package main
        import (
            "fmt"
            "time"
        )
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
        func main() {
            toYears := func(d time.Duration) int {
                return int(d.Hours() / (24 * 365))
            }
            
            future := time.Date(2051, 0, 0, 0, 0, 0, 0, time.Local)
            past := time.Date(1965, 0, 0, 0, 0, 0, 0, time.Local)
            
            Printfln("this year is %v.",time.Now().Year())
            Printfln("Future: %v is %v.", toYears(time.Until(future)), future.Year())
            Printfln("Past: %v is %v.", toYears(time.Since(past)), past.Year())
        }
    Output:
        this year is 2023.
        Future: 27 is 2050.
        Past: 58 is 1964.
████████████████████████████████████████████████████████████████████████
263.time.ParseDuration function
    This function returns a Duration and an error, indicating if there were problems
    parsing the specified string.
    The format of the strings supported by the ParseDuration function is a sequence of number values
    followed by the unit indicators:
    Unit        Description
    -----       --------------------
    h           This unit denotes hours.
    m           This unit denotes minutes.
    s           This unit denotes seconds.
    ms          This unit denotes milliseconds.
    us or μs    These units denotes microseconds.
    ns          This unit denotes nanoseconds.

    example:
        package main
        import (
            "fmt"
            "time"
        )
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
        func main() {
            d, err := time.ParseDuration("1h30m")
            if err == nil {
                Printfln("Hours: %v", d.Hours())
                Printfln("Mins: %v", d.Minutes())
                Printfln("Seconds: %v", d.Seconds())
                Printfln("Millseconds: %v", d.Milliseconds())
            } else {
                fmt.Println(err.Error())
            }
        }
    Output:
        Hours: 1.5
        Mins: 90
        Seconds: 5400
        Millseconds: 5400000
████████████████████████████████████████████████████████████████████████
264.Using the Time Features for Goroutines and Channels
    
    The time package provides a small set of functions that are useful for working with goroutines and channels.

    The time Package Functions:
    Name                        Description
    ----------------------      ----------------------------------------
    Sleep(duration)             This function pauses the current goroutine for at least the specified duration.
    AfterFunc(duration,func)    This function executes the specified function in its own goroutine after the
                                specified duration. The result is a *Timer, whose Stop method can be used to
                                cancel the execution of the function before the duration elapses.
    After(duration)             This function returns a channel that blocks for the specified duration and then
                                yields a Time value. See the “Receiving Timed Notifications” section for details.
    Tick(duration)              This function returns a channel that periodically sends a Time value, where the
                                period is specified as a duration.
████████████████████████████████████████████████████████████████████████
265.time.Sleep(time.Second * 1)
    Pausing a Goroutine
    The Sleep function pauses execution of the current goroutine for a specified duration

    example:
        package main
        import (
            "fmt"
            "time"
        )
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
        func writeToChannel(channel chan<- string) {
            names := []string{"Alice", "Bob", "Charlie", "Dora"}
            for _, name := range names {
                channel <- name
                time.Sleep(time.Second * 1)
            }
            close(channel)
        }
        func main() {
            nameChannel := make(chan string)
            
            go writeToChannel(nameChannel)
            
            for name := range nameChannel {
                Printfln("Read name: %v", name)
            }
        }
    Output:
    Read name: Alice
                        // 1 second delaying
    Read name: Bob
                        // 1 second delaying
    Read name: Charlie
                        // 1 second delaying
    Read name: Dora
████████████████████████████████████████████████████████████████████████
266.time.AfterFunc() function
    The AfterFunc function is used to defer the execution of a function for a specified period
    Deferring a Function:
    example:
        package main
        import (
            "fmt"
            "time"
        )
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
        func writeToChannel(channel chan<- string) {
            names := []string{"Alice", "Bob", "Charlie", "Dora"}
            for _, name := range names {
                channel <- name
                // time.Sleep(time.Second * 1)
            }
            close(channel)
        }
        func main() {
            nameChannel := make(chan string)
        
            time.AfterFunc(time.Second*5, func() {
                writeToChannel(nameChannel)
            })
        
            for name := range nameChannel {
                Printfln("Read name: %v", name)
            }
        }
    Output:
    // It waits for 5 seconds and then continues the program.
    Read name: Alice
    Read name: Bob
    Read name: Charlie
    Read name: Dora
████████████████████████████████████████████████████████████████████████
267.time.After(time.Second * 2)
    The result from the After function is a channel that carries Time values. The channel blocks for the
    specified duration, when a Time value is sent, indicating the duration has passed. In this example, the
    value sent over the channel acts as a signal and is not used directly, which is why it is assigned to the blank
    identifier, like this:

    example:
        package main
        import (
            "fmt"
            "time"
        )
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
        func writeToChannel(channel chan<- string) {
        
            Printfln("Waiting for initial duration...")
            _ = <-time.After(time.Second * 2)
            Printfln("Initial duration elapsed.")
        
            names := []string{"Alice", "Bob", "Charlie", "Dora"}
            for _, name := range names {
                channel <- name
                time.Sleep(time.Second * 1)
            }
            close(channel)
        }
        func main() {
            nameChannel := make(chan string)
            go writeToChannel(nameChannel)
            for name := range nameChannel {
                Printfln("Read name: %v", name)
            }
        }
    Output:
        Waiting for initial duration... // Wait for 2 seconds
        Initial duration elapsed.
        Read name: Alice    // wait for 1 second
        Read name: Bob      // wait for 1 second
        Read name: Charlie  // wait for 1 second
        Read name: Dora     // wait for 1 second
████████████████████████████████████████████████████████████████████████
268.time.Sleep(time.Second * 3) with select statement
    Using a Timeout in a Select Statement

    example:
        package main
        import (
            "fmt"
            "time"
        )
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
        func writeToChannel(channel chan<- string) {
            Printfln("Waiting for initial duration...")
            _ = <-time.After(time.Second * 2)
            Printfln("Initial duration elapsed.")
            names := []string{"Alice", "Bob", "Charlie", "Dora"}
            for _, name := range names {
                channel <- name
                time.Sleep(time.Second * 3)
            }
            close(channel)
        }
        func main() {
            nameChannel := make(chan string)
            go writeToChannel(nameChannel)
            channelOpen := true
            for channelOpen {
                Printfln("Starting channel read")
                select {
                case name, ok := <-nameChannel:
                    if !ok {
                        channelOpen = false
                    } else {
                        Printfln("Read name: %v", name)
                    }
                case <-time.After(time.Second * 2):
                    Printfln("Timeout")
                }
            }
        }
    Output:
        Starting channel read
        Waiting for initial duration...
        Timeout
        Starting channel read
        Initial duration elapsed.
        Read name: Alice
        Starting channel read
        Timeout
        Starting channel read
        Read name: Bob
        Starting channel read
        Timeout
        Starting channel read
        Read name: Charlie
        Starting channel read
        Timeout
        Starting channel read
        Read name: Dora
        Starting channel read
        Timeout
        Starting channel read
████████████████████████████████████████████████████████████████████████
269.NewTimer(duration)
    This function returns a *Timer with the specified period.
    Caution Be careful when stopping a timer. 
    The timer's channel is not closed, which means that reads from
    the channel will continue to block even after the timer has stopped.

    The Methods Defined by the Timer Struct:
    Name                Description
    ------------        -------------------------------------------
    C                   This field returns the channel over which the Time will send its Time value.
    Stop()              This method stops the timer. The result is a bool that will be true if the timer has been
                        stopped and false if the timer had already sent its message.
    Reset(duration)     This method stops a timer and resets it so that its interval is the specified Duration.
    
    example:
        package main
        import (
            "fmt"
            "time"
        )
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
        func writeToChannel(channel chan<- string) {
            timer := time.NewTimer(time.Minute * 10)
            go func() {
                time.Sleep(time.Second * 2)
                Printfln("Resetting timer")
                timer.Reset(time.Second)
            }()
            Printfln("Waiting for initial duration...")
            <-timer.C
            Printfln("Initial duration elapsed.")
            names := []string{"Alice", "Bob", "Charlie", "Dora"}
            for _, name := range names {
                channel <- name
            }
            close(channel)
        }
        func main() {
            nameChannel := make(chan string)
            go writeToChannel(nameChannel)
            for name := range nameChannel {
                Printfln("Read name: %v", name)
            }
        }
    Output:
        Waiting for initial duration...
        Resetting timer
        Initial duration elapsed.
        Read name: Alice
        Read name: Bob
        Read name: Charlie
        Read name: Dora
████████████████████████████████████████████████████████████████████████
270.time.Tick(time.Second)
    Receiving Recurring Notifications دریافت اعلان های مکرر
    The Tick function returns a channel over which Time values are sent at a specified interval
    
    Tick is a convenience wrapper for NewTicker providing access to the ticking channel only. 
    While Tick is useful for clients that have no need to shut down the Ticker, 
    be aware that without a way to shut it down the underlying
    Ticker cannot be recovered by the garbage collector; it "leaks".
    Unlike NewTicker, Tick will return nil if d <= 0.

    example:
        package main
        import (
            "fmt"
            "time"
        )
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
        func writeToChannel(nameChannel chan<- string) {
            names := []string{"Alice", "Bob", "Charlie", "Dora"}
            tickChannel := time.Tick(time.Second)
            index := 0
            for {
                <-tickChannel
                nameChannel <- names[index]
                index++
                if index == len(names) {
                    index = 0
                }
            }
        }
        func main() {
            nameChannel := make(chan string)

            go writeToChannel(nameChannel)

            for name := range nameChannel {
                Printfln("Read name: %v", name)
            }
            
        }
    Output:
        Read name: Alice
        Read name: Bob
        Read name: Charlie
        Read name: Dora
        Read name: Alice
        Read name: Bob
        Read name: Charlie
        Read name: Dora
        Read name: Alice
        Read name: Bob
        Read name: Charlie
        Read name: Dora
        ...
            
████████████████████████████████████████████████████████████████████████
271.NewTicker(duration)
    The result of the NewTicker function is a pointer to a Ticker struct, which defines the field and methods

    The time Function for Creating a Ticker:
    Name                    Description
    ---------------------   ---------------------------------
    NewTicker(duration)     This function returns a *Ticker with the specified period.

    The Field and Methods Defined by the Ticker Struct:
    Name                Description
    ----------------    --------------------------------
    C                   This field returns the channel over which the Ticker will send its Time values.
    Stop()              This method stops the ticker (but does not close the channel returned by the C field).
    Reset(duration)     This method stops a ticker and resets it so that its interval is the specified Duration.
████████████████████████████████████████████████████████████████████████
272.Creating a Ticker in the main.go
    example:
        package main
        import (
            "fmt"
            "time"
        )
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
        func writeToChannel(nameChannel chan<- string) {
            names := []string{"Alice", "Bob", "Charlie", "Dora"}
            ticker := time.NewTicker(time.Second / 10)
            index := 0
            for {
                <-ticker.C
                nameChannel <- names[index]
                index++
                if index == len(names) {
                    ticker.Stop()
                    close(nameChannel)
                    break
                }
            }
        }
        func main() {
            nameChannel := make(chan string)
        
            go writeToChannel(nameChannel)
        
            for name := range nameChannel {
                Printfln("Read name: %v", name)
            }
        }
    Output: 
        // This is printed after milliseconds
        Read name: Alice
        Read name: Bob
        Read name: Charlie
        Read name: Dora
████████████████████████████████████████████████████████████████████████
273.Reading and Writing Data
    These interfaces are used wherever data is read or written, which means that any
    source or destination for data can be treated in much the same way so that writing data to a file, for example,
    is just the same as writing data to a network connection.

    What are they?
    These interfaces define the basic methods required to read and write data.

    Why are they useful?
    This approach means that just about any data source can be used
    in the same way, while still allowing specialized features to be
    defined using the composition features.

    How is it used?
    The io package defines these interfaces, but the implementations
    are available from a range of other packages

    Are there any pitfalls or limitations?
    These interfaces don't entirely hide the detail of sources or
    destinations for data and additional methods are often required,
    provided by interfaces that build on Reader and Writer.

    Are there any alternatives?
    The use of these interfaces is optional, but they are hard to avoid
    because they are used throughout the standard library.

    The Reader and Writer interfaces are defined by the io package and 
    provide abstract ways to read and write data, 
    without being tied to where the data is coming from or going to.

    Preparing for This Chapter:
    product.go:
        package main
        type Product struct {
            Name, Category string
            Price          float64
        }
        var Kayak = Product{
            Name:     "Kayak",
            Category: "Watersports",
            Price:    279,
        }
        var Products = []Product{
            {"Kayak", "Watersports", 279},
            {"Lifejacket", "Watersports", 49.95},
            {"Soccer Ball", "Soccer", 19.50},
            {"Corner Flags", "Soccer", 34.95},
            {"Stadium", "Soccer", 79500},
            {"Thinking Cap", "Chess", 16},
            {"Unsteady Chair", "Chess", 75},
            {"Bling-Bling King", "Chess", 1200},
        }
    printer.go:
        package main
        import (
            "fmt"
        )
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
    main.go:
        package main
        func main() {
            Printfln("Product: %v, Price : %v", Kayak.Name, Kayak.Price)
        }
████████████████████████████████████████████████████████████████████████
274.The Reader interface
    The Reader interface doesn't include any detail about where data comes from or how it is obtained—it
    just defines the Read method. The details are left to the types that implement the interface, and there are
    reader implementations in the standard library for different data sources.

    defines a single method:
    Name                Description
    --------------      -------------------------
    Read(byteSlice)     This method reads data into the specified []byte. The method returns the number of
                        bytes that were read, expressed as an int, and an error.
████████████████████████████████████████████████████████████████████████
275.io.Reader package
    example:
        package main
        import (
            "io"
            "strings"
        )
        func processData(reader io.Reader) {
            b := make([]byte, 2)
            for {
                count, err := reader.Read(b)
                if count > 0 {
                    Printfln("Read %v bytes: %v", count, string(b[0:count]))
                }
                if err == io.EOF {
                    break
                }
            }
        }
        func main() {
            r := strings.NewReader("Kayak")
            processData(r)
        }
    Output:
        Read 2 bytes: Ka
        Read 2 bytes: ya
        Read 1 bytes: k
████████████████████████████████████████████████████████████████████████
276.Writer interface
    Write(byteSlice)
        This method writes the data from the specified byte slice. The method returns the
        number of bytes that were written and an error. The error will be non-nil if the
        number of bytes written is less than the length of the slice.
        The Writer interface doesn't include any details of how the written data is stored, transmitted, or
        processed, all of which is left to the types that implement the interface.
████████████████████████████████████████████████████████████████████████
277.io.Writer
    example:
        package main
        import(
            "io"
            "strings"
            "asd/asd"
        )
        func processData(reader io.Reader, writer io.Writer) {
            b := make([]byte, 2)
            for {
                count, err := reader.Read(b)
                if count > 0 {
                    writer.Write(b[0:count])
                    asd.Printfln("Read %v bytes: %v", count, string(b[0:count]))
                }
                if err == io.EOF {
                    break
                }
            }
        }
        func main() {
            r := strings.NewReader("Kayak")
            var builder strings.Builder
            processData(r, &builder)
            asd.Printfln("String builder contents: %s", builder.String())
        }
    Output:
        Read 2 bytes: Ka
        Read 2 bytes: ya
        Read 1 bytes: k
        String builder contents: Kayak
████████████████████████████████████████████████████████████████████████
278.io.EOF
    The io package defines a special error named EOF, 
    which is used to signal when the Reader reaches the end of the data. 
    If the error result from the Read function is equal to the EOF error, 
    then I break out of the for
    loop that has been reading data from the Reader:
    ...
    if err == io.EOF {
        break
    }
    ...
████████████████████████████████████████████████████████████████████████
279.the Utility Functions for Readers and Writers
    توابع مفید برای خوانندگان و نویسندگان
    
    Functions in the io Package for Readng and Writing Data
    Name                            Description
    ------------------------        -------------------------------------------------------
    Copy(w, r)                      This function copies data from a Reader to a Writer until EOF is returned or
                                    another error is encountered. The results are the number of bytes copies and an
                                    error used to describe any problems.
    CopyBuffer(w, r, buffer)        This function performs the same task as Copy but reads the data into the
                                    specified buffer before it is passed to the Writer.
    CopyN(w, r, count)              This function copies count bytes from the Reader to the Writer. The results are
                                    the number of bytes copies and an error used to describe any problems.
    ReadAll(r)                      This function reads data from the specified Reader until EOF is reached. The
                                    results are a byte slice containing the read data and an error, which is used to
                                    describe any problems.
    ReadAtLeast(r, byteSlice, min)  This function reads at least the specified number of bytes from the reader,
                                    placing them into the byte slice. An error is reported if fewer bytes than specified
                                    are read.
    ReadFull(r, byteSlice)          This function fills the specified byte slice with data. The result is the number
                                    of bytes read and an error. An error will be reported if EOF was encountered
                                    before enough bytes to fill the slice were read.
    WriteString(w, str)             This function writes the specified string to a writer.
████████████████████████████████████████████████████████████████████████
280.io.Copy(writer, reader)
    example:
        package main
        import(
            "io"
            "asd/asd"
            "strings"
        )
        func processData(reader io.Reader, writer io.Writer) {
            count, err := io.Copy(writer, reader)
                if (err == nil) {
                    asd.Printfln("Read %v bytes", count)
                } else {
                    asd.Printfln("Error: %v", err.Error())
                }
        }
        func main() {

            r := strings.NewReader("Kayak .")
            var builder strings.Builder
            processData(r, &builder)
            asd.Printfln("String builder contents: %s", builder.String())
        }
    Output:
        Read 7 bytes
        String builder contents: Kayak .
████████████████████████████████████████████████████████████████████████
281.The io Package Functions for Specialized Readers and Writers
    Name                        Description
    -----------                 ----------------------------------------
    Pipe()                      This function returns a PipeReader and a PipeWriter, which can be used to connect
                                functions that require a Reader and a Writer, as described in the “Using Pipes” section.
    MultiReader(...readers)     This function defines a variadic parameter that allows an arbitrary number of Reader
                                values to be specified. The result is a Reader that passes on the content from each of
                                its parameters in the sequence they are defined, as described in the “Concatenating
                                Multiple Readers” section.
    MultiWriter(...writers)     This function defines a variadic parameter that allows an arbitrary number of Writer
                                values to be specified. The result is a Writer that sends the same data to all the
                                specified writers, as described in the “Combining Multiple Writers” section.
    LimitReader(r, limit)       This function creates a Reader that will EOF after the specified number of bytes, as
                                described in the “Limiting Read Data” section.
████████████████████████████████████████████████████████████████████████
282.Pipe
    Pipes are used to connect code that consumes data through a Reader 
    and code that produces code through a Writer.

    The GenerateData function defines a Writer parameter, which it uses to write bytes from a string.
    example:
    data.go:
        package main
        import (
            "io"
            "asd/asd"
        )
        func GenerateData(writer io.Writer) {
            data := []byte("Kayak, Lifejacket")
            writeSize := 4
            for i := 0; i < len(data); i += writeSize {
                    end := i + writeSize;
                    if (end > len(data)) {
                        end = len(data)
                    }
                    count, err := writer.Write(data[i: end])
                    asd.Printfln("Wrote %v byte(s): %v", count, string(data[i: end]))
                    if (err != nil)  {
                        asd.Printfln("Error: %v", err.Error())
                    }
                }
            }
        func ConsumeData(reader io.Reader) {
            data := make([]byte, 0, 10)
            slice := make([]byte, 2)
            for {
                count, err := reader.Read(slice)
                if (count > 0) {
                    asd.Printfln("Read data: %v", string(slice[0:count]))
                    data = append(data, slice[0:count]...)
                }
                if (err == io.EOF) {
                    break
                }
            }
            asd.Printfln("Read data: %v", string(data))
        }

    Notice the parentheses at the end of this statement. These are required when creating a goroutine for an
    anonymous function, but it is easy to forget them.
    main.go:
        package main
        import (
            "io"
        )
        func main() {
        
            pipeReader, pipeWriter := io.Pipe()
            go func() {
                GenerateData(pipeWriter)
                pipeWriter.Close()
            }()
            ConsumeData(pipeReader)
        }



    The output highlights the fact that pipes are synchronous. The GenerateData function calls the writer's
    Write method and then blocks until the data is read. This is why the first message in the output is from the
    reader: the reader is consuming the data two bytes at a time, which means that two read operations are
    required before the initial call to the Write method, which is used to send four bytes, completes, and the
    message from the GenerateData function is displayed.
    Output:
        Read data: Ka
        Read data: ya
        Wrote 4 byte(s): Kaya
        Read data: k,
        Read data:  L
        Wrote 4 byte(s): k, L
        Read data: if
        Read data: ej
        Wrote 4 byte(s): ifej
        Read data: ac
        Read data: ke
        Wrote 4 byte(s): acke
        Read data: t
        Wrote 1 byte(s): t
        Read data: Kayak, Lifejacket
████████████████████████████████████████████████████████████████████████
283.PipeReader and a PipeWriter
    The io.Pipe function returns a PipeReader and a PipeWriter. The PipeReader and PipeWriter structs
    implement the Closer interface
    Name        Description
    -------     -----------------------
    Close()     This method closes the reader or writer. The details are implementation specific, but, in
                general, any subsequent reads from a closed Reader will return zero bytes and the EOF error,
                while any subsequent writes to a closed Writer will return an error.

    The PipeReader struct implements the Reader interface, which means I can use it as the argument to
    the ConsumeData function. The ConsumeData function is executed in the main goroutine, which means that
    the application won't exit until the function completes.
    The effect is that data is written into the pipe using the PipeWriter and read from the pipe using the
    PipeReader. When the GenerateData function is complete, the Close method is called on the PipeWriter,
    which causes the next read by the PipeReader to produce EOF.
████████████████████████████████████████████████████████████████████████
284.io.MultiReader()
    example:
    main.go:
        package main
        import (
            "io"
            "strings"
        )
        func main() {
            r1 := strings.NewReader("Kayak")
            r2 := strings.NewReader("Lifejacket")
            r3 := strings.NewReader("Canoe")
            concatReader := io.MultiReader(r1, r2, r3)
            ConsumeData(concatReader)
        }
    Output:
        Read data: Ka
        Read data: ya
        Read data: k
        Read data: Li
        Read data: fe
        Read data: ja
        Read data: ck
        Read data: et
        Read data: Ca
        Read data: no
        Read data: e
        Read data: KayakLifejacketCanoe
████████████████████████████████████████████████████████████████████████
285.io.MultiWriter()
example:
    package main
    import (
        "io"
        "strings"
        "asd/asd"
    )
    func main() {
        var w1 strings.Builder
        var w2 strings.Builder
        var w3 strings.Builder
        combinedWriter := io.MultiWriter(&w1, &w2, &w3)
        GenerateData(combinedWriter)
        asd.Printfln("Writer #1: %v", w1.String())
        asd.Printfln("Writer #2: %v", w2.String())
        asd.Printfln("Writer #3: %v", w3.String())
    }
Output:
    Wrote 4 byte(s): Kaya
    Wrote 4 byte(s): k, L
    Wrote 4 byte(s): ifej
    Wrote 4 byte(s): acke
    Wrote 1 byte(s): t
    Writer #1: Kayak, Lifejacket
    Writer #2: Kayak, Lifejacket
    Writer #3: Kayak, Lifejacket
████████████████████████████████████████████████████████████████████████
286.io.TeeReader(concatReader, &writer)
    Echoing Reads to a Writer

    The TeeReader function returns a Reader that echoes the data that it receives to a Writer.

    example:
        package main
        import (
            "asd/asd"
            "io"
            "strings"
        )
        func main() {
        
            r1 := strings.NewReader("Kayak")
            r2 := strings.NewReader("Lifejacket")
            r3 := strings.NewReader("Canoe")
            concatReader := io.MultiReader(r1, r2, r3)
            var writer strings.Builder
            teeReader := io.TeeReader(concatReader, &writer)
            ConsumeData(teeReader)
            asd.Printfln("Echo data: %v", writer.String())
        }
    Output:
        Read data: Ka
        Read data: ya
        Read data: k
        Read data: Li
        Read data: fe
        Read data: ja
        Read data: ck
        Read data: et
        Read data: Ca
        Read data: no
        Read data: e
        Read data: KayakLifejacketCanoe
        Echo data: KayakLifejacketCanoe
████████████████████████████████████████████████████████████████████████
287.io.LimitReader(concatReader, 5)
    The LimitReader function is used to restrict the amount of data that can be obtained from a Reader

    example:
        package main
        import (
            "io"
            "strings"
        )
        func main() {
            r1 := strings.NewReader("Kayak")
            r2 := strings.NewReader("Lifejacket")
            r3 := strings.NewReader("Canoe")
            concatReader := io.MultiReader(r1, r2, r3)
            limited := io.LimitReader(concatReader, 5)
            ConsumeData(limited)
        }
    Output:
        Read data: Ka
        Read data: ya
        Read data: k
        Read data: Kayak
████████████████████████████████████████████████████████████████████████
288.bufio
    The bufio package provides support for adding buffers to readers and writers.

    example:
        This code defined a struct type named CustomReader that acts as a wrapper around a Reader. 
        The implementation of the Read method generates output that reports how much data is read and
        how many read operations are performed overall. 
        custom.go:
            package main
            import (
                "io"
                "asd/asd"
            )
            type CustomReader struct {
                reader    io.Reader
                readCount int
            }
            func NewCustomReader(reader io.Reader) *CustomReader {
                return &CustomReader{reader, 0}
            }
            func (cr *CustomReader) Read(slice []byte) (count int, err error) {
                count, err = cr.reader.Read(slice)
                cr.readCount++
                asd.Printfln("Custom Reader: %v bytes", count)
                if err == io.EOF {
                    asd.Printfln("Total Reads: %v", cr.readCount)
                }
                return
            }

        main.go:
            package main
            import (
                "asd/asd"
                "io"
                "strings"
            )
            func main() {
                text := "It was a boat. A small boat."
                var reader io.Reader = NewCustomReader(strings.NewReader(text))
                var writer strings.Builder
                slice := make([]byte, 5)
                for {
                    count, err := reader.Read(slice)
                    if count > 0 {
                        writer.Write(slice[0:count])
                    }
                    if err != nil {
                        break
                    }
                }
                asd.Printfln("Read data: %v", writer.String())
            }
        The NewCustomreader function is used to create a CustomReader that reads from a string and uses a for
        loop to consume the data using a byte slice.
        Output:
            Custom Reader: 5 bytes
            Custom Reader: 5 bytes
            Custom Reader: 5 bytes
            Custom Reader: 5 bytes
            Custom Reader: 5 bytes
            Custom Reader: 3 bytes
            Custom Reader: 0 bytes
            Total Reads: 7
            Read data: It was a boat. A small boat.
████████████████████████████████████████████████████████████████████████
289.bufio Functions for Creating Buffered Readers
    Name                        Description
    ----------------------      ---------------------------------
    NewReader(r)                This function returns a buffered Reader with the default buffer size (which is
                                4,096 bytes at the time of writing).
    NewReaderSize(r, size)      This function returns a buffered Reader with the specified buffer size.
████████████████████████████████████████████████████████████████████████
290.bufio.NewReader(reader)
    The default buffer size is 4,096 bytes, which means that the buffered reader was able to read all the data
    in a single read operation, plus an additional read to produce the EOF result. Introducing the buffer reduces
    the overhead associated with the read operations, albeit at the cost of the memory used to buffer the data.

    example:
        package main
        import (
            "io"
            "strings"
            "bufio"
        )
        func main() {
            text := "It was a boat. A small boat."
            var reader io.Reader = NewCustomReader(strings.NewReader(text))
            var writer strings.Builder
            slice := make([]byte, 5)
            reader = bufio.NewReader(reader)
            for {
                count, err := reader.Read(slice)
                if (count > 0) {
                    writer.Write(slice[0:count])
                }
                if (err != nil) {
                    break
                }
            }
            Printfln("Read data: %v", writer.String())
        }
Output:
    Custom Reader: 28 bytes
    Custom Reader: 0 bytes
    Total Reads: 2
    Read data: It was a boat. A small boat.
████████████████████████████████████████████████████████████████████████
291.Additional Buffered Reader Methods
    The NewReader and NewReaderSize functions return bufio.Reader values, which implement the io.
    Reader interface and which can be used as drop-in wrappers for other types of Reader methods, seamlessly
    introducing a read buffer.


████████████████████████████████████████████████████████████████████████
292.The Methods Defined by the Buffered Reader
    Name                Description
    --------------      ------------------------------------------
    Buffered()          This method returns an int that indicates the number of bytes that can be read from the buffer.
    Discard(count)      This method discards the specified number of bytes.
    Peek(count)         This method returns the specified number of bytes without removing them from the
                        buffer, meaning they will be returned by subsequent calls to the Read method.
    Reset(reader)       This method discards the data in the buffer and performs subsequent reads from the
                        specified Reader.
    Size()              This method returns the size of the buffer, expressed int.
████████████████████████████████████████████████████████████████████████
293.buffered.Read(slice)
    example:
        package main
        import (
            "io"
            "strings"
            "bufio"
        )
        func main() {
            text := "It was a boat. A small boat."
            var reader io.Reader = NewCustomReader(strings.NewReader(text))
            var writer strings.Builder
            slice := make([]byte, 5)
            buffered := bufio.NewReader(reader)
            for {
                count, err := buffered.Read(slice)
                if (count > 0) {
                    Printfln("Buffer size: %v, buffered: %v",
                        buffered.Size(), buffered.Buffered())
                    writer.Write(slice[0:count])
                }
                if (err != nil) {
                        break
                    }
            }
            Printfln("Read data: %v", writer.String())
        }
    Output:
        Custom Reader: 28 bytes
        Buffer size: 4096, buffered: 23
        Buffer size: 4096, buffered: 18
        Buffer size: 4096, buffered: 13
        Buffer size: 4096, buffered: 8
        Buffer size: 4096, buffered: 3
        Buffer size: 4096, buffered: 0
        Custom Reader: 0 bytes
        Total Reads: 2
        Read data: It was a boat. A small boat.
████████████████████████████████████████████████████████████████████████
294.bufio Functions for Creating Buffered Writers
    Name                        Description
    -----------------------     --------------------------------------
    NewWriter(w)                This function returns a buffered Writer with the default buffer size (which is
                                4,096 bytes at the time of writing).
    NewWriterSize(w, size)      This function returns a buffered Writer with the specified buffer size.
████████████████████████████████████████████████████████████████████████
295.Methods Defined by the bufio.Writer Struct
    Name                Description
    --------------      -----------------------------------
    Available()         This method returns the number of available bytes in the buffer.
    Buffered()          This method returns the number of bytes that have been written to the buffer.
    Flush()             This method writes the contents of the buffer to the underlying Writer.
    Reset(writer)       This method discards the data in the buffer and performs subsequent writes to the
                        specified Writer.
    Size()              This method returns the capacity of the buffer in bytes.
████████████████████████████████████████████████████████████████████████
296.Defining a Custom Writer example
    The NewCustomWriter constructor wraps a Writer with a CustomWriter struct, which reports on its
    write operations.
    custom.go:
        package main
        import (
            "asd/asd"
            "io"
        )
        type CustomReader struct {
            reader    io.Reader
            readCount int
        }
        func NewCustomReader(reader io.Reader) *CustomReader {
            return &CustomReader{reader, 0}
        }
        func (cr *CustomReader) Read(slice []byte) (count int, err error) {
            count, err = cr.reader.Read(slice)
            cr.readCount++
            asd.Printfln("Custom Reader: %v bytes", count)
            if err == io.EOF {
                asd.Printfln("Total Reads: %v", cr.readCount)
            }
            return
        }
        type CustomWriter struct {
            writer     io.Writer
            writeCount int
        }
        func NewCustomWriter(writer io.Writer) *CustomWriter {
            return &CustomWriter{writer, 0}
        }
        func (cw *CustomWriter) Write(slice []byte) (count int, err error) {
            count, err = cw.writer.Write(slice)
            cw.writeCount++
            asd.Printfln("Custom Writer: %v bytes", count)
            return
        }
        func (cw *CustomWriter) Close() (err error) {
            if closer, ok := cw.writer.(io.Closer); ok {
                closer.Close()
            }
            asd.Printfln("Total Writes: %v", cw.writeCount)
            return
        }    
    main.go:
        package main
        import (
            "strings"
            "asd/asd"
        )
        func main() {
            text := "It was a boat. A small boat."
            var builder strings.Builder
            var writer = NewCustomWriter(&builder)
            for i := 0; true; {
                end := i + 5
                if end >= len(text) {
                    writer.Write([]byte(text[i:]))
                    break
                }
                writer.Write([]byte(text[i:end]))
                i = end
            }
            asd.Printfln("Written data: %v", builder.String())
        }
    Output:
        Custom Writer: 5 bytes
        Custom Writer: 5 bytes
        Custom Writer: 5 bytes
        Custom Writer: 5 bytes
        Custom Writer: 5 bytes
        Custom Writer: 3 bytes
        Written data: It was a boat. A small boat.
████████████████████████████████████████████████████████████████████████
297.Using a Buffered Writer in the main.go example
    example:
    main.go:
        package main
        import (
            "strings"
            "asd/asd"
            "bufio"
        )
        func main() {
            text := "It was a boat. A small boat."
            var builder strings.Builder
            var writer = bufio.NewWriterSize(NewCustomWriter(&builder), 20)
            for i := 0; true; {
                end := i + 5
                if end >= len(text) {
                    writer.Write([]byte(text[i:]))
                    writer.Flush()
                    break
                }
                writer.Write([]byte(text[i:end]))
                i = end
            }
            asd.Printfln("Written data: %v", builder.String())
        }
    Output:
        Custom Writer: 20 bytes
        Custom Writer: 8 bytes
        Written data: It was a boat. A small boat.
████████████████████████████████████████████████████████████████████████
298.Scanning from a Reader
    example:
    main.go
        package main
        import (
            "io"
            "strings"
            "asd/asd"
            "fmt"
        )
        func scanFromReader(reader io.Reader, template string,
            vals ...interface{}) (int, error) {
            return fmt.Fscanf(reader, template, vals...)
        }
        func main() {
            reader := strings.NewReader("Kayak Watersports $279.00")
            var name, category string
            var price float64
            scanTemplate := "%s %s $%f"
            _, err := scanFromReader(reader, scanTemplate, &name, &category, &price)
            if err != nil {
                asd.Printfln("Error: %v", err.Error())
            } else {
                asd.Printfln("Name: %v", name)
                asd.Printfln("Category: %v", category)
                asd.Printfln("Price: %.2f", price)
            }
        }
    Output:
        Name: Kayak
        Category: Watersports
        Price: 279.00
████████████████████████████████████████████████████████████████████████
299.Scanning Gradually اسکن به تدریج
    example:
    main.go
        package main
        import (
            "io"
            "strings"
            "asd/asd"
            "fmt"
        )
        func scanSingle(reader io.Reader, val interface{}) (int, error) {
            return fmt.Fscan(reader, val)
        }
        func main() {
            reader := strings.NewReader("Kayak Watersports $279.00")
            for {
                var str string
                _, err := scanSingle(reader, &str)
                if err != nil {
                    if err != io.EOF {
                        asd.Printfln("Error: %v", err.Error())
                    }
                    break
                }
                asd.Printfln("Value: %v", str)
            }
        }
    Output:
        Value: Kayak
        Value: Watersports
        Value: $279.00
████████████████████████████████████████████████████████████████████████
300.Writing Formatted Strings to a Writer
    The fmt package also provides functions for writing formatted strings to a Writer
    The writeFormatted function uses the fmt.Fprintf function to write a string formatted with a template
    to a Writer.

    example:
    main.go:
        package main
        import (
            "io"
            "strings"
            "fmt"
        )
        func writeFormatted(writer io.Writer, template string, vals ...interface{}) {
            fmt.Fprintf(writer, template, vals...)
        }
        func main() {
            var writer strings.Builder
            template := "Name: %s, Category: %s, Price: $%.2f"
            writeFormatted(&writer, template, "Kayak", "Watersports", float64(279))
            fmt.Println(writer.String())
        }
    Output:
        Name: Kayak, Category: Watersports, Price: $279.00
████████████████████████████████████████████████████████████████████████
301.strings.Replacer struct
    The strings.Replacer struct can be used to perform replacements on a string and output the modified
    result to a Writer.

    example:
        package main
        import (
            "fmt"
            "io"
            "strings"
        )
        func writeReplaced(writer io.Writer, str string, subs ...string) {
            replacer := strings.NewReplacer(subs...)
            replacer.WriteString(writer, str)
        }
        func main() {
            text := "It was a boat. A small boat."
            subs := []string{"boat", "kayak", "small", "huge"}
            var writer strings.Builder
            writeReplaced(&writer, text, subs...)
            fmt.Println(writer.String())
        }    
    Output:
        It was a kayak. A huge kayak.
████████████████████████████████████████████████████████████████████████
302.Working with JSON Data
    Putting Working with JSON Data in Context

    What is it?
    JSON data is the de facto standard for exchanging data, especially in HTTP applications.

    Why is it useful?
    JSON is simple enough to be supported by any language but can represent
    relatively complex data.

    How is it used?
    The encoding/json package provides support for encoding and decoding JSON data.

    Are there any pitfalls or limitations?
    Not all Go data types can be represented in JSON, which requires the developer
    to be mindful of how Go data types will be expressed.

    Are there any alternatives?
    There are many other data encodings available, some of which are supported by
    the Go standard library.

    The safest approach is to define a map with string keys and empty interface values, which ensures that
    all the key-value pairs in the JSON data can be decoded into the map


    Problem                     Solution
    -----------                 ------------------------
    Encode JSON                 dataCreate an Encoder with a Writer and invoke the Encode method
    Control struct encoding     Use JSON struct tags or implement the Mashaler interface
    Decode JSON data            Create a Decoder with a Reader and invoke the Decode method
    Control struct decoding     Use JSON struct tags or implement the Unmarshaler interface
████████████████████████████████████████████████████████████████████████
303.The encoding/json Constructor Functions for JSON Data
    Name                    Description
    -------------------     --------------------------------------------
    NewEncoder(writer)      This function returns an Encoder, which can be used to encode JSON data and
                            write it to the specified Writer.
    NewDecoder(reader)      This function returns a Decoder, which can be used to read JSON data from the
                            specified Reader and decode it.
████████████████████████████████████████████████████████████████████████
304.The Functions for Creating and Parsing JSON Data
    Name                            Description
    ------------------------        -----------------------------------------------
    Marshal(value)                  This function encodes the specified value as JSON. The results are the
                                    JSON content expressed in a byte slice and an error, which indicates any
                                    encoding problems.
    Unmarshal(byteSlice, val)       This function parses JSON data contained in the specified slice of bytes
                                    and assigns the result to the specified value.
████████████████████████████████████████████████████████████████████████
305.The Encoder Methods
    The NewEncoder constructor function is used to create an Encoder, which can be used to write JSON data to a Writer.

    Name                            Description
    -------------------------       --------------------------------------------------
    Encode(val)                     This method encodes the specified value as JSON and writes it to the Writer.
    SetEscapeHTML(on)               This method accepts a bool argument that, when true, encodes
                                    characters that would be dangerous in HTML to be escaped. The default
                                    behavior is to escape these characters.
    SetIndent(prefix, indent)       This method specifies a prefix and indentation that is applied to the name
                                    of each field in the JSON output.
████████████████████████████████████████████████████████████████████████
306.Expressing the Basic Go Data Types in JSON
    Data                TypeDescription
    ----------------    ------------------------------------
    bool                Go bool values are expressed as JSON true or false.
    string              Go string values are expressed as JSON strings. By default, unsafe HTML
                        characters are escaped.
    float32, float64    Go floating-point values are expressed as JSON numbers.
    int, int<size>      Go integer values are expressed as JSON numbers.
    uint, uint<size>    Go integer values are expressed as JSON numbers.
    byte                Go bytes are expressed as JSON numbers.
    rune                Go runes are expressed as JSON numbers.
    nil                 The Go nil value is expressed as the JSON null value.
    Pointers            The JSON encoder follows pointers and encodes the value at the pointer's location.
████████████████████████████████████████████████████████████████████████
307.encoding/json
    example:
        package main
        import (
            "encoding/json"
            "fmt"
            "strings"
        )
        func main() {
            var b bool = true
            var str string = "Hello"
            var fval float64 = 99.99
            var ival int = 200
            var pointer *int = &ival
            var writer strings.Builder
            encoder := json.NewEncoder(&writer)
            for _, val := range []interface{}{b, str, fval, ival, pointer} {
                encoder.Encode(val)
            }
            fmt.Print(writer.String())
        }
    Output:
        true
        "Hello"
        99.99
        200
        200
████████████████████████████████████████████████████████████████████████
308.Encoding Slices and Arrays
    Example:
        package main
        import (
            "encoding/json"
            "fmt"
            "strings"
        )
        func main() {
            names := []string{"Kayak", "Lifejacket", "Soccer Ball"}
            numbers := [3]int{10, 20, 30}
            var byteArray [5]byte
            copy(byteArray[0:], []byte(names[0]))
            byteSlice := []byte(names[0])
            
            var writer strings.Builder
            encoder := json.NewEncoder(&writer)
            
            encoder.Encode(names)
            encoder.Encode(numbers)
            encoder.Encode(byteArray)
            encoder.Encode(byteSlice)
            fmt.Print(writer.String())
    }
    Output:
        ["Kayak","Lifejacket","Soccer Ball"]
        [10,20,30]
        [75,97,121,97,107]
        "S2F5YWs="
████████████████████████████████████████████████████████████████████████
309.Encoding Maps
    Go maps are encoded as JSON objects, with the map keys used as the object keys. The values contained in
    the map are encoded based on their type.
    Maps can also be useful for creating custom JSON representations of Go data.


    encoder.Encode()
    example:
        package main
        import (
            "encoding/json"
            "fmt"
            "strings"
        )
        func main() {
            m := map[string]float64{
                "Kayak":      279,
                "Lifejacket": 49.95,
            }
            var writer strings.Builder
            encoder := json.NewEncoder(&writer)
            encoder.Encode(m)
            fmt.Print(writer.String())
        }
    Output:
            {"Kayak":279,"Lifejacket":49.95}

████████████████████████████████████████████████████████████████████████
310.Encoding Structs
    The Encoder expresses struct values as JSON objects, using the exported struct field names as the object's
    keys and the field values as the object's values

    example:
        package main
        import (
            "encoding/json"
            "fmt"
            "strings"
        )
        func main() {
            
            var writer strings.Builder
            encoder := json.NewEncoder(&writer)
            encoder.Encode(Kayak)
            fmt.Print(writer.String())
        }
    Output:
        {"Name":"Kayak","Category":"Watersports","Price":279}
████████████████████████████████████████████████████████████████████████
311.Effect of Promotion in JSON in Encoding
    When a struct defines an embedded field that is also a struct, the fields of the embedded struct are promoted
    and encoded as though they are defined by the enclosing type.
    discount.go:
        package main
        type DiscountedProduct struct {
            *Product
            Discount float64
        }
    main.go:
        package main
        import (
            "encoding/json"
            "fmt"
            "strings"
        )
        func main() {
            
            var writer strings.Builder
            encoder := json.NewEncoder(&writer)
            dp := DiscountedProduct {
                Product: &Kayak,
                Discount: 10.50,
            }
            encoder.Encode(&dp)
            fmt.Print(writer.String())
        }
    Output:
        {"Name":"Kayak","Category":"Watersports","Price":279,"Discount":10.5}
████████████████████████████████████████████████████████████████████████
312.Customizing the JSON Encoding of Structs
    How a struct is encoded can be customized using struct tags, which are string literals that follow fields. Struct
    tags are part of the Go support for reflection, 
    that tags follow fields and can be used to alter two aspects of how a field is encoded in JSON.

    example:
    discount.go:
        package main
        type DiscountedProduct struct {
            *Product ^json:"product"^           // ------------------------->   Use the symbol ^ above the Tab button instead

            Discount float64
        }
    
    Output:
        {"product":{"Name":"Kayak","Category":"Watersports","Price":279},"Discount":10.5}

████████████████████████████████████████████████████████████████████████
313.Omitting a Field حذف یک فیلد
    The Encoder skips fields decorated with a tag that specifies a hyphen (the - character) for the name
    The new tag tells the Encoder to skip the Discount field when creating the JSON representation of a
    DIscountedProduct value.

    exampe:
    discount.go:
        package main
        type DiscountedProduct struct {
            *Product ^json:"product"^           // ------------------------->   Use the symbol ^ above the Tab button instead

            Discount float64 ^json:"-"^         // ------------------------->   Use the symbol ^ above the Tab button instead

        }        
    Output:
        {"product":{"Name":"Kayak","Category":"Watersports","Price":279}}
████████████████████████████████████████████████████████████████████████
314.Omitting Unassigned Fields
    By default, the JSON Encoder includes struct fields, even when they have not been assigned a value
    
    example:
    main.go:
        package main
        import (
            "encoding/json"
            "fmt"
            "strings"
        )
        func main() {
            var writer strings.Builder
            encoder := json.NewEncoder(&writer)
            dp := DiscountedProduct{
                Product:  &Kayak,
                Discount: 10.50,
            }
            encoder.Encode(&dp)
            dp2 := DiscountedProduct { Discount: 10.50 }
            encoder.Encode(&dp2)
            fmt.Print(writer.String())
        }
    Output:
        {"product":{"Name":"Kayak","Category":"Watersports","Price":279}}
        {"product":null}

    To omit a nil field, the omitempty keyword is added to the tag for the field
    discount.go:
        package main
        type DiscountedProduct struct {
            // *Product ^json:"product"^                   // ------------------------->   Use the symbol ^ above the Tab button instead

            *Product ^json:"product,omitempty"^            // ------------------------->   Use the symbol ^ above the Tab button instead

            Discount float64 ^json:"-"^
        }
    Output:
        {"product":{"Name":"Kayak","Category":"Watersports","Price":279}}
        {}


    To skip a nil field without changing the name or field promotion, specify the omitempty keyword without a name
    discount.go:
        package main
        type DiscountedProduct struct {
            // *Product ^json:"product"^
            // *Product ^json:"product,omitempty"^
            *Product ^json:",omitempty"^                    // ------------------------->   Use the symbol ^ above the Tab button instead

            Discount float64 ^json:"-"^                     // ------------------------->   Use the symbol ^ above the Tab button instead

        }
    Output:
        {"Name":"Kayak","Category":"Watersports","Price":279}
        {}
████████████████████████████████████████████████████████████████████████
315.Forcing Fields to be Encoded as Strings
    Struct tags can be used to force a field value to be encoded as a string, overriding the normal encoding for
    the field type
    
    example:
    discount.go:
     package main
        type DiscountedProduct struct {
            *Product ^json:",omitempty"^
            // Discount float64 ^json:"-"^              // ------------------------->   Use the symbol ^ above the Tab button                

            Discount float64 ^json:",string"^           // ------------------------->   Use the symbol ^ above the Tab button    

        }
    Output:
        {"Name":"Kayak","Category":"Watersports","Price":279,"Discount":"10.5"}
        {"Discount":"10.5"}
████████████████████████████████████████████████████████████████████████
316.Encoding Interfaces
    The JSON encoder can be used on values assigned to interface variables, but it is the dynamic type that
    is encoded.

    No aspect(جنبه) of the interface is used to adapt the JSON, and all the exported fields of each value in the slice
    are included in the JSON. This can be a useful feature, but care must be taken when decoding this kind of
    JSON, because each value can have a different set of fields

    example:
    interface.go:
        package main
        type Named interface{ GetName() string }
        type Person struct{ PersonName string }
        func (p *Person) GetName() string { return p.PersonName }
        func (p *DiscountedProduct) GetName() string { return p.Name }

    main.go:
        package main
        import (
            "encoding/json"
            "fmt"
            "strings"
        )
        func main() {
            var writer strings.Builder
            encoder := json.NewEncoder(&writer)
            dp := DiscountedProduct{
                Product:  &Kayak,
                Discount: 10.50,
            }
            namedItems := []Named { &dp, &Person{ PersonName: "Alice"}}
            encoder.Encode(namedItems)
            fmt.Print(writer.String())
        }
    Output:
        [{"Name":"Kayak","Category":"Watersports","Price":279,"Discount":"10.5"},{"PersonName":"Alice"}]

████████████████████████████████████████████████████████████████████████
317.The Marshaler Method
    Creating Completely Custom JSON Encodings
    The Encoder checks to see whether a struct implements the Marshaler interface, which denotes a type that
    has a custom encoding and which defines the method.

    Name                Description
    --------------      ------------------------------------------
    MarshalJSON()       This method is invoked to create a JSON representation of a value and returns a byte
                        slice containing the JSON and an error indicating encoding problems.
████████████████████████████████████████████████████████████████████████
318.json.Marshal()
    The MarshalJSON method can generate JSON in any way that suits the project.

    I define a map with string keys and use the
    empty interface for the values. This allows me to build the JSON by adding key-value pairs to the map
    and then pass the map to the Marshal function, which uses the built-in support
    to encode each of the values contained in the map.

    example:
    discount.go:
        package main
        import "encoding/json"
        type DiscountedProduct struct {
            *Product ^json:",omitempty"^
            Discount float64 ^json:",string"^
        }
        func (dp *DiscountedProduct) MarshalJSON() (jsn []byte, err error) {
            if dp.Product != nil {
                m := map[string]interface{}{
                    "product": dp.Name,
                    "cost":    dp.Price - dp.Discount,
                }
                jsn, err = json.Marshal(m)
            }
            return
        }

    main.go:
        package main
        import (
            "encoding/json"
            "fmt"
            "strings"
        )
        func main() {
            var writer strings.Builder
            encoder := json.NewEncoder(&writer)
            dp := DiscountedProduct{
                Product:  &Kayak,
                Discount: 10.50,
            }
            namedItems := []Named { &dp, &Person{ PersonName: "Alice"}}
            encoder.Encode(namedItems)
            fmt.Print(writer.String())
        }
        
    Output:
        [{"cost":268.5,"product":"Kayak"},{"PersonName":"Alice"}]
    
████████████████████████████████████████████████████████████████████████
319.Decoding JSON Data
    The NewDecoder constructor function creates a Decoder, which can be used to decode JSON data obtained
    from a Reader.

    The Decoder Methods:

    Name                        Description
    ---------------------       --------------------------------------------
    Decode(value)               This method reads and decodes data, which is used to create the specified
                                value. The method returns an error that indicates problems decoding the
                                data to the required type or EOF.
    DisallowUnknownFields()     By default, when decoding a struct type, the Decoder ignores any key in
                                the JSON data for which there is no corresponding struct field. Calling this
                                method causes the Decode to return an error, rather than ignoring the key.
    UseNumber()                 By default, JSON number values are decoded into float64 values. Calling
                                this method uses the Number type instead, as described in the “Decoding
                                Number Values” section.
████████████████████████████████████████████████████████████████████████
320.Decoding Basic Data Types
    example:
    main.go:
        package main
        import (
            "encoding/json"
            "io"
            "strings"
            "asd/asd"
        )
        func main() {
            reader := strings.NewReader(^true "Hello" 99.99 200^)   // ------------------------->   Use the symbol ^ above the Tab button    

            vals := []interface{}{}
            decoder := json.NewDecoder(reader)
            for {
                var decodedVal interface{}
                err := decoder.Decode(&decodedVal)
                if err != nil {
                    if err != io.EOF {
                        asd.Printfln("Error: %v", err.Error())
                    }
                    break
                }
                vals = append(vals, decodedVal)
            }
            for _, val := range vals {
                asd.Printfln("Decoded (%T): %v", val, val)
            }
        }
    Output:
        Decoded (bool): true
        Decoded (string): Hello
        Decoded (float64): 99.99
        Decoded (float64): 200
    
████████████████████████████████████████████████████████████████████████
321.Decoding Number Values
    The Methods Defined by the Number Type
    Name            Description
    ---------       ------------------------------------
    Int64()         This method returns the decoded value as a int64 and an error that indicates if the value
                    cannot be converted.
    Float64()       This method returns the decoded value as a float64 and an error that indicates if the
                    value cannot be converted.
    String()        This method returns the unconverted string from the JSON data.

    Not all JSON number values can be expressed as
    Go int64 values, so this is the method that is typically called first. 
    If attempting to convert to an integer
    fails, then the Float64 method can be called. 
    If a number cannot be converted to either Go type, then the
    String method can be used to get the unconverted string from the JSON data.
████████████████████████████████████████████████████████████████████████
322.Decoding Numbers
    example:
    main.go:
        package main
        import (
            "encoding/json"
            "io"
            "strings"
        )
        func main() {
            reader := strings.NewReader(^true "Hello" 99.99 200^)   // ------------------------->   Use the symbol ^ above the Tab button 

            vals := []interface{}{}
            decoder := json.NewDecoder(reader)
            for {
                var decodedVal interface{}
                err := decoder.Decode(&decodedVal)
                if err != nil {
                    if err != io.EOF {
                        Printfln("Error: %v", err.Error())
                    }
                    break
                }
                vals = append(vals, decodedVal)
            }
            for _, val := range vals {
                if num, ok := val.(json.Number); ok {
                    if ival, err := num.Int64(); err == nil {
                        Printfln("Decoded Integer: %v", ival)
                    } else if fpval, err := num.Float64(); err == nil {
                        Printfln("Decoded Floating Point: %v", fpval)
                    } else {
                        Printfln("Decoded String: %v", num.String())
                    }
                } else {
                    Printfln("Decoded (%T): %v", val, val)
                }
            }
        }
    Output:
        Decoded (bool): true
        Decoded (string): Hello
        Decoded (float64): 99.99
        Decoded (float64): 200
████████████████████████████████████████████████████████████████████████
323.Specifying Types for Decoding
    example:
    main.go:
        package main
        import (
            "encoding/json"
            "strings"
        )
        func main() {
            reader := strings.NewReader(^true "Hello" 99.99 200^)   // ------------------------->   Use the symbol ^ above the Tab button 

            var bval bool
            var sval string
            var fpval float64
            var ival int
            vals := []interface{}{&bval, &sval, &fpval, &ival}
            decoder := json.NewDecoder(reader)
            for i := 0; i < len(vals); i++ {
                err := decoder.Decode(vals[i])
                if err != nil {
                    Printfln("Error: %v", err.Error())
                    break
                }
            }
            Printfln("Decoded (%T): %v", bval, bval)
            Printfln("Decoded (%T): %v", sval, sval)
            Printfln("Decoded (%T): %v", fpval, fpval)
            Printfln("Decoded (%T): %v", ival, ival)
        }
    Output:
        Decoded (bool): true
        Decoded (string): Hello
        Decoded (float64): 99.99
        Decoded (int): 200
████████████████████████████████████████████████████████████████████████
324.Decoding Arrays
    The Decoder processes arrays automatically, but care must be taken because JSON allows arrays to contain
    values of different types, which conflicts with the strict type rules enforced by Go.

    The source JSON data contains two arrays, one of which contains only numbers and one of which mixes
    numbers and strings. The Decoder doesn't try to figure out if a JSON array can be represented using a single
    Go type and decodes every array into an empty interface slice:
    example:
    main.go:
        package main
        import (
            "encoding/json"
            "io"
            "strings"
        )
        func main() {
            reader := strings.NewReader(^[10,20,30]["Kayak","Lifejacket",279]^)     // ------------------------->   Use the symbol ^ above the Tab button 

            vals := []interface{}{}
            decoder := json.NewDecoder(reader)
            for {
                var decodedVal interface{}
                err := decoder.Decode(&decodedVal)
                if err != nil {
                    if err != io.EOF {
                        Printfln("Error: %v", err.Error())
                    }
                    break
                }
                vals = append(vals, decodedVal)
            }
            for _, val := range vals {
                Printfln("Decoded (%T): %v", val, val)
            }
        }
    Output:
        Decoded ([]interface {}): [10 20 30]
        Decoded ([]interface {}): [Kayak Lifejacket 279]
████████████████████████████████████████████████████████████████████████
325.Specifying the Decoded Array Type
    The second array contains a mix of values, which means that I have to specify
    the empty interface as the target type. The literal slice syntax is awkward when using the empty interface
    because two sets of braces are required:
    ...
    mixed := []interface{} {}


    example:
    main.go:
        package main
        import (
            "encoding/json"
            "strings"
        )
        func main() {
            reader := strings.NewReader(^[10,20,30]["Kayak","Lifejacket",279]^) // ------------------------->   Use the symbol ^ above the Tab button 

            ints := []int {}
            mixed := []interface{} {}
            vals := []interface{} { &ints, &mixed}
            decoder := json.NewDecoder(reader)
            for i := 0; i < len(vals); i++ {
                err := decoder.Decode(vals[i])
                if err != nil {
                    Printfln("Error: %v", err.Error())
                    break
                }
            }
            Printfln("Decoded (%T): %v", ints, ints)
            Printfln("Decoded (%T): %v", mixed, mixed)
        }
    Output:
        Decoded ([]int): [10 20 30]
        Decoded ([]interface {}): [Kayak Lifejacket 279]
████████████████████████████████████████████████████████████████████████
326.Decoding Maps
    The safest approach is to define a map with string keys and empty interface values, which ensures that
    all the key-value pairs in the JSON data can be decoded into the map
    JavaScript objects are expressed as key-value pairs, which makes it easy to decode them into Go maps
    example:
    main.go:
        package main
        import (
            "encoding/json"
            "strings"
        )
        func main() {
            reader := strings.NewReader(^{"Kayak" : 279, "Lifejacket" : 49.95}^)    // ------------------------->   Use the symbol ^ above the Tab button 
 
            m := map[string]interface{}{}
            decoder := json.NewDecoder(reader)
            err := decoder.Decode(&m)
            if err != nil {
                Printfln("Error: %v", err.Error())
            } else {
                Printfln("Map: %T, %v", m, m)
                for k, v := range m {
                    Printfln("Key: %v, Value: %v", k, v)
                }
            }
        }
    Output:
        Map: map[string]interface {}, map[Kayak:279 Lifejacket:49.95]
        Key: Kayak, Value: 279
        Key: Lifejacket, Value: 49.95
████████████████████████████████████████████████████████████████████████
327.a Specific Value Type 
    A single JSON object can be used for multiple data types as values, but if you know in advance that you
    will be decoding a JSON object that has a single value type, then you can be more specific when defining the
    map into which the data will be decoded.

    example:
    main.go:
        package main
        import (
            "encoding/json"
            "strings"
        )
        func main() {
            reader := strings.NewReader(^{"Kayak" : 279, "Lifejacket" : 49.95}^)    // ------------------------->   Use the symbol ^ above the Tab button 
 
            m := map[string]float64 {}
            decoder := json.NewDecoder(reader)
            err := decoder.Decode(&m)
            if err != nil {
                Printfln("Error: %v", err.Error())
            } else {
                Printfln("Map: %T, %v", m, m)
                for k, v := range m {
                    Printfln("Key: %v, Value: %v", k, v)
                }
            }
        }
    Output:
        Map: map[string]float64, map[Kayak:279 Lifejacket:49.95]
        Key: Kayak, Value: 279
        Key: Lifejacket, Value: 49.95
████████████████████████████████████████████████████████████████████████
328.Decoding Structs

    The Decoder decodes the JSON object and uses the keys to set the values of the exported struct fields.
    The capitalization of the fields and JSON keys don't have to match, and the Decoder will ignore any JSON
    key for which there isn't a struct field and ignore any struct field for which there is no JSON key.
    The JSON objectscontain different capitalization and have more or fewer keys than the Product struct
    fields. The Decoder processes the data as best as it can.

    example:
    main.go:
        package main
        import (
            "strings"
            "encoding/json"
            "io"
        )
        func main() {
            reader := strings.NewReader(^    // ------------------------->   Use the symbol ^ above the Tab button 
 
                {"Name":"Kayak","Category":"Watersports","Price":279}
                {"Name":"Lifejacket","Category":"Watersports" }
                {"name":"Canoe","category":"Watersports", "price": 100, "inStock": true }
            ^)                                          // ------------------------->   Use the symbol ^ above the Tab button 
 
            decoder := json.NewDecoder(reader)
            for {
                var val Product
                err := decoder.Decode(&val)
                if err != nil {
                    if err != io.EOF {
                        Printfln("Error: %v", err.Error())
                    }
                    break
                } else {
                    Printfln("Name: %v, Category: %v, Price: %v",
                        val.Name, val.Category, val.Price)
                }
            }
        }
    Output:
        Name: Kayak, Category: Watersports, Price: 279
        Name: Lifejacket, Category: Watersports, Price: 0
        Name: Canoe, Category: Watersports, Price: 100

████████████████████████████████████████████████████████████████████████
329.Decoding to Interface Types
    As I explained earlier in the chapter, the JSON encoder deals with interfaces by encoding the value
    using the exported fields of the dynamic type. This is because JSON deals with key-value pairs and
    has no way to express methods. As a consequence, you cannot decode directly to an interface variable
    from JSON. Instead, you must decode to a struct or map and then assign the value that is created to an
    interface variable.
████████████████████████████████████████████████████████████████████████
330.Disallowing Unused Keys غیر مجاز کردن
    By default, the Decoder will ignore JSON keys for which there is no corresponding struct field. This behavior
    can be changed by calling the DisallowUnknownFields method

    example:
    Disallowing Unused Keys in the main.go
        ...
        decoder := json.NewDecoder(reader)
        decoder.DisallowUnknownFields()
        ...
    output:
        Name: Kayak, Category: Watersports, Price: 279
        Name: Lifejacket, Category: Watersports, Price: 0
        Error: json: unknown field "inStock"
████████████████████████████████████████████████████████████████████████
331.Struct Tags 
    The tag applied to the Discount field tells the Decoder that the value for this field should be obtained
    from the JSON key named offer and that the value will be parsed from a string, instead of the JSON
    number that would usually be expected for a Go float64 value.

    example:
    discount.go:
        package main
        import "encoding/json"
        type DiscountedProduct struct {
            *Product ^json:",omitempty"^                 // ------------------------->   Use the symbol ^ above the Tab button 
 
            Discount float64 ^json:"offer,string"^       // ------------------------->   Use the symbol ^ above the Tab button 
 
        }
        func (dp *DiscountedProduct) MarshalJSON() (jsn []byte, err error) {
            if (dp.Product != nil) {
                m := map[string]interface{} {
                    "product": dp.Name,
                    "cost": dp.Price - dp.Discount,
                }
                jsn, err = json.Marshal(m)
            }
            return
        }    
    main.go:
        package main
        import (
            "strings"
            "encoding/json"
            "io"
        )
        func main() {
            reader := strings.NewReader(^
                {"Name":"Kayak","Category":"Watersports","Price":279, "Offer": "10"}^)   // ------------------------->   Use the symbol ^ above the Tab button 
 
            decoder := json.NewDecoder(reader)
            for {
                var val DiscountedProduct
                err := decoder.Decode(&val)
                if err != nil {
                    if err != io.EOF {
                        Printfln("Error: %v", err.Error())
                    }
                    break
                } else {
                    Printfln("Name: %v, Category: %v, Price: %v, Discount: %v",
                        val.Name, val.Category, val.Price, val.Discount)
                }
            }
        }
    Output:
        Name: Kayak, Category: Watersports, Price: 279, Discount: 10
████████████████████████████████████████████████████████████████████████
332.Creating Completely Custom JSON Decoders
    The Unmarshaler Method
    Name                            Description
    ------------------------        ------------------------------------------
    UnmarshalJSON(byteSlice)        This method is invoked to decode JSON data contained in the specified
                                    byte slice. The result is an error indicating encoding problems.
████████████████████████████████████████████████████████████████████████
333.Defining a Custom Decoder
    This implementation of the UnmarshalJSON method uses the Unmarshal method to decode the JSON
    data into a map and then checks the type of each value required for the DiscountedProduct struct.

    example:
    discount.go:
        package main
        import (
            "encoding/json"
            "strconv"
        )
        type DiscountedProduct struct {
            *Product ^json:",omitempty"^               // ------------------------->   Use the symbol ^ above the Tab button 
    
            Discount float64 ^json:"offer,string"^         // ------------------------->   Use the symbol ^ above the Tab button 
    
        }
        func (dp *DiscountedProduct) MarshalJSON() (jsn []byte, err error) {
            if dp.Product != nil {
                m := map[string]interface{}{
                    "product": dp.Name,
                    "cost":    dp.Price - dp.Discount,
                }
                jsn, err = json.Marshal(m)
            }
            return
        }
        func (dp *DiscountedProduct) UnmarshalJSON(data []byte) (err error) {
            mdata := map[string]interface{}{}
            err = json.Unmarshal(data, &mdata)
            if dp.Product == nil {
                dp.Product = &Product{}
            }
            if err == nil {
                if name, ok := mdata["Name"].(string); ok {
                    dp.Name = name
                }
                if category, ok := mdata["Category"].(string); ok {
                    dp.Category = category
                }
                if price, ok := mdata["Price"].(float64); ok {
                    dp.Price = price
                }
                if discount, ok := mdata["Offer"].(string); ok {
                    fpval, fperr := strconv.ParseFloat(discount, 64)
                    if fperr == nil {
                        dp.Discount = fpval
                    }
                }
            }
            return
        }
    Output:
        Name: Kayak, Category: Watersports, Price: 279, Discount: 10
████████████████████████████████████████████████████████████████████████
334.Working with Files
    Putting Working with Files in Context
        
    Answer                                  Question
    ---------------------------             -------------------------------------------------------
    What are they?                          These features provide access to the file system so that files can be read and written.
    Why are they useful?                    Files are used for everything from logging to configuration files.
    How are they used?                      These features are accessed through the os package, which provides platform-
                                            neutral access to the file system.
    Are there any pitfallsor limitations?   Some consideration of the underlying file system must be made, especially when
                                            dealing with paths.
    Are there any alternatives?             Go supports alternative ways of storing data, such as databases, but there are no
                                            alternative mechanisms for accessing files.

    Preparing
    printer.go:
        package main
        import (
            "fmt"
        )
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
    product.go:
        package main
        type Product struct {
            Name, Category string
            Price          float64
        }
        var Kayak = Product{
            Name:     "Kayak",
            Category: "Watersports",
            Price:    279,
        }
        var Products = []Product{
            {"Kayak", "Watersports", 279},
            {"Lifejacket", "Watersports", 49.95},
            {"Soccer Ball", "Soccer", 19.50},
            {"Corner Flags", "Soccer", 34.95},
            {"Stadium", "Soccer", 79500},
            {"Thinking Cap", "Chess", 16},
            {"Unsteady Chair", "Chess", 75},
            {"Bling-Bling King", "Chess", 1200},
        }    
    main.go:
        package main
        func main() {
            for _, p := range Products {
                Printfln("Product: %v, Category: %v, Price: $%.2f",
                    p.Name, p.Category, p.Price)
            }
        }
    Output:
        Product: Kayak, Category: Watersports, Price: $279.00
        Product: Lifejacket, Category: Watersports, Price: $49.95
        Product: Soccer Ball, Category: Soccer, Price: $19.50
        Product: Corner Flags, Category: Soccer, Price: $34.95
        Product: Stadium, Category: Soccer, Price: $79500.00
        Product: Thinking Cap, Category: Chess, Price: $16.00
        Product: Unsteady Chair, Category: Chess, Price: $75.00
        Product: Bling-Bling King, Category: Chess, Price: $1200.00
████████████████████████████████████████████████████████████████████████
335.The os Package Functions for Reading Files
    Name                Description
    -------------       -------------------------------------
    ReadFile(name)      This function opens the specified file and reads its contents. The results are a byte
                        slice containing the file content and an error indicating problems opening or reading
                        the file.
    Open(name)          This function opens the specified file for reading. The result is a File struct and an
                        error that indicates problems opening the file.

    One of the most common reasons to read a file is to load configuration data. The JSON format is well-
    suited for configuration files because it is simple to process, has good support in the Go standard library
    The Contents of the config.json File in the files Folder:
        {
            "Username": "Alice",
            "AdditionalProducts": [
                {"name": "Hat", "category": "Skiing", "price": 10},
                {"name": "Boots", "category":"Skiing", "price": 220.51 },
                {"name": "Gloves", "category":"Skiing", "price": 40.20 }
            ]
        }
████████████████████████████████████████████████████████████████████████
336.os.ReadFile()
    The LoadConfig function uses the ReadFile function to read the contents of the config.json file. 
    The file will be read from the current working directory when the application is executed, 
    which means that I can open the file just with its name.

    The contents of the file are returned as a byte slice, which is converted to a string and written out. The
    LoadConfig function is invoked by an initialization function, which ensures the configuration file is read.

    example:
    readconfig.go:
        package main
        import "os"
        func LoadConfig() (err error) {
            data, err := os.ReadFile("config.json")
            if err == nil {
                Printfln(string(data))
            }
            return
        }
        func init() {
            err := LoadConfig()
            if err != nil {
                Printfln("Error Loading Config: %v", err.Error())
            }
        }
    Output:
        {
            "Username": "Alice",
            "AdditionalProducts": [
                {"name": "Hat", "category": "Skiing", "price": 10},
                {"name": "Boots", "category":"Skiing", "price": 220.51 },
                {"name": "Gloves", "category":"Skiing", "price": 40.20 }
            ]
        }
        Product: Kayak, Category: Watersports, Price: $279.00
        Product: Lifejacket, Category: Watersports, Price: $49.95
        Product: Soccer Ball, Category: Soccer, Price: $19.50
        Product: Corner Flags, Category: Soccer, Price: $34.95
        Product: Stadium, Category: Soccer, Price: $79500.00
        Product: Thinking Cap, Category: Chess, Price: $16.00
        Product: Unsteady Chair, Category: Chess, Price: $75.00
        Product: Bling-Bling King, Category: Chess, Price: $1200.00
████████████████████████████████████████████████████████████████████████
337.Decoding the JSON Data
    example:
    readconfig.go:
        package main
        import (
            "encoding/json"
            "os"
            "strings"
        )
        type ConfigData struct {
            UserName           string
            AdditionalProducts []Product
        }
        var Config ConfigData
        func LoadConfig() (err error) {
            data, err := os.ReadFile("config.json")
            if err == nil {
                decoder := json.NewDecoder(strings.NewReader(string(data)))
                err = decoder.Decode(&Config)
            }
            return
        }
        func init() {
            err := LoadConfig()
            if err != nil {
                Printfln("Error Loading Config: %v", err.Error())
            } else {
                Printfln("Username: %v", Config.UserName)
                Products = append(Products, Config.AdditionalProducts...)
            }
        }
    Output:
        Username: Alice
        Product: Kayak, Category: Watersports, Price: $279.00
        Product: Lifejacket, Category: Watersports, Price: $49.95
        Product: Soccer Ball, Category: Soccer, Price: $19.50
        Product: Corner Flags, Category: Soccer, Price: $34.95
        Product: Stadium, Category: Soccer, Price: $79500.00
        Product: Thinking Cap, Category: Chess, Price: $16.00
        Product: Unsteady Chair, Category: Chess, Price: $75.00
        Product: Bling-Bling King, Category: Chess, Price: $1200.00
        Product: Hat, Category: Skiing, Price: $10.00
        Product: Boots, Category: Skiing, Price: $220.51
        Product: Gloves, Category: Skiing, Price: $40.20
████████████████████████████████████████████████████████████████████████
338.os.Open("config.json")
    Using the File Struct to Read a File
    The Open function opens a file for reading and returns a File value, which represents the open file, and an
    error, which is used to indicate problems opening the file. The File struct implements the Reader interface,
    which makes it simple to read and process the example JSON data, without reading the entire file into a byte
    slice.
    The File struct also implements the Closer interface, which defines a Close method.
    
    The defer keyword can be used to call the Close method when the enclosing function completes,
    like this:
        defer file.Close()
    using the defer keyword ensures that the file is closed even when a function returns early.

    example:
    readconfig.go:
        package main
        import (
            "encoding/json"
            "os"
            "time"
        )
        type ConfigData struct {
            UserName           string
            AdditionalProducts []Product
        }
        var Config ConfigData
        func LoadConfig() (err error) {
            file, err := os.Open("config.json")
            if (err == nil) {
                defer file.Close()
                decoder := json.NewDecoder(file)
                err = decoder.Decode(&Config)
            }
            return
        }
        func init() {
            time.Sleep(time.Second)
        }
        func init() {
            err := LoadConfig()
            if err != nil {
                Printfln("Error Loading Config: %v", err.Error())
            } else {
                Printfln("Username: %v", Config.UserName)
                Products = append(Products, Config.AdditionalProducts...)
            }
        }
    Output:
        Username: Alice
        Product: Kayak, Category: Watersports, Price: $279.00
        Product: Lifejacket, Category: Watersports, Price: $49.95
        Product: Soccer Ball, Category: Soccer, Price: $19.50
        Product: Corner Flags, Category: Soccer, Price: $34.95
        Product: Stadium, Category: Soccer, Price: $79500.00
        Product: Thinking Cap, Category: Chess, Price: $16.00
        Product: Unsteady Chair, Category: Chess, Price: $75.00
        Product: Bling-Bling King, Category: Chess, Price: $1200.00
        Product: Hat, Category: Skiing, Price: $10.00
        Product: Boots, Category: Skiing, Price: $220.51
        Product: Gloves, Category: Skiing, Price: $40.20
████████████████████████████████████████████████████████████████████████
339.Methods Defined by the File Struct for Reading at a Specific Location
    Name                        Description
    --------------------        ------------------------------------------
    ReadAt(slice, offset)       This method is defined by the ReaderAt interface and performs a read into the
                                specific slice at the specified position offset in the file.
    Seek(offset, how)           This method is defined by the Seeker interface and moves the offset into
    جستجو کنید                  the file for the next read. The offset is determined by the combination of the
                                two arguments: the first argument specifies the number of bytes to offset,
                                and the second argument determines how the offset is applied—a value of 0
                                means the offset is relative to the start of the file, a value of 1 means the offset
                                is relative to the current read position, and a value of 2 means the offset is
                                relative to the end of the file.
    
    
    Reading from specific locations requires knowledge of the file structure.
    In this example, I know the location of the data I want to read, 
    which allows me to use the ReadAt method to read the username value
    and the Seek method to jump to the start of the product data.
    
    example:
    readconfig.go:
        package main
        import (
            "os"
            "encoding/json"
            //"strings"
        )
        type ConfigData struct {
            UserName string
            AdditionalProducts []Product
        }
        var Config ConfigData
        func LoadConfig() (err error) {
            file, err := os.Open("config.json")
            if (err == nil) {
                defer file.Close()
                nameSlice := make([]byte, 5)
                file.ReadAt(nameSlice, 19)
                Config.UserName = string(nameSlice)
                file.Seek(55, 0)
                decoder := json.NewDecoder(file)
                err = decoder.Decode(&Config.AdditionalProducts)
            }
            return
        }
        func init() {
            err := LoadConfig()
            if err != nil {
                Printfln("Username: %v", Config.UserName)
                Products = append(Products, Config.AdditionalProducts...)
            } else {
                Printfln("Error Loading Config: %v", err.Error())
            }
        }
    Output:
        Username: Alice
        Product: Kayak, Category: Watersports, Price: $279.00
        Product: Lifejacket, Category: Watersports, Price: $49.95
        Product: Soccer Ball, Category: Soccer, Price: $19.50
        Product: Corner Flags, Category: Soccer, Price: $34.95
        Product: Stadium, Category: Soccer, Price: $79500.00
        Product: Thinking Cap, Category: Chess, Price: $16.00
        Product: Unsteady Chair, Category: Chess, Price: $75.00
        Product: Bling-Bling King, Category: Chess, Price: $1200.00
████████████████████████████████████████████████████████████████████████
340.The os Package Function for Writing Files
    Name                                    Description
    -------------------------------         ----------------------------------------
    WriteFile(name,slice, modePerms)        This function creates a file with the specified name, mode, and permissions and
                                            writes the content of the specified byte slice. If the file already exists, its contents
                                            will be replaced with the byte slice. The result is an error that reports any problems
                                            creating the file or writing the data.
    OpenFile(name, flag, modePerms)         The function opens the file with the specified name, using the flags to control how
                                            the file is opened. If a new file is created, then the specified mode and permissions
                                            are applied. The result is a File value that provides access to the file contents and
                                            an error that indicates problems opening the file.
████████████████████████████████████████████████████████████████████████
341.the Write Convenience Function
    The file mode is used to specify special characteristics for the file, 
    but a value of zero is used for regular files, as in the example. 
    You can find a list of the file mode values and their settings at 
    
    https://golang.org/pkg/io/fs/#FileMode


    https://cs.opensource.google/go/go/+/go1.21.1:src/io/fs/fs.go;l=165
    
    type FileMode 
    type FileMode uint32
    A FileMode represents a file's mode and permission bits. 
    The bits have the same definition on all systems, 
    so that information about files can be moved from one system to another portably. 
    Not all bits apply to all systems. 
    The only required bit is ModeDir for directories.
    
    
    const (
        // The single letters are the abbreviations
        // used by the String method's formatting.
        ModeDir        FileMode = 1 << (32 - 1 - iota) // d: is a directory
        ModeAppend                                     // a: append-only
        ModeExclusive                                  // l: exclusive use
        ModeTemporary                                  // T: temporary file; Plan 9 only
        ModeSymlink                                    // L: symbolic link
        ModeDevice                                     // D: device file
        ModeNamedPipe                                  // p: named pipe (FIFO)
        ModeSocket                                     // S: Unix domain socket
        ModeSetuid                                     // u: setuid
        ModeSetgid                                     // g: setgid
        ModeCharDevice                                 // c: Unix character device, when ModeDevice is set
        ModeSticky                                     // t: sticky
        ModeIrregular                                  // ?: non-regular file; nothing else is known about this file
    
        // Mask for the type bits. For regular files, none will be set.
        ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice | ModeCharDevice | ModeIrregular
    
        ModePerm FileMode = 0777 // Unix permission bits
    )

    The defined file mode bits are the most significant bits of the FileMode. 
    The nine least-significant bits are the standard Unix rwxrwxrwx permissions. 
    The values of these bits should be considered part of the public API and 
    may be used in wire protocols or disk representations: they must not be changed, 
    although new bits might be added.


    example:
    main.go:
        package main
        import (
            "fmt"
            "os"
            "time"
        )
        func main() {
            total := 0.0
            for _, p := range Products {
                total += p.Price
            }
            dataStr := fmt.Sprintf("Time: %v, Total: $%.2f\n",time.Now().Format("Mon 15:04:05"), total)
            err := os.WriteFile("output.txt", []byte(dataStr), 0666)
            if err == nil {
                fmt.Println("Output file created")
            } else {
                Printfln("Error: %v", err.Error())
            }
        }
    Output: a file create with this name= output.txt
        Time: Thu 07:27:34, Total: $279.00
████████████████████████████████████████████████████████████████████████
342.Using the File Struct to Write to a File
    The OpenFile function opens a file and returns a File value. 
    Unlike the Open function, the OpenFile function accepts one or 
    more flags that specify how the file should be opened. 
    The flags are defined as constants in the os package,
    Care must be taken with these flags, not all of which are supported by every operating system.

████████████████████████████████████████████████████████████████████████
343.The File Opening Flags
    Name            Description
    --------        ------------------------------------------
    O_RDONLY        This flag opens the file read-only so that it can be read from but not written to.
    O_WRONLY        This flag opens the file write-only so that it can be written to but not read from.
    O_RDWR          This flag opens the file read-write so that it can be written to and read from.
    O_APPEND        This flag will append writes to the end of the file.
    O_CREATE        This flag will create the file if it doesn't exist.
    O_EXCL          This flag is used in conjunction with O_CREATE to ensure that a new file is created. If the file
                    already exists, this flag will trigger an error.
    O_SYNC          This flag enables synchronous writes, such that data is written to the storage device before
                    the write function/method returns.
    O_TRUNC         This flag truncates the existing content in the file.
████████████████████████████████████████████████████████████████████████
344.Writing to a File
    example:
    main.go:
        package main
        import (
            "fmt"
            "time"
            "os"
        )
        func main() {
            total := 0.0
            for _, p := range Products {
                total += p.Price
            }
            dataStr := fmt.Sprintf("Time: %v, Total: $%.2f\n",
                time.Now().Format("Mon 15:04:05"), total)
        
                
            file, err := os.OpenFile("output.txt",os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0666)
            
            if (err == nil) {
                defer file.Close()
                file.WriteString(dataStr)
            } else {
                Printfln("Error: %v", err.Error())
            }
        }
    

    I combined the O_WRONLY flag to open the file for writing, 
    the O_CREATE file to create if it doesn't already
    exist, and the O_APPEND flag to append any written data to the end of the file.
    Output:
        appended to file exist:
            Time: Thu 07:27:34, Total: $279.00
            Time: Thu 08:17:05, Total: $81174.40
            Time: Thu 08:17:09, Total: $81174.40
        
████████████████████████████████████████████████████████████████████████
345.The File Methods for Writing Data
    Name                        Description
    -----------------------     ---------------------------------------------------
    Seek(offset, how)           This method sets the location for subsequent operations.
    Write(slice)                This method writes the contents of the specified byte slice to the file.
                                The results are the number of bytes written and an error that indicates
                                problems writing the data.
    WriteAt(slice, offset)      This method writes the data in the slice at the specified location and is the
                                counterpart to the ReadAt method.
    WriteString(str)            This method writes a string to the file. This is a convenience method that
                                converts the string to a byte slice, invokes the Write method, and returns the
                                results it receives.
████████████████████████████████████████████████████████████████████████
346.Writing JSON Data to a File
    example:
    main.go:
        package main
        import (
            // "fmt"
            // "time"
            "encoding/json"
            "os"
        )
        func main() {
            cheapProducts := []Product{}
            for _, p := range Products {
                if p.Price < 100 {
                    cheapProducts = append(cheapProducts, p)
                }
            }
            file, err := os.OpenFile("cheap.json", os.O_WRONLY|os.O_CREATE, 0666)
            if err == nil {
                defer file.Close()
                encoder := json.NewEncoder(file)
                encoder.Encode(cheapProducts)
            } else {
                Printfln("Error: %v", err.Error())
            }
        }
    Output:
        create file = cheap.json
        content of this:
    
            [{"Name":"Lifejacket","Category":"Watersports","Price":49.95},{"Name":"Soccer Ball","Category":"Soccer","Price":19.5},{"Name":"Corner Flags","Category":"Soccer","Price":34.95},{"Name":"Thinking Cap","Category":"Chess","Price":16},{"Name":"Unsteady Chair","Category":"Chess","Price":75}]

████████████████████████████████████████████████████████████████████████
347.the Convenience Functions to Create New Files
    The os Package Functions for Creating Files
        The CreateTemp function can be useful, but it is important to understand that the purpose of this
        function is to generate a random filename and that in all other respects, the file that is created is just a
        regular file. The file that is created isn't removed automatically and will remain on the storage device after
        the application has been executed.
    Name                                Description
    -------------------------           --------------------------------------------
    Create(name)                        This function is equivalent to calling OpenFile with the O_RDWR, O_CREATE, and
                                        O_TRUNC flags. The results are the File, which can be used for reading and writing,
                                        and an error that is used to indicate problems creating the file. Note that this
                                        combination of flags means that if a file exists with the specified name, it will be
                                        opened, and its contents will be deleted.
    CreateTemp(dirName, fileName)       This function creates a new file in the directory with the specified name. If the
                                        name is the empty string, then the system temporary directory is used, obtained
                                        using the TempDir function. The file is created with a
                                        name that contains a random sequence of characters, as demonstrated in the text
                                        after the table. The file is opened with the O_RDWR, O_CREATE, and O_EXCL flags.
                                        The file isn't removed when it is closed.
████████████████████████████████████████████████████████████████████████
348.Creating a Temporary File
    The location of the temporary file is specified with a period, meaning the current working directory.
    if the empty string is used, then the file will be created in the default temporary
    directory, which is obtained using the TempDir function described.
    The name of the file can include an asterisk (the * character), 
    and if this is present, the random part of the filename will replace it. 
    If the filename does not contain an asterisk, 
    then the random part of the filename will be added to the end of the name.

    ompile and execute the project, and once execution is complete, you will see a new file in the files
    folder. The file in my project is named tempfile-1732419518.json, but your filename will be different, and
    you will see a new file and a unique name each time the program is executed.

    example:
    main.go:
        package main
        import (
            "os"
            "encoding/json"
        )
        func main() {
            cheapProducts := []Product {}
            for _, p := range Products {
                if (p.Price < 100) {
                    cheapProducts = append(cheapProducts, p)
                }
            }
            file, err := os.CreateTemp(".", "tempfile-*.json")
            if (err == nil) {
                defer file.Close()
                encoder := json.NewEncoder(file)
                encoder.Encode(cheapProducts)
            } else {
                Printfln("Error: %v", err.Error())
            }
        }
    Output:
    tempfile-1982129407.json:
        [{"Name":"Lifejacket","Category":"Watersports","Price":49.95},{"Name":"Soccer Ball","Category":"Soccer","Price":19.5},{"Name":"Corner Flags","Category":"Soccer","Price":34.95},{"Name":"Thinking Cap","Category":"Chess","Price":16},{"Name":"Unsteady Chair","Category":"Chess","Price":75}]
████████████████████████████████████████████████████████████████████████
349.Working with File Paths
    If you want to read and write files in
    other locations, then you must specify file paths. The issue is that not all of the operating systems that Go
    supports express file paths in the same way. For example, the path to a file named mydata.json in my home
    directory on a Linux system might be expressed like this:
        /home/adam/mydata.json

    where the path to the same in my home directory is expressed like this:
        C:\Users\adam\mydata.json
████████████████████████████████████████████████████████████████████████
350.The Common Location Functions Defined by the os Package
    Name                Description
    --------------      ------------------------------------
    Getwd()             This function returns the current working directory, expressed as a string, and an
                        error that indicates problems obtaining the value.
    UserHomeDir()       This function returns the user's home directory and an error that indicates problems
                        obtaining the path.
    UserCacheDir()      This function returns the default directory for user-specific cached data and an error
                        that indicates problems obtaining the path.
    UserConfigDir()     This function returns the default directory for user-specific configuration data and an
                        error that indicates problems obtaining the path.
    TempDir()           This function returns the default directory for temporary files and an error that
                        indicates problems obtaining the path.
████████████████████████████████████████████████████████████████████████
351.The path/filepath Functions for Paths
    Name                    Description
    ------------            ------------------------------------------
    Abs(path)               This function returns an absolute path, which is useful if you have a relative path,
                            such as a filename.
    IsAbs(path)             This function returns true if the specified path is absolute.
    Base(path)              This function returns the last element from the path.
    Clean(path)             This function tidies up path strings by removing duplicate separators and relative references.
    Dir(path)               This function returns all but the last element of the path.
    EvalSymlinks(path)      This function evaluates a symbolic link and returns the resulting path.
    Ext(path)               This function returns the file extension from the specified path, which is
                            assumed to be the suffix following the final period in the path string.
    FromSlash(path)         This function replaces each forward slash with the platform's file separator character.
    ToSlash(path)           This function replaces the platform's file separator with forward slashes.
    Join(...elements)       This function combines multiple elements using the platform's file separator.
    Match(pattern, path)    This function returns true if the path is matched by the specified pattern.
    Split(path)             This function returns the components on either side of the final path separator in
                            the specified path.
    SplitList(path)         This function splits a path into its components, which are returned as a string slice.
    VolumeName(path)        This function returns the volume component of the specified path or the empty
                            string if the path does not contain a volume.
████████████████████████████████████████████████████████████████████████
352.Working with a Path
    example:
    main.go
        package main
        import (
            // "fmt"
            // "time"
            "os"
            //"encoding/json"
            "path/filepath"
        )
        func main() {
            path, err := os.UserHomeDir()
            if (err == nil) {
                path = filepath.Join(path, "MyApp", "MyTempFile.json")
            }
            Printfln("Full path: %v", path)
            Printfln("Volume name: %v", filepath.VolumeName(path))
            Printfln("Dir component: %v", filepath.Dir(path))
            Printfln("File component: %v", filepath.Base(path))
            Printfln("File extension: %v", filepath.Ext(path))
        }
    Output:
        Username: Alice
        Full path: /home/sina/MyApp/MyTempFile.json
        Volume name: 
        Dir component: /home/sina/MyApp
        File component: MyTempFile.json
        File extension: .json
    
    received on the Windows machine:
        Username: Alice
        Full path: C:\Users\adam\MyApp\MyTempFile.json
        Volume name: C:
        Dir component: C:\Users\adam\MyApp
        File component: MyTempFile.json
        File extension: .json
████████████████████████████████████████████████████████████████████████
353.The os Package Functions for Managing Files and Directories
    Name                            Description
    -----------------               ------------------------------------------
    Chdir(dir)                      This function changes the current working directory to the specified directory.
                                    The result is an error that indicates problems making the change.
    Mkdir(name, modePerms)          This function creates a directory with the specified name and mode/
                                    permissions. The result is an error that is nil if the directory is created or that
                                    describes a problem if one arises.
    MkdirAll(name, modePerms)       This function performs the same task as Mkdir but creates any parent directories
                                    in the specified path.
    MkdirTemp(parentDir, name)      This function is similar to CreateTemp but creates a directory rather than a file.
                                    A random string is added to the end of the specified name or in place of an
                                    asterisk, and the new directory is created within the specified parent. The results
                                    are the name of the directory and an error indicating problems.
    Remove(name)                    This function removes the specified file or directory. The result is an error that
                                    describes any problems that arise.
    RemoveAll(name)                 This function removes the specified file or directory. If the name specifies a
                                    directory, then any children it contains are also removed. The result is an error
                                    that describes any problems that arise.
    Rename(old, new)                This function renames the specified file or folder. The result is an error that
                                    describes any problems that arise.
    Symlink(old, new)               This function creates a symbolic link to the specified file. The result is an error
                                    that describes any problems that arise.
████████████████████████████████████████████████████████████████████████
354.Creating Directories
    example:
    main.go:
        package main
        import (
            // "fmt"
            // "time"
            "encoding/json"
            "os"
            "path/filepath"
        )
        func main() {
            path, err := os.UserHomeDir()
            if err == nil {
                path = filepath.Join(path, "0-Repo/TEST-2/MyApp", "MyTempFile.json")
            }
            Printfln("Full path: %v", path)
            err = os.MkdirAll(filepath.Dir(path), 0766)
            if err == nil {
                file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
                if err == nil {
                    defer file.Close()
                    encoder := json.NewEncoder(file)
                    encoder.Encode(Products)
                }
            }
            if err != nil {
                Printfln("Error %v", err.Error())
            }
        }
    Output:
        print this:
            Username: Alice
            Full path: /home/sina/0-Repo/TEST-2/MyApp/MyTempFile.json

        Create this:
            MyTempFile.json in that directory with content:
                [{"Name":"Kayak","Category":"Watersports","Price":279},{"Name":"Lifejacket","Category":"Watersports","Price":49.95},{"Name":"Soccer Ball","Category":"Soccer","Price":19.5},{"Name":"Corner Flags","Category":"Soccer","Price":34.95},{"Name":"Stadium","Category":"Soccer","Price":79500},{"Name":"Thinking Cap","Category":"Chess","Price":16},{"Name":"Unsteady Chair","Category":"Chess","Price":75},{"Name":"Bling-Bling King","Category":"Chess","Price":1200}]
████████████████████████████████████████████████████████████████████████
355.ReadDir(name)
    The os Package Function for Listing Directories
    This function reads the specified directory and returns a DirEntry slice, each of which
    describes an item in the directory.
    The result of the ReadDir function is a slice of values that implement the DirEntry interface, which
    defines the methods.
████████████████████████████████████████████████████████████████████████
356.The Methods Defined by the DirEntry Interface
    Name        Description
    --------    --------------------------
    Name()      This method returns the name of the file or directory described by the DirEntry value.
    IsDir()     This method returns true if the DirEntry value represents a directory.
    Type()      This method returns a FileMode value, which is an alias to uint32, which describes the file more
                and the permissions of the file or directory represented by the DirEntry value.
    Info()      This method returns a FileInfo value that provides additional details about the file or directory
                represented by the DirEntry value.
████████████████████████████████████████████████████████████████████████
357.Useful Methods Defined by the FileInfo Interface
    Name            Description
    ----------      ------------------------------------
    Name()          This method returns a string containing the name of the file or directory.
    Size()          This method returns the size of the file, expressed as an int64 value.
    Mode()          This method returns the file mode and permission settings for the file or directory.
    ModTime()       This method returns the last modified time of the file or directory.
████████████████████████████████████████████████████████████████████████
358.Stat(path)
    The os Package Function for Inspecting a File
    This function accepts a path string. It returns a FileInfo value that describes the file and an
    error, which indicates problems inspecting the file.
████████████████████████████████████████████████████████████████████████
359.Enumerating Files
    example:
    main.go:
        package main
        import (
            "os"
        )
        func main() {
            path, err := os.Getwd()
            if err == nil {
                dirEntries, err := os.ReadDir(path)
                if err == nil {
                    for _, dentry := range dirEntries {
                        Printfln("Entry name: %v, IsDir: %v", dentry.Name(), dentry.IsDir())
                    }
                }
            }
            if err != nil {
                Printfln("Error %v", err.Error())
            }
        }
    Output:
        Username: Alice
        Entry name: .git, IsDir: true
        Entry name: .vscode, IsDir: true
        Entry name: README.md, IsDir: false
        Entry name: U.sh, IsDir: false
        Entry name: cheap.json, IsDir: false
        Entry name: config.json, IsDir: false
        Entry name: go.mod, IsDir: false
        Entry name: main.go, IsDir: false
        Entry name: output.txt, IsDir: false
        Entry name: printer.go, IsDir: false
        Entry name: product.go, IsDir: false
        Entry name: readconfig.go, IsDir: false
████████████████████████████████████████████████████████████████████████
360.Determining Whether a File Exists تعیین اینکه آیا یک فایل وجود دارد یا خیر
    Checking Whether a File Exists:
    main.go:
        package main
        import (
            "os"
        )
        func main() {
            targetFiles := []string { "no_such_file.txt", "config.json" }
            for _, name := range targetFiles {
                info, err := os.Stat(name)
                if os.IsNotExist(err) {
                    Printfln("File does not exist: %v", name)
                } else if err != nil  {
                    Printfln("Other error: %v", err.Error())
                } else {
                    Printfln("File %v, Size: %v", info.Name(), info.Size())
                }
            }
        }
    =============================
    Output:
        Username: Alice
        File does not exist: no_such_file.txt
        File config.json, Size: 253
████████████████████████████████████████████████████████████████████████
361.The path/filepath Function for Locating Files with a Pattern
    Name                        Description
    --------------------        -------------------------------
    Match(pattern, name)        This function matches a single path against a pattern. The results are a bool,
                                which indicates if there is a match, and an error, which indicates problems
                                with the pattern or with performing the match.
    Glob(pathPatten)            This function finds all the files that match the specified pattern. The results
                                are a string slice containing the matched paths and an error that indicates
                                problems with performing the search.
████████████████████████████████████████████████████████████████████████
362.The Search Pattern Syntax for the path/filepath Functions
    Term        Description
    --------    -------------------
     *          This term matches any sequence of characters, excluding the path separator.
     ?          This term matches any single character, excluding the path separator.
     [a-Z]      This term matches any character in the specified range.
████████████████████████████████████████████████████████████████████████
363.Locating Files
    example:
    main.go:
        package main
        import (
            "os"
            "path/filepath"
        )
        func main() {
            path, err := os.Getwd()
            if err == nil {
                matches, err := filepath.Glob(filepath.Join(path, "*.json"))
                if err == nil {
                    for _, m := range matches {
                        Printfln("Match: %v", m)
                    }
                }
            }
            if err != nil {
                Printfln("Error %v", err.Error())
            }
        }
    =============================
    Output:
        Username: Alice
        Match: /home/sina/0-Repo/TEST-2/cheap.json
        Match: /home/sina/0-Repo/TEST-2/config.json
████████████████████████████████████████████████████████████████████████
364.The Function Provided by the path/filepath Package
    Name                        Description
    -----------------------     ----------------------------------
    WalkDir(directory, func)    This function calls the specified function for each file and directory in the
                                specified directory.
████████████████████████████████████████████████████████████████████████
365.Walking a Directory قدم زدن در یک فهرست
    example:
    main.go:
        package main
        import (
            "os"
            "path/filepath"
        )
        func callback(path string, dir os.DirEntry, dirErr error) (err error) {
            info, _ := dir.Info()
            Printfln("Path %v, Size: %v", path, info.Size())
            return
        }
        func main() {
            path, err := os.Getwd()
            if err == nil {
                err = filepath.WalkDir(path, callback)
            } else {
                Printfln("Error %v", err.Error())
            }
        }
    =============================
    Output:
        Username: Alice
        Path /home/sina/0-Repo/TEST-2, Size: 4096
        Path /home/sina/0-Repo/TEST-2/.git, Size: 4096
        Path /home/sina/0-Repo/TEST-2/.git/COMMIT_EDITMSG, Size: 4
        Path /home/sina/0-Repo/TEST-2/.git/FETCH_HEAD, Size: 92
        Path /home/sina/0-Repo/TEST-2/.git/refs/remotes/origin/main, Size: 41
        Path /home/sina/0-Repo/TEST-2/.git/refs/tags, Size: 4096
        Path /home/sina/0-Repo/TEST-2/.vscode, Size: 4096
        Path /home/sina/0-Repo/TEST-2/.vscode/extensions.json, Size: 79
        Path /home/sina/0-Repo/TEST-2/.vscode/settings.json, Size: 405
        Path /home/sina/0-Repo/TEST-2/README.md, Size: 9
        Path /home/sina/0-Repo/TEST-2/U.sh, Size: 68
        Path /home/sina/0-Repo/TEST-2/cheap.json, Size: 487
        Path /home/sina/0-Repo/TEST-2/config.json, Size: 253
        Path /home/sina/0-Repo/TEST-2/go.mod, Size: 22
        Path /home/sina/0-Repo/TEST-2/main.go, Size: 8579
        Path /home/sina/0-Repo/TEST-2/output.txt, Size: 109
        Path /home/sina/0-Repo/TEST-2/printer.go, Size: 129
        Path /home/sina/0-Repo/TEST-2/product.go, Size: 474
        Path /home/sina/0-Repo/TEST-2/readconfig.go, Size: 704
████████████████████████████████████████████████████████████████████████
366.Putting HTML and Text Templates in Context
    Question Answer:

    What are they?
    These templates allow HTML and text content to be generated dynamically
    from Go data values.
    
    Why are they useful?
    Templates are useful when large amounts of content are required, such that
    defining the content as strings would be unmanageable.
    
    How are they used?
    The templates are HTML or text files, which are annotated with instructions
    for the template processing engine. When a template is rendered, the
    instructions are processed to generate HTML or text content.

    Are there any pitfalls or limitations?
    The template syntax is counterintuitive and is not checked by the Go
    compiler. This means that care must be taken to use the correct syntax,
    which can be a frustrating process.

    Are there any alternatives?
    Templates are optional, and smaller amounts of content can be produced
    using strings.

    Summary:
    Problem                                         Solution
    ------------------                              -----------------------------------
    Generate an HTML document                       Define an HTML template with actions that
                                                    incorporate data values into the output. Load and
                                                    execute the templates, providing data for the actions.
    Enumerate loaded templates                      Enumerate the results of the Templates method.
    Locate a specific template                      Use the Lookup method.
    Produce dynamic content                         Use a template action.
    Format a data value                             Use the formatting functions.
    Suppress whitespace                             Add hyphens to the template.
    Process a slice                                 Use the slice functions.
    Conditionally execute template content          Use the conditional actions and functions.
    Create a nested template                        Use the define and template actions.
    Define a default template                       Use the block and template actions.
    Create functions for use in a template          Define template functions.
    Disable encoding for function results           Return one of the type aliases defined by the html/template package.
    Store data values for later use in a template   Define template variables.
    Generate a text document                        Use the text/template package.



    Preparing for This Chapter:
    1- go mod init htmltext
    2- printer.go:
        package main
        import "fmt"
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template + "\n", values...)
        }
    3- product.go:
        package main
        type Product struct {
            Name, Category string
            Price float64
        }
        var Kayak = Product {
            Name: "Kayak",
            Category: "Watersports",
            Price: 279,
        }
        var Products = []Product {
            { "Kayak", "Watersports", 279 },
            { "Lifejacket", "Watersports", 49.95 },
            { "Soccer Ball", "Soccer", 19.50 },
            { "Corner Flags", "Soccer", 34.95 },
            { "Stadium", "Soccer", 79500 },
            { "Thinking Cap", "Chess", 16 },
            { "Unsteady Chair", "Chess", 75 },
            { "Bling-Bling King", "Chess", 1200 },
        }
        func (p *Product) AddTax() float64 {
            return p.Price * 1.2
        }
        func (p * Product) ApplyDiscount(amount float64) float64 {
            return p.Price - amount
        }
    4- main.go:
        package main
        func main() {
            for _, p := range Products {
                Printfln("Product: %v, Category: %v, Price: $%.2f",
                    p.Name, p.Category, p.Price)
            }
        }
    
████████████████████████████████████████████████████████████████████████
367.Creating HTML Templates
    The html/template package provides support for creating templates that are processed using a data
    structure to generate dynamic HTML output.
    Templates contain static content mixed with expressions that are enclosed in double curly braces,
    known as actions.

    The template uses the simplest action, which is a period (the . character)
    and which prints out the data used to execute the template, 
    which I explain in the next section.

    example:
    template.html:
        <h1>Template Value: {{ . }}</h1>
    
    A project can contain multiple templates files.
    extras.html:
        <h1>Extras Template Value: {{ . }}</h1>
    
    The new template uses the same action as the previous example but has different static content to make
    it clear which template has been executed in the next section. Once I have described the basic techniques for
    using templates, I'll introduce more complex template actions.
████████████████████████████████████████████████████████████████████████
368.Loading and Executing Templates
    Using templates is a two-step process. 
    First, the templates files are loaded and processed to create Template values.

    The html/template Functions for Loading Template Files:
    Name                        Description
    ---------------------       --------------------------------------------
    ParseFiles(...files)        This function loads one or more files, which are specified by name. The result
                                is a Template that can be used to generate content and an error that reports
                                problems loading the templates.
    ParseGlob(pattern)          This function loads one or more files, which are selected with a pattern. The
                                result is a Template that can be used to generate content and an error that
                                reports problems loading the templates.

    If you name your template files consistently, 
    then you can use the ParseGlob function to load them with a simple pattern. 
    If you want specific files—or the files are not named consistently—then you can specify
    individual files using the ParseFiles function.
████████████████████████████████████████████████████████████████████████
369.The Template Methods for Selecting and Executing Templates
    Name                                            Description
    ----------------------------                    -------------------------------------------
    Templates()                                     This function returns a slice containing pointers to the Template values that
                                                    have been loaded.
    Lookup(name)                                    This function returns a *Template for the specified loaded template.
    Name()                                          This method returns the name of the Template.
    Execute(writer, data)                           This function executes the Template, using the specified data and writes
                                                    the output to the specified Writer.
    ExecuteTemplate(writer, templateName, data)     This function executes the template with the specified name and data and
                                                    writes the output to the specified Writer.
████████████████████████████████████████████████████████████████████████
370.Loading and Executing a Template
    example:
    main.go:
        package main
        import (
            "fmt"
            "html/template"
            "os"
        )
        func main() {
            t, err := template.ParseFiles("templates/template.html")
            if (err == nil) {
                t.Execute(os.Stdout, &Kayak)
                fmt.Println()
            } else {
                Printfln("Error: %v", err.Error())
            }
        }
    =============================
    Output:
        <h1>Template Value: {Kayak Watersports 279}</h1>
████████████████████████████████████████████████████████████████████████
371.Loading Multiple Templates
    There are two approaches to working with multiple templates. 
    The first is to create a separate Template value
    for each of them and execute them separately.

    example:
    Using Separate Templates:
    main.go:
        package main
        import (
            "fmt"
            "html/template"
            "os"
        )
        func main() {
            t1, err1 := template.ParseFiles("templates/template.html")
            t2, err2 := template.ParseFiles("templates/extras.html")
            if (err1 == nil && err2 == nil) {
                t1.Execute(os.Stdout, &Kayak)
                os.Stdout.WriteString("\n")
                t2.Execute(os.Stdout, &Kayak)
                os.Stdout.WriteString("\n")
            } else {
                Printfln("Error: %v %v", err1.Error(), err2.Error())
            }
        }
    =============================
    Output:
        <h1>Template Value: {Kayak Watersports 279}</h1>
        <h1>Extras Template Value: {Kayak Watersports 279}</h1>
████████████████████████████████████████████████████████████████████████
372.Using a Combined Template
    When multiple files are loaded with the ParseFiles, 
    the result is a Template value on which the
    ExecuteTemplate method can be called to execute a specified template. 
    The filename is used as the template name, 
    which means that the templates in this example are named template.html and extras.html.

    You can call the Execute method on the Template returned by the ParseFiles or ParseGlob
    function, and the first template that was loaded will be selected 
    and used to produce the output. 
    Take care when using the ParseGlob function because 
    the first template loaded—and therefore the template that will be
    executed—may not be the file you expect.

    example:
    main.go:
        package main
        import (
            "html/template"
            "os"
        )
        func main() {
                allTemplates, err1 := template.ParseFiles("templates/template.html",
                    "templates/extras.html")
                if (err1 == nil) {
                    allTemplates.ExecuteTemplate(os.Stdout, "template.html", &Kayak)
                    os.Stdout.WriteString("\n")
                    allTemplates.ExecuteTemplate(os.Stdout, "extras.html", &Kayak)
                } else {
                    Printfln("Error: %v %v", err1.Error())
                }
            }
    =============================
    Output:
        <h1>Template Value: {Kayak Watersports 279}</h1>
        <h1>Extras Template Value: {Kayak Watersports 279}</h1>
████████████████████████████████████████████████████████████████████████
373.Enumerating Loaded Templates
    It can be useful to enumerate the templates that have been loaded, 
    especially when using the ParseGlob function, 
    to make sure that all the expected files have been discovered.

    example:
    main.go:
        package main
        import (
            "html/template"
        )
        func main() {
                allTemplates, err := template.ParseGlob("templates/*.html")
            if (err == nil) {
                for _, t := range allTemplates.Templates() {
                    Printfln("Template name: %v", t.Name())
                }
            } else {
                Printfln("Error: %v %v", err.Error())
            }
        }
    =============================
    Output:
        Template name: extras.html
        Template name: template.html
████████████████████████████████████████████████████████████████████████
374.Looking Up a Specific Template
    An alternative to specifying a name is to use the Lookup method to select a template, 
    which is useful when
    you want to pass a template as an argument to a function

    example:
    main.go:
        package main
        import (
            "html/template"
            "os"
        )
        func Exec(t *template.Template) error {
            return t.Execute(os.Stdout, &Kayak)
        }
        func main() {
            allTemplates, err := template.ParseGlob("templates/*.html")
            if err == nil {
                selectedTemplated := allTemplates.Lookup("template.html")
                err = Exec(selectedTemplated)
            }
            if err != nil {
                Printfln("Error: %v %v", err.Error())
            }
        }
    =============================
    Output:
        <h1>Template Value: {Kayak Watersports 279}</h1>
████████████████████████████████████████████████████████████████████████
375.The Template Actions
    Action                      Description
    ------------                -----------------------------------------
    {{ value }}                 This action inserts a data value or the result of an expression into the
    {{ expr }}                  template. A period is used to refer to the data value passed to the Execute or
                                ExecuteTemplate function. See the “Inserting Data Values” section for details.

    {{ value.fieldname }}       This action inserts the value of a struct field. See the “Inserting Data Values” section for details.
    {{ value.method arg }}      This action invokes a method and inserts the result into the template
                                output. Parentheses are not used, and arguments are separated by
                                spaces. See the “Inserting Data Values” section for details.

    {{ func arg }}              This action invokes a function and inserts the result into the output.
                                There are built-in functions for common tasks, such as formatting
                                data values, and custom functions can be defined, as described in the
                                “Defining Template Functions” section.

    {{ expr | value.method }}   Expressions can be chained together using a vertical bar so that the result
    {{ expr | func              of the first expression is used as the last argument in the second expression.
    {{ range value }}           This action iterates through the specified slice and adds the content
    ...                         between the range and end keyword for each element. The actions within
    {{ end }}                   the nested content are executed, with the current element accessible
                                through the period. See the “Using Slices in Templates” section for details.

    {{ range value }}           This action is similar to the range/end combination but defines a section
    ...                         of nested content that is used if the slice contains no elements.
    {{ else }}
    ...
    {{ end }}

    {{ if expr }}               This action evaluates an expression and executes the nested template
    ...                         content if the result is true, as demonstrated in the “Conditionally
    {{ end }}                   Executing Template Content” section. This action can be used with
                                optional else and else if clauses.

    {{ with expr }}             This action evaluates an expression and executes the nested template
    ...                         content if the result isn't nil or the empty string. This action can be used
    {{ end }}                   with optional clauses.

    {{ define "name" }}         This action defines a template with the specified name
    ...
    {{ end }}

    {{ template "name" expr }}  This action executes the template with the specified name and data and
                                inserts the result in the output.

    {{ block "name" expr }}     This action defines a template with the specified name and invokes it
    ...                         with the specified data. This is typically used to define a template that
    {{ end }}                   can be replaced by one loaded from another file, as demonstrated in the
                                “Defining Template Blocks” section.
████████████████████████████████████████████████████████████████████████
376.The Template Expressions for Inserting Values into Templates
    Inserting Data Values

    Expression          Description
    ------------        -----------------------------------------------
    .                   This expression inserts the value passed to the Execute or ExecuteTemplate method into the
                        template output.
    .Field              This expression inserts the value of the specified field into the template output.
    .Method             This expression calls the specified method without arguments and inserts the result into the
                        template output.
    .Method             This expression calls the specified method with the specified argument and inserts the result
    arg                 into the template output.
    call                This expression invokes a struct function field, using the specified arguments, which are
    .Field arg          separated by spaces. The result from the function is inserted into the template output.
    
████████████████████████████████████████████████████████████████████████
377.Inserting Data Values in the template.html
    Unlike Go code, methods are not invoked with parentheses, and arguments
    are simply specified after the name, separated by spaces. 
    It is the responsibility of the developer to ensure
    that arguments are of a type that can be used by the method or function.

    example:
    templates/template.html:
        <h1>Template Value: {{ . }}</h1>
        <h1>Name: {{ .Name }}</h1>
        <h1>Category: {{ .Category }}</h1>
        <h1>Price: {{ .Price }}</h1>
        <h1>Tax: {{ .AddTax }}</h1>
        <h1>Discount Price: {{ .ApplyDiscount 10 }}</h1>
    =============================
    Output:
        <h1>Template Value: {Kayak Watersports 279}</h1>
        <h1>Name: Kayak</h1>
        <h1>Category: Watersports</h1>
        <h1>Price: 279</h1>
        <h1>Tax: 334.8</h1>
        <h1>Discount Price: 269</h1>
████████████████████████████████████████████████████████████████████████
378.Understanding Contextual Escaping
    Values are automatically escaped to make them safe for inclusion in HTML, CSS, and JavaScript code,
    with the appropriate escaping rules applied based on context. For example, a string value such as
    "It was a <big> boat" used as the text content of an HTML element would be inserted into the
    template as "It was a <big> boat" but as "It was a \u003cbig\u003e boat" when used
    as a string literal value in JavaScript code. Full details of how values are escaped can be found at

    https://golang.org/pkg/html/template.

████████████████████████████████████████████████████████████████████████
379.The Built-in Templates Functions for Formatting Data
    Name        Description
    -------     --------------------------------
    print       This is an alias to the fmt.Sprint function.
    printf      This is an alias to the fmt.Sprintf function.
    println     This is an alias to the fmt.Sprintln function.
    html        This function encodes a value for safe inclusion in an HTML document.
    js          This function encodes a value for safe inclusion in a JavaScript document.
    urlquery    This function encodes a value for use in a URL query string.


    example:
    template.html:
        <h1>Template Value: {{ . }}</h1>
        <h1>Name: {{ .Name }}</h1>
        <h1>Category: {{ .Category }}</h1>
        <h1>Price: {{ printf "$%.3f" .Price }}</h1>
        <h1>Tax: {{ printf "$%.2f" .AddTax }}</h1>
        <h1>Discount Price: {{ .ApplyDiscount 10 }}</h1>
    =============================
    Output:
        <h1>Template Value: {Kayak Watersports 279}</h1>
        <h1>Name: Kayak</h1>
        <h1>Category: Watersports</h1>
        <h1>Price: $279.000</h1>
        <h1>Tax: $334.80</h1>
        <h1>Discount Price: 269</h1>
████████████████████████████████████████████████████████████████████████
380.Chaining Expressions
    Chaining expressions creates a pipeline for values, 
    which allows the output from one method or function
    to be used as the input for another.

    example:
    template.html:
        <h1>Template Value: {{ . }}</h1>
        <h1>Name: {{ .Name }}</h1>
        <h1>Category: {{ .Category }}</h1>
        <h1>Price: {{ printf "$%.2f" .Price }}</h1>
        <h1>Tax: {{ printf "$%.2f" .AddTax }}</h1>
        <h1>Discount Price: {{ .ApplyDiscount 10 | printf "$%.2f" }}</h1>
    =============================
    Output:
        <h1>Template Value: {Kayak Watersports 279}</h1>
        <h1>Name: Kayak</h1>
        <h1>Category: Watersports</h1>
        <h1>Price: $279.00</h1>
        <h1>Tax: $334.80</h1>
        <h1>Discount Price: $269.00</h1>
████████████████████████████████████████████████████████████████████████
381.Using Parentheses in html 
    Chaining can be used only for the last argument provided to a function. 
    An alternative approach—and
    one that can be used to set other function arguments is to use parentheses
    
    example:
    template.html:
        <h1>Template Value: {{ . }}</h1>
        <h1>Name: {{ .Name }}</h1>
        <h1>Category: {{ .Category }}</h1>
        <h1>Price: {{ printf "$%.2f" .Price }}</h1>
        <h1>Tax: {{ printf "$%.2f" .AddTax }}</h1>
        <h1>Discount Price: {{ printf "$%.2f" (.ApplyDiscount 10) }}</h1>
    =============================
    Output:
        <h1>Template Value: {Kayak Watersports 279}</h1>
        <h1>Name: Kayak</h1>
        <h1>Category: Watersports</h1>
        <h1>Price: $279.00</h1>
        <h1>Tax: $334.80</h1>
        <h1>Discount Price: $269.00</h1>
████████████████████████████████████████████████████████████████████████
382.Trimming Whitespace
    HTML isn't sensitive to the whitespace between elements, 
    but whitespace can still cause problems for text content and attribute values, 
    especially when you want to structure the content
    of a template to make it easy to read.

    example:
    template.html
        <h1>
            Name: {{ .Name }}, Category: {{ .Category }}, Price,
                {{ printf "$%.2f" .Price }}
        </h1>
    =============================
    Output:
        <h1>
            Name: Kayak, Category: Watersports, Price,
                $279.00
        </h1>
████████████████████████████████████████████████████████████████████████
383.The minus sign must
    The effect is to remove all of the whitespace to before or after the action.

    The minus sign can be used to trim whitespace, 
    applied immediately after or before the braces
    that open or close an action.

    example:
    template.html:
        <h1>
                Name: {{ .Name }}, Category: {{ .Category }}, Price,
                    {{ printf "$%.2f" .Price }}
        </h1>

        <h1>
                Name: {{ .Name }}, Category: {{ .Category }}, Price,
                    {{- printf "$%.2f" .Price -}}
        </h1>
    =============================
    Output:
        <h1>
                Name: Kayak, Category: Watersports, Price,
                    $279.00
        </h1>

        <h1>
                Name: Kayak, Category: Watersports, Price,$279.00</h1>
████████████████████████████████████████████████████████████████████████
384.Trimming Additional Whitespace
    The whitespace around the final action has been removed, 
    but there is still a newline character after
    the opening h1 tag because the whitespace trimming applies only to actions.

    example:
    template.html:
        <h1>
            {{- "" -}} Name: {{ .Name }}, Category: {{ .Category }}, Price,
                {{- printf "$%.2f" .Price -}}
        </h1>
    =============================
    Output:
        <h1>Name: Kayak, Category: Watersports, Price,$279.00</h1>
████████████████████████████████████████████████████████████████████████
385.Slices in Templates
    Template actions can be used to generate content for slices

    example:
    Processing a Slice in the template.html
        {{ range . -}}
            <h1>Name: {{ .Name }}, Category: {{ .Category }}, Price,
                {{- printf "$%.2f" .Price }}</h1>
        {{ end }}
    main.go:
        package main

        import (
            "html/template"
            "os"
        )
        
        func Exec(t *template.Template) error {
            return t.Execute(os.Stdout, Products)
        }
        func main() {
            allTemplates, err := template.ParseGlob("templates/*.html")
            if err == nil {
                selectedTemplated := allTemplates.Lookup("template.html")
                err = Exec(selectedTemplated)
            }
        
            
            if err != nil {
                Printfln("Error: %v %v", err.Error())
            }
        }
    =============================
    Output:
        <h1>Name: Kayak, Category: Watersports, Price,$279.00</h1>
        <h1>Name: Lifejacket, Category: Watersports, Price,$49.95</h1>
        <h1>Name: Soccer Ball, Category: Soccer, Price,$19.50</h1>
        <h1>Name: Corner Flags, Category: Soccer, Price,$34.95</h1>
        <h1>Name: Stadium, Category: Soccer, Price,$79500.00</h1>
        <h1>Name: Thinking Cap, Category: Chess, Price,$16.00</h1>
        <h1>Name: Unsteady Chair, Category: Chess, Price,$75.00</h1>
        <h1>Name: Bling-Bling King, Category: Chess, Price,$1200.00</h1>
████████████████████████████████████████████████████████████████████████
386.The Built-in Template Functions for Slices
    Go text templates support the built-in functions

    Name        Description
    ------      --------------
    slice       This function creates a new slice. Its arguments are the original slice, the start index, and the end index.
    index       This function returns the element at the specified index.
    len         This function returns the length of the specified slice.

    example:
    Built-in Functions in the template.html:
        <h1>There are {{ len . }} products in the source data.</h1>
        <h1>First product: {{ index . 0 }}</h1>
        {{ range slice . 3 6 -}} 
            <h1>Name: {{ .Name }}, Category: {{ .Category }}, Price,
                {{- printf "$%.2f" .Price }}</h1>
        {{ end }}
    =============================
    Output:
        <h1>There are 8 products in the source data.</h1>
        <h1>First product: {Kayak Watersports 279}</h1>
        <h1>Name: Corner Flags, Category: Soccer, Price,$34.95</h1>
        <h1>Name: Stadium, Category: Soccer, Price,$79500.00</h1>
        <h1>Name: Thinking Cap, Category: Chess, Price,$16.00</h1>
    
████████████████████████████████████████████████████████████████████████
387.Conditionally Executing Template Content
    Actions can be used to conditionally insert content into the output based on the evaluation of their
    expressions.

    The Template Conditional Functions:
        Function            Description
        ------------        -----------------------------------------------
        eq arg1 arg2        This function returns true if arg1 == arg2.
        ne arg1 arg2        This function returns true if arg1 != arg2.
        lt arg1 arg2        This function returns true if arg1 < arg2.
        le arg1 arg2        This function returns true if arg1 <= arg2.
        gt arg1 arg2        This function returns true if arg1 > arg2.
        ge arg1 arg2        This function returns true if arg1 >= arg2.
        and arg1 arg2       This function returns true if both arg1 and arg2 are true.
        not arg1            This function returns true if arg1 is false, and false if it is true.

    example:
    a Conditional Action in the template.html:
        <h1>There are {{ len . }} products in the source data.</h1>
        <h1>First product: {{ index . 0 }}</h1>
        {{ range . -}}
            {{ if lt .Price 100.00 -}}
                <h1>Name: {{ .Name }}, Category: {{ .Category }}, Price,
                    {{- printf "$%.2f" .Price }}</h1>
            {{ end -}}
        {{ end }}
    =============================
    Output:
        <h1>There are 8 products in the source data.</h1>
        <h1>First product: {Kayak Watersports 279}</h1>
        <h1>Name: Lifejacket, Category: Watersports, Price,$49.95</h1>
            <h1>Name: Soccer Ball, Category: Soccer, Price,$19.50</h1>
            <h1>Name: Corner Flags, Category: Soccer, Price,$34.95</h1>
            <h1>Name: Thinking Cap, Category: Chess, Price,$16.00</h1>
            <h1>Name: Unsteady Chair, Category: Chess, Price,$75.00</h1>        

    Despite the use of the minus sign to trim whitespace, 
    the output is oddly formatted because of the
    way I chose to structure the template.
    
████████████████████████████████████████████████████████████████████████
388.Using the Optional Conditional Actions
    The if action can be used with optional else and else if keywords

    example:
    template.html:
        <h1>There are {{ len . }} products in the source data.</h1>
        <h1>First product: {{ index . 0 }}</h1>
        {{ range . -}}
            {{ if lt .Price 100.00 -}}
                <h1>Name: {{ .Name }}, Category: {{ .Category }}, Price,
                    {{- printf "$%.2f" .Price }}</h1>
            {{ else if gt .Price 1500.00 -}}
                <h1>Expensive Product {{ .Name }} ({{ printf "$%.2f" .Price}})</h1>
            {{ else -}}
                <h1>Midrange Product: {{ .Name }} ({{ printf "$%.2f" .Price}})</h1>
            {{ end -}}
        {{ end }}
    main.go:
        package main
        import (
            "html/template"
            "os"
        )
        func Exec(t *template.Template) error {
            return t.Execute(os.Stdout, Products)
        }
        func main() {
            allTemplates, err := template.ParseGlob("templates/*.html")
            if err == nil {
                selectedTemplated := allTemplates.Lookup("template.html")
                err = Exec(selectedTemplated)
            }
            if err != nil {
                Printfln("Error: %v %v", err.Error())
            }
        }
    =============================
    Output:
        <h1>There are 8 products in the source data.</h1>
        <h1>First product: {Kayak Watersports 279}</h1>
        <h1>Midrange Product: Kayak ($279.00)</h1>
            <h1>Name: Lifejacket, Category: Watersports, Price,$49.95</h1>
            <h1>Name: Soccer Ball, Category: Soccer, Price,$19.50</h1>
            <h1>Name: Corner Flags, Category: Soccer, Price,$34.95</h1>
            <h1>Expensive Product Stadium ($79500.00)</h1>
            <h1>Name: Thinking Cap, Category: Chess, Price,$16.00</h1>
            <h1>Name: Unsteady Chair, Category: Chess, Price,$75.00</h1>
            <h1>Midrange Product: Bling-Bling King ($1200.00)</h1>
                
████████████████████████████████████████████████████████████████████████
389.Creating Named Nested Templates
    The define action is used to create a nested template that can be executed by name, 
    which allows content to
    be defined once and used repeatedly with the template action

    example:
    template.html:
        {{ define "currency" }}{{ printf "$%.2f" . }}{{ end }}
        {{ define "basicProduct" -}}
            Name: {{ .Name }}, Category: {{ .Category }}, Price,
                {{- template "currency" .Price }}
        {{- end }}
        {{ define "expensiveProduct" -}}
            Expensive Product {{ .Name }} ({{ template "currency" .Price }})
        {{- end }}
        <h1>There are {{ len . }} products in the source data.</h1>
        <h1>First product: {{ index . 0 }}</h1>
        {{ range . -}}
            {{ if lt .Price 100.00 -}}
                <h1>{{ template "basicProduct" . }}</h1>
            {{ else if gt .Price 1500.00 -}}
                <h1>{{ template "expensiveProduct" . }}</h1>
            {{ else -}}
                <h1>Midrange Product: {{ .Name }} ({{ printf "$%.2f" .Price}})</h1>
            {{ end -}}
        {{ end }}
    =============================
    Output:



        <h1>There are 8 products in the source data.</h1>
        <h1>First product: {Kayak Watersports 279}</h1>
        <h1>Midrange Product: Kayak ($279.00)</h1>
            <h1>Name: Lifejacket, Category: Watersports, Price,$49.95</h1>
            <h1>Name: Soccer Ball, Category: Soccer, Price,$19.50</h1>
            <h1>Name: Corner Flags, Category: Soccer, Price,$34.95</h1>
            <h1>Expensive Product Stadium ($79500.00)</h1>
            <h1>Name: Thinking Cap, Category: Chess, Price,$16.00</h1>
            <h1>Name: Unsteady Chair, Category: Chess, Price,$75.00</h1>
            <h1>Midrange Product: Bling-Bling King ($1200.00)</h1>


    The define keyword is followed by the template name in quotes, 
    and the template is terminated by the end keyword. 
    The template keyword is used to execute a named template, 
    specifying the template name and a data value:
        ...
        {{- template "currency" .Price }}
        ...
████████████████████████████████████████████████████████████████████████
390.Selecting a Named Template in the main.go
    Using the define and end keywords for the main template content excludes the whitespace used to
    separate the other named templates.

    Adding a Named Template in the template.html:
        {{ define "currency" }}{{ printf "$%.2f" . }}{{ end }}
        {{ define "basicProduct" -}}
            Name: {{ .Name }}, Category: {{ .Category }}, Price,
                {{- template "currency" .Price }}
        {{- end }}
        {{ define "expensiveProduct" -}}
            Expensive Product {{ .Name }} ({{ template "currency" .Price }})
        {{- end }}
        {{ define "mainTemplate" -}}
            <h1>There are {{ len . }} products in the source data.</h1>
            <h1>First product: {{ index . 0 }}</h1>
            {{ range . -}}
                {{ if lt .Price 100.00 -}}
                    <h1>{{ template "basicProduct" . }}</h1>
                {{ else if gt .Price 1500.00 -}}
                    <h1>{{ template "expensiveProduct" . }}</h1>
                {{ else -}}
                    <h1>Midrange Product: {{ .Name }} ({{ printf "$%.2f" .Price}})</h1>
                {{ end -}}
            {{ end }}
        {{- end}}    
    =============================
    main.go:
        package main
        import (
            "html/template"
            "os"
        )
        func Exec(t *template.Template) error {
            return t.Execute(os.Stdout, Products)
        }
        func main() {
            allTemplates, err := template.ParseGlob("templates/*.html")
            if (err == nil) {
                selectedTemplated := allTemplates.Lookup("mainTemplate")
                err = Exec(selectedTemplated)
            }
            if (err != nil) {
                Printfln("Error: %v %v", err.Error())
            }
        }
    =============================
    Any of the named templates can be executed directly, but I have selected the mainTemplate
    Output:
        <h1>There are 8 products in the source data.</h1>
        <h1>First product: {Kayak Watersports 279}</h1>
        <h1>Midrange Product: Kayak ($279.00)</h1>
            <h1>Name: Lifejacket, Category: Watersports, Price,$49.95</h1>
            <h1>Name: Soccer Ball, Category: Soccer, Price,$19.50</h1>
            <h1>Name: Corner Flags, Category: Soccer, Price,$34.95</h1>
            <h1>Expensive Product Stadium ($79500.00)</h1>
            <h1>Name: Thinking Cap, Category: Chess, Price,$16.00</h1>
            <h1>Name: Unsteady Chair, Category: Chess, Price,$75.00</h1>
            <h1>Midrange Product: Bling-Bling King ($1200.00)</h1>
████████████████████████████████████████████████████████████████████████
391.Defining Template Blocks
    Template blocks are used to define a template with default content that can be overridden in another
    template file, which requires multiple templates to be loaded and executed together. 

    This is often used to common content, such as a layout.

    The templates must be loaded so that the file that contains the block action is loaded before the file that
    contains the define action that redefines the template. 
    
    When the templates are loaded, the template defined
    in the list.html file redefines the template named body so that the content in the list.html file replaces
    the content in the template.html file.


    example:
    template.html File in the templates Folder
        {{ define "mainTemplate" -}}
            <h1>This is the layout header</h1>
            {{ block "body" . }}
                <h2>There are {{ len . }} products in the source data.</h2>
            {{ end }}
            <h1>This is the layout footer</h1>
        {{ end }}

    Output:
        <h1>This is the layout header</h1>
        
                <h2>There are 8 products in the source data.</h2>

            <h1>This is the layout footer</h1>
    

    When used alone, the output from the template file includes the content in the block. 
    But this content can be redefined by another template file.
    example:
    list.html:
        {{ define "body" }}
            {{ range . }}
                <h2>Product: {{ .Name }} ({{ printf "$%.2f" .Price}})</h2>
            {{ end -}}
        {{ end }}
    ========================
    main.go:
        package main

        import (
            "html/template"
            "os"
        )
        
        func Exec(t *template.Template) error {
            return t.Execute(os.Stdout, Products)
        }
        func main() {
            
            allTemplates, err := template.ParseFiles("templates/template.html","templates/list.html")
        
        
        
            if err == nil {
                selectedTemplated := allTemplates.Lookup("mainTemplate")
                err = Exec(selectedTemplated)
            }
            if err != nil {
                Printfln("Error: %v %v", err.Error())
            }
        }
    =============================
    Output:
        <h1>This is the layout header</h1>
        
            
            <h2>Product: Kayak ($279.00)</h2>

            <h2>Product: Lifejacket ($49.95)</h2>

            <h2>Product: Soccer Ball ($19.50)</h2>

            <h2>Product: Corner Flags ($34.95)</h2>

            <h2>Product: Stadium ($79500.00)</h2>

            <h2>Product: Thinking Cap ($16.00)</h2>

            <h2>Product: Unsteady Chair ($75.00)</h2>

            <h2>Product: Bling-Bling King ($1200.00)</h2>

        <h1>This is the layout footer</h1>    
████████████████████████████████████████████████████████████████████████
392.Defining Template Functions
    example:
    main.go File in the htmltext Folder:
        package main
        import (
            "html/template"
            "os"
        )
        func GetCategories(products []Product) (categories []string) {
            catMap := map[string]string {}
            for _, p := range products {
                if (catMap[p.Category] == "") {
                    catMap[p.Category] = p.Category
                    categories = append(categories, p.Category)
                }
            }
            return
        }
        func Exec(t *template.Template) error {
            return t.Execute(os.Stdout, Products)
        }
        func main() {
            allTemplates := template.New("allTemplates")
            allTemplates.Funcs(map[string]interface{} {
                "getCats": GetCategories,
            })
            allTemplates, err := allTemplates.ParseGlob("templates/*.html")
            if (err == nil) {
                selectedTemplated := allTemplates.Lookup("mainTemplate")
                err = Exec(selectedTemplated)
            }
            if (err != nil) {
                Printfln("Error: %v %v", err.Error())
            }
        }
    =============================
    Output:
        <h1>This is the layout header</h1>
        
            <h2>There are 8 products in the source data.</h2>

        <h1>This is the layout footer</h1>    
████████████████████████████████████████████████████████████████████████
393.Using a Custom Function in the template.html
    example:
    template.html:
        {{ define "mainTemplate" -}}
            <h1>There are {{ len . }} products in the source data.</h1>
            {{ range getCats .  -}}
                <h1>Category: {{ . }}</h1>
            {{ end }}
        {{- end }}
    =============================================
    Output:
        <h1>There are 8 products in the source data.</h1>
            <h1>Category: Watersports</h1>
            <h1>Category: Soccer</h1>
            <h1>Category: Chess</h1>
████████████████████████████████████████████████████████████████████████
394.Creating an HTML Fragment in the main.go
    example:
    main.go:
        ...
        func GetCategories(products []Product) (categories []string) {
            catMap := map[string]string {}
            for _, p := range products {
                if (catMap[p.Category] == "") {
                    catMap[p.Category] = p.Category
                    categories = append(categories, "<b>p.Category</b>")
                }
            }
            return
        }
        ...
    ===============================================================
    Output:
        <h1>There are 8 products in the source data.</h1>
            <h1>Category: &lt;b&gt;p.Category&lt;/b&gt;</h1>
            <h1>Category: &lt;b&gt;p.Category&lt;/b&gt;</h1>
            <h1>Category: &lt;b&gt;p.Category&lt;/b&gt;</h1>
████████████████████████████████████████████████████████████████████████
395.The Types Aliases Used to Denote Content Types
    Name        Description
    --------    -----------
    CSS         This type denotes CSS content.
    HTML        This type denotes a fragment of HTML.
    HTMLAttr    This type denotes a value that will be used as the value for an HTML attribute.
    JS          This type denotes a fragment of JavaScript code.
    JSStr       This type denotes a value that is intended to appear between quotes in a JavaScript expression.
    Srcset      This type denotes a value that can be used in the srcset attribute of an img element.
    URL         This type denotes a URL.
████████████████████████████████████████████████████████████████████████
396.Returning HTML Content in the main.go
    example:
    main.go:
        ...
        func GetCategories(products []Product) (categories []template.HTML) {
            catMap := map[string]string {}
            for _, p := range products {
                if (catMap[p.Category] == "") {
                    catMap[p.Category] = p.Category
                    categories = append(categories, "<b>p.Category</b>")
                }
            }
            return
        }
        ...
    =======================================================
    Output:
        <h1>There are 8 products in the source data.</h1>
            <h1>Category: <b>p.Category</b></h1>
            <h1>Category: <b>p.Category</b></h1>
            <h1>Category: <b>p.Category</b></h1>
████████████████████████████████████████████████████████████████████████
397.Providing Access to Standard Library Functions
    Adding a Function Mapping in the main.go
    example:
    main.go:
        package main
        import (
            "html/template"
            "os"
            "strings"
        )
        func GetCategories(products []Product) (categories []string) {
            catMap := map[string]string{}
            for _, p := range products {
                if catMap[p.Category] == "" {
                    catMap[p.Category] = p.Category
                    categories = append(categories, p.Category)
                }
            }
            return
        }
        func Exec(t *template.Template) error {
            return t.Execute(os.Stdout, Products)
        }
        func main() {
            allTemplates := template.New("allTemplates")
            allTemplates.Funcs(map[string]interface{}{
                "getCats": GetCategories,
                "lower":   strings.ToLower,
            })
            allTemplates, err := allTemplates.ParseGlob("templates/*.html")
            if err == nil {
                selectedTemplated := allTemplates.Lookup("mainTemplate")
                err = Exec(selectedTemplated)
            }
            if err != nil {
                Printfln("Error: %v %v", err.Error())
            }
        }
    =============================================
    Output:
        <h1>There are 8 products in the source data.</h1>
            <h1>Category: Watersports</h1>
            <h1>Category: Soccer</h1>
            <h1>Category: Chess</h1>
████████████████████████████████████████████████████████████████████████
398.Using a Template Function in the template.html
    example:
    template.html:
        {{ define "mainTemplate" -}}
            <h1>There are {{ len . }} products in the source data.</h1>
            {{ range getCats .  -}}
                <h1>Category: {{ lower . }}</h1>
            {{ end }}
        {{- end }}
    =============================================================
    Output:
        <h1>There are 8 products in the source data.</h1>
            <h1>Category: watersports</h1>
            <h1>Category: soccer</h1>
            <h1>Category: chess</h1>
████████████████████████████████████████████████████████████████████████
399.Defining Template Variables
    Defining and Using a Template Variable in the template.html
    example:
    template.html:
        {{ define "mainTemplate" -}}
            {{ $length := len . }}
            <h1>There are {{ $length }} products in the source data.</h1>
            {{ range getCats .  -}}
                <h1>Category: {{ lower . }}</h1>
            {{ end }}
        {{- end }}
    =============================================================
    Output:
            
        <h1>There are 8 products in the source data.</h1>
        <h1>Category: watersports</h1>
        <h1>Category: soccer</h1>
        <h1>Category: chess</h1>
████████████████████████████████████████████████████████████████████████
400.Defining and Using a Template Variable in the template.html
    example:
    template.html:
        {{ define "mainTemplate" -}}
            <h1>There are {{ len . }} products in the source data.</h1>
            {{- range getCats .  -}}
                {{ if ne ($char := slice (lower .) 0 1) "s"  }}
                    <h1>{{$char}}: {{.}}</h1>
                {{- end }}
            {{- end }}
        {{- end }}
    ==============================================
    Output:
        <h1>There are 8 products in the source data.</h1>
                    <h1>w: Watersports</h1>
                    <h1>c: Chess</h1>
████████████████████████████████████████████████████████████████████████
401.Using Template Variables in Range Actions
    Enumerating a Map in the template.html

    example:
    main.go:
        ...
        func Exec(t *template.Template) error {
            productMap := map[string]Product {}
            for _, p := range Products {
                productMap[p.Name] = p
            }
            return t.Execute(os.Stdout, &productMap)
        }
        ...
    
    template.html:
        {{ define "mainTemplate" -}}
            {{ range $key, $value := . -}}
                <h1>{{ $key }}: {{ printf "$%.2f" $value.Price }}</h1>
            {{ end }}
        {{- end }}
    Output:
        <h1>Bling-Bling King: $1200.00</h1>
            <h1>Corner Flags: $34.95</h1>
            <h1>Kayak: $279.00</h1>
            <h1>Lifejacket: $49.95</h1>
            <h1>Soccer Ball: $19.50</h1>
            <h1>Stadium: $79500.00</h1>
            <h1>Thinking Cap: $16.00</h1>
            <h1>Unsteady Chair: $75.00</h1>
████████████████████████████████████████████████████████████████████████
402.Creating Text Templates
    Loading and Executing a Text Template in the main.go

    The Contents of the template.txt
    example:
    templates/template.txt:
        {{ define "mainTemplate" -}}
            {{ range $key, $value := . -}}
                {{ $key }}: {{ printf "$%.2f" $value.Price }}
            {{ end }}
        {{- end }}
    ---------------------------------------
    main.go:
        package main
        import (
            "text/template"
            "os"
            "strings"
        )
        func GetCategories(products []Product) (categories []string) {
            catMap := map[string]string {}
            for _, p := range products {
                if (catMap[p.Category] == "") {
                    catMap[p.Category] = p.Category
                    categories = append(categories, p.Category)
                }
            }
            return
        }
        func Exec(t *template.Template) error {
            productMap := map[string]Product {}
            for _, p := range Products {
                productMap[p.Name] = p
            }
            return t.Execute(os.Stdout, &productMap)
        }
        func main() {
            allTemplates := template.New("allTemplates")
            allTemplates.Funcs(map[string]interface{} {
                "getCats": GetCategories,
                "lower": strings.ToLower,
            })
            allTemplates, err := allTemplates.ParseGlob("templates/*.txt")
            if (err == nil) {
                selectedTemplated := allTemplates.Lookup("mainTemplate")
                err = Exec(selectedTemplated)
            }
            if (err != nil) {
                Printfln("Error: %v %v", err.Error())
            }
        }
    Output:
        Bling-Bling King: $1200.00
            Corner Flags: $34.95
            Kayak: $279.00
            Lifejacket: $49.95
            Soccer Ball: $19.50
            Stadium: $79500.00
            Thinking Cap: $16.00
            Unsteady Chair: $75.00
████████████████████████████████████████████████████████████████████████
403.Creating HTTP Servers
    What are they?
    The features described in this chapter make it easy for Go applications to create HTTP servers.

    Why are they useful?
    HTTP is one of the most widely used protocols and is useful for both user-facing
    applications and web services.

    How is it used?
    The features of the net/http package are used to create a server and handle requests.

    Are there any pitfalls or limitations?
    These features are well-designed and easy to use.

    Are there any alternatives?
    The standard library includes support for other network protocols and also for
    opening and using lower-level network connections. 
    See https://pkg.go.dev/net@go1.17.1 for details of the net package and its subpackages, 
    such as net/smtp, for example, which implements the SMTP protocol.

    Problem                                 Solution
    --------                                -------------
    Create an HTTP or HTTPS server          Use the ListenAndServe or ListenAndServeTLS functions
    Inspect an HTTP request                 Use the features of the Request struct
    Produce a response                      Use the ResponseWriter interface or the
                                            convenience functions

    Handle requests to specific URLs        Use the integrated router
    Serve static content                    Use the FileServer and StripPrefix function
    Use a template to produce a response    Write the content to the ResponseWriter
    or produce a JSON response

    Handle form data                        Use the Request methods
    Set or read cookies                     Use the Cookie, Cookies, and SetCookie methods

    Preparing for This Chapter:
    1- go mod init httpserver
    2- printer.go:
        package main
        import "fmt"
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
    3- product.go:
        package main
        type Product struct {
            Name, Category string
            Price          float64
        
        var Products = []Product{
            {"Kayak", "Watersports", 279},
            {"Lifejacket", "Watersports", 49.95},
            {"Soccer Ball", "Soccer", 19.50},
            {"Corner Flags", "Soccer", 34.95},
            {"Stadium", "Soccer", 79500},
            {"Thinking Cap", "Chess", 16},
            {"Unsteady Chair", "Chess", 75},
            {"Bling-Bling King", "Chess", 1200},
        }
    4- main.go:
        package main
        func main() {
            for _, p := range Products {
                Printfln("Product: %v, Category: %v, Price: $%.2f",
                    p.Name, p.Category, p.Price)
            }
        }
    =================================
    Output:
        Product: Kayak, Category: Watersports, Price: $279.00
        Product: Lifejacket, Category: Watersports, Price: $49.95
        Product: Soccer Ball, Category: Soccer, Price: $19.50
        Product: Corner Flags, Category: Soccer, Price: $34.95
        Product: Stadium, Category: Soccer, Price: $79500.00
        Product: Thinking Cap, Category: Chess, Price: $16.00
        Product: Unsteady Chair, Category: Chess, Price: $75.00
        Product: Bling-Bling King, Category: Chess, Price: $1200.00
████████████████████████████████████████████████████████████████████████
404.Creating a Simple HTTP Server
    example:
    main.go:
        package main
        import (
            "net/http"
            "io"
        )
        type StringHandler struct {
            message string
        }
        func (sh StringHandler) ServeHTTP(writer http.ResponseWriter,
                request *http.Request) {
            io.WriteString(writer, sh.message)
        }
        func main() {
            err := http.ListenAndServe(":5000", StringHandler{ message: "Hello, World"})
            if (err != nil) {
                Printfln("Error: %v", err.Error())
            }
        }
        =================================
        Output: go run .
        in Web Browser, search localhost:5000/
            Hello, World

████████████████████████████████████████████████████████████████████████
405.Creating the HTTP Listener and Handler
    The net/http Convenience Functions

    The ListenAndServe function starts listening for HTTP requests on a specified network address. 
    The ListenAndServeTLS function does the same for HTTP requests.
    The addresses accepted by the functions can be used to restrict the HTTP server so that it
    only accepts requests on a specific interface or to listen for requests on any interface.

    Name                                            Description
    ---------                                       ----------------------------------------
    ListenAndServe(addr, handler)                   This function starts listening for HTTP requests on a specified
                                                    address and passes requests onto the specified handler.
    ListenAndServeTLS(addr, cert, key, handler)     This function starts listening for HTTPS requests. The arguments are the address

████████████████████████████████████████████████████████████████████████
406.The Method Defined by the Handler Interface
    Name                            Description
    ServeHTTP(writer, request)      This method is invoked to process a HTTP request. The request is described by a
                                    Request value, and the response is written using a ResponseWriter, both of which
                                    are received as parameters.
████████████████████████████████████████████████████████████████████████
407.Inspecting the Request
    The Basic Fields Defined by the Request Struct
    Name        Description
    -------     ------------------------
    Method      This field provides the HTTP method (GET, POST, etc.) as a string. The net/http package defines
                constants for the HTTP methods, such as MethodGet and MethodPost.
    URL         This field returns the requested URL, expressed as a URL value.
    Proto       This field returns a string that indicates the version of HTTP used for the request.
    Host        This field returns a string containing the requested hos.
    Header      This field returns a Header value, which is an alias to map[string][]string and contains the
                request headers. The map keys are the names of the headers, and the values are string slices
                containing the header values.
    Trailer     This field returns a map[string]string that contains any additional headers that are included in
                the request after the body.
    Body        This filed returns a ReadCloser, which is an interface that combines the Read method of the
                Reader interface with the Close method of the Closer interface
████████████████████████████████████████████████████████████████████████
408.Writing Request Fields in the main.go
    example:
    main.go:
        package main
        import (
            "io"
            "net/http"
        )
        type StringHandler struct {
            message string
        }
        func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
            Printfln("Method: %v", request.Method)
            Printfln("URL: %v", request.URL)
            Printfln("HTTP Version: %v", request.Proto)
            Printfln("Host: %v", request.Host)
            for name, val := range request.Header {
                Printfln("Header: %v, Value: %v", name, val)
            }
            Printfln("---")
            io.WriteString(writer, sh.message)
        }
        func main() {
            err := http.ListenAndServe(":5000", StringHandler{message: "Hello, World"})
            if err != nil {
                Printfln("Error: %v", err.Error())
            }
        }
    ===================================================================================
    Output: Compile and execute the project and request http://localhost:5000
        Method: GET
        URL: /
        HTTP Version: HTTP/1.1
        Host: localhost:5000
        Header: Accept, Value: [text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8]
        Header: Accept-Language, Value: [en-US,en;q=0.5]
        Header: Accept-Encoding, Value: [gzip, deflate, br]
        Header: Sec-Fetch-Mode, Value: [navigate]
        Header: Sec-Fetch-Site, Value: [none]
        Header: User-Agent, Value: [Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/118.0]
        Header: Connection, Value: [keep-alive]
        Header: Upgrade-Insecure-Requests, Value: [1]
        Header: Sec-Fetch-Dest, Value: [document]
        Header: Sec-Fetch-User, Value: [?1]
        ---
        Method: GET
        URL: /favicon.ico
        HTTP Version: HTTP/1.1
        Host: localhost:5000
        Header: Accept-Encoding, Value: [gzip, deflate, br]
        Header: Connection, Value: [keep-alive]
        Header: Sec-Fetch-Dest, Value: [image]
        Header: Sec-Fetch-Mode, Value: [no-cors]
        Header: Sec-Fetch-Site, Value: [same-origin]
        Header: User-Agent, Value: [Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/118.0]
        Header: Accept, Value: [image/avif,image/webp,*/*]
        Header: Accept-Language, Value: [en-US,en;q=0.5]
        Header: Referer, Value: [http://localhost:5000/]
        ---
████████████████████████████████████████████████████████████████████████
409.Save All Logs in a txt file:
    Server.go:
        package server
        import (
            "io"
            "log"
            "net/http"
            "os"
        )
        type StringHandler struct {
            Message string
        }
        func MyServer() {
            logFile, err := os.OpenFile("Server/LogFile/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
            if err != nil {
                log.Fatal(err)
            }
            defer logFile.Close()
            log.SetOutput(logFile)
        
            err = http.ListenAndServe(":5000", StringHandler{Message: "Hello, World"})
            if err != nil {
                log.Printf("Error: %v", err.Error())
            }
        }
        func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
            log.Printf("Method: %v", request.Method)
            log.Printf("URL: %v", request.URL)
            log.Printf("HTTP Version: %v", request.Proto)
            log.Printf("Host: %v", request.Host)
            for name, val := range request.Header {
                log.Printf("Header: %v, Value: %v", name, val)
            }
            io.WriteString(writer, sh.Message)
            log.Printf("============================================◉🧭🧭🧭🧭🧭🧭🧭◉==========================================")
        }
    Output: Create a TXT file in Server/LogFile/logs.txt and save appendation in to it.
████████████████████████████████████████████████████████████████████████
410.Filtering Requests and Generating Responses
    Useful Fields and Methods Defined by the URL Struct
    The ResponseWriter interface defines the methods that are available when creating a response.
    Name        Description
    -------     -----------------------------
    Scheme      This field returns the scheme component of the URL.
    Host        This field returns the host component of the URL, which may include the port.
    RawQuery    This field returns the query string from the URL. Use the Query method to process the query
                string into a map.
    Path        This field returns the path component of the URL.
    Fragment    This field returns the fragment component of the URL, without the # character.
    Hostname()  This method returns the hostname component of the URL as a string.
    Port()T     his method returns the port component of the URL as a string.
    Query()     This method returns a map[string][]string (a map with string keys and string slice
                values), containing the query string fields.
    User()      This method returns the user information associated with the request.
    String()    This method returns a string representation of the URL.
████████████████████████████████████████████████████████████████████████
411.The ResponseWriter Methods
    Name                Description
    Header()            This method returns a Header, which is an alias to map[string][]string, that can be
                        used to set the response headers.
    WriteHeader(code)   This method sets the status code for the response, specified as an int. The net/http
                        package defines constants for most status codes.
    Write(data)         This method writes data to the response body and implements the Writer interface.
████████████████████████████████████████████████████████████████████████
412.Producing Difference Responses in the main.go
    example:
    main.go:
        package main
        import (
            "net/http"
            "io"
        )
        type StringHandler struct {
            message string
        }
        func (sh StringHandler) ServeHTTP(writer http.ResponseWriter,
                request *http.Request) {
            if (request.URL.Path == "/favicon.ico") {
                Printfln("Request for icon detected - returning 404")
                writer.WriteHeader(http.StatusNotFound)
                return
            }
            Printfln("Request for %v", request.URL.Path)
            io.WriteString(writer, sh.message)
        }
        func main() {
            err := http.ListenAndServe(":5000", StringHandler{ message: "Hello, World"})
            if (err != nil) {
                Printfln("Error: %v", err.Error())
            }
        }
    Output: in Terminal
        Request for /
        Request for icon detected - returning 404
████████████████████████████████████████████████████████████████████████
413.Get Logs and Handle request if URL not Exist:
        example:
        main.go:
            package main
            import (
                "io"
                "log"
                "net/http"
                "os"
            )
            type StringHandler struct {
                Message string
            }
            func main() {
                logFile, err := os.OpenFile("LogFile/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
                if err != nil {
                    log.Fatal(err)
                }
                defer logFile.Close()
                log.SetOutput(logFile)
            
                err = http.ListenAndServe(":5000", StringHandler{Message: "Hello, World"})
                if err != nil {
                    log.Printf("Error: %v", err.Error())
                }
            }
            func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
                if (request.URL.Path != "/") {
                    Printfln("Request for icon detected - returning 404")
                    writer.WriteHeader(http.StatusNotFound)
                    return
                }
                Printfln("Request for %v", request.URL.Path)
                io.WriteString(writer, sh.Message)
            
             
            
                log.Printf("Method: %v", request.Method)
                log.Printf("URL: %v", request.URL)
                log.Printf("HTTP Version: %v", request.Proto)
                log.Printf("Host: %v", request.Host)
                for name, val := range request.Header {
                    log.Printf("Header: %v, Value: %v", name, val)
                }
                io.WriteString(writer, sh.Message)
                log.Printf("============================================◉🧭🧭🧭🧭🧭🧭🧭◉==========================================")
            }
        Output: so we have to logs, one of that is logs about user request,
        another that is about path URL in Terminal
████████████████████████████████████████████████████████████████████████
414.Using the Response Convenience Functions
    Name                                    Description
    -----------------------                 ----------------------------
    Error(writer, message, code)            This function sets the header to the specified code, sets the Content-Type header
                                            to text/plain, and writes the error message to the response. The X-Content-
                                            Type-Options header is also set to stop browsers from interpreting the response as
                                            anything other than text.
    NotFound(writer, request)               This function calls Error and specifies a 404 error code.
    Redirect(writer, request, url, code)    This function sends a redirection response to the specified URL and with the
                                            specified status code.
    ServeFile(writer, request, fileName)    This function sends a response containing the contents of the specified file. The
                                            Content-Type header is set based on the file name but can be overridden by
                                            explicitly setting the header before calling the function. See the “Creating a Static
                                            HTTP Server” section for an example that serves files.
████████████████████████████████████████████████████████████████████████
415.Convenience Functions in the main.go
    example:
    main.go:
        package main
        import (
            "io"
            "log"
            "net/http"
            "os"
        )
        type StringHandler struct {
            Message string
        }
        func main() {
            logFile, err := os.OpenFile("LogFile/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
            if err != nil {
                log.Fatal(err)
            }
            defer logFile.Close()
            log.SetOutput(logFile)
            err = http.ListenAndServe(":5000", StringHandler{Message: "Hello, World"})
            if err != nil {
                log.Printf("Error: %v", err.Error())
            }
        }
        func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
            Printfln("Request for %v", request.URL.Path)
            switch request.URL.Path {
            case "/favicon.ico":
                http.NotFound(writer, request)
            case "/message":
                io.WriteString(writer, sh.Message)
            default:
                http.Redirect(writer, request, "/message", http.StatusTemporaryRedirect)
            }
            log.Printf("Method: %v", request.Method)
            log.Printf("URL: %v", request.URL)
            log.Printf("HTTP Version: %v", request.Proto)
            log.Printf("Host: %v", request.Host)
            for name, val := range request.Header {
                log.Printf("Header: %v, Value: %v", name, val)
            }
            io.WriteString(writer, sh.Message)
            log.Printf("============================================◉🧭🧭🧭🧭🧭🧭🧭◉==========================================")
        }
    Output: Will write in Terminal and File for feture analyzing.
    
████████████████████████████████████████████████████████████████████████
416.Using the Convenience Functions in the main.go
    example:
    main.go:
        package main
        import (
            "io"
            "net/http"
        )
        type StringHandler struct {
            message string
        }
        func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
            Printfln("Request for %v", request.URL.Path)
            switch request.URL.Path {
            case "/favicon.ico":
                http.NotFound(writer, request)
            case "/message":
                io.WriteString(writer, sh.message)
            default:
                http.Redirect(writer, request, "/message", http.StatusTemporaryRedirect)
            }
        }
        func main() {
            err := http.ListenAndServe(":5000", StringHandler{message: "Hello, World"})
            if err != nil {
                Printfln("Error: %v", err.Error())
            }
        }
    Output:
        every wrong links redirect to specific link.
        uses a switch statement to decide how to respond to a request.
████████████████████████████████████████████████████████████████████████
417.Using the Convenience Routing Handler
    The process of inspecting the URL and selecting a response can produce complex code that is difficult to
    read and maintain. 
    To simplify the process, the net/http package provides a Handler implementation that
    allows matching the URL to be separated from producing a request.

    example:
    main.go:
        package main
        import (
            "io"
            "net/http"
        )
        type StringHandler struct {
            message string
        }
        func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
            Printfln("Request for %v", request.URL.Path)
            io.WriteString(writer, sh.message)
        }
        func main() {
            http.Handle("/message", StringHandler{"Hello, World"})
            http.Handle("/favicon.ico", http.NotFoundHandler())
            http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))
            err := http.ListenAndServe(":5000", nil)
            if err != nil {
                Printfln("Error: %v", err.Error())
            }
        }
    Output:
        don't show wrong search.
████████████████████████████████████████████████████████████████████████
418.The net/http Functions for Creating Routing Rules
    Name                                Description
    -----------------                   ---------------------------------
    Handle(pattern, handler)            This function creates a rule that invokes the specified ServeHTTP method of the
                                        specified Hander for requests that match the pattern.
    HandleFunc(pattern, handlerFunc)    This function creates a rule that invokes the specified function for requests that match
                                        the pattern. The function is invoked with ResponseWriter and Request arguments.
████████████████████████████████████████████████████████████████████████
419.he net/http Functions for Creating Request Handlers
    Name                                        Description
    FileServer(root)                            This function creates a Handler that produces responses using the ServeFile
                                                function. See the “Creating a Static HTTP Server” section for an example that
                                                serves files.
    NotFoundHandler()                           This function creates a Handler that produces responses using the NotFound function.
    RedirectHandler(url, code)                  This function creates a Handler that produces responses using the Redirect function.
    StripPrefix(prefix, handler)                This function creates a Handler that removes the specified prefix from the
                                                request URL and passes on the request to the specified Handler. See the
                                                “Creating a Static HTTP Server” section for details.
    TimeoutHandler(handler, duration, message)  This function passes on the request to the specified Handler but generates an error
                                                response if the response hasn't been produced within the specified duration.

████████████████████████████████████████████████████████████████████████
420.Supporting HTTPS Requests
    The net/http package provides integrated support for HTTPS. 
    To prepare for HTTPS, you will need to add
    two files to the httpserver folder: 
        a certificate file and a private key file.

████████████████████████████████████████████████████████████████████████
421.Getting Certificates for HTTPS
    A good way to get started with HTTPS is with a self-signed certificate, 
    which can be used for development and testing. 
    If you don't already have a self-signed certificate, 
    then you can create one online using sites such as 
    
    https://getacert.com 
    or 
    https://www.selfsignedcertificate.com
    
    both of which will let you create a self-signed certificate easily 
    and without charge.

    Two files are required to use HTTPS, regardless of whether your certificate 
    is self-signed or not. The first is the certificate file, 
    which usually has a cer or cert file extension. 
    The second is the private key file, which usually has a key file extension.

    after ready to deploy:
        https://letsencrypt.org

    The ListenAndServeTLS function is used to enable HTTPS, 
    where the additional arguments specify the
    certificate and private key files, 
    which are named certificate.cer and certificate.key in my project
████████████████████████████████████████████████████████████████████████
422.Enabling HTTPS in the main.go
    example:
    main.go:
        package main
        import (
            "io"
            "net/http"
        )
        type StringHandler struct {
            message string
        }
        func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
            Printfln("Request for %v", request.URL.Path)
            io.WriteString(writer, sh.message)
        }
        func main() {
            http.Handle("/message", StringHandler{"Hello, World"})
            http.Handle("/favicon.ico", http.NotFoundHandler())
            http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))
        
            go func() {
                err := http.ListenAndServeTLS(":5500", "certificate.cer",
                    "certificate.key", nil)
                if err != nil {
                    Printfln("HTTPS Error: %v", err.Error())
                }
            }()
        
            err := http.ListenAndServe(":5000", nil)
            if err != nil {
                Printfln("Error: %v", err.Error())
            }
        }
    Output:
        HTTPS Error: open certificate.cer: no such file or directory
        Request for /message
████████████████████████████████████████████████████████████████████████
423.Redirecting HTTP Requests to HTTPS
    A common requirement when creating web servers is to redirect HTTP requests to the HTTPS port. 
    This can be done by creating a custom handler
    
    example:
    Redirecting to HTTPS in the main.go:
        package main
        import (
            "net/http"
            "io"
            "strings"
        )
        type StringHandler struct {
            message string
        }
        func (sh StringHandler) ServeHTTP(writer http.ResponseWriter,
                request *http.Request) {
            Printfln("Request for %v", request.URL.Path)
            io.WriteString(writer, sh.message)
        }
        func HTTPSRedirect(writer http.ResponseWriter,
                request *http.Request) {
            host := strings.Split(request.Host, ":")[0]
            target := "https://" + host + ":5500" + request.URL.Path
            if len(request.URL.RawQuery) > 0 {
                target += "?" + request.URL.RawQuery
            }
            http.Redirect(writer, request, target, http.StatusTemporaryRedirect)
        }
        func main() {
            http.Handle("/message", StringHandler{ "Hello, World"})
            http.Handle("/favicon.ico", http.NotFoundHandler())
            http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))
            go func () {
                err := http.ListenAndServeTLS(":5520", "certificate.cer",
                    "certificate.key", nil)
                if (err != nil) {
                    Printfln("HTTPS Error: %v", err.Error())
                }
            }()
            err := http.ListenAndServe(":5000", http.HandlerFunc(HTTPSRedirect))
            if (err != nil) {
                Printfln("Error: %v", err.Error())
            }
        }
    Output:
        Not work for my Browser but you have to try, maybe work it for you.
████████████████████████████████████████████████████████████████████████
424.Creating a Static HTTP Server
    The net/http package includes built-in support for responding to requests with the contents of files. 
    To prepare for the static HTTP server, 
    create the httpserver/static folder and add to it a file named index.html

    example:
    static/index.html:
        <!DOCTYPE html>
        <html>
        <head>
            <title>Pro Go</title>
            <meta name="viewport" content="width=device-width" />
            <link href="bootstrap.min.css" rel="stylesheet" />
        </head>
        <body>
            <div class="m-1 p-2 bg-primary text-white h2">
                Hello, World
            </div>
        </body>
        </html>
    
    store.html:
        <!DOCTYPE html>
        <html>
        <head>
            <title>Pro Go</title>
            <meta name="viewport" content="width=device-width" />
            <link href="bootstrap.min.css" rel="stylesheet" />
        </head>
        <body>
            <div class="m-1 p-2 bg-primary text-white h2 text-center">
                Products
            </div>
            <table class="table table-sm table-bordered table-striped">
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>Category</th>
                        <th>Price</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td>Kayak</td>
                        <td>Watersports</td>
                        <td>$279.00</td>
                    </tr>
                    <tr>
                        <td>Lifejacket</td>
                        <td>Watersports</td>
                        <td>$49.95</td>
                    </tr>
                </tbody>
            </table>
        </body>
        </html>

        The HTML files depend on the Bootstrap CSS package to style the HTML content. Run the command
        shown in here in the httpserver folder to download the Bootstrap CSS file into the static folder.
        (You may have to install the curl command.):

            curl https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css --output static/bootstrap.min.css

        Output:
            ootstrap.min.css
            % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                            Dload  Upload   Total   Spent    Left  Speed
            100  152k    0  152k    0     0   163k      0 --:--:-- --:--:-- --:--:--  163k
████████████████████████████████████████████████████████████████████████
425.Creating the Static File Route
    example:
    Defining a Route in the main.go:
        package main
        import (
            "net/http"
            "io"
        )
        type StringHandler struct {
            message string
        }
        func (sh StringHandler) ServeHTTP(writer http.ResponseWriter,
                request *http.Request) {
            Printfln("Request for %v", request.URL.Path)
            io.WriteString(writer, sh.message)
        }
        func main() {
            http.Handle("/message", StringHandler{ "Hello, World"})
            http.Handle("/favicon.ico", http.NotFoundHandler())
            http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))
            fsHandler := http.FileServer(http.Dir("./static"))
            http.Handle("/files/", http.StripPrefix("/files", fsHandler))
            err := http.ListenAndServe(":5000", nil)
            if (err != nil) {
                Printfln("Error: %v", err.Error())
            }
        }
    Output:
        redirect from => https://localhost:5500/files/store.html
                   to => static/store.html
                   
████████████████████████████████████████████████████████████████████████
426.Using Templates to Generate Responses
    example:
    templates/products.html:
        <!DOCTYPE html>
        <html>
        <head>
            <meta name="viewport" content="width=device-width" />
            <title>Pro Go</title>
            <link rel="stylesheet" href="/files/bootstrap.min.css">
        </head>
        <body>
            <h3 class="bg-primary text-white text-center p-2 m-2">Products</h3>
            <div class="p-2">
                <table class="table table-sm table-striped table-bordered">
                    <thead>
                        <tr>
                            <th>Index</th>
                            <th>Name</th>
                            <th>Category</th>
                            <th class="text-end">Price</th>
                        </tr>
                    </thead>
                    <tbody>
                        {{ range $index, $product := .Data }}
                        <tr>
                            <td>{{ $index }}</td>
                            <td>{{ $product.Name }}</td>
                            <td>{{ $product.Category }}</td>
                            <td class="text-end">
                                {{ printf "$%.2f" $product.Price }}
                            </td>
                        </tr>
                        {{ end }}
                    </tbody>
                </table>
            </div>
        </body>
        </html>
    
    dynamic.go:
        package main
        import (
            "html/template"
            "net/http"
            "strconv"
        )
        type Context struct {
            Request *http.Request
            Data []Product
        }
        var htmlTemplates *template.Template
        func HandleTemplateRequest(writer http.ResponseWriter, request *http.Request) {
            path := request.URL.Path
            if (path == "") {
                path = "products.html"
            }
            t := htmlTemplates.Lookup(path)
            if (t == nil) {
                http.NotFound(writer, request)
            } else {
                err := t.Execute(writer, Context{  request, Products})
                if (err != nil) {
                    http.Error(writer, err.Error(), http.StatusInternalServerError)
                }
            }
        }
        func init() {
            var err error
            htmlTemplates = template.New("all")
            htmlTemplates.Funcs(map[string]interface{} {
                "intVal": strconv.Atoi,
            })
            htmlTemplates, err = htmlTemplates.ParseGlob("templates/*.html")
            if (err == nil) {
                http.Handle("/templates/", http.StripPrefix("/templates/",
                    http.HandlerFunc(HandleTemplateRequest)))
            } else {
                panic(err)
            }
        }

    Output:
        The initialization function loads all the templates with the html extension in the templates folder and
        sets up a route so that requests that start with /templates/ are processed by the HandleTemplateRequest
        function. This function looks up the template, falling back to the products.html file if no file path is
        specified, executes the template, and writes the response.


████████████████████████████████████████████████████████████████████████
427.Understanding Content Type Sniffing
    which implements the MIME Sniffing algorithm defined by: 
        
        https://mimesniff.spec.whatwg.org 
    
    The sniffing process can't detect every content type, 
    but it does well with standard web types, such as HTML, CSS, and JavaScript.
    The DetectContentType function returns a MIME type, 
    which is used as the value for the Content-Type header.

████████████████████████████████████████████████████████████████████████
428.Responding with JSON Data
    JSON responses are widely used in web services, 
    which provide access to an application's data for clients
    that don't want to receive HTML, such as Angular or React JavaScript clients.

    example:
    json.go:
        package main
        import (
            "net/http"
            "encoding/json"
        )
        func HandleJsonRequest(writer http.ResponseWriter, request *http.Request) {
            writer.Header().Set("Content-Type", "application/json")
            json.NewEncoder(writer).Encode(Products)
        }
        func init() {
            http.HandleFunc("/json", HandleJsonRequest)
        }
    =========================================================================
    main.go:
        package main
        import (
            "net/http"
            "io"
        )
        type StringHandler struct {
            message string
        }
        func (sh StringHandler) ServeHTTP(writer http.ResponseWriter,
                request *http.Request) {
            Printfln("Request for %v", request.URL.Path)
            io.WriteString(writer, sh.message)
        }
        func main() {
            http.Handle("/message", StringHandler{ "Hello, World"})
            http.Handle("/favicon.ico", http.NotFoundHandler())
            http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))

            fsHandler := http.FileServer(http.Dir("./static"))
            http.Handle("/files/", http.StripPrefix("/files", fsHandler))

            err := http.ListenAndServe(":5000", nil)
            if (err != nil) {
                Printfln("Error: %v", err.Error())
            }
        }
    ===================================================================
    Output:
        The initialization function creates a route, 
        which means that requests for /json will be processed by the
        HandleJsonRequest function.
████████████████████████████████████████████████████████████████████████
429.Handling Form Data
    The net/http package provides support for easily receiving and processing form data.

    This template makes use of template variables, expressions, 
    and functions to get the query string from
    the request and select the first index value, 
    which is converted to an int and used to retrieve a Product value
    from the data provided to the template.

    example:
    edit.html:
        <!DOCTYPE html>
        <html>
        <head>
            <meta name="viewport" content="width=device-width" />
            <title>Pro Go</title>
            <link rel="stylesheet" href="/files/bootstrap.min.css">
        </head>
        <body>
            {{ $index := intVal (index (index .Request.URL.Query "index") 0) }}
            {{ if lt $index (len .Data)}}
            {{ with index .Data $index}}
            <h3 class="bg-primary text-white text-center p-2 m-2">Product</h3>
            <form method="POST" action="/forms/edit" class="m-2">
                <div class="form-group">
                    <label>Index</label>
                    <input name="index" value="{{$index}}" class="form-control" disabled />
                    <input name="index" value="{{$index}}" type="hidden" />
                </div>
                <div class="form-group">
                    <label>Name</label>
                    <input name="name" value="{{.Name}}" class="form-control" />
                </div>
                <div class="form-group">
                    <label>Category</label>
                    <input name="category" value="{{.Category}}" class="form-control" />
                </div>
                <div class="form-group">
                    <label>Price</label>
                    <input name="price" value="{{.Price}}" class="form-control" />
                </div>
                <div class="mt-2">
                    <button type="submit" class="btn btn-primary">Save</button>
                    <a href="/templates/" class="btn btn-secondary">Cancel</a>
                </div>
            </form>
            {{ end }}
            {{ else }}
            <h3 class="bg-danger text-white text-center p-2">
                No Product At Specified Index
            </h3>
            {{end }}
        </body>
        </html>
████████████████████████████████████████████████████████████████████████
430.Reading Form Data from Requests
    The Request Form Data Fields and Methods


    Name                Description
    --------            ------------------------------
    Form                This field returns a map[string][]string containing the parsed form data and the
                        query string parameters. The ParseForm method must be called before this field is read.
    PostForm            This field is similar to Form but excludes the query string parameters so that only
                        data from the request body is contained in the map. The ParseForm method must
                        be called before this field is read.
    MultipartForm       This field returns a multipart form represented using the Form struct defined in the
                        mime/multipart package. The ParseMultipartForm method must be called before
                        this field is read.
    FormValue(key)      This method returns the first value for the specified form key and returns the
                        empty string if there is no value. The source of data for this method is the Form
                        field, and calling the FormValue method automatically calls ParseForm or
                        ParseMultipartForm to parse the form.
    PostFormValue(key)  This method returns the first value for the specified form key and returns the
                        empty string if there is no value. The source of data for this method is the PostForm
                        field, and calling the PostFormValue method automatically calls ParseForm or
                        ParseMultipartForm to parse the form.
    FormFile(key)       This method provides access to the first file with the specified key in the form.
                        The results are a File and FileHeader, both of which are defined in the mime/
                        multipart package, and an error. Calling this function causes the ParseForm or
                        ParseMultipartForm functions to be invoked to parse the form.
    ParseForm()         This method parses a form and populates the Form and PostForm fields. The result
                        is an error that describes any parsing problems.
    ParseMultipart      This method parses a MIME multipart form and populates the MultipartForm field.
    Form(max)           The argument specifies the maximum number of bytes to allocate to the form data,
                        and the result is an error that describes any problems processing the form.



    The init function sets up a new route so that the ProcessFormData function handles requests whose
    path is /forms/edit. Within the ProcessFormData function, the request method is checked, and the form
    data in the request is used to create a Product struct and replace the existing data value. In a real project,
    validating the data submitted in the form is essential, but for this chapter I trust that the form contains
    valid data.
    Processing form data:


    https://localhost:5500/templates/edit.html?index=2

    example:
    The Contents of the forms.go File in the httpserver Folder:
        package main
        import (
            "net/http"
            "strconv"
        )
        func ProcessFormData(writer http.ResponseWriter, request *http.Request) {
            if (request.Method == http.MethodPost) {
                index, _ := strconv.Atoi(request.PostFormValue("index"))
                p := Product {}
                p.Name = request.PostFormValue("name")
                p.Category = request.PostFormValue("category")
                p.Price, _ = strconv.ParseFloat(request.PostFormValue("price"), 64)
                Products[index] = p
            }
            http.Redirect(writer, request, "/templates", http.StatusTemporaryRedirect)
        }
        func init() {
            http.HandleFunc("/forms/edit", ProcessFormData)
        }
    
    Output:
        without view of upload you can add data to templates URL.
        
████████████████████████████████████████████████████████████████████████
431.Reading Multipart Forms
    example:
    upload.html:
        <!DOCTYPE html>
        <html>
        <head>
            <title>Pro Go</title>
            <meta name="viewport" content="width=device-width" />
            <link href="bootstrap.min.css" rel="stylesheet" />
        </head>
        <body>
            <div class="m-1 p-2 bg-primary text-white h2 text-center">
                Upload File
            </div>
            <form method="POST" action="/forms/upload" class="p-2"
                    enctype="multipart/form-data">
                <div class="form-group">
                    <label class="form-label">Name</label>
                    <input class="form-control" type="text" name="name">
                </div>
                <div class="form-group">
                    <label class="form-label">City</label>
                    <input class="form-control" type="text" name="city">
                </div>
                <div class="form-group">
                    <label class="form-label">Choose Files</label>
                    <input class="form-control" type="file" name="files" multiple>
                </div>
                <button type="submit" class="btn btn-primary mt-2">Upload</button>
            </form>
        </body>
        </html>
    ====================================================================================
    upload.go:
        package main
        import (
            "fmt"
            "io"
            "net/http"
        )
        func HandleMultipartForm(writer http.ResponseWriter, request *http.Request) {
            fmt.Fprintf(writer, "Name: %v, City: %v\n", request.FormValue("name"),
                request.FormValue("city"))
            fmt.Fprintln(writer, "------")
            file, header, err := request.FormFile("files")
            if err == nil {
                defer file.Close()
                fmt.Fprintf(writer, "Name: %v, Size: %v\n", header.Filename, header.Size)
                for k, v := range header.Header {
                    fmt.Fprintf(writer, "Key: %v, Value: %v\n", k, v)
                }
                fmt.Fprintln(writer, "------")
                io.Copy(writer, file)
            } else {
                http.Error(writer, err.Error(), http.StatusInternalServerError)
            }
        }
        func init() {
            http.HandleFunc("/forms/upload", HandleMultipartForm)
        }
    Output: Search in Browser: http://localhost:5000/files/upload.html
        You Can Upload

████████████████████████████████████████████████████████████████████████
432.The FileHeader Fields and Method
    Name        Description
    -------     ------------------------------
    Name        This field returns a string containing the name of the file.
    Size        This field returns an int64 containing the size of the file.
    Header      This field returns a map[string][]string, which contains the headers for the MIME part that
                contains the file.
    Open()      This method returns a File that can be used to read the content associated with the header, as
                demonstrated in the next section.
████████████████████████████████████████████████████████████████████████
433.Reading and Setting Cookies
    The net/http Function for Setting Cookies
    Name                            Description
    ---------------                 ------------------------------------
    SetCookie(writer, cookie)       This function adds a Set-Cookie header to the specified ResponseWriter. The
                                    cookie is described using a pointer to a Cookie struct, which is described next.



    Cookies can be complex, and care must be taken to configure them correctly. 
    The detail of how cookies work is beyond the scope of this book, 
    but there is a good description available at 
        
        https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies 

    and a detailed breakdown of the cookie
    fields at: 

        https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie.

████████████████████████████████████████████████████████████████████████
434.The Fields Defined by the Cookie Struct
    Name        Description
    -------     --------------------------------
    Name        This field represents the name of the cookie, expressed as a string.
    Value       This field represents the cookie value, expressed as a string.
    Path        This optional field specifies the cookie path.
    Domain      This optional field specifies the host/domain to which the cookie will be set.
    Expires     This field specifies the cookie expiry, expressed as a time.Time value.
    MaxAge      This field specifies the number of seconds until the cookie expires, expressed as an int.
    Secure      When this bool field is true, the client will only send the cookie over HTTPS connections.
    HttpOnly    When this bool field is true, the client will prevent JavaScript code from accessing the cookie.
    SameSite    This field specifies the cross-origin policy for the cookie using the SameSite constants, which
                defines SameSiteDefaultMode, SameSiteLaxMode, SameSiteStrictMode, and SameSiteNoneMode.
████████████████████████████████████████████████████████████████████████
435.The Request Methods for Cookies
    Name            Description
    -----------     ----------------
    Cookie(name)    This method returns a pointer to the Cookie value with the specified name and an error
                    that indicates when there is no matching cookie.
    Cookies()       This method returns a slice of Cookie pointers.


    example:
    cookies.go:
        package main
        import (
            "net/http"
            "fmt"
            "strconv"
        )
        func GetAndSetCookie(writer http.ResponseWriter, request *http.Request) {
            counterVal := 1
            counterCookie, err := request.Cookie("counter")
            if (err == nil) {
                counterVal, _ = strconv.Atoi(counterCookie.Value)
                counterVal++
            }
        http.SetCookie(writer, &http.Cookie{
                Name: "counter", Value: strconv.Itoa(counterVal),
            })
            if (len(request.Cookies()) > 0) {
                for _, c := range request.Cookies() {
                    fmt.Fprintf(writer, "Cookie Name: %v, Value: %v", c.Name, c.Value)
                }
            } else {
                fmt.Fprintln(writer, "Request contains no cookies")
            }
        }
        func init() {
            http.HandleFunc("/cookies", GetAndSetCookie)
        }
    ======================================================================================
    Compile and execute the project and use a browser to request http://localhost:5000/cookies
    Output:
        search-1:
            Request contains no cookies
        search-2:
            Cookie Name: counter, Value: 1
        search-3:
            Cookie Name: counter, Value: 2
        ...
████████████████████████████████████████████████████████████████████████
436.Creating HTTP Clients
    Putting HTTP Clients in Context

    What are they?
        HTTP requests are used to retrieve data from HTTP servers
    
    Why are they useful?
        HTTP is one of the most widely used protocols and is commonly used to provide
        access to content that can be presented to the user as well as data that is consumed
        programmatically.

    How is it used?
        The features of the net/http package are used to create and send requests and
        process responses.
    
    Are there any pitfalls or limitations?
        These features are well-designed and easy to use, although some features require a
        specific sequence to use.
    
    Are there any alternatives?
        The standard library includes support for other network protocols and also for
        opening and using lower-level network connections. 
        See the 
        
        https://pkg.go.dev/net@go1.17.1 

        https://pkg.go.dev/net
        
        for details of the net package and its subpackages, such as net/smtp,
        for example, which implements the SMTP protocol.


    Chapter Summary:
    Problem                                     Solution
    --------                                    ------------
    Send HTTP requests                          Use the convenience methods for specific HTTP methods
    Configure HTTP requests                     Use the fields and methods defined by the Client struct
    Create a preconfigured request              Use the NewRequest convenience functions
    Use cookies in a request                    Use a cookie jar
    Configure how redirections are processed    Use the CheckRedirect field to register a function that is
                                                invoked to deal with a redirection
    Send multipart forms                        Use the mime/multipart package


    Preparing for This Chapter
    1- go mod init httpclient
    2- printer.go
        package main
        import "fmt"
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template + "\n", values...)
        }
    3- httpclient folder -> file named product.go
        package main
        type Product struct {
            Name, Category string
            Price float64
        }
        var Products = []Product {
            { "Kayak", "Watersports", 279 },
            { "Lifejacket", "Watersports", 49.95 },
            { "Soccer Ball", "Soccer", 19.50 },
            { "Corner Flags", "Soccer", 34.95 },
            { "Stadium", "Soccer", 79500 },
            { "Thinking Cap", "Chess", 16 },
            { "Unsteady Chair", "Chess", 75 },
            { "Bling-Bling King", "Chess", 1200 },
        }
    4- The Contents of the index.html File in the httpclient Folder
        <!DOCTYPE html>
        <html>
        <head>
            <title>Pro Go</title>
            <meta name="viewport" content="width=device-width" />
        </head>
        <body>
            <h1>Hello, World</div>
        </body>
        </html>    
    5- The Contents of the server.go File in the httpclient Folder
        package main
        import (
            "encoding/json"
            "fmt"
            "io"
            "net/http"
            "os"
        )
        func init() {
            http.HandleFunc("/html",
                func(writer http.ResponseWriter, request *http.Request) {
                    http.ServeFile(writer, request, "./index.html")
                })
            http.HandleFunc("/json",
                func(writer http.ResponseWriter, request *http.Request) {
                    writer.Header().Set("Content-Type", "application/json")
                    json.NewEncoder(writer).Encode(Products)
                })
            http.HandleFunc("/echo",
                func(writer http.ResponseWriter, request *http.Request) {
                    writer.Header().Set("Content-Type", "text/plain")
                    fmt.Fprintf(writer, "Method: %v\n", request.Method)
                    for header, vals := range request.Header {
                        fmt.Fprintf(writer, "Header: %v: %v\n", header, vals)
                    }
                    fmt.Fprintln(writer, "----")
                    data, err := io.ReadAll(request.Body)
                    if err == nil {
                        if len(data) == 0 {
                            fmt.Fprintln(writer, "No body")
                        } else {
                            writer.Write(data)
                        }
                    } else {
                        fmt.Fprintf(os.Stdout, "Error reading body: %v\n", err.Error())
                    }
                })
        }
    6- The Contents of the main.go File in the httpclient Folder
        package main
        import (
            "net/http"
        )
        func main() {
            Printfln("Starting HTTP Server")
            http.ListenAndServe(":5000", nil)
        }
    =======================================================================================
    Output:
        The code in httpclient folder will be compiled and executed. 
        Use a web browser to request http://localhost:5000/html and http://localhost:5000/json,
        To see the echo result, request http://localhost:5000/echo
████████████████████████████████████████████████████████████████████████
437.Sending Simple HTTP Requests
    The net/http package provides a set of convenience functions that make basic HTTP requests. 
    The functions are named after the HTTP method of the request they created.
    
    The Convenience Methods for HTTP Requests
    Name                                Description
    -----------------                   -----------------------------
    Get(url)                            This function sends a GET request to the specified HTTP or HTTPS URL. The
                                        results are a Response and an error that reports problems with the request.
    Head(url)                           This function sends a HEAD request to the specified HTTP or HTTPS URL.
                                        A HEAD request returns the headers that would be returned for a GET request.
                                        The results are a Response and an error that reports problems with the request.
    Post(url, contentType, reader)      This function sends a POST request to the specified HTTP or HTTPS URL, with
                                        the specified Content-Type header value. The content for the form is provided
                                        by the specified Reader. The results are a Response and an error that reports
                                        problems with the request.
    PostForm(url, data)                 This function sends a POST request to the specified HTTP or HTTPS URL, with
                                        the Content-Type header set to application/x-www-form-urlencoded. The
                                        content for the form is provided by a map[string][]string. The results are a
                                        Response and an error that reports problems with the request.

████████████████████████████████████████████████████████████████████████
438.Sending a GET Request in the main.go
    main.go:
        package main
        import (
            "net/http"
            "os"
            "time"
        )
        func main() {
            go http.ListenAndServe(":5000", nil)
            time.Sleep(time.Second)
            response, err := http.Get("http://localhost:5000/html")
            if (err == nil) {
                response.Write(os.Stdout)
            } else {
                Printfln("Error: %v", err.Error())
            }
        }
    ===================================================================================================
    The argument to the Get function is a string that contains the URL to request. 
    The results are a Response value and an error that reports any problems sending the request.

    Output:
        HTTP/1.1 200 OK
        Content-Length: 171
        Accept-Ranges: bytes
        Content-Type: text/html; charset=utf-8
        Date: Thu, 12 Oct 2023 16:17:10 GMT
        Last-Modified: Wed, 11 Oct 2023 12:44:50 GMT
        
        <!DOCTYPE html>
        <html>
        <head>
            <title>Pro Go</title>
            <meta name="viewport" content="width=device-width" />
        </head>
        <body>
            <h1>Hello, World</div>
        </body>
        </html>
    
    ========================================================================================================
    example-2:
    main.go:
    package main
    import (
        "net/http"
        "os"
        // "time"
    )
    func main() {
        // go http.ListenAndServe(":5000", nil)
        // time.Sleep(time.Second)
        response, err := http.Get("https://www.google.com")
        if (err == nil) {
            response.Write(os.Stdout)
        } else {
            Printfln("Error: %v", err.Error())
        }
    }

Output:
    HTTP/2.0 403 Forbidden
    Content-Length: 1579
    Alt-Svc: h3=":443"; ma=2592000,h3-29=":443"; ma=2592000
    Content-Type: text/html; charset=UTF-8
    Date: Thu, 12 Oct 2023 16:15:23 GMT
    Referrer-Policy: no-referrer
    
    <!DOCTYPE html>
    <html lang=en>
    <meta charset=utf-8>
    <meta name=viewport content="initial-scale=1, minimum-scale=1, width=device-width">
    <title>Error 403 (Forbidden)!!1</title>
    <style>
        *{margin:0;padding:0}html,code{font:15px/22px arial,sans-serif}html{background:#fff;color:#222;padding:15px}body{margin:7% auto 0;max-width:390px;min-height:180px;padding:30px 0 15px}* > body{background:url(//www.google.com/images/errors/robot.png) 100% 5px no-repeat;padding-right:205px}p{margin:11px 0 22px;overflow:hidden}ins{color:#777;text-decoration:none}a img{border:0}@media screen and (max-width:772px){body{background:none;margin-top:0;max-width:none;padding-right:0}}#logo{background:url(//www.google.com/images/branding/googlelogo/1x/googlelogo_color_150x54dp.png) no-repeat;margin-left:-5px}@media only screen and (min-resolution:192dpi){#logo{background:url(//www.google.com/images/branding/googlelogo/2x/googlelogo_color_150x54dp.png) no-repeat 0% 0%/100% 100%;-moz-border-image:url(//www.google.com/images/branding/googlelogo/2x/googlelogo_color_150x54dp.png) 0}}@media only screen and (-webkit-min-device-pixel-ratio:2){#logo{background:url(//www.google.com/images/branding/googlelogo/2x/googlelogo_color_150x54dp.png) no-repeat;-webkit-background-size:100% 100%}}#logo{display:inline-block;height:54px;width:150px}
    </style>
    <a href=//www.google.com/><span id=logo aria-label=Google></span></a>
    <p><b>403.</b> <ins>That's an error.</ins>
    <p>Your client does not have permission to get URL <code>/</code> from this server.  <ins>That's all we know.</ins>

████████████████████████████████████████████████████████████████████████
439.The Fields and Methods Defined by the Response Struct
    Name            Description
    ----------      -------------------------------------
    StatusCode      This field returns the response status code, expressed as an int.
    Status          This field returns a string containing the status description.
    Proto           This field returns a string containing the response HTTP protocol.
    Header          This field returns a map[string][]string that contains the response headers.
    Body            This field returns a ReadCloser, which is a Reader that defines a Close method and
                    which provides access to the response body.
    Trailer         This field returns a map[string][]string that contains the response trailers.
    ContentLength   This field returns the value of the Content-Length header, parsed into an int64 value.
                    TransferEncoding This field returns the set of Transfer-Encoding header values.
    Close           This bool field returns true if the response contains a Connection header set to close,
                    which indicates that the HTTP connection should be closed.
    Uncompressed    This field returns true if the server sent a compressed response that was
                    decompressed by the net/http package.
    Request         This field returns the Request that was used to obtain the response. The Request struct
                    is described in Chapter 24.
    TLS             This field provides details of the HTTPS connection.
    Cookies()       This method returns a []*Cookie, which contains the Set-Cookie headers in the
                    response. The Cookie struct is described in Chapter 24.
    Location()      This method returns the URL from the response Location header and an error that
                    indicates when the response does not contain this header.
    Write(writer)   This method writes a summary of the response to the specified Writer.
████████████████████████████████████████████████████████████████████████
440.Reading the Response Body in the main.go
    main.go:
        package main
        import (
            "net/http"
            "os"
            "time"
            "io"
        )
        func main() {
            go http.ListenAndServe(":5000", nil)
            time.Sleep(time.Second)
            response, err := http.Get("http://localhost:5000/html")
            if (err == nil && response.StatusCode == http.StatusOK) {
                data, err := io.ReadAll(response.Body)
                if (err == nil) {
                    defer response.Body.Close()
                    os.Stdout.Write(data)
                }
            } else {
                Printfln("Error: %v, Status Code: %v", err.Error(), response.StatusCode)
            }
        }
    ====================================================================
    Output: in Terminal
        <!DOCTYPE html>
        <html>
        <head>
            <title>Pro Go</title>
            <meta name="viewport" content="width=device-width" />
        </head>
        <body>
            <h1>Hello, World</div>
        </body>
        </html>
████████████████████████████████████████████████████████████████████████
441.Reading and Parsing Data in the main.go
    main.go:
        package main
        import (
            "net/http"
            //"os"
            "time"
            //"io"
            "encoding/json"
        )
        func main() {
            go http.ListenAndServe(":5000", nil)
            time.Sleep(time.Second)
            response, err := http.Get("http://localhost:5000/json")
            if (err == nil && response.StatusCode == http.StatusOK) {
                defer response.Body.Close()
                data := []Product {}
                err = json.NewDecoder(response.Body).Decode(&data)
                if (err == nil) {
                    for _, p := range data {
                        Printfln("Name: %v, Price: $%.2f", p.Name, p.Price)
                    }
                } else {
                    Printfln("Decode error: %v", err.Error())
                }
            } else {
                Printfln("Error: %v, Status Code: %v", err.Error(), response.StatusCode)
            }
        }
    ====================================================================
    Output: in Terminal
        Name: Kayak,Price: $279.00
        Name: Lifejacket,Price: $49.95
        Name: Soccer Ball,Price: $19.50
        Name: Corner Flags,Price: $34.95
        Name: Stadium,Price: $79500.00
        Name: Thinking Cap,Price: $16.00
        Name: Unsteady Chair,Price: $75.00
        Name: Bling-Bling King,Price: $1200.00
████████████████████████████████████████████████████████████████████████
442.Sending POST Requests
    The Post and PostForm functions are used to send POST requests. 
    The PostForm function encodes a map of values as form data.

    Sending a Form in the main.go
    main.go:
        package main
        import (
            "net/http"
            "os"
            "time"
            "io"
            //"encoding/json"
        )
        func main() {
            go http.ListenAndServe(":5000", nil)
            time.Sleep(time.Second)
            formData := map[string][]string {
                "name":  { "Kayak "},
                "category": { "Watersports"},
                "price":  { "279"},
            }
            response, err := http.PostForm("http://localhost:5000/echo", formData)
            if (err == nil && response.StatusCode == http.StatusOK) {
                io.Copy(os.Stdout, response.Body)
                defer response.Body.Close()
            } else {
                Printfln("Error: %v, Status Code: %v", err.Error(), response.StatusCode)
            }
        }
    ====================================================================
    Output: 
        Method: POST
        Header: Accept-Encoding: [gzip]
        Header: User-Agent: [Go-http-client/1.1]
        Header: Content-Length: [42]
        Header: Content-Type: [application/x-www-form-urlencoded]
        ----
        category=Watersports&name=Kayak+&price=279
████████████████████████████████████████████████████████████████████████
443.Posting a Form Using a Reader
    The Post function sends a POST request 
    to the server and creates the request body by reading content from a Reader.

    Posting from a Reader in the main.go:
        package main
        import (
            "net/http"
            "os"
            "time"
            "io"
            "encoding/json"
            "strings"
        )
        func main() {
            go http.ListenAndServe(":5000", nil)
            time.Sleep(time.Second)
            var builder strings.Builder
            err := json.NewEncoder(&builder).Encode(Products[0])
            if (err == nil) {
                response, err := http.Post("http://localhost:5000/echo",
                    "application/json",
                    strings.NewReader(builder.String()))
                if (err == nil && response.StatusCode == http.StatusOK) {
                    io.Copy(os.Stdout, response.Body)
                    defer response.Body.Close()
                } else {
                    Printfln("Error: %v", err.Error())
                }
            } else {
                Printfln("Error: %v", err.Error())
            }
        }
    ====================================================================
    Output:
        Method: POST
        Header: User-Agent: [Go-http-client/1.1]
        Header: Content-Length: [54]
        Header: Content-Type: [application/json]
        Header: Accept-Encoding: [gzip]
        ----
        {"Name":"Kayak","Category":"Watersports","Price":279}
████████████████████████████████████████████████████████████████████████
444.
████████████████████████████████████████████████████████████████████████
445.
████████████████████████████████████████████████████████████████████████
446.
████████████████████████████████████████████████████████████████████████
447.
████████████████████████████████████████████████████████████████████████
448.
████████████████████████████████████████████████████████████████████████
449.
████████████████████████████████████████████████████████████████████████
450.
████████████████████████████████████████████████████████████████████████
451.
████████████████████████████████████████████████████████████████████████
452.
████████████████████████████████████████████████████████████████████████
453.
████████████████████████████████████████████████████████████████████████
454.
████████████████████████████████████████████████████████████████████████
455.
████████████████████████████████████████████████████████████████████████
456.
████████████████████████████████████████████████████████████████████████
457.
████████████████████████████████████████████████████████████████████████
458.
████████████████████████████████████████████████████████████████████████
459.
████████████████████████████████████████████████████████████████████████
460.
████████████████████████████████████████████████████████████████████████
461.
████████████████████████████████████████████████████████████████████████
462.
████████████████████████████████████████████████████████████████████████
463.
████████████████████████████████████████████████████████████████████████
464.
████████████████████████████████████████████████████████████████████████
465.
████████████████████████████████████████████████████████████████████████
466.
████████████████████████████████████████████████████████████████████████
467.
████████████████████████████████████████████████████████████████████████
468.
████████████████████████████████████████████████████████████████████████
469.
████████████████████████████████████████████████████████████████████████
470.
████████████████████████████████████████████████████████████████████████
471.
████████████████████████████████████████████████████████████████████████
472.
████████████████████████████████████████████████████████████████████████
473.
████████████████████████████████████████████████████████████████████████
474.
████████████████████████████████████████████████████████████████████████
475.
████████████████████████████████████████████████████████████████████████
476.
████████████████████████████████████████████████████████████████████████
477.
████████████████████████████████████████████████████████████████████████
478.
████████████████████████████████████████████████████████████████████████
479.
████████████████████████████████████████████████████████████████████████
480.
████████████████████████████████████████████████████████████████████████
481.
████████████████████████████████████████████████████████████████████████
482.
████████████████████████████████████████████████████████████████████████
483.
████████████████████████████████████████████████████████████████████████
484.
████████████████████████████████████████████████████████████████████████
485.
████████████████████████████████████████████████████████████████████████
486.
████████████████████████████████████████████████████████████████████████
487.
████████████████████████████████████████████████████████████████████████
488.
████████████████████████████████████████████████████████████████████████
489.
████████████████████████████████████████████████████████████████████████
490.
████████████████████████████████████████████████████████████████████████
491.
████████████████████████████████████████████████████████████████████████
492.
████████████████████████████████████████████████████████████████████████
493.
████████████████████████████████████████████████████████████████████████
494.
████████████████████████████████████████████████████████████████████████
495.
████████████████████████████████████████████████████████████████████████
496.
████████████████████████████████████████████████████████████████████████
497.
████████████████████████████████████████████████████████████████████████
498.
████████████████████████████████████████████████████████████████████████
499.
████████████████████████████████████████████████████████████████████████
500.
████████████████████████████████████████████████████████████████████████
501.
████████████████████████████████████████████████████████████████████████
502.
████████████████████████████████████████████████████████████████████████
503.
████████████████████████████████████████████████████████████████████████
504.
████████████████████████████████████████████████████████████████████████
505.
████████████████████████████████████████████████████████████████████████
506.
████████████████████████████████████████████████████████████████████████
507.
████████████████████████████████████████████████████████████████████████
508.
████████████████████████████████████████████████████████████████████████
509.
████████████████████████████████████████████████████████████████████████
510.
████████████████████████████████████████████████████████████████████████
511.
████████████████████████████████████████████████████████████████████████
512.
████████████████████████████████████████████████████████████████████████
513.
████████████████████████████████████████████████████████████████████████
514.
████████████████████████████████████████████████████████████████████████
515.
████████████████████████████████████████████████████████████████████████
516.
████████████████████████████████████████████████████████████████████████
517.
████████████████████████████████████████████████████████████████████████
518.
████████████████████████████████████████████████████████████████████████
519.
████████████████████████████████████████████████████████████████████████
420.
████████████████████████████████████████████████████████████████████████
521.
████████████████████████████████████████████████████████████████████████
522.
████████████████████████████████████████████████████████████████████████
523.
████████████████████████████████████████████████████████████████████████
524.
████████████████████████████████████████████████████████████████████████
525.
████████████████████████████████████████████████████████████████████████
526.
████████████████████████████████████████████████████████████████████████
527.
████████████████████████████████████████████████████████████████████████
528.
████████████████████████████████████████████████████████████████████████
529.
████████████████████████████████████████████████████████████████████████
530.
████████████████████████████████████████████████████████████████████████
531.
████████████████████████████████████████████████████████████████████████
532.
████████████████████████████████████████████████████████████████████████
533.
████████████████████████████████████████████████████████████████████████
534.
████████████████████████████████████████████████████████████████████████
535.
████████████████████████████████████████████████████████████████████████
536.
████████████████████████████████████████████████████████████████████████
537.
████████████████████████████████████████████████████████████████████████
538.
████████████████████████████████████████████████████████████████████████
539.
████████████████████████████████████████████████████████████████████████
540.
████████████████████████████████████████████████████████████████████████
541.
████████████████████████████████████████████████████████████████████████
542.
████████████████████████████████████████████████████████████████████████
543.
████████████████████████████████████████████████████████████████████████
544.
████████████████████████████████████████████████████████████████████████
545.
████████████████████████████████████████████████████████████████████████
546.
████████████████████████████████████████████████████████████████████████






































████████████████████████████████████████████████████████████████████████
████████████████████████████████████████████████████████████████████████
████████████████████████████████████████████████████████████████████████
████████████████████████████████████████████████████████████████████████
████████████████████████████████████████████████████████████████████████
████████████████████████████████████████████████████████████████████████
████████████████████████████████████████████████████████████████████████
████████████████████████████████████████████████████████████████████████
████████████████████████████████████████████████████████████████████████
████████████████████████████████████████████████████████████████████████


`,
}
