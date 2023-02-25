package main

import "fmt"

var n, m, k int
var seats [1001][1001]int

// dir 0:N, 1:E, 2:S, 3:W
var dx = [4]int{0, 1, 0, -1}
var dy = [4]int{-1, 0, 1, 0}
var resx, resy = 0, 0

func turnNext(dir int) int {
	ndir := 0
	switch dir {
	case 0:
		ndir = 1
	case 1:
		ndir = 2
	case 2:
		ndir = 3
	case 3:
		ndir = 0
	}
	return ndir
}

func makeASeats(y, x, tnum, dir int) {
	if tnum > n*m {
		return
	}
	seats[y][x] = tnum
	if tnum == k {
		resx = x
		resy = n - y + 1
		return
	}

	ny, nx := y+dy[dir], x+dx[dir]
	ndir := dir
	if ny <= 0 || ny > n || nx <= 0 || nx > m || seats[ny][nx] != 0 {
		ndir = turnNext(dir)
		ny, nx = y+dy[ndir], x+dx[ndir]
	}
	makeASeats(ny, nx, tnum+1, ndir)
}

func printSeats() {
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Printf("%02d ", seats[i][j])
		}
		fmt.Println()
	}
}

func main() {
	fmt.Scanln(&m, &n)
	fmt.Scanln(&k)

	makeASeats(n, 1, 1, 0)
	if resx == resy && resx == 0 {
		fmt.Println("0")
	} else {
		fmt.Printf("%d %d\n", resx, resy)
	}
	//printSeats()
}
