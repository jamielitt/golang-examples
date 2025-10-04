package main

import "fmt"

func main() {

	// Declaring an array
	var a [5]int
	fmt.Println("emp:", a)

	// Setting an element of the array
	a[2] = 100
	fmt.Println("set: ", a)
	fmt.Println("get:", a[2])
	fmt.Println("len: ", len(a))

	// Declare and initialise an array in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	// the compiler can count the number of elements too with ... :)
	c := [...]int{2, 3, 4, 5, 6, 7, 8}
	fmt.Println("dcl: ", c)

	// if you specify the index with :, the element will be zeroed
	// This "fills" that position with zeros up to the count specified, e.g. the below will output:
	// idx:  [100 3 400 0 0 500]
	d := [...]int{100, 3, 400, 5: 500}
	fmt.Println("idx: ", d)

	// Multi dimensional arrays
	e := [2][2]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println("2d: ", e)
}
