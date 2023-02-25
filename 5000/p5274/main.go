package main

import (
	"fmt"
)

func solve(n int) {
	var sum = 0
	for i := n; i >= 1; i-- {
		j := n - i + 1
		sum += (j * j)
	}
	fmt.Println(sum)
}

func main() {

	for {
		var v int
		fmt.Scanf("%d\n", &v)
		if v == 0 {
			break
		}
		solve(v)
	}

}
