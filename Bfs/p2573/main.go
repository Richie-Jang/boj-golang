package main

import (
	"bufio"
	"fmt"
	"os"
)

type pos struct {
	y, x int
}

var r, c int
var im [300][300]int
var dp [300][300]int
var hm = [4]int{-1, 1, 0, 0}
var vm = [4]int{0, 0, -1, 1}

func meltDownPos(y, x int) int {
	res := 0
	if im[y][x] != 0 {
		for i := 0; i < 4; i++ {
			h := hm[i] + x
			v := vm[i] + y
			if h >= 0 && h < c && v >= 0 && v < r {
				cc := im[v][h]
				if cc == 0 {
					res++
				}
			}
		}
	}
	return res
}

func printMountain() {
	fmt.Println("========Cur Mountain=======")
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Printf("%d ", im[i][j])
		}
		fmt.Println()
	}
	fmt.Println("===========================")
}

func printMcountArray() {
	fmt.Println("========dp arr=============")
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Printf("%d ", dp[i][j])
		}
		fmt.Println()
	}
	fmt.Println("===========================")
}

func meltDown() {
	// dp is also used for storing meltdown numbers
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			dp[i][j] = meltDownPos(i, j)
		}
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			im[i][j] -= dp[i][j]
			if im[i][j] < 0 {
				im[i][j] = 0
			}
		}
	}
}

func searching(y, x int, mn int) {
	q := make([]pos, 1)
	visited := make(map[pos]bool)
	q[0] = pos{y, x}
	for len(q) > 0 {
		vv := q[0]
		dp[vv.y][vv.x] = mn
		visited[vv] = true
		q = q[1:]

		for i := 0; i < 4; i++ {
			np := pos{vv.y + vm[i], vv.x + hm[i]}
			_, ok := visited[np]
			if np.y >= 0 && np.y < r && np.x >= 0 && np.x < c && im[np.y][np.x] != 0 &&
				dp[np.y][np.x] == 0 && !ok {
				q = append(q, np)
			}
		}
	}
}

func checkMountainCount() int {
	res := 0
	for k := 0; k < r; k++ {
		for l := 0; l < c; l++ {
			dp[k][l] = 0
		}
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			cc := im[i][j]
			if cc != 0 && dp[i][j] == 0 {
				// searching
				res++
				searching(i, j, res)
			}
		}
	}
	return res
}

func solve() {
	var res = 1
	for {
		meltDown()
		//printMountain()
		ccc := checkMountainCount()
		//printMcountArray()
		if ccc >= 2 {
			break
		} else if ccc == 0 {
			res = 0
			break
		}
		res++
	}

	fmt.Println(res)
}

func main() {
	rr := bufio.NewReader(os.Stdin)
	fmt.Fscanf(rr, "%d %d\n", &r, &c)
	var a int
	var tmp string
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Fscanf(rr, "%d", &a)
			im[i][j] = a
		}
		fmt.Fscanln(rr, &tmp)
	}
	solve()
}
