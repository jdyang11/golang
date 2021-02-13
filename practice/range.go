package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
        // i is index, v is value, range pow is for all entries in slice pow 
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
