package main

import "fmt"

var n, a, b int

func main() {
	fmt.Scan(&n)
	a = 100
	b = 100
	var c1, c2 int
	for i := 0; i < n; i++ {
		fmt.Scan(&c1, &c2)
		if c1 < c2 {
			a -= c2
		} else if c1 > c2 {
			b -= c1
		}
	}
	fmt.Println(a)
	fmt.Println(b)
}
