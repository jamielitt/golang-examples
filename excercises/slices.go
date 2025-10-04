package main

import (
	"fmt"
	"slices"
	"strconv"
)

func main() {
	var s []string
	fmt.Println("uninitialised: ", s == nil, len(s) == 0)

	// Making a slice of 3 elements (which will be non zero length)
	// (initially zero value)
	// We can pass a capacity value to the slice if required
	s = make([]string, 3)
	fmt.Println("emp: ", s, "len: ", len(s), "cap: ", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])
	fmt.Println("len: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("app: ", s)

	// Slices can also be copied
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ", c)

	// Support for a "slice" operator - slice[low:high]
	// (starts at c for the count of high - slice 2 and 3
	// cpy:  [a b c d e f]
	//. slc:  [c d]
	d := s[2:4]
	fmt.Println("slc: ", d)

	// Slice up to (but excluding last)
	upto := s[:4]
	fmt.Println("upt: ", upto)

	// Slice up from
	upfrom := s[3:]
	fmt.Println("upf: ", upfrom)

	// Declare and initialise a slice in a single line
	dcl := []string{"1", "2", "3"}
	fmt.Println("dcl: ", dcl)

	// Make a copy so we can do something with it
	dcl2 := make([]string, len(dcl))
	copy(dcl2, dcl)

	// Using slices package, we can check for equality
	equality := slices.Equal(dcl, dcl2)
	fmt.Println("eq: ", equality)

	// Multi dimension slices, inner dimension can be different in length which
	// is unlike normal multi dimensional arrays. This is demonstrated here
	multi := make([][]string, 3)
	for i := range 3 {
		inner := i + 1
		multi[i] = make([]string, inner)
		for j := range inner {
			multi[i][j] = strconv.Itoa((i + j))
		}
	}
	fmt.Println("mul: ", multi)
}
