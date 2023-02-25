package main

import "fmt"

var n int

func printHistogram(v int) {
	for i := 1; i <= v; i++ {
		fmt.Print("=")
	}
	fmt.Println()
}

func main() {
	fmt.Scanln(&n)
	var v int
	for i := 1; i <= n; i++ {
		fmt.Scanln(&v)
		printHistogram(v)
	}
}
