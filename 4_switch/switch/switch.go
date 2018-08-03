package _switch

import (
	"fmt"
	"time"
)

// basic switch statements
func BasicSwitch(num int){
	switch num{
	case 5:
		fmt.Println("The value is five.")
	case 10:
		fmt.Println("The value is Ten.")
	default:
		fmt.Println("The value is niether five nor ten.")
	}
}

// switch statement with multiple cases
func GetDay(){
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("WEEKENDS...CHEERS...!")
	default:
		fmt.Println("ohh... weeeeeeekkkdaaayysss...! -_-")
	}
}

// switch without statement
func GetTime(){
	t := time.Now()
	switch {
	case t.Hour() < 22:
		fmt.Println("Still have time to work.")
	case t.Hour() > 22 && t.Hour() <= 24:
		fmt.Println("Movies and TV series time!")
	default:
		fmt.Println("Get to bed.")
	}
}