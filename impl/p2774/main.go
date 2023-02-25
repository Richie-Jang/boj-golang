package main

import (
	"fmt"
)

var n int

func main() {
	fmt.Scanln(&n)
	for r := 0; r < n; r++ {
		var s string
		fmt.Scanln(&s)
		dict := make(map[byte]int)
		for i := 0; i <= 9; i++ {
			dict[byte(i)] = 0
		}
		for i := 0; i < len(s); i++ {
			vv := byte(s[i] - '0')
			dict[vv] = (dict[vv] + 1)
		}

		count := 0
		for i := 0; i <= 9; i++ {
			if dict[byte(i)] > 0 {
				count++
			}
		}

		fmt.Println(count)
	}
}
