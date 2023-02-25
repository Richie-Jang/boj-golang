package main

import (
	"bufio"
	"fmt"
	"os"
)

var n int
var axis, q1, q2, q3, q4 int

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)
	var x, y int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &x, &y)
		if x == 0 || y == 0 {
			axis++
		} else {
			if x > 0 && y > 0 {
				q1++
			} else if x > 0 && y < 0 {
				q4++
			} else if x < 0 && y > 0 {
				q2++
			} else {
				q3++
			}
		}
	}
	fmt.Printf("Q1: %d\n", q1)
	fmt.Printf("Q2: %d\n", q2)
	fmt.Printf("Q3: %d\n", q3)
	fmt.Printf("Q4: %d\n", q4)
	fmt.Printf("AXIS: %d\n", axis)
}
