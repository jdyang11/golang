package main

import "fmt"

func main() {
	a := make([]int, 5)
        for i := range a {
            a[i] = i
        }
	printSlice("a", a)

	b := make([]int, 0, 5)
        f := b[:cap(b)]
        for i := range f {
            f[i] = i + 10
        }
	printSlice("b", b)
	printSlice("f", f)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
	
	printSlice("b", b[:cap(b)])
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

/*
a len=5 cap=5 [0 1 2 3 4]
b len=0 cap=5 []
f len=5 cap=5 [10 11 12 13 14]
c len=2 cap=5 [10 11]
d len=3 cap=3 [12 13 14]
b len=5 cap=5 [10 11 12 13 14]
*/

