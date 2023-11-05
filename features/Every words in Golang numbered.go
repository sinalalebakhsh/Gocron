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
		"17":`17.continue
Useful Debugger Commands for Controlling Execution
This command resumes execution of the application.`,
// ====================================================================================
		"18":`18.next
This command moves to the next statement.`,
// ====================================================================================
		"19":`19.step
This command steps into the current statement.`,
// ====================================================================================
		"20":`20.stepout
This command steps out of the current statement.`,
// ====================================================================================
		"21":`21.restart
This command restarts the process. Use the continue command to begin execution.`,
// ====================================================================================
		"22":`22.exit
This command exits the debugger.`,
// ====================================================================================
		"23":`23.int = integers
Understanding the Basic Data Types
This type represents a whole number, which can be positive or negative. The
int type size is platform-dependent and will be either 32 or 64 bits. There are
also integer types that have a specific size, such as int8, int16, int32, and
int64, but the int type should be used unless you need a specific size.

Literal Value Examples:
	20, -20. Values can also be expressed in hex (0x14), octal (0o24), and binary notation
	(0b0010100).`,
// ====================================================================================
		"24":`24.unit
uint8,  uint16,  uint32,  uint64,int8,  int16,  int32
This type represents a positive whole number. The uint type size is platform-
dependent and will be either 32 or 64 bits. There are also unsigned integer
types that have a specific size, such as uint8, uint16, uint32, and uint64, but
the uint type should be used unless you need a specific size.
Literal Value Examples:
	There are no uint literals. All literal whole numbers are treated as int values.`,
// ====================================================================================
		"25":`25.byte = uint8
This type is an alias for uint8 and is typically used to represent a byte of data.
byte = 8 bits, 
1024bytes = 1 kilobyte
1024 kilobytes = 1 megabyte
Literal Value Examples:
	There are no byte literals. Bytes are typically expressed as integer literals (such as 101) or run
	literals ('e') since the byte type is an alias for the uint8 type.`,
// ====================================================================================
		"26":`26.float32, float64
Floating Point Numbers. e.g 3.14, 123.541, 100.4020401
These types represent numbers with a fraction. These types allocate 32 or 64
bits to store the value.
Literal Value Examples:
	20.2, -20.2, 1.2e10, 1.2e-10. Values can also be expressed in hex notation (0x2p10), although
	the exponent is expressed in decimal digits.`,
// ====================================================================================
		"27":`27.complex64, complex128
These types represent numbers that have real and imaginary components.
These types allocate 64 or 128 bits to store the value.`,
// ====================================================================================
		"28":`28.bool 
bolean
This type represents a Boolean truth with the values true and false.

Literal Value Examples:
	true, false.`,
// ====================================================================================
		"29":`29.string
This type represents a sequence of characters.
Literal Value Examples:
	"Hello". Character sequences escaped with a backslash are interpreted if the value is enclosed
	in double quotes ("Hello\n"). Escape sequences are not interpreted if the value is enclosed in
	backquotes ('Hello\n').`,
// ====================================================================================
		"30":`30.rune = int32
This type represents a single Unicode code point. Unicode is complicated,
but—loosely—this is the representation of a single character. The rune type is
an alias for int32.

Literal Value Examples:
	'A', '\n', '\u00A5', '¥'. Characters, glyphs, and escape sequences are enclosed in single
	quotes (the ' character).`,
// ====================================================================================
		"31":`31.var = variable 
Variables are defined using the var keyword, and, unlike constants, the value assigned to a variable can be
changed`,
// ====================================================================================
		"32":`32.const = constant
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
	rune        0`,
// ====================================================================================
		"33":`33.+, -, *, /, %
Operations and Conversions
These operators are used to perform arithmetic using numeric values.`,
// ====================================================================================
		"34":`34.==, !=, <, <=, >, >=
These operators compare two values.`,
// ====================================================================================
		"35":`35., &&, !
These are the logical operators, which are applied to bool values and return a bool value.`,
// ====================================================================================
		"36":`36.=, :=
These are the assignment operators. The standard assignment operator (=) is used to set
the initial value when a constant or variable is defined, or to change the value assigned to
a previously defined variable. The shorthand operator (:=) is used to define a variable and
assign a value.`,
// ====================================================================================
		"37":`37.-=, +=, ++, --
These operators increment and decrement numeric values.`,
// ====================================================================================
		"38":`38.&, , ^, &^, <<, >>
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
	||			This operator returns true if either operand is true. 
				If the first operand is true, then the second
				operand will not be evaluated.

	&&       This operator returns true if both operands are true. 
				If the first operand is false, then the
				second operand will not be evaluated.

	!        This operator is used with a single operand. 
				It returns true if the operand is false and false if
				the operand is true.`,
