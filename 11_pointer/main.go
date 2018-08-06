package main

import "fmt"

func main(){

	val := 1

	fmt.Println("value :", val)		// prints the value
	zeroVal(val)						// passes the value which actually doesn't return anything
	fmt.Println("value :", val)		// prints the value
	zeroPtr(&val)						// passes the address of 'value' and that function assigns 0 to that address
	fmt.Println("value :", val)		// prints the new value
	fmt.Println("value :", &val)		// prints the memory address of the value
}

// with integer parameter as a copy of the value, therefore changes doesn't matter
func zeroVal(val int){
	val = 0
}

// with pointer integer parameter so that whatever address is passed here will reassign 0 to that address
func zeroPtr(val *int){
	*val = 0
}
