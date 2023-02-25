package main

import (
	"bufio"
	"fmt"
	"os"
)

var n int
var data []int

func lower_bound(search, left, right int) int {
	for left < right {
		mid := (left + right) / 2
		if data[mid] >= search {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	fmt.Fscanln(rd, &n)
	var d int
	fmt.Fscan(rd, &d)
	data = append(data, d)
	res := 0
	for i := 1; i < n; i++ {
		fmt.Fscan(rd, &d)
		l := data[len(data)-1]
		if l > d {
			data = append(data, d)
			continue
		}

		lb := lower_bound(d, 0, len(data)-1)
		data[lb] = d
		res++
	}

	fmt.Println(res)
}
