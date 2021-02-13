package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
	
	s = append(s, 5, 6, 7)
	printSlice(s)
	
	s = append(s, 8, 9, 10, 11, 12)
	printSlice(s)
	
	a := make([]int, 5)
	printSlice(a)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

/*
output:

len=0 cap=0 []
len=1 cap=1 [0]
len=2 cap=2 [0 1]
len=5 cap=6 [0 1 2 3 4]
len=8 cap=12 [0 1 2 3 4 5 6 7]
len=13 cap=24 [0 1 2 3 4 5 6 7 8 9 10 11 12]
len=5 cap=5 [0 0 0 0 0]
*/
