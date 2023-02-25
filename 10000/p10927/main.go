package main

import (
	"fmt"
	"strconv"
	"strings"
)

var data []int

func recursive(pos, limit int, buf []int) {
	if pos == limit {
		strs := make([]string, len(buf))
		for index, v := range buf {
			strs[index] = strconv.Itoa(v)
		}
		fmt.Println(strings.Join(strs, " "))
		return
	}

	var swap = func(i, j int) {
		t := data[i]
		data[i] = data[j]
		data[j] = t
	}

	for i := pos; i < len(data); i++ {
		swap(pos, i)
		recursive(pos+1, limit, append(buf, data[pos]))
		swap(pos, i)
	}
}

func main() {
	var n int
	//fmt.Scanln(&n)
	n = 4

	data = make([]int, n)

	// for i := 0; i < n; i++ {
	// 	fmt.Scanf("%d", &data[i])
	// }
	data[0] = 1
	data[1] = 2
	data[2] = 3
	data[3] = 4

	recursive(0, len(data), make([]int, 0))
}
