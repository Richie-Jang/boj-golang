package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var arr []int

func main() {
	rd := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(rd, &n)
	arr = make([]int, 1)
	// init
	arr[0] = 0
	for i := 0; i < n; i++ {
		var c int
		fmt.Fscan(rd, &c)
		last := arr[len(arr)-1]

		if last < c {
			arr = append(arr, c)
		} else {
			bs := sort.SearchInts(arr, c)
			arr[bs] = c
		}
	}

	fmt.Println(len(arr) - 1)
}
