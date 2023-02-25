package main

import "fmt"

var (
	a, b int
)

func main() {
	fmt.Scanf("%d %d\n", &a, &b)

	if a <= 1 {
		fmt.Println("0")
		return
	}

	c := a / 2
	if c <= b {
		fmt.Printf("%d\n", c)
	} else {
		fmt.Printf("%d\n", b)
	}
}
