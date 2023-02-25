package main

import "fmt"

var (
	a, b int
)

func abs(c int) int {
	if c < 0 {
		return -1 * c
	}
	return c
}

func main() {
	fmt.Scan(&a, &b)
	yd1 := a % 4
	xc1 := false
	if yd1 == 0 {
		yd1 = 4
		xc1 = true
	}
	yd2 := b % 4
	xc2 := false
	if yd2 == 0 {
		yd2 = 4
		xc2 = true
	}
	y := abs(yd1 - yd2)
	xd1 := a / 4
	xd2 := b / 4

	if xc1 {
		xd1--
	}
	if xc2 {
		xd2--
	}

	x := abs(xd1 - xd2)
	fmt.Println(x + y)
}
