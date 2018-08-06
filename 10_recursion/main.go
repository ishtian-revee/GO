package main

import "fmt"

func main(){
	fmt.Println("Factorial of 10 is:", factorial(5))
}

// this is a recursion function calculating factorial
func factorial(num int) int {
	if num == 0 {
		return 1
	}

	return num * factorial(num-1)
}
