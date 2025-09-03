// Interesting point, there are no ternary operators in Go, a full statement
// is expected, even for basic conditions
package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is an Even number")
	} else {
		fmt.Println("7 is an odd number")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// In if statements, you can declare a variable too
	if num := 9; num < 0 {
		fmt.Println("num is negative")
	} else if num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("num is greater than 10")
	}
}
