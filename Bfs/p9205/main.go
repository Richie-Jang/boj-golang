package main

import (
	"bufio"
	"fmt"
	"os"
)

type t2 struct {
	x, y int
}

var (
	t, n            int
	data            []t2
	homePos, endPos t2
)

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func dist(a, b t2) int {
	x1 := abs(a.x - b.x)
	y1 := abs(a.y - b.y)
	return x1 + y1
}

func solve(bw *bufio.Writer) {

	checkRule := func(a, b t2) bool {
		d := dist(a, b)
		return d <= 1000
	}

	okPrint := func() {
		fmt.Fprintln(bw, "happy")
	}
	failPrint := func() {
		fmt.Fprintln(bw, "sad")
	}

	q := []int{0}
	ck := make([]bool, n+2)
	ck[0] = true

	for len(q) > 0 {
		c := q[0]
		q = q[1:]
		cur := data[c]

		if checkRule(cur, endPos) {
			okPrint()
			return
		}

		for i := 1; i < n+1; i++ {
			if !ck[i] {
				d := dist(cur, data[i])
				if d <= 1000 {
					q = append(q, i)
					ck[i] = true
				}
			}
		}
	}
	failPrint()
}

func main() {
	br := bufio.NewReader(os.Stdin)
	bw := bufio.NewWriter(os.Stdout)
	fmt.Fscan(br, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(br, &n)
		data = make([]t2, 0)
		for j := 0; j < n+2; j++ {
			a, b := 0, 0
			fmt.Fscan(br, &a, &b)
			data = append(data, t2{a, b})
			switch j {
			case 0:
				homePos = data[j]
			case n + 1:
				endPos = data[j]
			}
		}
		solve(bw)
	}

	bw.Flush()
}
