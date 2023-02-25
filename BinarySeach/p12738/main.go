package main

import (
	"bufio"
	"fmt"
	"os"
)

func lowerBound(arr []int, target int) int {
	low := 0
	upper := len(arr) - 1
	for low < upper {
		mid := (low + upper) / 2
		if target <= arr[mid] {
			upper = mid
		} else {
			low = mid + 1
		}
	}
	return upper
}

func main() {

	var n int
	rd := bufio.NewReader(os.Stdin)
	fmt.Fscanln(rd, &n)

	var v int
	arr := make([]int, 1)
	arr[0] = -1000000001
	for i := 0; i < n; i++ {
		fmt.Fscanf(rd, "%d", &v)
		last := arr[len(arr)-1]
		if last < v {
			arr = append(arr, v)
		} else {
			f := lowerBound(arr, v)
			arr[f] = v
		}
	}
	fmt.Println(len(arr) - 1)

}
