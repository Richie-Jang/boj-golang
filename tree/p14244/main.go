package main

import "fmt"

var n, m int

func main() {

	fmt.Scan(&n, &m)

	lastNode := 0
	k := n - m
	for i := 1; i < n; i++ {
		fmt.Println(lastNode, i)
		if i == k && m != 2 {
			lastNode = i - 1
		} else {
			lastNode = i
		}
	}
}
