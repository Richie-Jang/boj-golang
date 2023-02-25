package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n, m                   int
	grid                   [][]int
	dx                     = []int{-1, 1, 0, 0}
	dy                     = []int{0, 0, -1, 1}
	vs                     [][]bool
	whitePower, blackPower int
)

type t2 struct {
	y, x int
}

func searchPower(y, x int) {
	q := []t2{{y, x}}
	vs[y][x] = true
	isWhite := true
	count := 1
	if grid[y][x] == 0 {
		isWhite = false
	}
	for len(q) > 0 {
		c := q[0]
		q = q[1:]
		// next
		for i := 0; i < 4; i++ {
			ddx := dx[i] + c.x
			ddy := dy[i] + c.y
			if ddx >= 0 && ddx < n && ddy >= 0 && ddy < m && !vs[ddy][ddx] {
				if (isWhite && grid[ddy][ddx] == 1) || (!isWhite && grid[ddy][ddx] == 0) {
					q = append(q, t2{ddy, ddx})
					vs[ddy][ddx] = true
					count++
				}
			}
		}
	}

	power := count * count
	if isWhite {
		whitePower += power
		//fmt.Println("searched white =", y, x, power)
	} else {
		blackPower += power
		//fmt.Println("searched black =", y, x, power)
	}
}

func main() {
	br := bufio.NewReader(os.Stdin)
	fmt.Fscan(br, &n, &m)
	grid = make([][]int, m)
	vs = make([][]bool, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]int, n)
		vs[i] = make([]bool, n)
	}
	var a string
	for i := 0; i < m; i++ {
		fmt.Fscan(br, &a)
		for j := 0; j < len(a); j++ {
			if a[j] == 'W' {
				grid[i][j] = 1
			} else if a[j] == 'B' {
				grid[i][j] = 0
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !vs[i][j] {
				searchPower(i, j)
			}
		}
	}
	fmt.Println(whitePower, blackPower)
}
