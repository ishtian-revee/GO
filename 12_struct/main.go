package main

import "fmt"

// basic struct type
type person struct {
	name string
	age int
}

func main(){
	fmt.Println(person{"revee", 22})

	fmt.Println(person{name:"rev", age:18})

	fmt.Println(person{name:"ishtian"})

	s := person{name:"rahim", age:30}
	fmt.Println("Name:", s.name)
	fmt.Println("Age:", s.age)

	st := &s
	fmt.Println("Name:", st.name)
	fmt.Println("Age:", st.age)
}
