package main

import (
	"fmt"
)

var (
	l, r string
)

func main() {
	fmt.Scanf("%s %s\n", &l, &r)

	res := 0

	if len(l) == len(r) {

		for i := 0; i < len(l); i++ {
			if l[i] != r[i] {
				break
			}
			if l[i] == r[i] && l[i] == '8' {
				res++
				continue
			}

		}
	}

	fmt.Println(res)
}
