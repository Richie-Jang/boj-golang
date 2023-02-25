package main

import (
	"bufio"
	"fmt"
	"os"
)

func lower_bound(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left < right {
		mid := (left + right) / 2
		//fmt.Println("left", left, "right", right, "mid", mid, "target", target)
		if arr[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	//fmt.Println("Return =>", left)
	return left
}

func main() {
	rd := bufio.NewReader(os.Stdin)

	for {
		var n int
		_, err := fmt.Fscanln(rd, &n)
		if err != nil {
			break
		}

		data := make([]int, 0)
		for i := 0; i < n; i++ {
			var v int
			fmt.Fscan(rd, &v)
			if i == 0 {
				data = append(data, v)
			} else {
				last := data[len(data)-1]
				if last < v {
					data = append(data, v)
				} else {
					// lower bound
					lb := lower_bound(data, v)
					data[lb] = v
				}
			}
		}
		fmt.Fscanln(rd)
		fmt.Println(len(data))
		//fmt.Println(data)
	}

}
