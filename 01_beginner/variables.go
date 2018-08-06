package main

import "fmt"

// package scope/global variables
var outside = "This is outside of the functions scope"
var outDec int		// declared outside

func main(){
	outDec = 10		// initialized inside of the function scope

	fmt.Println(outside)
	fmt.Println(outDec)

	var message string
	message = "Yaaoo Rev"

	fmt.Println("\n" + message)
	fmt.Print(message)
	fmt.Println()

	var a, b, c int		// declaration
	a = 11				// assignment/initialization

	// default values of int type will be assigned in b and c
	fmt.Println(message, a, b, c)

	// declaration and initialization at once
	var msg string = "REV"
	var x, y, z int = 11, 19, 40

	fmt.Println("\n" + msg, x, y, z)

	// multiple types
	var p, q, r = true, 10.0, "YO"

	fmt.Println(p, q, r)

	// implicit types
	// can only be used in functions
	text := "Width and height is"
	width, height := 100.89, 50.77

	fmt.Println("\n" + text, width, height)

	// constant variables
	const Size = 10

	fmt.Println("Size:", Size)
}
