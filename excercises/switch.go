package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("text not defined")
	}

	switch time.Now().Weekday() {

	// use of comma to separate multiple expressions
	case time.Saturday, time.Sunday:
		fmt.Println("Today is the weekend")
	default:
		fmt.Println("Today is a weekday")
	}

	t := time.Now()
	// switch without expression is an alternative to if/then/else logic
	switch {
	case t.Hour() < 12:
		fmt.Println("It's the morning")
	default:
		fmt.Println("It's the afternoon / Evening")
	}
}
