package main

import (
	"bufio"
	"fmt"
	"os"
)

var n int

func lower_bound(data []int, left, right, search int) int {
	for left < right {
		mid := (left + right) / 2
		if data[mid] < search {
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
	data := []int{d}

	var res int

	for i := 1; i < n; i++ {
		fmt.Fscan(rd, &d)
		last := data[len(data)-1]
		if d > last {
			data = append(data, d)
			continue
		}
		lb := lower_bound(data, 0, len(data)-1, d)
		data[lb] = d
		res++
	}

	fmt.Println(res)
}
