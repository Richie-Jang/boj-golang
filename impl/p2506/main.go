package main

import (
	"fmt"
)

var n int

func main() {
	fmt.Scanln(&n)
	var v int
	var sum int
	lastV := 0
	for i := 1; i <= n; i++ {
		fmt.Scan(&v)
		if lastV == 0 {
			if v == 1 {
				lastV = 1
			} else {
				lastV = 0
			}
		} else {
			if v == 1 {
				lastV++
			} else {
				lastV = 0
			}
		}
		sum += lastV
	}
	fmt.Println(sum)
}
