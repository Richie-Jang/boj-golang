package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	n    int
	grid [][]int
	res  int
)

func getNumOn2Grid(g []int) int {
	sort.Ints(g)
	return g[2]
}

func rec(g [][]int, size int, x, y int, g2 [][]int) {
	if size == 2 {
		res = getNumOn2Grid([]int{
			g[y][x],
			g[y][x+1],
			g[y+1][x],
			g[y+1][x+1],
		})
		return
	}

	if x >= size {
		rec(g, size, 0, y+2, g2)
		return
	}

	if y >= size {
		hh := len(g2) / 2
		g3 := make([][]int, hh)
		for i := 0; i < hh; i++ {
			g3[i] = make([]int, hh)
		}
		rec(g2, len(g2), 0, 0, g3)
		return
	}

	ny := y / 2
	nx := x / 2

	var gs []int = []int{
		g[y][x],
		g[y][x+1],
		g[y+1][x],
		g[y+1][x+1],
	}

	g2[ny][nx] = getNumOn2Grid(gs)
	rec(g, size, x+2, y, g2)

}

func main() {
	br := bufio.NewReader(os.Stdin)
	fmt.Fscanf(br, "%d\n", &n)
	grid = make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscanf(br, "%d", &grid[i][j])
		}
		fmt.Fscanln(br)
	}

	g2 := make([][]int, n/2)
	for i := 0; i < n/2; i++ {
		g2[i] = make([]int, n/2)
	}
	rec(grid, n, 0, 0, g2)
	fmt.Println(res)
}
