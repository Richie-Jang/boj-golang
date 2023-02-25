package main

import (
	"fmt"
)

var n int

func main() {
	fmt.Scanln(&n)
	m := n*2 - 1
	for i := 1; i <= n; i++ {
		var count = 0
		for j := 1; j <= i-1; j++ {
			fmt.Print(" ")
			count++
		}
		for j := 1; j <= m-count*2; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
