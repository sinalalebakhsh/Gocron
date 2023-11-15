package features

type SingleDefinitionExamples struct {
	MapSingleDefEx map[string]string
}

var OriginalSingleDefExamples = SingleDefinitionExamples{
	MapSingleDefEx: map[string]string{
		"for":`56.for
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

	"http servers":`403.Creating HTTP Servers
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
"simple server":`403.Creating HTTP Servers
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
"web application":`
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
"simple web application":`main.go:
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
"simple web app":`main.go:
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
"bool":`Boolean
var closed bool    // boolean variable 'closed' implicitly initialized with 'false'
speeding := true   // boolean variable 'speeding' initialized with 'true'
hasError := false  // boolean variable 'hasError' initialized with 'false' 
`,
"boolean":`Boolean
var closed bool    // boolean variable 'closed' implicitly initialized with 'false'
speeding := true   // boolean variable 'speeding' initialized with 'true'
hasError := false  // boolean variable 'hasError' initialized with 'false' 
`,
"booleans":`Boolean
var closed bool    // boolean variable 'closed' implicitly initialized with 'false'
speeding := true   // boolean variable 'speeding' initialized with 'true'
hasError := false  // boolean variable 'hasError' initialized with 'false' 
`,
"package comments":`kelvin.go
// Package kelvin provides tools to convert
// temperatures to and from Kelvin.
package kelvin
`,
"package comment":`kelvin.go
// Package kelvin provides tools to convert
// temperatures to and from Kelvin.
package kelvin`,
"function comment":`example:
// CelsiusFreezingTemp returns an integer value equal to the temperature at which water freezes in degrees Celsius.
func CelsiusFreezingTemp() int {
	return 0
}`,
"function comments":`example:
// CelsiusFreezingTemp returns an integer value equal to the temperature at which water freezes in degrees Celsius.
func CelsiusFreezingTemp() int {
	return 0
}`,
"operator":`Operator 	Example
--------	-------------
+ 		4 + 6 == 10
- 		15 - 10 == 5
* 		2 * 3 == 6
/ 		13 / 3 == 4
% 		13 % 3 == 1`,
"operators":`Operator 	Example
--------	-------------
+ 		4 + 6 == 10
- 		15 - 10 == 5
* 		2 * 3 == 6
/ 		13 / 3 == 4
% 		13 % 3 == 1`,
"arithmetic operators":`Operator 	Example
--------	-------------
+ 		4 + 6 == 10
- 		15 - 10 == 5
* 		2 * 3 == 6
/ 		13 / 3 == 4
% 		13 % 3 == 1`,
"converting between types":`var x int = 42 // x has type int
f := float64(x) // f has type float64 (ie. 42.0)
var y float64 = 11.9 // y has type float64
i := int(y) // i has type int (ie. 11)`,
"arithmetic operations on different types":`var x int = 42

// this line produces an error
value := float32(2.0) * x // invalid operation: mismatched types float32 and int

// you must convert int type to float32 before performing arithmetic operation
value := float32(2.0) * float32(x)`,

"arithmetic operations on different type":`var x int = 42

// this line produces an error
value := float32(2.0) * x // invalid operation: mismatched types float32 and int

// you must convert int type to float32 before performing arithmetic operation
value := float32(2.0) * float32(x)`,

	"godoc":`An Example Program With Godoc Comments
The code below adheres to the Go way, in this case using single-line comments.
===================================

`,
	"string":`A string literal is defined between double quotes:

	const name = "Jane"
	
==================================================	
Strings can be concatenated via the + operator:

	"Jane" + " " + "Austen"
	// => "Jane Austen"
	
==================================================	
`,

"beego":`Quick Start
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

"beego1":`beego-2
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

"beego framework":`Quick Start
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

"generics" :`main.go:
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

"generic" :`main.go:
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

"reflection":`example:
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


"reflection ":`example:
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


"reflection 1":`example:
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


"reflection1":`example:
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

"server":`server1 = you can use server2 example, server3 or more 
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

"server1":`server1 = you can use server2 example, server3 or more
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

"beego2":`beego-2
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
"goal or none sense game":`game-1 Goal Or None Sense Game

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
"game":`game-1 Goal Or None Sense Game
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
"game1":`game-1 Goal Or None Sense Game
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
"game goal or none sense":`game-1 Goal Or None Sense Game
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
"goal or none sense":`game-1 Goal Or None Sense Game
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
},

}


