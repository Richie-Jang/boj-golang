package main

import (
	"fmt"
)

func main() {

	var n, k int
	fmt.Scanln(&n, &k)
	res := make([]int, 0)
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			res = append(res, i)
		}
	}
	if len(res) <= k-1 {
		fmt.Println("0")
	} else {
		fmt.Println(res[k-1])
	}
}
