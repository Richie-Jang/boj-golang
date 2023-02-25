package main

import (
	"bufio"
	"fmt"
	"os"
)

var n, m int
var grid [][]byte

// 48 -> 0
type point struct {
	y, x int
}

var queue []point

var next_poss = [4]point{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}

func main() {
	rd := bufio.NewReader(os.Stdin)
	fmt.Fscanln(rd, &n, &m)

	grid = make([][]byte, n)

	for i := 0; i < n; i++ {
		grid[i] = make([]byte, m)
		var line string
		fmt.Fscanln(rd, &line)
		for j := 0; j < m; j++ {
			grid[i][j] = line[j]
		}
	}

	for i := 0; i < m; i++ {
		queue = append(queue, point{0, i})
	}

	var is_found = false

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		grid[p.y][p.x] = '2'
		if p.y == n-1 {
			is_found = true
			break
		}
		for j := 0; j < 4; j++ {
			np := next_poss[j]
			npx := np.x + p.x
			npy := np.y + p.y
			if npx < 0 || npx >= m || npy < 0 || npy >= n {
				continue
			}
			if grid[npy][npx] == '0' {
				queue = append(queue, point{npy, npx})
			}
		}
	}

	if is_found {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
