package main

import "fmt"

func main(){
	// array declaration
	var arr [5]int
	fmt.Println(arr)

	var arrStr [5]string
	fmt.Println(arrStr)
	arrStr = [5]string{"Legen...", "wait", "for", "it", "daryyyy."}
	fmt.Println(arrStr)

	// assign value at a certain index
	arr[4] = 11
	fmt.Println("at index 5:", arr[4])
	fmt.Println(arr)

	// finding length of an array
	fmt.Println("length of the array:", len(arr))

	// initialize an array in 1 line
	arr2 := [5]int {2, 4, 6, 8, 10}
	fmt.Println(arr2)

	// multi dimensional array [2D]
	var mulArr[2][3]int
	for i:=0; i<2; i++ {
		for j:=0; j<3; j++ {
			mulArr[i][j] = i+j
		}
	}
	fmt.Println("multi-dimensional array:", mulArr)
}
