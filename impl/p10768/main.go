package main

import (
	"fmt"
)

var m, d int

func main() {
	fmt.Scanln(&m)
	fmt.Scanln(&d)
	if m < 2 {
		fmt.Println("Before")
	} else if m > 2 {
		fmt.Println("After")
	} else {
		// m == 2
		if d < 18 {
			fmt.Println("Before")
		} else if d > 18 {
			fmt.Println("After")
		} else {
			fmt.Println("Special")
		}
	}
}