// ====================================================================================
		"39":`39.Ceil(value)
Converting Floating-Point Values to Integers
Functions in the math Package for Converting Numeric Types
This function returns the smallest integer that is greater than the specified floating-
point value. The smallest integer that is greater than 27.1, for example, is 28.`,
// ====================================================================================
		"40":`40.Floor(value)
This function returns the largest integer that is less than the specified floating-point
value. The largest integer that is less than 27.1, for example, is 28.`,
// ====================================================================================
		"41":`41.Round(value)
This function rounds the specified floating-point value to the nearest integer.`,
// ====================================================================================
		"42":`42.RoundToEven(value)
This function rounds the specified floating-point value to the nearest even integer.`,
// ====================================================================================
		"43":`43.ParseBool(str)
Parsing from Strings
Functions for Parsing Strings into Other Data Types
This function parses a string into a bool value. Recognized string values are "true",
"false", "TRUE", "FALSE", "True", "False", "T", "F", "0", and "1".`,
// ====================================================================================
		"44":`44.ParseFloat(str,size)
This function parses a string into a floating-point value with the specified size, as
described in the “Parsing Floating-Point Numbers” section.`,
// ====================================================================================
		"45":`45.ParseInt(str,base, size)
This function parses a string into an int64 with the specified base and size. Acceptable
base values are 2 for binary, 8 for octal, 16 for hex, and 10.`,
// ====================================================================================
		"46":`46.ParseUint(str,base, size)
This function parses a string into an unsigned integer value with the specified base and size.`,
// ====================================================================================
		"47":`47.Atoi(str)
This function parses a string into a base 10 int and is equivalent to calling
ParseInt(str, 10, 0)`,
// ====================================================================================
		"48":`48.FormatBool(val)
Formatting Values as Strings
The strconv Functions for Converting Values into Strings
This function returns the string true or false based on the value of the
specified bool.`,
// ====================================================================================
		"49":`49.FormatInt(val, base)
This function returns a string representation of the specified int64 value,
expressed in the specified base.`,
// ====================================================================================
		"50":`50.FormatUint(val, base)
This function returns a string representation of the specified uint64 value,
expressed in the specified base.`,
// ====================================================================================
		"51":`51.FormatFloat(val, format, precision, size) 
This function returns a string representation of the specified float64 value,
expressed using the specified format, precision, and size.`,
// ====================================================================================
		"52":`52.Itoa(val)
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
	smaller values.`,
// ====================================================================================
		"53":`53.if
Understanding Flow Control
The if keyword is followed by the expression and then the group of statements to be executed,
surrounded by braces`,
// ====================================================================================
		"54":`54.else
The else keyword can be used to create additional clauses in an if statement`,
// ====================================================================================
		"55":`55.else if
The else/if combination can be repeated to create a sequence of clauses`,
// ====================================================================================
		"56":`56.for
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
for name := range sliceOfNames`,
// ====================================================================================
		"57":`57.switch
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
	}`,
// ====================================================================================
		"58":`58.Falling Through
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
		}`,
// ====================================================================================
		"59":`59.default
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
		}`,
// ====================================================================================
		"60":`60.goto label 
Label statements allow execution to jump to a different point, 
giving greater flexibility than other flow control features
example:
	counter := 0
	target: fmt.Println("Counter", counter)
	counter++
	if (counter < 5) {
			goto target
	}`,
// ====================================================================================
		"61":`61.Array
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
	}`,
// ====================================================================================
		"62":`62.Slices
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
	names := []string {"Kayak", "Lifejacket", "Paddle"}`,
// ====================================================================================
		"63":`63.append
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
	.\main.go:13:30: invalid operation: p1 == p2 (slice can only be compared to nil)`,
// ====================================================================================
		"64":`64.copy
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
	Index: 1 Value: Hat`,
// ====================================================================================
		"65":`65.sort
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
	Index: 3 Value: Paddle`,
// ====================================================================================
		"66":`66.DeepEqual
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
	Equal: true`,
