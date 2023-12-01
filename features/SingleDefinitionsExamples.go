package features

type SingleDefinitionExamples struct {
	MapSingleDefEx map[string]string
}

var OriginalSingleDefExamples = SingleDefinitionExamples{
	MapSingleDefEx: map[string]string{
		"for": `56.for
=================
example-1 
------------
main.go
	package main

	import (
		"fmt"
	)

	func main() {

		counter := "EXAMPLE"

		// Enumerating Sequences:
		for index, character := range counter {
			fmt.Println("Index:", index, "Character:", string(character))
		}

	}
===========================
in Terminal: go run .
===========================
Output:
	Index: 0 Character: E
	Index: 1 Character: X
	Index: 2 Character: A
	Index: 3 Character: M
	Index: 4 Character: P
	Index: 5 Character: L
	Index: 6 Character: E`,

		"http servers": `403.Creating HTTP Servers
=================
example-1 
------------
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
4- in -> 404.Creating a Simple HTTP Server
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
`,
		"simple server": `403.Creating HTTP Servers
=================
example-1 
------------
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
4- in -> 404.Creating a Simple HTTP Server
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
`,
		"web application": `
main.go:
	package main
	import (
		"fmt"
		"net/http"
	)
	func handlerFunction(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "<h1>Welcome to AcronProject Site</h1>")
	}
	func main() {
		http.HandleFunc("/", handlerFunction)
		fmt.Println("Starting the Web Server on http://locanhost:3000")
		http.ListenAndServe(":3000", nil)
	}
========================================
Output: go run main.go
	in Web Browser, search localhost:3000
	Welcome to AcronProject Site
`,
		"simple web application": `main.go:
	package main
	import (
		"fmt"
		"net/http"
	)
	func handlerFunction(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "<h1>Welcome to AcronProject Site</h1>")
	}
	func main() {
		http.HandleFunc("/", handlerFunction)
		fmt.Println("Starting the Web Server on http://locanhost:3000")
		http.ListenAndServe(":3000", nil)
	}
========================================
Output: go run main.go
	in Web Browser, search localhost:3000
	Welcome to AcronProject Site

	`,
		"simple web app": `main.go:
	package main
	import (
		"fmt"
		"net/http"
	)
	func handlerFunction(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "<h1>Welcome to AcronProject Site</h1>")
	}
	func main() {
		http.HandleFunc("/", handlerFunction)
		fmt.Println("Starting the Web Server on http://locanhost:3000")
		http.ListenAndServe(":3000", nil)
	}
========================================
Output: go run main.go
	in Web Browser, search localhost:3000
	Welcome to AcronProject Site
		`,
		"bool": `Boolean
var closed bool    // boolean variable 'closed' implicitly initialized with 'false'
speeding := true   // boolean variable 'speeding' initialized with 'true'
hasError := false  // boolean variable 'hasError' initialized with 'false' 
`,
		"boolean": `Boolean
var closed bool    // boolean variable 'closed' implicitly initialized with 'false'
speeding := true   // boolean variable 'speeding' initialized with 'true'
hasError := false  // boolean variable 'hasError' initialized with 'false' 
`,
		"booleans": `Boolean
var closed bool    // boolean variable 'closed' implicitly initialized with 'false'
speeding := true   // boolean variable 'speeding' initialized with 'true'
hasError := false  // boolean variable 'hasError' initialized with 'false' 
`,
		"package comments": `kelvin.go
// Package kelvin provides tools to convert
// temperatures to and from Kelvin.
package kelvin
`,
		"package comment": `kelvin.go
// Package kelvin provides tools to convert
// temperatures to and from Kelvin.
package kelvin`,
		"function comment": `example:
// CelsiusFreezingTemp returns an integer value equal to the temperature at which water freezes in degrees Celsius.
func CelsiusFreezingTemp() int {
	return 0
}`,
		"function comments": `example:
// CelsiusFreezingTemp returns an integer value equal to the temperature at which water freezes in degrees Celsius.
func CelsiusFreezingTemp() int {
	return 0
}`,
		"operator": `Operator 	Example
--------	-------------
+ 		4 + 6 == 10
- 		15 - 10 == 5
* 		2 * 3 == 6
/ 		13 / 3 == 4
% 		13 % 3 == 1`,
		"operators": `Operator 	Example
--------	-------------
+ 		4 + 6 == 10
- 		15 - 10 == 5
* 		2 * 3 == 6
/ 		13 / 3 == 4
% 		13 % 3 == 1`,
		"arithmetic operators": `Operator 	Example
--------	-------------
+ 		4 + 6 == 10
- 		15 - 10 == 5
* 		2 * 3 == 6
/ 		13 / 3 == 4
% 		13 % 3 == 1`,
		"converting between types": `var x int = 42 // x has type int
f := float64(x) // f has type float64 (ie. 42.0)
var y float64 = 11.9 // y has type float64
i := int(y) // i has type int (ie. 11)`,
		"arithmetic operations on different types": `var x int = 42

// this line produces an error
value := float32(2.0) * x // invalid operation: mismatched types float32 and int

// you must convert int type to float32 before performing arithmetic operation
value := float32(2.0) * float32(x)`,

		"arithmetic operations on different type": `var x int = 42

// this line produces an error
value := float32(2.0) * x // invalid operation: mismatched types float32 and int

// you must convert int type to float32 before performing arithmetic operation
value := float32(2.0) * float32(x)`,

		"godoc": `An Example Program With Godoc Comments
The code below adheres to the Go way, in this case using single-line comments.
===================================

`,
		"string": `A string literal is defined between double quotes:

	const name = "Jane"
	
==================================================	
Strings can be concatenated via the + operator:

	"Jane" + " " + "Austen"
	// => "Jane Austen"
	
==================================================	
`,

		"beego": `Quick Start
Create hello directory, cd hello directory
	mkdir hello
	cd hello

Init module
	go mod init

Download and install
	go get github.com/beego/beego

Create file hello.go
	package main
	import "github.com/beego/beego"
	func main(){
		beego.Run()
	}

Build and run
	go build hello.go
	./hello
`,

		"beego1": `beego-2
this is first example, you can write => beggo or beego1 or beego2 and more for another examples:
Quick Start

Create hello directory, cd hello directory
	mkdir hello
	cd hello

Init module
	go mod init

Download and install
	go get github.com/beego/beego

Create file hello.go
	package main
	import "github.com/beego/beego"
	func main(){
		beego.Run()
	}

Build and run
	go build hello.go
	./hello
`,

		"beego framework": `Quick Start
Create hello directory, cd hello directory
	mkdir hello
	cd hello

Init module
	go mod init

Download and install
	go get github.com/beego/beego

Create file hello.go
	package main
	import "github.com/beego/beego"
	func main(){
		beego.Run()
	}

Build and run
	go build hello.go
	./hello
`,

		"generics": `main.go:
package main

import "fmt"

type CustomMap[T comparable, V int | string] map[T]V

func main() {
	Printfln("Hello, Data")

	obj := make(CustomMap[int, string])
	obj[3] = "3"
	obj[1] = "1"

	fmt.Println(obj)
}`,

		"generic": `main.go:
package main

import "fmt"

type CustomMap[T comparable, V int | string] map[T]V

func main() {
	Printfln("Hello, Data")

	obj := make(CustomMap[int, string])
	obj[3] = "3"
	obj[1] = "1"

	fmt.Println(obj)
}`,

		"reflection": `example:
package main
func printDetails(values ...interface{}) {
	for _, elem := range values {
		switch val := elem.(type) {
			case Product:
				Printfln("Product: Name: %v, Category: %v, Price: %v",
					val.Name, val.Category, val.Price)
			case Customer:
				Printfln("Customer: Name: %v, City: %v", val.Name, val.City)
		}
	}
}
func main() {
	product := Product {
		Name: "Kayak", Category: "Watersports", Price: 279,
	}
	customer := Customer { Name: "Alice", City: "New York" }
	printDetails(product, customer)
}
====================================================================
Output:
Product: Name: Kayak, Category: Watersports, Price: 279
Customer: Name: Alice, City: New York`,

		"reflection ": `example:
package main
func printDetails(values ...interface{}) {
	for _, elem := range values {
		switch val := elem.(type) {
			case Product:
				Printfln("Product: Name: %v, Category: %v, Price: %v",
					val.Name, val.Category, val.Price)
			case Customer:
				Printfln("Customer: Name: %v, City: %v", val.Name, val.City)
		}
	}
}
func main() {
	product := Product {
		Name: "Kayak", Category: "Watersports", Price: 279,
	}
	customer := Customer { Name: "Alice", City: "New York" }
	printDetails(product, customer)
}
====================================================================
Output:
Product: Name: Kayak, Category: Watersports, Price: 279
Customer: Name: Alice, City: New York`,

		"reflection 1": `example:
package main
func printDetails(values ...interface{}) {
	for _, elem := range values {
		switch val := elem.(type) {
			case Product:
				Printfln("Product: Name: %v, Category: %v, Price: %v",
					val.Name, val.Category, val.Price)
			case Customer:
				Printfln("Customer: Name: %v, City: %v", val.Name, val.City)
		}
	}
}
func main() {
	product := Product {
		Name: "Kayak", Category: "Watersports", Price: 279,
	}
	customer := Customer { Name: "Alice", City: "New York" }
	printDetails(product, customer)
}
====================================================================
Output:
Product: Name: Kayak, Category: Watersports, Price: 279
Customer: Name: Alice, City: New York`,

		"reflection1": `example:
package main
func printDetails(values ...interface{}) {
	for _, elem := range values {
		switch val := elem.(type) {
			case Product:
				Printfln("Product: Name: %v, Category: %v, Price: %v",
					val.Name, val.Category, val.Price)
			case Customer:
				Printfln("Customer: Name: %v, City: %v", val.Name, val.City)
		}
	}
}
func main() {
	product := Product {
		Name: "Kayak", Category: "Watersports", Price: 279,
	}
	customer := Customer { Name: "Alice", City: "New York" }
	printDetails(product, customer)
}
====================================================================
Output:
Product: Name: Kayak, Category: Watersports, Price: 279
Customer: Name: Alice, City: New York`,

		"server": `server1 = you can use server2 example, server3 or more 
main.go:
	package main
	import (
		"net/http"
		"os"
	)
	func main()  {
		dir, _ := os.Getwd()
		http.ListenAndServe(":3000", http.FileServer(http.Dir(dir)))
	}
`,

		"server1": `server1 = you can use server2 example, server3 or more
main.go:
	package main
	import (
		"net/http"
		"os"
	)
	func main()  {
		dir, _ := os.Getwd()
		http.ListenAndServe(":3000", http.FileServer(http.Dir(dir)))
	}
`,

		"beego2": `beego-2
this is first example, you can write => beggo or beego1 or beego2 and more for another examples:
=========================================================================
main.go:
	package main
	import (
		"github.com/beego/beego/v2/server/web"
	)
	func main() {
		web.Run()
	}
========================================================================
go.mod:
	module github.com/sinalalebakhsh/WebApplication1
	go 1.21.3
	require github.com/beego/beego/v2 v2.1.3
	require (
		github.com/beorn7/perks v1.0.1 // indirect
		github.com/cespare/xxhash/v2 v2.2.0 // indirect
		github.com/golang/protobuf v1.5.3 // indirect
		github.com/hashicorp/golang-lru v0.5.4 // indirect
		github.com/kr/text v0.2.0 // indirect
		github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
		github.com/mitchellh/mapstructure v1.5.0 // indirect
		github.com/pkg/errors v0.9.1 // indirect
		github.com/prometheus/client_golang v1.16.0 // indirect
		github.com/prometheus/client_model v0.3.0 // indirect
		github.com/prometheus/common v0.42.0 // indirect
		github.com/prometheus/procfs v0.10.1 // indirect
		github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
		golang.org/x/crypto v0.10.0 // indirect
		golang.org/x/net v0.10.0 // indirect
		golang.org/x/sys v0.9.0 // indirect
		golang.org/x/text v0.10.0 // indirect
		google.golang.org/protobuf v1.30.0 // indirect
		gopkg.in/yaml.v3 v3.0.1 // indirect
	)
=========================================================================
go.sum:
	github.com/beego/beego/v2 v2.1.3 h1:x436yz6jrSasYBzfOP39S097kvq5/5fBTFfEvVA456M=
	github.com/beego/beego/v2 v2.1.3/go.mod h1:0J0RQVIpepnRUfu6ax+kLVVB1FcdYryHK9lpRl5wvbY=
	github.com/beorn7/perks v1.0.1 h1:VlbKKnNfV8bJzeqoa4cOKqO6bYr3WgKZxO8Z16+hsOM=
	github.com/beorn7/perks v1.0.1/go.mod h1:G2ZrVWU2WbWT9wwq4/hrbKbnv/1ERSJQ0ibhJ6rlkpw=
	github.com/cespare/xxhash/v2 v2.2.0 h1:DC2CZ1Ep5Y4k3ZQ899DldepgrayRUGE6BBZ/cd9Cj44=
	github.com/cespare/xxhash/v2 v2.2.0/go.mod h1:VGX0DQ3Q6kWi7AoAeZDth3/j3BFtOZR5XLFGgcrjCOs=
	github.com/creack/pty v1.1.9/go.mod h1:oKZEueFk5CKHvIhNR5MUki03XCEU+Q6VDXinZuGJ33E=
	github.com/davecgh/go-spew v1.1.1 h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=
	github.com/davecgh/go-spew v1.1.1/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
	github.com/elazarl/go-bindata-assetfs v1.0.1 h1:m0kkaHRKEu7tUIUFVwhGGGYClXvyl4RE03qmvRTNfbw=
	github.com/elazarl/go-bindata-assetfs v1.0.1/go.mod h1:v+YaWX3bdea5J/mo8dSETolEo7R71Vk1u8bnjau5yw4=
	github.com/golang/protobuf v1.2.0/go.mod h1:6lQm79b+lXiMfvg/cZm0SGofjICqVBUtrP5yJMmIC1U=
	github.com/golang/protobuf v1.3.5/go.mod h1:6O5/vntMXwX2lRkT1hjjk0nAC1IDOTvTlVgjlRvqsdk=
	github.com/golang/protobuf v1.5.0/go.mod h1:FsONVRAS9T7sI+LIUmWTfcYkHO4aIWwzhcaSAoJOfIk=
	github.com/golang/protobuf v1.5.3 h1:KhyjKVUg7Usr/dYsdSqoFveMYd5ko72D+zANwlG1mmg=
	github.com/golang/protobuf v1.5.3/go.mod h1:XVQd3VNwM+JqD3oG2Ue2ip4fOMUkwXdXDdiuN0vRsmY=
	github.com/google/go-cmp v0.5.5/go.mod h1:v8dTdLbMG2kIc/vJvl+f65V22dbkXbowE6jgT/gNBxE=
	github.com/google/go-cmp v0.5.9 h1:O2Tfq5qg4qc4AmwVlvv0oLiVAGB7enBSJ2x2DqQFi38=
	github.com/google/go-cmp v0.5.9/go.mod h1:17dUlkBOakJ0+DkrSSNjCkIjxS6bF9zb3elmeNGIjoY=
	github.com/hashicorp/golang-lru v0.5.4 h1:YDjusn29QI/Das2iO9M0BHnIbxPeyuCHsjMW+lJfyTc=
	github.com/hashicorp/golang-lru v0.5.4/go.mod h1:iADmTwqILo4mZ8BN3D2Q6+9jd8WM5uGBxy+E8yxSoD4=
	github.com/kr/pretty v0.3.1 h1:flRD4NNwYAUpkphVc1HcthR4KEIFJ65n8Mw5qdRn3LE=
	github.com/kr/pretty v0.3.1/go.mod h1:hoEshYVHaxMs3cyo3Yncou5ZscifuDolrwPKZanG3xk=
	github.com/kr/text v0.2.0 h1:5Nx0Ya0ZqY2ygV366QzturHI13Jq95ApcVaJBhpS+AY=
	github.com/kr/text v0.2.0/go.mod h1:eLer722TekiGuMkidMxC/pM04lWEeraHUUmBw8l2grE=
	github.com/matttproud/golang_protobuf_extensions v1.0.4 h1:mmDVorXM7PCGKw94cs5zkfA9PSy5pEvNWRP0ET0TIVo=
	github.com/matttproud/golang_protobuf_extensions v1.0.4/go.mod h1:BSXmuO+STAnVfrANrmjBb36TMTDstsz7MSK+HVaYKv4=
	github.com/mitchellh/mapstructure v1.5.0 h1:jeMsZIYE/09sWLaz43PL7Gy6RuMjD2eJVyuac5Z2hdY=
	github.com/mitchellh/mapstructure v1.5.0/go.mod h1:bFUtVrKA4DC2yAKiSyO/QUcy7e+RRV2QTWOzhPopBRo=
	github.com/pkg/errors v0.9.1 h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=
	github.com/pkg/errors v0.9.1/go.mod h1:bwawxfHBFNV+L2hUp1rHADufV3IMtnDRdf1r5NINEl0=
	github.com/pmezard/go-difflib v1.0.0 h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=
	github.com/pmezard/go-difflib v1.0.0/go.mod h1:iKH77koFhYxTK1pcRnkKkqfTogsbg7gZNVY4sRDYZ/4=
	github.com/prometheus/client_golang v1.16.0 h1:yk/hx9hDbrGHovbci4BY+pRMfSuuat626eFsHb7tmT8=
	github.com/prometheus/client_golang v1.16.0/go.mod h1:Zsulrv/L9oM40tJ7T815tM89lFEugiJ9HzIqaAx4LKc=
	github.com/prometheus/client_model v0.3.0 h1:UBgGFHqYdG/TPFD1B1ogZywDqEkwp3fBMvqdiQ7Xew4=
	github.com/prometheus/client_model v0.3.0/go.mod h1:LDGWKZIo7rky3hgvBe+caln+Dr3dPggB5dvjtD7w9+w=
	github.com/prometheus/common v0.42.0 h1:EKsfXEYo4JpWMHH5cg+KOUWeuJSov1Id8zGR8eeI1YM=
	github.com/prometheus/common v0.42.0/go.mod h1:xBwqVerjNdUDjgODMpudtOMwlOwf2SaTr1yjz4b7Zbc=
	github.com/prometheus/procfs v0.10.1 h1:kYK1Va/YMlutzCGazswoHKo//tZVlFpKYh+PymziUAg=
	github.com/prometheus/procfs v0.10.1/go.mod h1:nwNm2aOCAYw8uTR/9bWRREkZFxAUcWzPHWJq+XBB/FM=
	github.com/rogpeppe/go-internal v1.10.0 h1:TMyTOH3F/DB16zRVcYyreMH6GnZZrwQVAoYjRBZyWFQ=
	github.com/rogpeppe/go-internal v1.10.0/go.mod h1:UQnix2H7Ngw/k4C5ijL5+65zddjncjaFoBhdsK/akog=
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 h1:DAYUYH5869yV94zvCES9F51oYtN5oGlwjxJJz7ZCnik=
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18/go.mod h1:nkxAfR/5quYxwPZhyDxgasBMnRtBZd0FCEpawpjMUFg=
	github.com/stretchr/testify v1.8.1 h1:w7B6lhMri9wdJUVmEZPGGhZzrYTPvgJArz7wNPgYKsk=
	github.com/stretchr/testify v1.8.1/go.mod h1:w2LPCIKwWwSfY2zedu0+kehJoqGctiVI29o6fzry7u4=
	golang.org/x/crypto v0.10.0 h1:LKqV2xt9+kDzSTfOhx4FrkEBcMrAgHSYgzywV9zcGmM=
	golang.org/x/crypto v0.10.0/go.mod h1:o4eNf7Ede1fv+hwOwZsTHl9EsPFO6q6ZvYR8vYfY45I=
	golang.org/x/net v0.10.0 h1:X2//UzNDwYmtCLn7To6G58Wr6f5ahEAQgKNzv9Y951M=
	golang.org/x/net v0.10.0/go.mod h1:0qNGK6F8kojg2nk9dLZ2mShWaEBan6FAoqfSigmmuDg=
	golang.org/x/sync v0.0.0-20181221193216-37e7f081c4d4/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
	golang.org/x/sys v0.9.0 h1:KS/R3tvhPqvJvwcKfnBHJwwthS11LRhmM5D59eEXa0s=
	golang.org/x/sys v0.9.0/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
	golang.org/x/text v0.10.0 h1:UpjohKhiEgNc0CSauXmwYftY1+LlaC75SJwh0SgCX58=
	golang.org/x/text v0.10.0/go.mod h1:TvPlkZtksWOMsz7fbANvkp4WM8x/WCo/om8BMLbz+aE=
	golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543/go.mod h1:I/5z698sn9Ka8TeJc9MKroUUfqBBauWjQqLJ2OPfmY0=
	google.golang.org/protobuf v1.26.0-rc.1/go.mod h1:jlhhOSvTdKEhbULTjvd4ARK9grFBp09yW+WbY/TyQbw=
	google.golang.org/protobuf v1.26.0/go.mod h1:9q0QmTI4eRPtz6boOQmLYwt+qCgq0jsYwAQnmE0givc=
	google.golang.org/protobuf v1.30.0 h1:kPPoIgf3TsEvrm0PFe15JQ+570QVxYzEvvHqChK+cng=
	google.golang.org/protobuf v1.30.0/go.mod h1:HV8QOd/L58Z+nl8r43ehVNZIU/HEI6OcFqwMG9pJV4I=
	gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c h1:Hei/4ADfdWqJk1ZMxUNpqntNwaWcugrBjAiHlqqRiVk=
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c/go.mod h1:JHkPIbrfpd72SG/EVd6muEfDQjcINNoR0C8j2r3qZ4Q=
	gopkg.in/yaml.v3 v3.0.1 h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=
	gopkg.in/yaml.v3 v3.0.1/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
=========================================================================
Output:
	in broswer  search:
		localhost:8080 + Enter the key

	You will see that
==========================================================================
/*
1- Write in Terminal:
	go mod tidy

2- Output:
	go: finding module for package github.com/beego/beego/v2/server/web
	go: downloading github.com/beego/beego/v2 v2.1.3
	go: found github.com/beego/beego/v2/server/web in github.com/beego/beego/v2 v2.1.3
	go: downloading github.com/prometheus/client_golang v1.16.0
	go: downloading golang.org/x/crypto v0.10.0
	go: downloading google.golang.org/protobuf v1.30.0
	go: downloading github.com/elazarl/go-bindata-assetfs v1.0.1
	go: downloading github.com/stretchr/testify v1.8.1
	go: downloading github.com/pkg/errors v0.9.1
	go: downloading github.com/mitchellh/mapstructure v1.5.0
	go: downloading github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18
	go: downloading github.com/prometheus/client_model v0.3.0
	go: downloading github.com/prometheus/common v0.42.0
	go: downloading github.com/cespare/xxhash/v2 v2.2.0
	go: downloading github.com/prometheus/procfs v0.10.1
	go: downloading golang.org/x/net v0.10.0
	go: downloading github.com/golang/protobuf v1.5.3
	go: downloading github.com/matttproud/golang_protobuf_extensions v1.0.4
	go: downloading golang.org/x/text v0.10.0
	go: downloading gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c
	go: downloading github.com/kr/pretty v0.3.1
	go: downloading github.com/rogpeppe/go-internal v1.10.0
	go: finding module for package github.com/kr/text
	go: downloading github.com/kr/text v0.2.0
	go: found github.com/kr/text in github.com/kr/text v0.2.0

*/

/*
1- After that, write:
	go run .

2- Output in Terminal is:
	2023/11/11 08:10:26.835 [D]  init global config instance failed. If you do not use this, just ignore it.  open conf/app.conf: no such file or directory
	2023/11/11 08:10:26.839 [I]  http server Running on http://:8080

3- write in browser:
	localhost:8080

4- Output:
	NOT Found blue text with another information.
*/
`,
		"goal or none sense game": `game-1 Goal Or None Sense Game

main.go:
	package main
	import (
		"fmt"
		"math/rand"
		"time"
	)
	func main() {
		rand.Seed(time.Now().UnixNano())

		var cups, chances int

		fmt.Print("How many cups? ")
		fmt.Scan(&cups)

		fmt.Print("How many chances? ")
		fmt.Scan(&chances)

		aiGoal := rand.Intn(cups) + 1
		fmt.Println("----------")
		var userGuess int
		for i := 0; i < chances; i++ {
			fmt.Printf("%d chances left.\n", chances-i)
			fmt.Print("Guess [1-", cups, "]: ")
			fmt.Scan(&userGuess)
			if userGuess == aiGoal {
				fmt.Println("You guessed right.")
				break
			} else {
				fmt.Println("🥶 🫨 🥶 🫨")
				fmt.Println("🥶 🫥 🥶 🫨 🫨")
				fmt.Println("----------")
			}
		}
		if userGuess == aiGoal {
			fmt.Println("==============")
			fmt.Println("💓 😍 💓 😍 💓 😍 💓 😍💗💗💗")
		} else {
			fmt.Printf("The right answer is %d.\n", aiGoal)
			fmt.Println("You lost. Sorry!")
		}
	}
`,
		"game": `game-1 Goal Or None Sense Game
main.go:
package main
import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	rand.Seed(time.Now().UnixNano())
	var cups, chances int
	fmt.Print("How many cups? ")
	fmt.Scan(&cups)
	fmt.Print("How many chances? ")
	fmt.Scan(&chances)
	aiGoal := rand.Intn(cups) + 1
	fmt.Println("----------")
	var userGuess int
	for i := 0; i < chances; i++ {
		fmt.Printf("%d chances left.\n", chances-i)
		fmt.Print("Guess [1-", cups, "]: ")
		fmt.Scan(&userGuess)

		if userGuess == aiGoal {
			fmt.Println("You guessed right.")
			break
		} else {
			fmt.Println("🥶 🫨 🥶 🫨")
			fmt.Println("🥶 🫥 🥶 🫨 🫨")
			fmt.Println("----------")
		}
	}
	if userGuess == aiGoal {
		fmt.Println("==============")
		fmt.Println("💓 😍 💓 😍 💓 😍 💓 😍💗💗💗")
	} else {
		fmt.Printf("The right answer is %d.\n", aiGoal)
		fmt.Println("You lost. Sorry!")
	}
}
`,
		// =======================================================================================
		// =======================================================================================
		// =======================================================================================
		"game1": `game-1 Goal Or None Sense Game
main.go:
package main
import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	rand.Seed(time.Now().UnixNano())
	var cups, chances int
	fmt.Print("How many cups? ")
	fmt.Scan(&cups)
	fmt.Print("How many chances? ")
	fmt.Scan(&chances)
	aiGoal := rand.Intn(cups) + 1
	fmt.Println("----------")
	var userGuess int
	for i := 0; i < chances; i++ {
		fmt.Printf("%d chances left.\n", chances-i)
		fmt.Print("Guess [1-", cups, "]: ")
		fmt.Scan(&userGuess)

		if userGuess == aiGoal {
			fmt.Println("You guessed right.")
			break
		} else {
			fmt.Println("🥶 🫨 🥶 🫨")
			fmt.Println("🥶 🫥 🥶 🫨 🫨")
			fmt.Println("----------")
		}
	}
	if userGuess == aiGoal {
		fmt.Println("==============")
		fmt.Println("💓 😍 💓 😍 💓 😍 💓 😍💗💗💗")
	} else {
		fmt.Printf("The right answer is %d.\n", aiGoal)
		fmt.Println("You lost. Sorry!")
	}
}
`,
		// =======================================================================================
		// =======================================================================================
		// =======================================================================================
		"game goal or none sense": `game-1 Goal Or None Sense Game
main.go:
package main
import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	rand.Seed(time.Now().UnixNano())
	var cups, chances int
	fmt.Print("How many cups? ")
	fmt.Scan(&cups)
	fmt.Print("How many chances? ")
	fmt.Scan(&chances)
	aiGoal := rand.Intn(cups) + 1
	fmt.Println("----------")
	var userGuess int
	for i := 0; i < chances; i++ {
		fmt.Printf("%d chances left.\n", chances-i)
		fmt.Print("Guess [1-", cups, "]: ")
		fmt.Scan(&userGuess)

		if userGuess == aiGoal {
			fmt.Println("You guessed right.")
			break
		} else {
			fmt.Println("🥶 🫨 🥶 🫨")
			fmt.Println("🥶 🫥 🥶 🫨 🫨")
			fmt.Println("----------")
		}
	}
	if userGuess == aiGoal {
		fmt.Println("==============")
		fmt.Println("💓 😍 💓 😍 💓 😍 💓 😍💗💗💗")
	} else {
		fmt.Printf("The right answer is %d.\n", aiGoal)
		fmt.Println("You lost. Sorry!")
	}
}
`,
		// =======================================================================================
		// =======================================================================================
		// =======================================================================================
		"goal or none sense": `game-1 Goal Or None Sense Game
main.go:
package main
import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	rand.Seed(time.Now().UnixNano())
	var cups, chances int
	fmt.Print("How many cups? ")
	fmt.Scan(&cups)
	fmt.Print("How many chances? ")
	fmt.Scan(&chances)
	aiGoal := rand.Intn(cups) + 1
	fmt.Println("----------")
	var userGuess int
	for i := 0; i < chances; i++ {
		fmt.Printf("%d chances left.\n", chances-i)
		fmt.Print("Guess [1-", cups, "]: ")
		fmt.Scan(&userGuess)

		if userGuess == aiGoal {
			fmt.Println("You guessed right.")
			break
		} else {
			fmt.Println("🥶 🫨 🥶 🫨")
			fmt.Println("🥶 🫥 🥶 🫨 🫨")
			fmt.Println("----------")
		}
	}
	if userGuess == aiGoal {
		fmt.Println("==============")
		fmt.Println("💓 😍 💓 😍 💓 😍 💓 😍💗💗💗")
	} else {
		fmt.Printf("The right answer is %d.\n", aiGoal)
		fmt.Println("You lost. Sorry!")
	}
}
`,
		// ******************************************************************************************************
		"encryption": `encryption1 simple encryption
main.go:
	package main

	import "fmt"

	func main() {
		userInput := "sina"

		var Encryption string

		for _, char := range userInput {
			Encryption += fmt.Sprintf("%d ", char)
		}

		fmt.Println(Encryption)
	}
`,
		// *********************************//////////////////////////////**********************************
		"encryption1": `encryption1 simple encryption
main.go:
	package main

	import "fmt"

	func main() {
		userInput := "sina"

		var Encryption string

		for _, char := range userInput {
			Encryption += fmt.Sprintf("%d ", char)
		}

		fmt.Println(Encryption)
	}
`,
		// ************************************************************************************************
		// =================================================================================================
		"simple encryption": `encryption1 simple encryption
main.go:
	package main

	import "fmt"

	func main() {
		userInput := "sina"

		var Encryption string

		for _, char := range userInput {
			Encryption += fmt.Sprintf("%d ", char)
		}

		fmt.Println(Encryption)
	}
`,
		// ////////////////////////////////////////////////////////////////////////////////////////////////
		// ================================================================================================
		"first encryption": `encryption1 simple encryption
main.go:
	package main

	import "fmt"

	func main() {
		userInput := "sina"

		var Encryption string

		for _, char := range userInput {
			Encryption += fmt.Sprintf("%d ", char)
		}

		fmt.Println(Encryption)
	}
`,
		// ////////////////////////////////////////////////////////////////////////////////////////////////
		// ================================================================================================
		"go tool dist list": `
488🚀 This will provide you a list of operating systems and architectures separated by / characters:
Output:
	aix/ppc64
	android/386
	android/amd64
	android/arm
	android/arm64
	darwin/amd64
	darwin/arm64
	dragonfly/amd64
	freebsd/386
	freebsd/amd64
	freebsd/arm
	freebsd/arm64
	freebsd/riscv64
	illumos/amd64
	ios/amd64
	ios/arm64
	js/wasm
	linux/386
	linux/amd64
	linux/arm
	linux/arm64
	linux/loong64
	linux/mips
	linux/mips64
	linux/mips64le
	linux/mipsle
	linux/ppc64
	linux/ppc64le
	linux/riscv64
	linux/s390x
	netbsd/386
	netbsd/amd64
	netbsd/arm
	netbsd/arm64
	openbsd/386
	openbsd/amd64
	openbsd/arm
	openbsd/arm64
	plan9/386
	plan9/amd64
	plan9/arm
	solaris/amd64
	wasip1/wasm
	windows/386
	windows/amd64
	windows/arm
	windows/arm64`,
		// ////////////////////////////////////////////////////////////////////////////////////////////////
		// ===============================================================================================
		"tool dist list": `
488🚀 This will provide you a list of operating systems and architectures separated by / characters:
Output:
	aix/ppc64
	android/386
	android/amd64
	android/arm
	android/arm64
	darwin/amd64
	darwin/arm64
	dragonfly/amd64
	freebsd/386
	freebsd/amd64
	freebsd/arm
	freebsd/arm64
	freebsd/riscv64
	illumos/amd64
	ios/amd64
	ios/arm64
	js/wasm
	linux/386
	linux/amd64
	linux/arm
	linux/arm64
	linux/loong64
	linux/mips
	linux/mips64
	linux/mips64le
	linux/mipsle
	linux/ppc64
	linux/ppc64le
	linux/riscv64
	linux/s390x
	netbsd/386
	netbsd/amd64
	netbsd/arm
	netbsd/arm64
	openbsd/386
	openbsd/amd64
	openbsd/arm
	openbsd/arm64
	plan9/386
	plan9/amd64
	plan9/arm
	solaris/amd64
	wasip1/wasm
	windows/386
	windows/amd64
	windows/arm
	windows/arm64`,
		// ///////////////////////////////////////////////////////////////////////////////////
		// ==================================================================================
		"dist list": `
488🚀 This will provide you a list of operating systems and architectures separated by / characters:
Output:
	aix/ppc64
	android/386
	android/amd64
	android/arm
	android/arm64
	darwin/amd64
	darwin/arm64
	dragonfly/amd64
	freebsd/386
	freebsd/amd64
	freebsd/arm
	freebsd/arm64
	freebsd/riscv64
	illumos/amd64
	ios/amd64
	ios/arm64
	js/wasm
	linux/386
	linux/amd64
	linux/arm
	linux/arm64
	linux/loong64
	linux/mips
	linux/mips64
	linux/mips64le
	linux/mipsle
	linux/ppc64
	linux/ppc64le
	linux/riscv64
	linux/s390x
	netbsd/386
	netbsd/amd64
	netbsd/arm
	netbsd/arm64
	openbsd/386
	openbsd/amd64
	openbsd/arm
	openbsd/arm64
	plan9/386
	plan9/amd64
	plan9/arm
	solaris/amd64
	wasip1/wasm
	windows/386
	windows/amd64
	windows/arm
	windows/arm64`,
		// =====================================================================================
		"random generator": `
	// random generator1
	// go get github.com/Knetic/govaluate

	package main
	
	import (
		"fmt"
		"math/rand"
		"time"
	
		"github.com/Knetic/govaluate"
	)
	
	func generateRandomFormula() string {
		// Get the current time with millisecond precision
		currentTime := time.Now()
	
		// Get the current time with microsecond precision
		microseconds := currentTime.UnixNano() / int64(time.Microsecond) % 1000000
	
		microsecondsPlus := currentTime.UnixNano() / int64(time.Microsecond) % 1000000
	
		// Generate a 6-digit random number
		randomNumber := microseconds
	
	
		randomNumberPlus := (microsecondsPlus / 12) * (2 + 9) - (10 * 41) * 2 
		
	
		// Generate random operations (you can customize this based on your needs)
		operations := []string{"+", "-", "*", "/"}
		randomOperation := operations[rand.Intn(len(operations))]
	
		// Create the random formula
		formula := fmt.Sprintf("%d %s %d", randomNumber, randomOperation, randomNumberPlus)
	
		return formula
	}
	
	func evaluateFormula(formula string) (float64, error) {
		expression, err := govaluate.NewEvaluableExpression(formula)
		if err != nil {
			return 0, err
		}
	
		result, err := expression.Evaluate(nil)
		if err != nil {
			return 0, err
		}
	
		return result.(float64), nil
	}
	
	func main() {
		// Generate and print a random formula
		randomFormula := generateRandomFormula()
		fmt.Println("Random Formula:", randomFormula)
	
		result, err := evaluateFormula(randomFormula)
		if err != nil {
			fmt.Println("Error evaluating formula:", err)
			return
		}
	
		fmt.Println("Result:", result)
	
	}
	`,
		// ==================================================================
		"random generator1": `
	// random generator1
	// go get github.com/Knetic/govaluate

	package main
	
	import (
		"fmt"
		"math/rand"
		"time"
	
		"github.com/Knetic/govaluate"
	)
	
	func generateRandomFormula() string {
		// Get the current time with millisecond precision
		currentTime := time.Now()
	
		// Get the current time with microsecond precision
		microseconds := currentTime.UnixNano() / int64(time.Microsecond) % 1000000
	
		microsecondsPlus := currentTime.UnixNano() / int64(time.Microsecond) % 1000000
	
		// Generate a 6-digit random number
		randomNumber := microseconds
	
	
		randomNumberPlus := (microsecondsPlus / 12) * (2 + 9) - (10 * 41) * 2 
		
	
		// Generate random operations (you can customize this based on your needs)
		operations := []string{"+", "-", "*", "/"}
		randomOperation := operations[rand.Intn(len(operations))]
	
		// Create the random formula
		formula := fmt.Sprintf("%d %s %d", randomNumber, randomOperation, randomNumberPlus)
	
		return formula
	}
	
	func evaluateFormula(formula string) (float64, error) {
		expression, err := govaluate.NewEvaluableExpression(formula)
		if err != nil {
			return 0, err
		}
	
		result, err := expression.Evaluate(nil)
		if err != nil {
			return 0, err
		}
	
		return result.(float64), nil
	}
	
	func main() {
		// Generate and print a random formula
		randomFormula := generateRandomFormula()
		fmt.Println("Random Formula:", randomFormula)
	
		result, err := evaluateFormula(randomFormula)
		if err != nil {
			fmt.Println("Error evaluating formula:", err)
			return
		}
	
		fmt.Println("Result:", result)
	
	}
	`,

		"random generator numbers": `
	// random generator1
	// go get github.com/Knetic/govaluate

	package main
	
	import (
		"fmt"
		"math/rand"
		"time"
	
		"github.com/Knetic/govaluate"
	)
	
	func generateRandomFormula() string {
		// Get the current time with millisecond precision
		currentTime := time.Now()
	
		// Get the current time with microsecond precision
		microseconds := currentTime.UnixNano() / int64(time.Microsecond) % 1000000
	
		microsecondsPlus := currentTime.UnixNano() / int64(time.Microsecond) % 1000000
	
		// Generate a 6-digit random number
		randomNumber := microseconds
	
	
		randomNumberPlus := (microsecondsPlus / 12) * (2 + 9) - (10 * 41) * 2 
		
	
		// Generate random operations (you can customize this based on your needs)
		operations := []string{"+", "-", "*", "/"}
		randomOperation := operations[rand.Intn(len(operations))]
	
		// Create the random formula
		formula := fmt.Sprintf("%d %s %d", randomNumber, randomOperation, randomNumberPlus)
	
		return formula
	}
	
	func evaluateFormula(formula string) (float64, error) {
		expression, err := govaluate.NewEvaluableExpression(formula)
		if err != nil {
			return 0, err
		}
	
		result, err := expression.Evaluate(nil)
		if err != nil {
			return 0, err
		}
	
		return result.(float64), nil
	}
	
	func main() {
		// Generate and print a random formula
		randomFormula := generateRandomFormula()
		fmt.Println("Random Formula:", randomFormula)
	
		result, err := evaluateFormula(randomFormula)
		if err != nil {
			fmt.Println("Error evaluating formula:", err)
			return
		}
	
		fmt.Println("Result:", result)
	
	}
	`,
		// =======================================================================================
		// =======================================================================================
		// =======================================================================================
		"crawler": `crawler1
	In Terminal in current directory write:
		go get golang.org/x/net/html

	main.go:
		package main
		import (
			"fmt"
			"golang.org/x/net/html"
			"net/http"
			"os"
			"strconv"
			"time"
		)
		// Crawl function takes a URL and recursively crawls the pages
		func Crawl(url string, depth int, searchDir string) {
			if depth <= 0 {
				return
			}
			// Create a directory for the search with the current date, minutes, and seconds
			now := time.Now()
			searchDir = fmt.Sprintf("%s-search-%02d%02d%02d", now.Format("2006-01-02"), now.Minute(), now.Second())
			err := os.Mkdir(searchDir, 0755)
			if err != nil {
				fmt.Println("Error creating directory:", err)
				return
			}
			// Create a file to store search results
			resultFile, err := os.Create(fmt.Sprintf("%s/searchResult.txt", searchDir))
			if err != nil {
				fmt.Println("Error creating result file:", err)
				return
			}
			defer resultFile.Close()
			// Make an HTTP request
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("Error making request:", err)
				return
			}
			defer resp.Body.Close()
			// Parse the HTML content
			doc, err := html.Parse(resp.Body)
			if err != nil {
				fmt.Println("Error parsing HTML:", err)
				return
			}
			// Process the links on the current page and write them to the result file
			processLinks(doc, resultFile)
			// Recursively crawl the linked pages
			links := extractLinks(doc)
			for _, link := range links {
				Crawl(link, depth-1, searchDir)
			}
		}
		// processLinks extracts and prints the links on the current page
		func processLinks(n *html.Node, resultFile *os.File) {
			if n.Type == html.ElementNode && n.Data == "a" {
				for _, a := range n.Attr {
					if a.Key == "href" {
						link := fmt.Sprintf("Link: %s\n", a.Val)
						resultFile.WriteString(link)
					}
				}
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				processLinks(c, resultFile)
			}
		}
		// extractLinks returns a slice of unique links from the HTML document
		func extractLinks(n *html.Node) []string {
			var links []string
			visited := make(map[string]bool)
			var visitNode func(*html.Node)
			visitNode = func(n *html.Node) {
				if n.Type == html.ElementNode && n.Data == "a" {
					for _, a := range n.Attr {
						if a.Key == "href" {
							link := a.Val
							if !visited[link] {
								links = append(links, link)
								visited[link] = true
							}
						}
					}
				}
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					visitNode(c)
				}
			}
			visitNode(n)
			return links
		}
		func main() {
			if len(os.Args) != 3 {
				fmt.Println("Usage: go run crawler.go <url> <depth>")
				return
			}
			url := os.Args[1]
			depthStr := os.Args[2]
			depth, err := strconv.Atoi(depthStr)
			if err != nil {
				fmt.Println("Error converting depth to integer:", err)
				return
			}
			var searchDir string // searchDir will be created automatically based on the current date, minutes, and seconds
			Crawl(url, depth, searchDir)
		}
`,
		// =======================================================================================
		// =======================================================================================
		// =======================================================================================
		"crawler2": `crawler2
in Terminal in current directory
	go get golang.org/x/net/html

main.go:
	package main

	import (
		"fmt"
		"net/http"
		"os"
		"strconv"

		"golang.org/x/net/html"
	)

	// Crawl function takes a URL and recursively crawls the pages
	func Crawl(url string, depth int) {
		if depth <= 0 {
			return
		}

		// Make an HTTP request
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error making request:", err)
			return
		}
		defer resp.Body.Close()

		// Parse the HTML content
		doc, err := html.Parse(resp.Body)
		if err != nil {
			fmt.Println("Error parsing HTML:", err)
			return
		}

		// Process the links on the current page
		processLinks(doc)

		// Recursively crawl the linked pages
		links := extractLinks(doc)
		for _, link := range links {
			Crawl(link, depth-1)
		}
	}

	// processLinks extracts and prints the links on the current page
	func processLinks(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println("Link:", a.Val)
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			processLinks(c)
		}
	}

	// extractLinks returns a slice of unique links from the HTML document
	func extractLinks(n *html.Node) []string {
		var links []string
		visited := make(map[string]bool)

		var visitNode func(*html.Node)
		visitNode = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "a" {
				for _, a := range n.Attr {
					if a.Key == "href" {
						link := a.Val
						if !visited[link] {
							links = append(links, link)
							visited[link] = true
						}
					}
				}
			}

			for c := n.FirstChild; c != nil; c = c.NextSibling {
				visitNode(c)
			}
		}

		visitNode(n)
		return links
	}

	func main() {
		if len(os.Args) != 3 {
			fmt.Println("Usage: go run crawler.go <url> <depth>")
			return
		}

		url := os.Args[1]
		depthstr := os.Args[2]

		depth, err := strconv.Atoi(depthstr)
		if err != nil {
			fmt.Println("Error converting depth to integer:", err)
			return
		}

		Crawl(url, depth)
	}

`,
		// =======================================================================================
		// =======================================================================================
		// =======================================================================================
		"simple crawler": `crawler2
in Terminal in current directory
	go get golang.org/x/net/html

main.go:
	package main

	import (
		"fmt"
		"net/http"
		"os"
		"strconv"

		"golang.org/x/net/html"
	)

	// Crawl function takes a URL and recursively crawls the pages
	func Crawl(url string, depth int) {
		if depth <= 0 {
			return
		}

		// Make an HTTP request
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error making request:", err)
			return
		}
		defer resp.Body.Close()

		// Parse the HTML content
		doc, err := html.Parse(resp.Body)
		if err != nil {
			fmt.Println("Error parsing HTML:", err)
			return
		}

		// Process the links on the current page
		processLinks(doc)

		// Recursively crawl the linked pages
		links := extractLinks(doc)
		for _, link := range links {
			Crawl(link, depth-1)
		}
	}

	// processLinks extracts and prints the links on the current page
	func processLinks(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println("Link:", a.Val)
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			processLinks(c)
		}
	}

	// extractLinks returns a slice of unique links from the HTML document
	func extractLinks(n *html.Node) []string {
		var links []string
		visited := make(map[string]bool)

		var visitNode func(*html.Node)
		visitNode = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "a" {
				for _, a := range n.Attr {
					if a.Key == "href" {
						link := a.Val
						if !visited[link] {
							links = append(links, link)
							visited[link] = true
						}
					}
				}
			}

			for c := n.FirstChild; c != nil; c = c.NextSibling {
				visitNode(c)
			}
		}

		visitNode(n)
		return links
	}

	func main() {
		if len(os.Args) != 3 {
			fmt.Println("Usage: go run crawler.go <url> <depth>")
			return
		}

		url := os.Args[1]
		depthstr := os.Args[2]

		depth, err := strconv.Atoi(depthstr)
		if err != nil {
			fmt.Println("Error converting depth to integer:", err)
			return
		}

		Crawl(url, depth)
	}

`,
		// =======================================================================================
		// =======================================================================================
		// =======================================================================================
		"simple crawl": `crawler2
in Terminal in current directory
	go get golang.org/x/net/html

main.go:
	package main

	import (
		"fmt"
		"net/http"
		"os"
		"strconv"

		"golang.org/x/net/html"
	)

	// Crawl function takes a URL and recursively crawls the pages
	func Crawl(url string, depth int) {
		if depth <= 0 {
			return
		}

		// Make an HTTP request
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error making request:", err)
			return
		}
		defer resp.Body.Close()

		// Parse the HTML content
		doc, err := html.Parse(resp.Body)
		if err != nil {
			fmt.Println("Error parsing HTML:", err)
			return
		}

		// Process the links on the current page
		processLinks(doc)

		// Recursively crawl the linked pages
		links := extractLinks(doc)
		for _, link := range links {
			Crawl(link, depth-1)
		}
	}

	// processLinks extracts and prints the links on the current page
	func processLinks(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println("Link:", a.Val)
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			processLinks(c)
		}
	}

	// extractLinks returns a slice of unique links from the HTML document
	func extractLinks(n *html.Node) []string {
		var links []string
		visited := make(map[string]bool)

		var visitNode func(*html.Node)
		visitNode = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "a" {
				for _, a := range n.Attr {
					if a.Key == "href" {
						link := a.Val
						if !visited[link] {
							links = append(links, link)
							visited[link] = true
						}
					}
				}
			}

			for c := n.FirstChild; c != nil; c = c.NextSibling {
				visitNode(c)
			}
		}

		visitNode(n)
		return links
	}

	func main() {
		if len(os.Args) != 3 {
			fmt.Println("Usage: go run crawler.go <url> <depth>")
			return
		}

		url := os.Args[1]
		depthstr := os.Args[2]

		depth, err := strconv.Atoi(depthstr)
		if err != nil {
			fmt.Println("Error converting depth to integer:", err)
			return
		}

		Crawl(url, depth)
	}

`,
		// =======================================================================================
		// =======================================================================================
		// =======================================================================================
		"simple crawlling": `crawler2
in Terminal in current directory
	go get golang.org/x/net/html

main.go:
	package main

	import (
		"fmt"
		"net/http"
		"os"
		"strconv"

		"golang.org/x/net/html"
	)

	// Crawl function takes a URL and recursively crawls the pages
	func Crawl(url string, depth int) {
		if depth <= 0 {
			return
		}

		// Make an HTTP request
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error making request:", err)
			return
		}
		defer resp.Body.Close()

		// Parse the HTML content
		doc, err := html.Parse(resp.Body)
		if err != nil {
			fmt.Println("Error parsing HTML:", err)
			return
		}

		// Process the links on the current page
		processLinks(doc)

		// Recursively crawl the linked pages
		links := extractLinks(doc)
		for _, link := range links {
			Crawl(link, depth-1)
		}
	}

	// processLinks extracts and prints the links on the current page
	func processLinks(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println("Link:", a.Val)
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			processLinks(c)
		}
	}

	// extractLinks returns a slice of unique links from the HTML document
	func extractLinks(n *html.Node) []string {
		var links []string
		visited := make(map[string]bool)

		var visitNode func(*html.Node)
		visitNode = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "a" {
				for _, a := range n.Attr {
					if a.Key == "href" {
						link := a.Val
						if !visited[link] {
							links = append(links, link)
							visited[link] = true
						}
					}
				}
			}

			for c := n.FirstChild; c != nil; c = c.NextSibling {
				visitNode(c)
			}
		}

		visitNode(n)
		return links
	}

	func main() {
		if len(os.Args) != 3 {
			fmt.Println("Usage: go run crawler.go <url> <depth>")
			return
		}

		url := os.Args[1]
		depthstr := os.Args[2]

		depth, err := strconv.Atoi(depthstr)
		if err != nil {
			fmt.Println("Error converting depth to integer:", err)
			return
		}

		Crawl(url, depth)
	}

`,

		// =======================================================================================
		// =======================================================================================
		// =======================================================================================
		"crawlling": `crawler2
in Terminal in current directory
	go get golang.org/x/net/html

main.go:
	package main

	import (
		"fmt"
		"net/http"
		"os"
		"strconv"

		"golang.org/x/net/html"
	)

	// Crawl function takes a URL and recursively crawls the pages
	func Crawl(url string, depth int) {
		if depth <= 0 {
			return
		}

		// Make an HTTP request
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error making request:", err)
			return
		}
		defer resp.Body.Close()

		// Parse the HTML content
		doc, err := html.Parse(resp.Body)
		if err != nil {
			fmt.Println("Error parsing HTML:", err)
			return
		}

		// Process the links on the current page
		processLinks(doc)

		// Recursively crawl the linked pages
		links := extractLinks(doc)
		for _, link := range links {
			Crawl(link, depth-1)
		}
	}

	// processLinks extracts and prints the links on the current page
	func processLinks(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println("Link:", a.Val)
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			processLinks(c)
		}
	}

	// extractLinks returns a slice of unique links from the HTML document
	func extractLinks(n *html.Node) []string {
		var links []string
		visited := make(map[string]bool)

		var visitNode func(*html.Node)
		visitNode = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "a" {
				for _, a := range n.Attr {
					if a.Key == "href" {
						link := a.Val
						if !visited[link] {
							links = append(links, link)
							visited[link] = true
						}
					}
				}
			}

			for c := n.FirstChild; c != nil; c = c.NextSibling {
				visitNode(c)
			}
		}

		visitNode(n)
		return links
	}

	func main() {
		if len(os.Args) != 3 {
			fmt.Println("Usage: go run crawler.go <url> <depth>")
			return
		}

		url := os.Args[1]
		depthstr := os.Args[2]

		depth, err := strconv.Atoi(depthstr)
		if err != nil {
			fmt.Println("Error converting depth to integer:", err)
			return
		}

		Crawl(url, depth)
	}

`,

		// =======================================================================================
		// =======================================================================================
		// =======================================================================================

		"cms": `Content Management System with GO language
========================================================
Directory architecture:
└─templates/
|   └─index.gohtml
└─cmd/
|   └─main.go
└─go.mod
└─template.go
========================================================
go.mod:
	module cmsOne
	go 1.21.3
========================================================
cmd/main.go:
	package main
	import (
		cms "cmsOne"
		"os"
	)
	func main()  {
		p := cms.Page{	
			Title: "Hello, world!",
			Content: "This is body of our website",	
		} 
		cms.Templ.ExecuteTemplate(os.Stdout, "index", p)
	}
========================================================
template.go:
	package cms
	import (
		"html/template"
	)
	// Templ is an exported variable. 
	// Templ is stand for Templates
	var Templ = template.Must(template.ParseGlob("../templates/*"))
	type Page struct {
		Title string
		Content string
	}
========================================================
templates/index.gohtml:
	{{ define "index" }}
	<!DOCTYPE html>
	asdasd
	<html>
		<head>
			<title>
				{{.Title}}
			</title>
		</head>
		<body>
			{{ .Content }}
		</body>
	</html>
	{{ end }}
========================================================
	`,

		// =======================================================================================
		// =======================================================================================
		// =======================================================================================

		"cms-go": `Content Management System with GO language
	Content Management System with GO language
========================================================
Directory architecture:
└─templates/
|   └─index.gohtml
└─cmd/
|   └─main.go
└─go.mod
└─template.go
========================================================
go.mod:
	module cmsOne
	go 1.21.3
========================================================
cmd/main.go:
	package main
	import (
		cms "cmsOne"
		"os"
	)
	func main()  {
		p := cms.Page{	
			Title: "Hello, world!",
			Content: "This is body of our website",	
		} 
		cms.Templ.ExecuteTemplate(os.Stdout, "index", p)
	}
========================================================
template.go:
	package cms
	import (
		"html/template"
	)
	// Templ is an exported variable. 
	// Templ is stand for Templates
	var Templ = template.Must(template.ParseGlob("../templates/*"))
	type Page struct {
		Title string
		Content string
	}
========================================================
templates/index.gohtml:
	{{ define "index" }}
	<!DOCTYPE html>
	asdasd
	<html>
		<head>
			<title>
				{{.Title}}
			</title>
		</head>
		<body>
			{{ .Content }}
		</body>
	</html>
	{{ end }}
========================================================`,

		// =======================================================================================
		// =======================================================================================
		// =======================================================================================

		"cms with go": `Content Management System with GO language
	Content Management System with GO language
========================================================
Directory architecture:
└─templates/
|   └─index.gohtml
└─cmd/
|   └─main.go
└─go.mod
└─template.go
========================================================
go.mod:
	module cmsOne
	go 1.21.3
========================================================
cmd/main.go:
	package main
	import (
		cms "cmsOne"
		"os"
	)
	func main()  {
		p := cms.Page{	
			Title: "Hello, world!",
			Content: "This is body of our website",	
		} 
		cms.Templ.ExecuteTemplate(os.Stdout, "index", p)
	}
========================================================
template.go:
	package cms
	import (
		"html/template"
	)
	// Templ is an exported variable. 
	// Templ is stand for Templates
	var Templ = template.Must(template.ParseGlob("../templates/*"))
	type Page struct {
		Title string
		Content string
	}
========================================================
templates/index.gohtml:
	{{ define "index" }}
	<!DOCTYPE html>
	asdasd
	<html>
		<head>
			<title>
				{{.Title}}
			</title>
		</head>
		<body>
			{{ .Content }}
		</body>
	</html>
	{{ end }}
========================================================`,

		// =======================================================================================
		// =======================================================================================
		// =======================================================================================

		"content management system": `Content Management System with GO language
	Content Management System with GO language
========================================================
Directory architecture:
└─templates/
|   └─index.gohtml
└─cmd/
|   └─main.go
└─go.mod
└─template.go
========================================================
go.mod:
	module cmsOne
	go 1.21.3
========================================================
cmd/main.go:
	package main
	import (
		cms "cmsOne"
		"os"
	)
	func main()  {
		p := cms.Page{	
			Title: "Hello, world!",
			Content: "This is body of our website",	
		} 
		cms.Templ.ExecuteTemplate(os.Stdout, "index", p)
	}
========================================================
template.go:
	package cms
	import (
		"html/template"
	)
	// Templ is an exported variable. 
	// Templ is stand for Templates
	var Templ = template.Must(template.ParseGlob("../templates/*"))
	type Page struct {
		Title string
		Content string
	}
========================================================
templates/index.gohtml:
	{{ define "index" }}
	<!DOCTYPE html>
	asdasd
	<html>
		<head>
			<title>
				{{.Title}}
			</title>
		</head>
		<body>
			{{ .Content }}
		</body>
	</html>
	{{ end }}
========================================================`,

		// =======================================================================================
		// =======================================================================================
		// =======================================================================================
		"illustrates this process is IP": `illustrate-1
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}
	os.Exit(0)
}
`,
		// =======================================================================================
		// =======================================================================================
		// =======================================================================================
		"illustrate process is IP": `illustrate-1
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}
	os.Exit(0)
}
`,

		// =======================================================================================
		// =======================================================================================
		// =======================================================================================
		"illustrate process IP": `illustrate-1
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}
	os.Exit(0)
}
`,

		// =======================================================================================
		// =======================================================================================
		// =======================================================================================

		"illustrate IP": `illustrate-1
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}
	os.Exit(0)
}
`,

		// =======================================================================================
		// =======================================================================================
		// =======================================================================================

		"illustrator IP": `illustrate-1
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}
	os.Exit(0)
}
`,

		// =======================================================================================
		// =======================================================================================
		// =======================================================================================

		"resolveip.go": `ResolveIP.go-1
main.go:
	/* ResolveIP
	*/
	package main

	import (
		"fmt"
		"net"
		"os"
	)

	func main() {
		if len(os.Args) != 2 {
			fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
			fmt.Println("Usage: ", os.Args[0], "hostname")
			os.Exit(1)
		}
		name := os.Args[1]
		addr, err := net.ResolveIPAddr("ip", name)
		if err != nil {
			fmt.Println("Resolution error", err.Error())
			os.Exit(1)
		}
		fmt.Println("Resolved address is ", addr.String())
		os.Exit(0)
	}

`,

		// =======================================================================================
		// =======================================================================================
		// =======================================================================================

		"resolveip go": `ResolveIP.go-1
main.go:
	/* ResolveIP
	*/
	package main

	import (
		"fmt"
		"net"
		"os"
	)

	func main() {
		if len(os.Args) != 2 {
			fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
			fmt.Println("Usage: ", os.Args[0], "hostname")
			os.Exit(1)
		}
		name := os.Args[1]
		addr, err := net.ResolveIPAddr("ip", name)
		if err != nil {
			fmt.Println("Resolution error", err.Error())
			os.Exit(1)
		}
		fmt.Println("Resolved address is ", addr.String())
		os.Exit(0)
	}

`,

		// =======================================================================================
		// =======================================================================================
		// =======================================================================================
		"resolveip": `ResolveIP.go-1
main.go:
	/* ResolveIP
	*/
	package main

	import (
		"fmt"
		"net"
		"os"
	)

	func main() {
		if len(os.Args) != 2 {
			fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
			fmt.Println("Usage: ", os.Args[0], "hostname")
			os.Exit(1)
		}
		name := os.Args[1]
		addr, err := net.ResolveIPAddr("ip", name)
		if err != nil {
			fmt.Println("Resolution error", err.Error())
			os.Exit(1)
		}
		fmt.Println("Resolved address is ", addr.String())
		os.Exit(0)
	}

`,

		// =======================================================================================
		// =======================================================================================
		// =======================================================================================
		"resolver ip": `ResolveIP.go-1
main.go:
	/* ResolveIP
	*/
	package main

	import (
		"fmt"
		"net"
		"os"
	)

	func main() {
		if len(os.Args) != 2 {
			fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
			fmt.Println("Usage: ", os.Args[0], "hostname")
			os.Exit(1)
		}
		name := os.Args[1]
		addr, err := net.ResolveIPAddr("ip", name)
		if err != nil {
			fmt.Println("Resolution error", err.Error())
			os.Exit(1)
		}
		fmt.Println("Resolved address is ", addr.String())
		os.Exit(0)
	}

`,

		// =======================================================================================
		// =======================================================================================
		// =======================================================================================
		"resolverip": `ResolveIP.go-1
main.go:
	/* ResolveIP
	*/
	package main

	import (
		"fmt"
		"net"
		"os"
	)

	func main() {
		if len(os.Args) != 2 {
			fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
			fmt.Println("Usage: ", os.Args[0], "hostname")
			os.Exit(1)
		}
		name := os.Args[1]
		addr, err := net.ResolveIPAddr("ip", name)
		if err != nil {
			fmt.Println("Resolution error", err.Error())
			os.Exit(1)
		}
		fmt.Println("Resolved address is ", addr.String())
		os.Exit(0)
	}

`,

// =======================================================================================
"lookuphost":         `LookupHost-1
main.go:
	/* LookupHost
	*/
	package main

	import (
		"fmt"
		"net"
		"os"
	)

	func main() {
		if len(os.Args) != 2 {
			fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
			os.Exit(1)
		}
		name := os.Args[1]
		addrs, err := net.LookupHost(name)
		if err != nil {
			fmt.Println("Error: ", err.Error())
			os.Exit(2)
		}
		for _, s := range addrs {
			fmt.Println(s)
		}
		os.Exit(0)
	}

`,
// =======================================================================================
"find lookuphost":    `LookupHost-1
main.go:
	/* LookupHost
	*/
	package main

	import (
		"fmt"
		"net"
		"os"
	)

	func main() {
		if len(os.Args) != 2 {
			fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
			os.Exit(1)
		}
		name := os.Args[1]
		addrs, err := net.LookupHost(name)
		if err != nil {
			fmt.Println("Error: ", err.Error())
			os.Exit(2)
		}
		for _, s := range addrs {
			fmt.Println(s)
		}
		os.Exit(0)
	}

`,
// =======================================================================================
"lookuphost code":    `LookupHost-1
main.go:
	/* LookupHost
	*/
	package main

	import (
		"fmt"
		"net"
		"os"
	)

	func main() {
		if len(os.Args) != 2 {
			fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
			os.Exit(1)
		}
		name := os.Args[1]
		addrs, err := net.LookupHost(name)
		if err != nil {
			fmt.Println("Error: ", err.Error())
			os.Exit(2)
		}
		for _, s := range addrs {
			fmt.Println(s)
		}
		os.Exit(0)
	}

`,
// =======================================================================================
"lookuphost program": `LookupHost-1
main.go:
	/* LookupHost
	*/
	package main

	import (
		"fmt"
		"net"
		"os"
	)

	func main() {
		if len(os.Args) != 2 {
			fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
			os.Exit(1)
		}
		name := os.Args[1]
		addrs, err := net.LookupHost(name)
		if err != nil {
			fmt.Println("Error: ", err.Error())
			os.Exit(2)
		}
		for _, s := range addrs {
			fmt.Println(s)
		}
		os.Exit(0)
	}

`,
// =======================================================================================
// =======================================================================================
// =======================================================================================
// =======================================================================================
},
}
