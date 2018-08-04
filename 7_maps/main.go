package main

import "fmt"

func main(){
	// map (associative data/key-value pair) declaration
	mp := make(map[string]int)
	fmt.Println("map:", mp)

	// set key/value
	mp["key1"] = 11
	mp["key2"] = 19
	mp["key3"] = 61
	fmt.Println("key 1:", mp["key1"])
	fmt.Println("map:", mp)

	// getting value from a map
	val := mp["key3"]
	fmt.Println("value 3:", val)

	// length of a map
	fmt.Println("length of a map:", len(mp))

	// deleting a key-value pair
	delete(mp, "key2")
	fmt.Println("after deleting map:", mp)

	// optional second return value when getting a value from a map indicates if the key was present in the map
	_, prs := mp["key3"]
	fmt.Println("prs:", prs)

	// declare and initialization in one line
	mp2 := map[string]int{"rev": 11, "pro": 2}
	fmt.Println("map 2:", mp2)
}
