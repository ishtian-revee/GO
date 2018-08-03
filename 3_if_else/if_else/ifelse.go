package if_else

import "fmt"

// basic if/else statement
func IsEven(num int){

	if num % 2 == 0{
		fmt.Println("The number is even.")
	}else{
		fmt.Println("The number is odd.")
	}
}

// short condition in if statement
func IsSmall(num int){

	if temp := 5; temp < num{
		fmt.Println("The number is large.")
	}else{
		fmt.Println("The number is small.")
	}
}