// ====================================================================================
		"67":`67.map
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
	}`,
// ====================================================================================
		"68":`68.strconv
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
	Length: 8`,
// ====================================================================================
		"69":`69.rune
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
	7 53`,
// ====================================================================================
		"70":``,
// ====================================================================================
		"71":``,
// ====================================================================================
		"72":``,
// ====================================================================================
		"73":``,
// ====================================================================================
		"74":``,
// ====================================================================================
		"75":``,
// ====================================================================================
		"76":``,
// ====================================================================================
		"77":``,
// ====================================================================================
		"78":``,
// ====================================================================================
		"79":``,
// ====================================================================================
		"80":``,
// ====================================================================================
		"81":``,
// ====================================================================================
		"82":``,
// ====================================================================================
		"83":``,
// ====================================================================================
		"84":``,
// ====================================================================================
		"85":``,
// ====================================================================================
		"86":``,
// ====================================================================================
		"87":``,
// ====================================================================================
		"88":``,
// ====================================================================================
		"89":``,
// ====================================================================================
		"90":``,
// ====================================================================================
		"91":``,
// ====================================================================================
		"92":``,
// ====================================================================================
		"93":``,
// ====================================================================================
		"94":``,
// ====================================================================================
		"95":``,
// ====================================================================================
		"96":``,
// ====================================================================================
		"97":``,
// ====================================================================================
		"98":``,
// ====================================================================================
		"99":``,
// ====================================================================================
		"100":``,
// ====================================================================================
		"101":``,
// ====================================================================================
		"102":``,
// ====================================================================================
		"103":``,
// ====================================================================================
		"104":``,
// ====================================================================================
		"105":``,
// ====================================================================================
		"106":``,
// ====================================================================================
		"107":``,
// ====================================================================================
		"108":``,
// ====================================================================================
		"109":``,
// ====================================================================================
		"110":``,
// ====================================================================================
		"111":``,
// ====================================================================================
		"112":``,
// ====================================================================================
		"113":``,
// ====================================================================================
		"114":``,
// ====================================================================================
		"115":``,
// ====================================================================================
		"116":``,
// ====================================================================================
		"117":``,
// ====================================================================================
		"118":``,
// ====================================================================================
		"119":``,
// ====================================================================================
		"120":``,
// ====================================================================================
		"121":``,
// ====================================================================================
		"122":``,
// ====================================================================================
		"123":``,
// ====================================================================================
		"124":``,
// ====================================================================================
		"125":``,
// ====================================================================================
		"126":``,
// ====================================================================================
		"127":``,
// ====================================================================================
		"128":``,
// ====================================================================================
		"129":``,
// ====================================================================================
		"130":``,
// ====================================================================================
		"131":``,
// ====================================================================================
		"132":``,
// ====================================================================================
		"133":``,
// ====================================================================================
		"134":``,
// ====================================================================================
		"135":``,
// ====================================================================================
		"136":``,
// ====================================================================================
		"137":``,
// ====================================================================================
		"138":``,
// ====================================================================================
		"139":``,
// ====================================================================================
		"140":``,
// ====================================================================================
		"141":``,
// ====================================================================================
		"142":``,
// ====================================================================================
		"143":``,
// ====================================================================================
		"144":``,
// ====================================================================================
		"145":``,
// ====================================================================================
		"146":``,
// ====================================================================================
		"147":``,
// ====================================================================================
		"148":``,
// ====================================================================================
		"149":``,
// ====================================================================================
		"150":``,
// ====================================================================================
		"151":``,
// ====================================================================================
		"152":``,
// ====================================================================================
		"153":``,
// ====================================================================================
		"154":``,
// ====================================================================================
		"155":``,
// ====================================================================================
		"156":``,
// ====================================================================================
		"157":``,
// ====================================================================================
		"158":``,
// ====================================================================================
		"159":``,
// ====================================================================================
		"160":``,
// ====================================================================================
		"161":``,
// ====================================================================================
		"162":``,
// ====================================================================================
		"163":``,
// ====================================================================================
		"164":``,
// ====================================================================================
		"165":``,
// ====================================================================================
		"166":``,
// ====================================================================================
		"167":``,
// ====================================================================================
		"168":``,
// ====================================================================================
		"169":``,
// ====================================================================================
		"170":``,
// ====================================================================================
		"171":``,
// ====================================================================================
		"172":``,
// ====================================================================================
		"173":``,
