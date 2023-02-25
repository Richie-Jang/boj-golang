package main

import (
	"fmt"
	"math"
)

var (
	abit, bbit int
	n          int
	arr        []int
)

func converting() {
	count := 0
	sum := int64(0)
	for v := len(arr) - 1; v >= 0; v-- {
		tt := int64(math.Pow(float64(abit), float64(count)))
		sum += int64(arr[v]) * tt
		count++
	}
	barr := make([]int, 0)
	for {
		vv1 := sum / int64(bbit)
		vv2 := sum % int64(bbit)
		barr = append(barr, int(vv2))
		if vv1 <= 0 {
			break
		}
		sum = vv1
	}
	for j := len(barr) - 1; j >= 0; j-- {
		fmt.Print(barr[j], " ")
	}
	fmt.Println()
}

func main() {
	fmt.Scanln(&abit, &bbit)
	fmt.Scanln(&n)

	arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	converting()
}
