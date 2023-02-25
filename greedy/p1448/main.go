package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	n   int
	arr []int
	res = -1
)

func main() {
	br := bufio.NewReader(os.Stdin)
	fmt.Fscanf(br, "%d\n", &n)
	arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(br, "%d\n", &arr[i])
	}
	sort.Ints(arr)

	i1 := n - 1
	i2 := n - 2
	i3 := n - 3

	for i1 >= 2 {
		a := arr[i1]
		if i2 < 1 {
			break
		}
		b := arr[i2]
		if i3 < 0 {
			i1--
			i2 = i1 - 1
			i3 = i2 - 1
		} else {
			c := arr[i3]
			if a < b+c {
				res = a + b + c
				break
			}
			i3--
		}
	}

	fmt.Println(res)
}
