package main

import (
	"fmt"
)

func solve(w, h, p, q, t int) {
	var xdir = 1
	var ydir = 1
	for t >= 1 {
		t--
		if p == w || p == 0 {
			xdir *= -1
		}
		if q == 0 || q == h {
			ydir *= -1
		}
		p += xdir
		q += ydir
	}

	fmt.Println(p, q)
}

func main() {

	var w, h int
	fmt.Scanln(&w, &h)
	var p, q int
	fmt.Scanln(&p, &q)
	var t int
	fmt.Scanln(&t)

	solve(w, h, p, q, t)
}
