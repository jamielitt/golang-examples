package main

import "fmt"

func main() {

	i := 3

	// Simplest of conditions
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Standard look we all know
	for j := 0; j <= 3; j++ {
		fmt.Println(j)
	}

	// "Do this n times, starting at 0"
	for k := range 3 {
		fmt.Println("range", k)
	}

	// Infinite look that needs breaking out of
	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
