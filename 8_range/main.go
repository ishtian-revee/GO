package main

import "fmt"

func main(){
	// basic usage of range in slice without the need of the index value
	slc := []int{11, 12, 13}
	for _, val := range slc{
		fmt.Println("Value are:", val)
	}

	// basic usage of range in slice with the index and as well as value
	slc2 := []string{"captain", "batman", "flash"}
	for i, val := range slc2{
		fmt.Printf("Value in %d: %s\n", i, val)
	}

	// range can iterate over maps
	map1 := map[int]string{1:"one", 2:"two"}
	for key, val := range map1 {
		fmt.Printf("Key %d: %s\n", key, val)
	}

	// range can also iterate over just key of a map
	for key := range map1{
		fmt.Println("Key:", key)
	}

	// range on strings iterates over unicode code points
	for i, c := range "revee"{
		fmt.Println(i, c)
	}
}
