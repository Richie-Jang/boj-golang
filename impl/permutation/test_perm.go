package main1

import "fmt"

var loop func(int, int)

func main() {
	var data = []int{1, 2, 3, 4, 5, 6}

	swap := func(i1, i2 int) {
		var t = data[i1]
		data[i1] = data[i2]
		data[i2] = t
	}

	var totalCount = 0
	loop = func(start, end int) {
		if start == end {
			for i := 0; i < len(data); i++ {
				fmt.Print(data[i], " ")
			}
			fmt.Println()
			totalCount++
			return
		}
		for i := start; i <= end; i++ {
			swap(start, i)
			loop(start+1, end)
			swap(start, i)
		}
	}

	loop(0, len(data)-1)
	fmt.Println("Total Count :", totalCount)

}
