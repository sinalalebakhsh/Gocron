package features

type Features struct {
	EveryWordsInGolang string
}


var OriginalFeatures Features = Features{
	EveryWordsInGolang: `
 ┌────────────────────────────────────────────────────────────────────┐ 
 │                       GoCron v2.49.1                               │
 │ █████████  █████████  ████████  █████████  █████████  ███      ██  │
 │ █          █       █  █         █       █  █       █  ████     ██  │
 │ █          █       █  █         █       █  █       █  ██ ██    ██  │
 │ █     ███  █       █  █         █████████  █       █  ██  ██   ██  │
 │ █       █  █       █  █         █    ██    █       █  ██   ██  ██  │
 | █       █  █       █  █         █     ██   █       █  ██    ██ ██  │
 | █████████  █████████  ████████  █      ██  █████████  ██     ████  │
 │       is a API for learning GO language with example               │
 └────────────────────────────────────────────────────────────────────┘ 
 

████████████████████████████████████████████████████████████████████████
0.Create Environment GO
    in Command Line Interface Go:
    1- go mod init YOURNAME
    2- go work init YOURWORKDIRECTORY
    3- go run main.go  OR  go run projectName.go
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
    built-in make function

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
    maps can use other data types as keys

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

    
████████████████████████████████████████████████████████████████████████
244.
████████████████████████████████████████████████████████████████████████
245.
████████████████████████████████████████████████████████████████████████
246.
████████████████████████████████████████████████████████████████████████
247.
████████████████████████████████████████████████████████████████████████
248.
████████████████████████████████████████████████████████████████████████
249.
████████████████████████████████████████████████████████████████████████
250.
████████████████████████████████████████████████████████████████████████
251.
████████████████████████████████████████████████████████████████████████
252.
████████████████████████████████████████████████████████████████████████
253.
████████████████████████████████████████████████████████████████████████
254.
████████████████████████████████████████████████████████████████████████
255.
████████████████████████████████████████████████████████████████████████
256.
████████████████████████████████████████████████████████████████████████
257.
████████████████████████████████████████████████████████████████████████
258.
████████████████████████████████████████████████████████████████████████
259.
████████████████████████████████████████████████████████████████████████
260.
████████████████████████████████████████████████████████████████████████
261.
████████████████████████████████████████████████████████████████████████
262.
████████████████████████████████████████████████████████████████████████
263.
████████████████████████████████████████████████████████████████████████
264.
████████████████████████████████████████████████████████████████████████
265.
████████████████████████████████████████████████████████████████████████
266.
████████████████████████████████████████████████████████████████████████
267.
████████████████████████████████████████████████████████████████████████
268.
████████████████████████████████████████████████████████████████████████
269.
████████████████████████████████████████████████████████████████████████
270.
████████████████████████████████████████████████████████████████████████
271.
████████████████████████████████████████████████████████████████████████
272.
████████████████████████████████████████████████████████████████████████
273.
████████████████████████████████████████████████████████████████████████
274.
████████████████████████████████████████████████████████████████████████
275.
████████████████████████████████████████████████████████████████████████
276.
████████████████████████████████████████████████████████████████████████
277.
████████████████████████████████████████████████████████████████████████
278.
████████████████████████████████████████████████████████████████████████
279.
████████████████████████████████████████████████████████████████████████
280.
████████████████████████████████████████████████████████████████████████
281.
████████████████████████████████████████████████████████████████████████
282.
████████████████████████████████████████████████████████████████████████
283.
████████████████████████████████████████████████████████████████████████
284.
████████████████████████████████████████████████████████████████████████
285.
████████████████████████████████████████████████████████████████████████
286.
████████████████████████████████████████████████████████████████████████
287.
████████████████████████████████████████████████████████████████████████
288.
████████████████████████████████████████████████████████████████████████
289.
████████████████████████████████████████████████████████████████████████
290.
████████████████████████████████████████████████████████████████████████
291.
████████████████████████████████████████████████████████████████████████
292.
████████████████████████████████████████████████████████████████████████
293.
████████████████████████████████████████████████████████████████████████
294.
████████████████████████████████████████████████████████████████████████
295.
████████████████████████████████████████████████████████████████████████
296.
████████████████████████████████████████████████████████████████████████
297.
████████████████████████████████████████████████████████████████████████
298.
████████████████████████████████████████████████████████████████████████
299.
████████████████████████████████████████████████████████████████████████
300.
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
████████████████████████████████████████████████████████████████████████
████████████████████████████████████████████████████████████████████████
████████████████████████████████████████████████████████████████████████
████████████████████████████████████████████████████████████████████████
████████████████████████████████████████████████████████████████████████
████████████████████████████████████████████████████████████████████████
████████████████████████████████████████████████████████████████████████
████████████████████████████████████████████████████████████████████████
████████████████████████████████████████████████████████████████████████






























































































































































































































    package main = first executable file main.go
func main() {} = every execute files in "package main"
    fmt = for printing in CLI
import = for importing another package
    string
e.g
    Format String like: fmt.Printf("%v %s %f", variable, string, float32)
    Scope
    Go is lexically scoped using blocks.

    Defining Multiple Variables
    fmt.Scanf("%f", &input)
    We use another function from the fmt package to readthe   user   input   (Scanf).
    for 
    loop
    if
    else
    switch
    default
    array
    An array is a numbered sequence of elements of a sin-gle type with a fixed length.
    target
    value
    range
    len(x)
    length

27.
    slice
    A slice is a segment of an array. Like arrays slices areindexable and have a length. 
    Unlike arrays this lengthis allowed to change.

28.
    make
    f you want to create a slice you should use the built-inmake function

29.
    slice functions:
    append
    add to slice
    [low : high]
    Another way to create slices is to use the [low : high]expression

30.
    map
    A map is an unordered collection of key-value pairs.
    Also known as an associative array, a hash table or dictionary, 
    maps are used to look up a value by its associated key.

31.
    slices maps 

32.
    delete function
    We can also delete items from a map using the built-indelete function.

33.
    make map

34.
    function
    A function is an independent section of code 
    that mapszero or more input parameters to zero or more outputparameters.   
    Functions (also known as procedures orsubroutines) 
    are often represented as a black box: (theblack box represents the function)
    func FunctionName(ParameterName ParameterType) ReturnType { body of function}
    e.g
    func getUserName(UserInput string) string { return UserInput }

35.
    returning multiple values
    Go is also capable of returning multiple values from afunction

36.
    variadic functions
    func add(args ...int) int {}

37.
    closure
    It is possible to create functions inside of functions

38.
    increment adds 1 to the variable x which is defined inthe main function's scope. 
    
39.
    return function from a function

40.
    recursion
    Finally a function is able to call itself. 
    Here is one way to compute the factorial of a number

41.
    defer
    Go has a special statement called defer which schedules 
    a function call to be run after the function completes.
    is often used when resources need to be freed insome way. 
    For example when we open a file we need to make sure to close it later. With defer:

42.
    panic
    the panicfunction to cause a run time error.  

43.
    recover
    We can handle arun-time panic with the built-in recover function.
    recover stops the panic and returns the value that waspassed to the call to panic.

44.
    Pointers    
    * &

45.
    new
    Another way to get a pointer is to use the built-in newfunction.

46.
    type
    In the provided Go code, "type" is a keyword used 
    to define a new named type or alias. 
    It allows you to create custom types with specific behaviors and characteristics.
    By using the "type" keyword, 
    you can create your own named types or aliases that can be used throughout your code, 
    providing better abstraction and type safety.

47.
    struct
    A struct can mix regular and embedded field types
    A struct is a type which contains named fields.
    the keyword structto indicate that we are defining 
    a struct type and alist of fields inside of curly braces.
    Each field has a name and a type. 
    For a struct zero means each of the fields is set to their correspond-ing zero value:
    (0 for ints, 0.0 for floats, "" for strings,nil for pointers, ...)
48.
    new function struct
    e.g:
    c := new(Circle)
    This allocates memory for all the fields,
    sets each of them to their zero value and returns a pointer.
    (*Circle) More often we want to give each of the fieldsa value.
49.
    method
    function receiver VS. function parameter
    The receiver can be of two types: 
    value receiver (func (t T) methodName()) 
    or pointer receiver (func (t *T) methodName()), 
    depending on whether the method modifies the receiver or not. 
50.
    interface
51.
    Concurrency
52.
    Goroutines
53.
    Channels
    chan
54.
    Channel Direction
55.
    Select
    case
56.
    goto
57.
    buffered channels
58.
    packages
    creating packages
59.
    testing
60.
    the core packages
61.
    iput / output
    Go's io package.
62.
    files & folders
    To open a file in Go use the Open function from the os package.
63.
    errors
64.
    containers & sort
    Go has several more collections available underneath the container package.
65.
    hashes & cryptography
    Hash functions in Go are broken into two categories:
    cryptographic and non-crypto-graphic.
    The non-cryptographic hash functions 
    can be found underneath the hash package and include adler32, crc32, crc64 and fnv.
66.
    servers
67.
    HTTP servers
    RPC
    The  net/rpc (remote procedure call) 
    and net/rpc/jsonrpc packages provide an easy way to expose methods 
    so they can be invoked over a network. (rather than just in the program running them)
68.
    parsing command line arguments
69.
    synchronization primitives
70.
    Mutexes
    A mutex (mutal exclusive lock) locks a section of code 
    to a single thread at a time and is used to protect 
    shared resources from non-atomic operations.
71.
    Collections
72.
    function recover()
73.
    function
        min(1,2)
        max(1,2)
74.
    Currying
    Function currying is the practice of writing a function 
    that takes a function (or functions) as input, and returns a new function.
75.
    Replace(s)
    This method returns a string for which all the replacements specified with the
    constructor have been performed on the string s.
76.
    WriteString(writer, s)	
    This method is used to perform the replacements specified with the constructor
    and write the results to an io.Writer

77.
    TrimSpace(s)			
    This function returns the string s without leading or trailing whitespace
    characters.
78.
    Trim(s, set)
    This function returns a string from which any leading or trailing characters
    contained in the string set are removed from the string s.
79.
    TrimLeft(s, set)
    This function returns the string s without any leading character contained
    in the string set. This function matches any of the specified characters—use
    the TrimPrefix function to remove a complete substring.
80.
    TrimRight(s, set)
    This function returns the string s without any trailing character contained
    in the string set. This function matches any of the specified characters—use
    the TrimSuffix function to remove a complete substring.
81.
    TrimPrefix(s, prefix)
    function returns the string s after removing the specified prefix string.
    This function removes the complete prefix string—use the TrimLeft
    function to remove characters from a set.
82.
    TrimSuffix(s, suffix)
    function returns the string s after removing the specified suffix string.
    This function removes the complete suffix string—use the TrimRight
    function to remove characters from a set.
83.
    TrimFunc(s, func)
    This function returns the string s from which any leading or trailing
    character for which a custom function returns true are removed.
84.
    TrimLeftFunc(s, func)
    function returns the string s from which any leading character for
    which a custom function returns true are removed.
85.
    TrimRightFunc(s, func)
    function returns the string s from which any trailing character for
    which a custom function returns true are removed.
86.
    Join(slice, sep)
    function combines the elements in the specified string slice, with the specified
    separator string placed between elements.
87.
    Repeat(s, count)
    function generates a string by repeating the string s for a specified number
    of times.
88.

89.

90.

91.

92.

93.

94.

95.

96.

97.

98.

99.

100.

101.

102.

103.

104.

105.

106.

107.

108.

109.

110.

111.

112.

113.

114.

115.
    `,
}

