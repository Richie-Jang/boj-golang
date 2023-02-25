package main

import (
	"fmt"
	"strings"
)

var (
	inp string
)

func checkPal(i int, l int) bool {
	for j := i; j < l; j++ {
		if inp[j] != inp[l+i-j-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Scanf("%s\n", &inp)
	l := len(strings.TrimSpace(inp))
	ans := l

	for i := 0; i < l; i++ {
		if checkPal(i, l) {
			ans += i
			break
		}
	}

	fmt.Println(ans)
}