// ====================================================================================
		"174":``,
// ====================================================================================
		"175":``,
// ====================================================================================
		"176":``,
// ====================================================================================
		"177":``,
// ====================================================================================
		"178":``,
// ====================================================================================
		"179":``,
// ====================================================================================
		"180":``,
// ====================================================================================
		"181":``,
// ====================================================================================
		"182":``,
// ====================================================================================
		"183":``,
// ====================================================================================
		"184":``,
// ====================================================================================
		"185":``,
// ====================================================================================
		"186":``,
// ====================================================================================
		"187":``,
// ====================================================================================
		"188":``,
// ====================================================================================
		"189":``,
// ====================================================================================
		"190":``,
// ====================================================================================
		"191":``,
// ====================================================================================
		"192":``,
// ====================================================================================
		"193":``,
// ====================================================================================
		"194":``,
// ====================================================================================
		"195":``,
// ====================================================================================
		"196":``,
// ====================================================================================
		"197":``,
// ====================================================================================
		"198":``,
// ====================================================================================
		"199":``,
// ====================================================================================
		"200":``,
// ====================================================================================
		"201":``,
// ====================================================================================
		"202":``,
// ====================================================================================
		"203":``,
// ====================================================================================
		"204":``,
// ====================================================================================
		"205":``,
// ====================================================================================
		"206":``,
// ====================================================================================
		"207":``,
// ====================================================================================
		"208":``,
// ====================================================================================
		"209":``,
// ====================================================================================
		"210":``,
// ====================================================================================
		"211":``,
// ====================================================================================
		"212":``,
// ====================================================================================
		"213":``,
// ====================================================================================
		"214":``,
// ====================================================================================
		"215":``,
// ====================================================================================
		"216":``,
// ====================================================================================
		"217":``,
// ====================================================================================
		"218":``,
// ====================================================================================
		"219":``,
// ====================================================================================
		"220":``,
// ====================================================================================
		"221":``,
// ====================================================================================
		"222":``,
// ====================================================================================
		"223":``,
// ====================================================================================
		"224":``,
// ====================================================================================
		"225":``,
// ====================================================================================
		"226":``,
// ====================================================================================
		"227":``,
// ====================================================================================
		"228":``,
// ====================================================================================
		"229":``,
// ====================================================================================
		"230":``,
// ====================================================================================
		"231":``,
// ====================================================================================
		"232":``,
// ====================================================================================
		"233":``,
// ====================================================================================
		"234":``,
// ====================================================================================
		"235":``,
// ====================================================================================
		"236":``,
// ====================================================================================
		"237":``,
// ====================================================================================
		"238":``,
// ====================================================================================
		"239":``,
// ====================================================================================
		"240":``,
// ====================================================================================
		"241":``,
// ====================================================================================
		"242":``,
// ====================================================================================
		"243":``,
// ====================================================================================
		"244":``,
// ====================================================================================
		"245":``,
// ====================================================================================
		"246":``,
// ====================================================================================
		"247":``,
// ====================================================================================
		"248":``,
// ====================================================================================
		"249":``,
// ====================================================================================
		"250":``,
// ====================================================================================
		"251":``,
// ====================================================================================
		"252":``,
// ====================================================================================
		"253":``,
// ====================================================================================
		"254":``,
// ====================================================================================
		"255":``,
// ====================================================================================
		"256":``,
// ====================================================================================
		"257":``,
// ====================================================================================
		"258":``,
// ====================================================================================
		"259":``,
// ====================================================================================
		"260":``,
// ====================================================================================
		"261":``,
// ====================================================================================
		"262":``,
// ====================================================================================
		"263":``,
// ====================================================================================
		"264":``,
// ====================================================================================
		"265":``,
// ====================================================================================
		"266":``,
// ====================================================================================
		"267":``,
// ====================================================================================
		"268":``,
// ====================================================================================
		"269":``,
// ====================================================================================
		"270":``,
// ====================================================================================
		"271":``,
// ====================================================================================
		"272":``,
// ====================================================================================
		"273":``,
// ====================================================================================
		"274":``,
// ====================================================================================
		"275":``,
// ====================================================================================
		"276":``,
// ====================================================================================
		"277":``,
// ====================================================================================
		"279":``,
// ====================================================================================
		"280":``,
// ====================================================================================
		"281":``,
// ====================================================================================
		"282":``,
// ====================================================================================
		"283":``,
