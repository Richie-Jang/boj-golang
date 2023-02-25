package main

import (
	"fmt"
	"sort"
)

var n int
var data [5]int

func main() {
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		for j := 0; j < 5; j++ {
			fmt.Scan(&data[j])
		}
		// sort
		sort.Ints(data[:])
		v1 := data[1]
		v2 := data[3]
		if v2-v1 >= 4 {
			fmt.Println("KIN")
		} else {
			sum := 0
			for j := 1; j <= 3; j++ {
				sum += data[j]
			}
			fmt.Println(sum)
		}
	}
}
