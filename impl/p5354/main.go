package main

import (
	"fmt"
)

var n, v int

func printVal(v int) {
	for j := 1; j <= v; j++ {
		if j == 1 || j == v {
			for i := 1; i <= v; i++ {
				fmt.Print("#")
			}
			fmt.Println()
		} else {
			fmt.Print("#")
			for k := 1; k <= v-2; k++ {
				fmt.Print("J")
			}
			fmt.Print("#")
			fmt.Println()
		}
	}
}

func main() {
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&v)
		printVal(v)
		fmt.Println()
	}
}
