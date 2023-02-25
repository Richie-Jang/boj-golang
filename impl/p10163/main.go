package main

import "fmt"

var n int
var grid [101][101]int

func main() {
	fmt.Scanln(&n)
	var y, x, w, h int
	for i := 1; i <= n; i++ {
		fmt.Scanln(&x, &y, &w, &h)
		for r := y; r <= y+h-1; r++ {
			for c := x; c <= x+w-1; c++ {
				grid[r][c] = i
			}
		}
	}

	dict := make(map[int]int)
	for i := 0; i <= n; i++ {
		dict[i] = 0
	}

	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100; j++ {
			hh := grid[i][j]
			dict[hh] = (dict[hh] + 1)
		}
	}

	for i := 1; i <= n; i++ {
		fmt.Println(dict[i])
	}
}
