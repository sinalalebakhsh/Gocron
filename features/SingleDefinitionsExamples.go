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
	},
}


