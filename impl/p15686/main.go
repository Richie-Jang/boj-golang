package main

import (
	"bufio"
	"fmt"
	"os"
)

type pos struct {
	y, x int
}

var n, m int
var grid [50][50]int
var chs, selectedChs []pos
var hs []pos
var res = 1 << 30

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func calculateDist(c_pos, h_pos pos) int {
	return abs(c_pos.y-h_pos.y) + abs(c_pos.x-h_pos.x)
}

func chooseChickenHouse(cs, k int) {
	if cs == m {
		// check distance
		sum := 0
		for _, h := range hs {
			minDist := 1 << 30
			for i := 0; i < m; i++ {
				d := calculateDist(selectedChs[i], h)
				if minDist > d {
					minDist = d
				}
			}
			sum += minDist
		}
		if res > sum {
			res = sum
		}
		return
	}
	for i := k; i < len(chs); i++ {
		selectedChs[cs] = chs[i]
		chooseChickenHouse(cs+1, i+1)
	}
}

func main() {
	chs = make([]pos, 0)
	hs = make([]pos, 0)
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n, &m)
	selectedChs = make([]pos, m)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &grid[i][j])
			if grid[i][j] == 2 {
				chs = append(chs, pos{i, j})
			} else if grid[i][j] == 1 {
				hs = append(hs, pos{i, j})
			}
		}
	}

	chooseChickenHouse(0, 0)
	fmt.Println(res)
}
