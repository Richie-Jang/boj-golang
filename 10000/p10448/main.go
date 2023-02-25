package main

import (
	"fmt"
)

const MAX = 1000

var m []int

func solve(k int) int {
	for i := 1; i < len(m); i++ {
		v1 := k - m[i]
		if v1 <= 0 {
			break
		}
		for j := 1; j < len(m); j++ {
			v2 := v1 - m[j]
			if v2 <= 0 {
				break
			}
			for l := len(m) - 1; l >= 1; l-- {
				v3 := v2 - m[l]
				if v3 == 0 {
					//fmt.Println(i, "=", m[i], j, "=", m[j], l, "=", m[l])
					return 1
				}
				if v3 > 0 {
					break
				}
			}
		}
	}
	return 0
}
func main() {

	m = make([]int, 45)
	for i := 1; i <= 44; i++ {
		v := i * (i + 1) / 2
		m[i] = v
	}

	var n int
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {

		var k int
		fmt.Scanln(&k)

		fmt.Println(solve(k))

	}
}
