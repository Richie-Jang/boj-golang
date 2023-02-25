package test1

import (
	"fmt"
)

func lower_bound(arr []int, key int) int {
	var start = 0
	var end = len(arr)
	for end > start {
		mid := (start + end) / 2
		if arr[mid] < key {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return end
}

func main() {

	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(lower_bound(arr, 7))
}
