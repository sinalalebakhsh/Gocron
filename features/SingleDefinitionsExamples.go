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
"simple web application":`
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
	},
}


