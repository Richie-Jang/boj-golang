package main

import (
	"fmt"
)

type pair struct {
	dir  int
	dist int
}

var data [6]pair

func main() {
	var n int
	fmt.Scanln(&n)

	for i := 1; i <= 6; i++ {
		var dir, v int
		fmt.Scanln(&dir, &v)
		data[i-1] = pair{dir, v}
	}

	maxwidth := 0
	maxheight := 0

	for i := 0; i <= 5; i++ {
		cur := data[i].dist
		if i%2 == 0 {
			if maxwidth < cur {
				maxwidth = cur
			}
		} else {
			if maxheight < cur {
				maxheight = cur
			}
		}
	}

	small_width := 0
	small_height := 0

	for i := 0; i <= 5; i++ {
		if i%2 == 0 {
			if data[(i+5)%6].dist+data[(i+1)%6].dist == maxheight {
				small_width = data[i].dist
			}
		} else {
			if data[(i+5)%6].dist+data[(i+1)%6].dist == maxwidth {
				small_height = data[i].dist
			}
		}
	}

	res := ((maxheight * maxwidth) - (small_height * small_width)) * n

	fmt.Println(res)

}
