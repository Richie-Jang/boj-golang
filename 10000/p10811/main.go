package main

import (
	"fmt"
	"strconv"
	"strings"
)

var n, m int

type inp struct {
	a int
	b int
}

var (
	inps []inp
	arr  []int
)

func reverse(d inp) {

	a := d.a
	b := d.b

	for {
		if a >= b {
			break
		}

		a1 := arr[a]
		a2 := arr[b]
		arr[a] = a2
		arr[b] = a1
		a++
		b--
	}

}

func main() {
	fmt.Scanln(&n, &m)
	inps = make([]inp, m)
	arr = make([]int, n+1)

	for i := 1; i <= n; i++ {
		arr[i] = i
	}
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scanln(&a, &b)
		inps[i] = inp{a, b}
	}

	for _, v := range inps {
		reverse(v)
	}

	strs := make([]string, n+1)
	for i, v := range arr {
		strs[i] = strconv.FormatInt(int64(v), 10)
	}
	res := strings.Join(strs[1:], " ")
	fmt.Println(res)

}
