package main

import "fmt"

func main(){
	num1, num2, num3 := 10, 12, 20

	// function call
	fmt.Printf("%d + %d = %d", num1, num2, add(num1, num2))
	fmt.Printf("\n%d + %d + %d = %d", num1, num2, num3, multipleAdd(num1, num2, num3))
}

// basic function in go
func add(num1 int, num2 int) int {
	return  num1 + num2
}

// multiple consecutive parameters in function
func multipleAdd(num1, num2, num3 int) int {
	return num1 + num2 + num3
}
