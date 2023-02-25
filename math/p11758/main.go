package main

import (
	"fmt"
)

type pair struct {
	x, y int
}

func main() {

	var ps [3]pair

	for i := 0; i < 3; i++ {
		var x, y int
		fmt.Scanln(&x, &y)
		ps[i] = pair{x, y}
	}

	v1 := (ps[1].x - ps[0].x) * ps[2].y
	v2 := (ps[1].y-ps[0].y)*(ps[2].x-ps[0].x) + ps[0].y*(ps[1].x-ps[0].x)

	var res int
	if v1 == v2 {
		res = 0
	} else if v1 > v2 {
		res = 1
	} else {
		res = -1
	}
	fmt.Println(res)
}
