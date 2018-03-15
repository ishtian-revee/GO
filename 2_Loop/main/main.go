package main

import(
	"fmt"
	"github.com/ishtian-revee/2_Loop/loop"
)

func main(){
	sum := 10

	ret := loop.Basic(sum)
	fmt.Println("Sum:", ret)

	ret2 := loop.Optional(sum)
	fmt.Println("Sum: ", ret2)

	ret3 := loop.Forever(sum)
	fmt.Println("Sum: ", ret3)
}