// ====================================================================================
		"284":``,
// ====================================================================================
		"285":``,
// ====================================================================================
		"286":``,
// ====================================================================================
		"287":``,
// ====================================================================================
		"288":``,
// ====================================================================================
		"289":``,
// ====================================================================================
		"290":``,
// ====================================================================================
		"291":``,
// ====================================================================================
		"292":``,
// ====================================================================================
		"293":``,
// ====================================================================================
		"294":``,
// ====================================================================================
		"295":``,
// ====================================================================================
		"296":``,
// ====================================================================================
		"297":``,
// ====================================================================================
		"298":``,
// ====================================================================================
		"299":``,
// ====================================================================================
		"300":``,
// ====================================================================================
		"301":``,
// ====================================================================================
		"302":``,
// ====================================================================================
		"303":``,
// ====================================================================================
		"304":``,
// ====================================================================================
		"305":``,
// ====================================================================================
		"306":``,
// ====================================================================================
		"307":``,
// ====================================================================================
		"308":``,
// ====================================================================================
		"309":``,
// ====================================================================================
		"310":``,
// ====================================================================================
		"311":``,
// ====================================================================================
		"312":``,
// ====================================================================================
		"313":``,
// ====================================================================================
		"314":``,
// ====================================================================================
		"315":``,
// ====================================================================================
		"316":``,
// ====================================================================================
		"317":``,
// ====================================================================================
		"318":``,
// ====================================================================================
		"319":``,
// ====================================================================================
		"320":``,
// ====================================================================================
		"321":``,
// ====================================================================================
		"322":``,
// ====================================================================================
		"323":``,
// ====================================================================================
		"324":``,
// ====================================================================================
		"325":``,
// ====================================================================================
		"326":``,
// ====================================================================================
		"327":``,
// ====================================================================================
		"328":``,
// ====================================================================================
		"329":``,
// ====================================================================================
		"330":``,
// ====================================================================================
		"331":``,
// ====================================================================================
		"332":``,
// ====================================================================================
		"333":``,
// ====================================================================================
		"334":``,
// ====================================================================================
		"335":``,
// ====================================================================================
		"336":``,
// ====================================================================================
		"337":``,
// ====================================================================================
		"338":``,
// ====================================================================================
		"339":``,
// ====================================================================================
		"340":``,
// ====================================================================================
		"341":``,
// ====================================================================================
		"342":``,
// ====================================================================================
		"343":``,
// ====================================================================================
		"344":``,
// ====================================================================================
		"345":``,
// ====================================================================================
		"346":``,
// ====================================================================================
		"347":``,
// ====================================================================================
		"348":``,
// ====================================================================================
		"349":``,
// ====================================================================================
		"350":``,
// ====================================================================================
		"351":``,
// ====================================================================================
		"352":``,
// ====================================================================================
		"353":``,
// ====================================================================================
		"354":``,
// ====================================================================================
		"355":``,
// ====================================================================================
		"356":``,
// ====================================================================================
		"357":``,
// ====================================================================================
		"358":``,
// ====================================================================================
		"359":``,
// ====================================================================================
		"360":``,
// ====================================================================================
		"361":``,
// ====================================================================================
		"362":``,
// ====================================================================================
		"363":``,
// ====================================================================================
		"364":``,
// ====================================================================================
		"365":``,
// ====================================================================================
		"366":``,
// ====================================================================================
		"367":``,
// ====================================================================================
		"368":``,
// ====================================================================================
		"369":``,
// ====================================================================================
		"370":``,
// ====================================================================================
		"371":``,
// ====================================================================================
		"372":``,
// ====================================================================================
		"373":``,
// ====================================================================================
		"374":``,
// ====================================================================================
		"375":``,
// ====================================================================================
		"376":``,
// ====================================================================================
		"377":``,
// ====================================================================================
		"378":``,
// ====================================================================================
		"379":``,
// ====================================================================================
		"380":``,
// ====================================================================================
		"381":``,
// ====================================================================================
		"382":``,
// ====================================================================================
		"383":``,
// ====================================================================================
		"384":``,
// ====================================================================================
		"385":``,
// ====================================================================================
		"386":``,
// ====================================================================================
		"387":``,
// ====================================================================================
		"388":``,
// ====================================================================================
		"389":``,
// ====================================================================================
		"390":``,
// ====================================================================================
		"391":``,
