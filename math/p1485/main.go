package main

import (
	"fmt"
	"sort"
)

var (
	n   int
	xs  [4]int
	ys  [4]int
	res []int64
)

func distance(x1, y1, x2, y2 int) int64 {
	xx := int64((x2 - x1) * (x2 - x1))
	yy := int64((y2 - y1) * (y2 - y1))
	return xx + yy
}

func main() {
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		for j := 0; j < 4; j++ {
			var x, y int
			fmt.Scanln(&x, &y)
			xs[j] = x
			ys[j] = y
		}
		res = make([]int64, 0)
		for j := 0; j < 4; j++ {
			for k := j + 1; k < 4; k++ {
				res = append(res, distance(xs[j], ys[j], xs[k], ys[k]))
			}
		}
		sort.Slice(res, func(i1, i2 int) bool {
			return res[i1] < res[i2]
		})

		if res[0] == res[1] && res[1] == res[2] && res[2] == res[3] && res[4] == res[5] {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	}

}
