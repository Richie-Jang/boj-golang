package main

import "fmt"

var n, m int
var buckets [101]int

func main() {
	fmt.Scanln(&n, &m)
	for i := 1; i <= n; i++ {
		buckets[i] = i
	}
	var a, b int
	for i := 0; i < m; i++ {
		fmt.Scanln(&a, &b)
		t := buckets[a]
		buckets[a] = buckets[b]
		buckets[b] = t
	}

	for i := 1; i <= n; i++ {
		if i == n {
			fmt.Print(buckets[i])
		} else {
			fmt.Print(buckets[i], " ")
		}
	}
	fmt.Println()
}
