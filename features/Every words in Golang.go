package features

type Features struct {
	EveryWordsInGolang string
}

var OriginalFeatures Features = Features{
	EveryWordsInGolang: `

in Command Line Interface Go:
----------------------------------------------------------------
1- go mod init YOUR_NAME
2- go work init YOUR_WORK_DIRECTORY
3- go run main.go  OR  go run projectName.go


Using the Go Command
------------------------------------------------------------------------------------------------
1.build
    The go build command compiles the source code in the current directory 
    and generates an executable file.

2.clean
    The go clean command removes the output produced by the go build command, 
    including the executable and any temporary files that were created during the build.

3.doc
    The go doc command generates documentation from source code.

4.fmt
    The go fmt command ensures consistent indentation and alignment in source code files.

5.get
    The go get command downloads and installs external packages.

6.install
    The go install command downloads packages and is usually used to install tool packages.

7.help
    The go help command displays help information for other Go features. The command go
    help build, for example, displays information about the build argument.

8.mod
    The go mod command is used to create and manage a Go module.

9.run
    The go run command builds and executes the source code in a specified folder without
    creating an executable output

10.test
    The go test command executes unit tests

11.version
    The go version command writes out the Go version number.

12.vet
    The go vet command detects common problems in Go code

    

Useful Debugger State Commands
------------------------------------------------------------------------------------------------
13.print <expr>
    This command evaluates an expression and displays the result. It can
    be used to display a value (print i) or perform a more complex test
    (print i > 0).

14.set <variable> = <value>
    This command changes the value of the specified variable.
    این دستور مقدار متغیر مشخص شده را تغییر می دهد.

15.locals
    This command prints the value of all local variables.

16.whatis <expr>
    This command prints the type of the specified expression such as whatis


Useful Debugger Commands for Controlling Execution
------------------------------------------------------------------------------------------------
17.continue
    This command resumes execution of the application.

18.next
    This command moves to the next statement.

19.step
    This command steps into the current statement.

20.stepout
    This command steps out of the current statement.

21.restart
    This command restarts the process. Use the continue command to begin execution.
    
22.exit
    This command exits the debugger.



Understanding the Basic Data Types
------------------------------------------------------------------------------------------------
23.int = integers
    This type represents a whole number, which can be positive or negative. The
    int type size is platform-dependent and will be either 32 or 64 bits. There are
    also integer types that have a specific size, such as int8, int16, int32, and
    int64, but the int type should be used unless you need a specific size.
    
    Literal Value Examples:
        20, -20. Values can also be expressed in hex (0x14), octal (0o24), and binary notation
        (0b0010100).

24.unit
    uint8,  uint16,  uint32,  uint64,int8,  int16,  int32
    This type represents a positive whole number. The uint type size is platform-
    dependent and will be either 32 or 64 bits. There are also unsigned integer
    types that have a specific size, such as uint8, uint16, uint32, and uint64, but
    the uint type should be used unless you need a specific size.

    Literal Value Examples:
        There are no uint literals. All literal whole numbers are treated as int values.

25.byte = uint8
    This type is an alias for uint8 and is typically used to represent a byte of data.
    byte = 8 bits, 
    1024bytes = 1 kilobyte
    1024 kilobytes = 1 megabyte

    Literal Value Examples:
        There are no byte literals. Bytes are typically expressed as integer literals (such as 101) or run
        literals ('e') since the byte type is an alias for the uint8 type.

26.float32, float64
    Floating Point Numbers. e.g 3.14, 123.541, 100.4020401
    These types represent numbers with a fraction. These types allocate 32 or 64
    bits to store the value.

    Literal Value Examples:
        20.2, -20.2, 1.2e10, 1.2e-10. Values can also be expressed in hex notation (0x2p10), although
        the exponent is expressed in decimal digits.

27.complex64, complex128
    These types represent numbers that have real and imaginary components.
    These types allocate 64 or 128 bits to store the value.

28.bool 
    bolean
    This type represents a Boolean truth with the values true and false.

    Literal Value Examples:
        true, false.

29.string
    This type represents a sequence of characters.

    Literal Value Examples:
        "Hello". Character sequences escaped with a backslash are interpreted if the value is enclosed
        in double quotes ("Hello\n"). Escape sequences are not interpreted if the value is enclosed in
        backquotes ('Hello\n').

30.rune = int32
    This type represents a single Unicode code point. Unicode is complicated,
    but—loosely—this is the representation of a single character. The rune type is
    an alias for int32.
   
    Literal Value Examples:
        'A', '\n', '\u00A5', '¥'. Characters, glyphs, and escape sequences are enclosed in single
        quotes (the ' character).

31.var = variable 
    Variables are defined using the var keyword, and, unlike constants, the value assigned to a variable can be
    changed

32.const = constant
    Constants are basically variables whose values cannot be changed later.


The Zero Values for the Basic Data Types
------------------------------------------------------------------------------------------------
Type    Zero Value
------  -----------
int         0
unit        0
byte        0
float64     0
bool        false
string      “” (the empty string)
rune        0





Operations and Conversions
------------------------------------------------------------------------------------------------
33.+, -, *, /, %
    These operators are used to perform arithmetic using numeric values.

34.==, !=, <, <=, >, >=
    These operators compare two values.

35.||, &&, !
    These are the logical operators, which are applied to bool values and return a bool value.

36.=, :=
    These are the assignment operators. The standard assignment operator (=) is used to set
    the initial value when a constant or variable is defined, or to change the value assigned to
    a previously defined variable. The shorthand operator (:=) is used to define a variable and
    assign a value.

37.-=, +=, ++, --
    These operators increment and decrement numeric values.

38.&, |, ^, &^, <<, >>
    These are the bitwise operators, which can be applied to integer values. These operators are
    not often required in mainstream development
    where the | operator is used to configure the Go logging features.

The Arithmetic Operators
------------------------------------------------------------------------------------------------
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
------------------------------------------------------------------------------------------------
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
------------------------------------------------------------------------------------------------
Operator    Description
-------     ---------------------
   ||       This operator returns true if either operand is true. 
            If the first operand is true, then the second
            operand will not be evaluated.

   &&       This operator returns true if both operands are true. 
            If the first operand is false, then the
            second operand will not be evaluated.

   !        This operator is used with a single operand. 
            It returns true if the operand is false and false if
            the operand is true.

            
Converting Floating-Point Values to Integers
Functions in the math Package for Converting Numeric Types
------------------------------------------------------------------------------------------------
39.Ceil(value)
    This function returns the smallest integer that is greater than the specified floating-
    point value. The smallest integer that is greater than 27.1, for example, is 28.

40.Floor(value)
    This function returns the largest integer that is less than the specified floating-point
    value. The largest integer that is less than 27.1, for example, is 28.

41.Round(value)
    This function rounds the specified floating-point value to the nearest integer.

42.RoundToEven(value)
    This function rounds the specified floating-point value to the nearest even integer.



Parsing from Strings
Functions for Parsing Strings into Other Data Types
------------------------------------------------------------------------------------------------
43.ParseBool(str)
    This function parses a string into a bool value. Recognized string values are "true",
    "false", "TRUE", "FALSE", "True", "False", "T", "F", "0", and "1".

44.ParseFloat(str,size)
    This function parses a string into a floating-point value with the specified size, as
    described in the “Parsing Floating-Point Numbers” section.

45.ParseInt(str,base, size)
    This function parses a string into an int64 with the specified base and size. Acceptable
    base values are 2 for binary, 8 for octal, 16 for hex, and 10.

46.ParseUint(str,base, size)
    This function parses a string into an unsigned integer value with the specified base and size.

47.Atoi(str)
    This function parses a string into a base 10 int and is equivalent to calling
    ParseInt(str, 10, 0)


Formatting Values as Strings
The strconv Functions for Converting Values into Strings
------------------------------------------------------------------------------------------------
48.FormatBool(val)
    This function returns the string true or false based on the value of the
    specified bool.

49.FormatInt(val, base)
    This function returns a string representation of the specified int64 value,
    expressed in the specified base.

50.FormatUint(val, base)
    This function returns a string representation of the specified uint64 value,
    expressed in the specified base.

51.FormatFloat(val, format, precision, size) 
    This function returns a string representation of the specified float64 value,
    expressed using the specified format, precision, and size.

52.Itoa(val)
    This function returns a string representation of the specified int value,
    expressed using base 10.


Commonly Used Format Options for Floating-Point String Formatting
------------------------------------------------------------------------------------------------
Function
---------
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


Understanding Flow Control
------------------------------------------------------------------------------------------------
53.if
    The if keyword is followed by the expression and then the group of statements to be executed,
    surrounded by braces

54.else
    The else keyword can be used to create additional clauses in an if statement
    
55.else if
    The else/if combination can be repeated to create a sequence of clauses

56.for
    The for keyword is used to create loops that repeatedly execute statements. The most basic for loops will
    repeat indefinitely unless interrupted by the break keyword

    Incorporating the Condition into the Loop
    
    example: 
        for (counter <= 3) {}

    Enumerating Sequences:
        for index, character := range product {
                    fmt.Println("Index:", index, "Character:", string(character))
                }

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
        switch (character) {
            case 'K', 'k':
                fmt.Println("K or k at position", index)
            case 'y':
                fmt.Println("y at position", index)
        }

    Terminate case Statement Execution:
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

61.Array
    Go arrays are a fixed length and contain elements of a single type, which are accessed by index
    var names [3]string

    Array Literal Syntax:
        names := [3]string { "Kayak", "Lifejacket", "Paddle" }

    The number of elements specified with the literal syntax can be less than the capacity of the array. 
    Any position in the array for which a value is not provided will be assigned the zero value for the array type.

    Enumerating Arrays:
        names := [3]string { "Kayak", "Lifejacket", "Paddle" }
        for index, value := range names {
            fmt.Println("Index:", index, "Value:", value)
        }
    
62.Slices
    The slice type in this example is []string
    The best way to think of slices is as a variable-length array because they are useful when you don’t know how
    many values you need to store or when the number changes over time. One way to define a slice is to use the
    built-in make function

    example:
        names := make([]string, 3)
        names[0] = "Kayak"
        names[1] = "Lifejacket"
        names[2] = "Paddle"

    Literal Syntax:
        names := []string {"Kayak", "Lifejacket", "Paddle"}

63.append
    If you define a slice variable but don’t initialize it, then the result is a slice that has a length of zero
    and a capacity of zero, and this will cause an error when an element is appended to it.
    One of the key advantages of slices is that they can be expanded to accommodate additional elements
    
    Appending Elements to a Slice:
        names := []string {"Kayak", "Lifejacket", "Paddle"}
        names = append(names, "Hat", "Gloves")

    Appending Items to a Slice:
        names := []string {"Kayak", "Lifejacket", "Paddle"}
        appendedNames := append(names, "Hat", "Gloves")
        names[0] = "Canoe"

    Allocating Additional Slice Capacity:
        names := make([]string, 3, 6)
        names[0] = "Kayak"
        names[1] = "Lifejacket"
        names[2] = "Paddle"

    Adding Elements to a Slice:
        names := make([]string, 3, 6)
        names[0] = "Kayak"
        names[1] = "Lifejacket"
        names[2] = "Paddle"
        appendedNames := append(names, "Hat", "Gloves")
        names[0] = "Canoe"
        fmt.Println("names:",names)
        fmt.Println("appendedNames:", appendedNames)


    Appending One Slice to Another:
        names := make([]string, 3, 6)
        names[0] = "Kayak"
        names[1] = "Lifejacket"
        names[2] = "Paddle"
        moreNames := []string { "Hat Gloves"}
        appendedNames := append(names, moreNames...)
        fmt.Println("appendedNames:", appendedNames)

    Creating Slices from Existing Arrays:
        products := [4]string { "Kayak", "Lifejacket", "Paddle", "Hat"}
        someNames := products[1:3]
        allNames := products[:]

    Specifying Capacity When Creating a Slice from an Array:
        products := [4]string { "Kayak", "Lifejacket", "Paddle", "Hat"}
        someNames := products[1:3:3]
        allNames := products[:]
        someNames = append(someNames, "Gloves")

    Creating Slices from Other Slices:
        products := [4]string { "Kayak", "Lifejacket", "Paddle", "Hat"}
        allNames := products[1:]
        someNames := allNames[1:3]
        allNames = append(allNames, "Gloves")
        allNames[1] = "Canoe"

    Comparing Slices:
        Go restricts the use of the comparison operator so that slices can be compared only to the nil value.
            p1 := []string { "Kayak", "Lifejacket", "Paddle", "Hat"}
            p2 := p1
            fmt.Println("Equal:", p1 == p2)
        Output:
            .\main.go:13:30: invalid operation: p1 == p2 (slice can only be compared to nil)


64.copy
    The copy function is used to copy elements between slices. This function can be used to ensure that slices
    have separate arrays and to create slices that combine elements from different sources.

         Destination   Source
         ---------     --------
    copy(someNames,    allNames)


    Duplicating a Slice:
    The copy function accepts two arguments:
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
        products := [4]string { "Kayak", "Lifejacket", "Paddle", "Hat"}
        allNames := products[1:]
        var someNames []string
        copy(someNames, allNames)
    Output:
        someNames: []
        allNames [Lifejacket Paddle Hat]

    Specifying Ranges When Copying Slices:
    Fine-grained control over the elements that are copied can be achieved using ranges
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
        products := []string { "Kayak", "Lifejacket", "Paddle", "Hat"}
        for index, value := range products[2:] {
            fmt.Println("Index:", index, "Value:", value)
        }
    Output:
        Index: 0 Value: Paddle
        Index: 1 Value: Hat



65.sort
    There is no built-in support for sorting slices, but the standard library includes the sort package, which
    defines functions for sorting different types of slice.
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

66.DeepEqual
    The DeepEqual function can be used to compare a wider range of data types than the
    equality operator, including slices.
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
        products := map[string]float64 {
            "Kayak" : 279,
            "Lifejacket": 48.95,
            "Hat": 0,
        }
        delete(products, "Hat")            



    Enumerating the Contents of a Map:
        products := map[string]float64 {
                "Kayak" : 279,
                "Lifejacket": 48.95,
                "Hat": 0,
        }
        for key, value := range products {
            fmt.Println("Key:", key, "Value:", value)
        }

    Enumerating a Map in Order:
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
            for key, _ := range products {
                keys = append(keys, key)
            }
            sort.Strings(keys)
            for _, key := range keys {
                fmt.Println("Key:", key, "Value:", products[key])
            }
        }

        

Understanding the Dual Nature of Strings
------------------------------------------------------------------------------------------------        
68.strconv
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

Converting a String to Runes
------------------------------------------------------------------------------------------------        
69.rune
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
------------------------------------------------------------------------------------------------        
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
------------------------------------------------------------------------------------------------        
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

70.func
    Functions are groups of code statements that are executed only when the function is
    invoked during the flow of execution.

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


72.Defining Variadic Parameters
    example:
        func printSuppliers(product string, suppliers []string ) {
            for _, supplier := range suppliers {
                fmt.Println("Product:", product, "Supplier:", supplier)
            }
        }

    example:
        func printSuppliers(product string, suppliers ...string ) {
            for _, supplier := range suppliers {
                fmt.Println("Product:", product, "Supplier:", supplier)
            }
        }

73.Dealing with No Arguments for a Variadic Parameter
    example:
        func printSuppliers(product string, suppliers ...string ) {
            for _, supplier := range suppliers {
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

74.return Function Results
    func calcTax(price float64) float64 {
        return price + (price * 0.2)
    }

75.Returning Multiple Function Results
    func swapValues(first, second int) (int, int) {
        return second, first
    }
    func main() {
        val1, val2 := 10, 20
        fmt.Println("Before calling function", val1, val2)
        val1, val2 = swapValues(val1, val2)
        fmt.Println("After calling function", val1, val2)
    }

    





































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
