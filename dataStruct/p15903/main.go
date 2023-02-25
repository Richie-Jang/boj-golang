package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type pq []int

// sort interface
func (p pq) Len() int           { return len(p) }
func (p pq) Less(i, j int) bool { return p[i] < p[j] }
func (p pq) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// heap interface
func (p *pq) Pop() any {
	old := *p
	l := len(old)
	x := old[l-1]
	*p = old[:l-1]
	return x
}
func (p *pq) Push(x any) {
	*p = append(*p, x.(int))
}

func main() {
	var n, m, g int

	br := bufio.NewReader(os.Stdin)
	fmt.Fscan(br, &n, &m)

	p := make(pq, 0)
	heap.Init(&p)

	for i := 0; i < n; i++ {
		fmt.Fscan(br, &g)
		heap.Push(&p, g)
	}

	// mix
	for i := 0; i < m; i++ {
		a := heap.Pop(&p).(int)
		b := heap.Pop(&p).(int)
		c := a + b
		for j := 0; j < 2; j++ {
			heap.Push(&p, c)
		}
	}

	// sum
	var sum int
	for _, r := range p {
		sum += r
	}

	fmt.Println(sum)

}
