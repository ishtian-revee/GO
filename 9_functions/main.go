package main

import "fmt"

func main(){
	num1, num2, num3 := 10, 12, 20

	// function call
	fmt.Printf("%d + %d = %d", num1, num2, add(num1, num2))
	fmt.Printf("\n%d + %d + %d = %d", num1, num2, num3, multipleAdd(num1, num2, num3))

	// calling multiple return value function
	a, b := multipleReturn()
	fmt.Println("\nvalue of a:", a)
	fmt.Println("value of b:", b)

	// only a subset of the return
	_, c := multipleReturn()
	fmt.Println("value of c:", c)

	// calling variadic functions
	sum(1, 3)
	sum(1, 3, 4, 6)

	// slice can be pass as argument of a variadic function
	slc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum(slc...)
}

// basic function in go
func add(num1 int, num2 int) int {
	return  num1 + num2
}

// multiple consecutive parameters in function
func multipleAdd(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

// multiple return value function [takes no parameter here]
func multipleReturn() (int, int) {
	return 61, 19
}

// variadic function
func sum(nums ...int){
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums{
		total += num
	}
	fmt.Println("total:", total)
}