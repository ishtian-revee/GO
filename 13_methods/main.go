package main

import "fmt"

type rect struct {
	width, height int
}

func main(){
	r := rect{width:20, height:30}
	fmt.Println("Area:", r.area())
	fmt.Println("Perimeter:", r.perimeter())

	r2 := &r
	fmt.Println("Area:", r2.area())
	fmt.Println("Perimeter:", r2.perimeter())
}

// this method has a receiver type of *rect
func (r *rect) area() int{
	return r.width * r.height
}

// this method has a receiver type value receiver
func (r rect) perimeter() int{
	return 2*r.width + 2*r.height
}
