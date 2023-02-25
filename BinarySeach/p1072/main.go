package main

import (
	"fmt"
	"math"
)

var x, y int64

const MAX = 1000000000

func percent(a, b int64) int {
	a1 := float64(a)
	b1 := float64(b)
	c := b1 / a1 * 100.0
	return int(math.Trunc(c))
}

func solve(curPercent int) {
	low := 0
	high := MAX
	for low <= high {
		mid := int64(low + (high-low)/2)
		c := percent(x+mid, y+mid)
		//fmt.Println("low:", low, "high:", high, "mid:", mid, "cur:", curPercent, "new:", c)
		if c > curPercent {
			high = int(mid - 1)
		} else {
			low = int(mid + 1)
		}
	}
	fmt.Println(low)
}

func main() {

	fmt.Scanln(&x, &y)

	c := percent(x, y)
	if c >= 99 {
		fmt.Println("-1")
	} else {
		solve(c)
	}

}
