package main

import (
	"fmt"
)

var n int
var data []int

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func solve() {
	var gcd_val = 0
	switch n {
	case 2:
		gcd_val = gcd(data[0], data[1])
	case 3:
		g1 := gcd(data[0], data[1])
		gcd_val = gcd(g1, data[2])
	}

	for i := 1; i <= gcd_val; i++ {
		if gcd_val%i == 0 {
			fmt.Println(i)
		}
	}
}

func main() {
	fmt.Scanln(&n)
	data = make([]int, n)
	for i := 0; i < n; i++ {
		var s int
		fmt.Scan(&s)
		data[i] = s
	}
	solve()
}
