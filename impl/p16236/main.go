package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

// Ti2 is tuple X and Y
type Ti2 struct {
	x, y int
}

// Grid : Tuple xy and depth count
type Grid struct {
	p     Ti2
	depth int
}

// Shark : shark
type Shark struct {
	size     int
	gridPos  Ti2
	eatCount int
	runTime  int
}

// Pq : priority queue
type Pq []Grid

// Len : for heap & sort.slice
func (p Pq) Len() int { return len(p) }

// Less : for heap
func (p Pq) Less(i, j int) bool {
	a1 := p[i]
	a2 := p[j]
	if a1.depth == a2.depth {
		if a1.p.y == a2.p.y {
			return a1.p.x < a2.p.x
		}
		return a1.p.y < a2.p.y
	}
	return a1.depth < a2.depth
}

// Swap : for heap
func (p *Pq) Swap(i, j int) {
	pp := *p
	pp[i], pp[j] = pp[j], pp[i]
}

// Push : for heap
func (p *Pq) Push(i any) {
	*p = append(*p, i.(Grid))
}

// Pop : for heap
func (p *Pq) Pop() any {
	d := *p
	l := len(d)
	res := d[l-1]
	*p = d[:l-1]
	return res
}

var dir []Ti2 = []Ti2{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

var (
	n     int
	grid  [][]int
	nGrid [][]Grid
	sk    Shark
	pq    Pq // priority queue
)

func run() bool {
	// search small fish existence
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			nGrid[y][x] = Grid{Ti2{x, y}, math.MaxInt}
		}
	}

	nGrid[sk.gridPos.y][sk.gridPos.x].depth = 0
	queue := []Grid{{Ti2{sk.gridPos.x, sk.gridPos.y}, 0}}

	// refresh board
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, v := range dir {
			nx := v.x + cur.p.x
			ny := v.y + cur.p.y
			if 0 <= nx && nx < n && 0 <= ny && ny < n && grid[ny][nx] <= sk.size && cur.depth+1 < nGrid[ny][nx].depth {
				item := Grid{Ti2{nx, ny}, cur.depth + 1}
				nGrid[ny][nx].depth = item.depth
				queue = append(queue, item)
			}
		}
	}

	//fmt.Println(nGrid)

	pq = make(Pq, 0)

	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			g := nGrid[y][x]
			dd := g.depth
			if dd == math.MaxInt || dd == 0 {
				continue
			}
			s := grid[g.p.y][g.p.x]
			if s > 0 && s < sk.size {
				pq = append(pq, nGrid[y][x])
			}
		}
	}

	if len(pq) == 0 {
		return false
	}

	heap.Init(&pq)
	small := heap.Pop(&pq).(Grid)
	// update
	grid[small.p.y][small.p.x] = 0
	sk.gridPos.x = small.p.x
	sk.gridPos.y = small.p.y
	sk.eatCount++
	if sk.eatCount >= sk.size {
		sk.eatCount = 0
		sk.size++
	}
	sk.runTime += small.depth
	pq = nil
	//fmt.Println(grid)
	//fmt.Println("Shark", sk)
	//fmt.Println("===============================================================================")
	return true
}

func main() {

	br := bufio.NewReader(os.Stdin)
	fmt.Fscan(br, &n)
	sk.size = 2
	grid = make([][]int, n)
	nGrid = make([][]Grid, n)

	for y := 0; y < n; y++ {
		grid[y] = make([]int, n)
		nGrid[y] = make([]Grid, n)
		for x := 0; x < n; x++ {
			fmt.Fscan(br, &(grid[y][x]))
			if grid[y][x] == 9 {
				sk.gridPos = Ti2{x, y}
				grid[y][x] = 0
			}
		}
	}

	//fmt.Println("Start Shark : ", sk)

	for {
		b := run()
		if !b {
			break
		}
	}

	fmt.Println(sk.runTime)

}
