package main

import (
	"fmt"
)

func main() {

	var a,b int
	fmt.Scanln(&a, &b)

	res := a
	
	func doing (c int) {
		if c >= b {
			aa := c / b
			res += aa
			doing (aa)
		}
	}

	

}