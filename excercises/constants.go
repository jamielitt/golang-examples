package main

import (
	"fmt"
	"math"
)

const constantString = "Hello, this is a constant string"

func main() {

	fmt.Println(constantString)

	const n = 500000000
	const d = 3e20 / n

	fmt.Println(n)

	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
