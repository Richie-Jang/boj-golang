package main

import (
	"fmt"
)

var n, l, r int
var grid [][]int
var vs [][]bool

type pos struct {
	x, y int
}

func searchUnionAndUpdateGrid() bool {

	vs = make([][]bool, n)

	for i := 0; i < n; i++ {
		vs[i] = make([]bool, n)
	}

	nextMoves := [4][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	bfs := func(cx, cy int, ps *[]pos) {
		queue := make([]pos, 0)
		vs[cy][cx] = true
		queue = append(queue, pos{cx, cy})
		//while
		for {
			if len(queue) == 0 {
				//fmt.Printf("now break, %d,%d => %v\n", cx, cy, *ps)
				break
			}
			cur := queue[0]
			queue = queue[1:]
			cv1 := grid[cur.y][cur.x]

			for _, nn := range nextMoves {
				nx := nn[1] + cur.x
				ny := nn[0] + cur.y
				if nx >= 0 && nx < n && ny >= 0 && ny < n && !vs[ny][nx] {
					cv2 := grid[ny][nx]
					cv3 := cv1 - cv2
					if cv3 < 0 {
						cv3 = cv3 * -1
					}
					if l <= cv3 && cv3 <= r {
						//fmt.Printf("adding: %v\n", *ps)
						*ps = append(*ps, pos{x: nx, y: ny})
						vs[ny][nx] = true
						queue = append(queue, pos{nx, ny})
					}
				}
			}
		}
	}

	// new
	ngrid := make([][]int, n)
	for i := 0; i < n; i++ {
		ngrid[i] = make([]int, n)
		copy(ngrid[i], grid[i])
	}

	mergePeople := func(ps []pos) {
		var sum int
		for _, p := range ps {
			sum += grid[p.y][p.x]
		}
		nv := sum / len(ps)
		for _, p := range ps {
			ngrid[p.y][p.x] = nv
		}
	}

	res := false

	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			if !vs[y][x] {
				// search
				//fmt.Println("searching : ", x, y)
				ps := make([]pos, 0)
				ps = append(ps, pos{x, y})
				bfs(x, y, &ps)

				if len(ps) >= 2 {
					if !res {
						res = true
					}
					mergePeople(ps)
				}
			}
		}
	}

	if res {
		// change
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				grid[i][j] = ngrid[i][j]
			}
		}
	}

	return res
}

func main() {

	fmt.Scanln(&n, &l, &r)
	grid = make([][]int, n)
	vs = make([][]bool, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
		var c int
		for j := 0; j < n; j++ {
			fmt.Scan(&c)
			grid[i][j] = c
		}
	}

	count := 0
	for {
		fnd := searchUnionAndUpdateGrid()
		if !fnd {
			break
		}
		count++
	}

	fmt.Println(count)
}
