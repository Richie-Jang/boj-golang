package main

import "fmt"

func lower_bound(arr []int, target int) int {
	start := 0
	end := len(arr) - 1
	for start < end {
		mid := start + (end-start)/2
		if arr[mid] >= target {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return start
}

func main() {

	t1 := []int{10, 20, 30, 40, 50}

	fmt.Println(lower_bound(t1, 40))
	//t2 := []int {10,20,30,30,30,40,50}

}