// ====================================================================================
		"392":``,
// ====================================================================================
		"393":``,
// ====================================================================================
		"394":``,
// ====================================================================================
		"395":``,
// ====================================================================================
		"396":``,
// ====================================================================================
		"397":``,
// ====================================================================================
		"398":``,
// ====================================================================================
		"399":``,
// ====================================================================================
		"400":``,
// ====================================================================================
		"401":``,
// ====================================================================================
		"402":``,
// ====================================================================================
		"403":``,
// ====================================================================================
		"404":``,
// ====================================================================================
		"405":``,
// ====================================================================================
		"406":``,
// ====================================================================================
		"407":``,
// ====================================================================================
		"408":``,
// ====================================================================================
		"409":``,
// ====================================================================================
		"410":``,
// ====================================================================================
		"411":``,
// ====================================================================================
		"412":``,
// ====================================================================================
		"413":``,
// ====================================================================================
		"414":``,
// ====================================================================================
		"415":``,
// ====================================================================================
		"416":``,
// ====================================================================================
		"417":``,
// ====================================================================================
		"418":``,
// ====================================================================================
		"419":``,
// ====================================================================================
		"420":``,
// ====================================================================================
		"421":``,
// ====================================================================================
		"422":``,
// ====================================================================================
		"423":``,
// ====================================================================================
		"424":``,
// ====================================================================================
		"425":``,
// ====================================================================================
		"426":``,
// ====================================================================================
		"427":``,
// ====================================================================================
		"428":``,
// ====================================================================================
		"429":``,
// ====================================================================================
		"430":``,
// ====================================================================================
		"431":``,
// ====================================================================================
		"432":``,
// ====================================================================================
		"433":``,
// ====================================================================================
		"434":``,
// ====================================================================================
		"435":``,
// ====================================================================================
		"436":``,
// ====================================================================================
		"437":``,
// ====================================================================================
		"438":``,
// ====================================================================================
		"439":``,
// ====================================================================================
		"440":``,
// ====================================================================================
		"441":``,
// ====================================================================================
		"442":``,
// ====================================================================================
		"443":``,
// ====================================================================================
		"444":``,
// ====================================================================================
		"445":``,
// ====================================================================================
		"446":``,
// ====================================================================================
		"447":``,
// ====================================================================================
		"448":``,
// ====================================================================================
		"449":``,
// ====================================================================================
		"450":``,
// ====================================================================================
		"451":``,
// ====================================================================================
		"452":``,
// ====================================================================================
		"453":``,
// ====================================================================================
		"454":``,
// ====================================================================================
		"455":``,
// ====================================================================================
		"456":``,
// ====================================================================================
		"457":``,
// ====================================================================================
		"458":``,
// ====================================================================================
		"459":``,
// ====================================================================================
		"460":``,
// ====================================================================================
		"461":``,
// ====================================================================================
		"462":``,
// ====================================================================================
		"463":``,
// ====================================================================================
		"464":``,
// ====================================================================================
		"465":``,
// ====================================================================================
		"466":``,
// ====================================================================================
		"467":``,
// ====================================================================================
		"468":``,
// ====================================================================================
		"469":``,
// ====================================================================================
		"470":``,
// ====================================================================================
		"471":``,
// ====================================================================================
		"472":``,
// ====================================================================================
		"473":``,
// ====================================================================================
		"474":``,
// ====================================================================================
		"475":``,
// ====================================================================================
		"476":``,
// ====================================================================================
		"477":``,
// ====================================================================================
		"478":``,
// ====================================================================================
		"479":``,
// ====================================================================================
		"480":``,
// ====================================================================================
		"481":``,
// ====================================================================================
		"482":``,
// ====================================================================================
		"483":``,
// ====================================================================================
		"484":``,
// ====================================================================================
		"485":``,
// ====================================================================================
		"486":``,
// ====================================================================================
		"487":``,
// ====================================================================================
		"488":``,
// ====================================================================================
		"489":``,
// ====================================================================================
		"490":``,
// ====================================================================================
		"491":``,
// ====================================================================================
		"492":``,
// ====================================================================================
		"493":``,
// ====================================================================================
		"494":``,
// ====================================================================================
		"495":``,
// ====================================================================================
		"496":``,
// ====================================================================================
		"497":``,
// ====================================================================================
		"498":``,
// ====================================================================================
		"499":``,
// ====================================================================================



	},
}