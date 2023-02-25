package main

import (
	"fmt"
)

var n, m, result int
var y, x, d int
var grid [50][50]int

//  NextTrun     N  E  S  W
var dy = [4]int{-1, 0, 1, 0}
var dx = [4]int{0, 1, 0, -1}

func leftTurn(d int) int {
	return (d + 3) % 4
}

func start() {
	for {
		isClean := false
		for i := 0; i < 4; i++ {
			d = leftTurn(d)
			yy := y + dy[d]
			xx := x + dx[d]
			if grid[yy][xx] == 0 {
				grid[yy][xx] = 2
				result++
				y = yy
				x = xx
				isClean = true
				break
			}
		}
		if !isClean {
			if grid[y-dy[d]][x-dx[d]] == 1 {
				break
			}
			y = y - dy[d]
			x = x - dx[d]
		}
	}
}

func main() {
	fmt.Scanln(&n, &m)
	fmt.Scanln(&y, &x, &d)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&grid[i][j])
		}
	}
	//start
	grid[y][x] = 2
	result++
	start()
	fmt.Println(result)
}
