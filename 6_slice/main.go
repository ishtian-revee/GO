package main

import "fmt"

func main(){
	// slice declaration
	slc := make([]string, 3)
	fmt.Println("slice:", slc)

	// assign slice indexes
	slc[0] = "batman"
	slc[1] = "ironman"
	slc[2] = "thor"
	fmt.Println("slice:", slc)
	fmt.Println("slice:", slc[2])

	// length of a slice
	fmt.Println("length of a slice:", len(slc))

	// adding/append more values to a certain slice
	slc = append(slc, "superman")
	slc = append(slc, "hulk")
	fmt.Println("after appending slice:", slc)

	// slice[low:high]
	a := slc[2:4]
	fmt.Println("part of slice:", a)
	b := slc[:3]
	fmt.Println("part of slice:", b)
	c := slc[2:]
	fmt.Println("part of slice:", c)

	// copy another slice
	slc2 := make([]string, len(slc))
	copy(slc2, slc)
	fmt.Println("First Slice:", slc)
	fmt.Println("Copy Slice:", slc2)

	// declare and initialize in one line
	slc3 := []string{"legen", "darry"}
	fmt.Println("slice 3:", slc3)

	// multi dimensional slice
	mulSlc := make([][]int, 3)
	for i:=0; i<len(mulSlc); i++ {
		innerSlice := i+1
		mulSlc[i] = make([]int, innerSlice)
		for j:=0; j<innerSlice; j++ {
			mulSlc[i][j] = i+j
		}
	}
	fmt.Println("Multi dimensional array:", mulSlc)
}
