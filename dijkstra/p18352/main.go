package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

// heap
type Node struct {
	to, dist int
}

type HeapNode []Node

func (h HeapNode) Len() int {
	return len(h)
}

func (h HeapNode) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h HeapNode) Less(i, j int) bool {
	return h[i].dist < h[j].dist
}

func (h *HeapNode) Push(v any) {
	*h = append(*h, v.(Node))
}

func (h *HeapNode) Pop() any {
	old := *h
	size := len(old)
	res := old[size-1]
	*h = old[0 : size-1]
	return res
}

var (
	n, m, k, x int
	adjs       [][]Node
)

const IMAX = 999_999_999

func main() {

	input := os.Stdin
	defer input.Close()
	br := bufio.NewReader(input)

	fmt.Fscanf(br, "%d %d %d %d\n", &n, &m, &k, &x)
	adjs = make([][]Node, n+1)
	for i := 1; i <= n; i++ {
		adjs[i] = make([]Node, 0)
	}

	var a, b int
	for i := 0; i < m; i++ {
		fmt.Fscanf(br, "%d %d\n", &a, &b)
		adjs[a] = append(adjs[a], Node{b, 1})
	}

	// start dijkstra compute base on x
	dist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = IMAX
	}
	dist[x] = 0
	pq := &HeapNode{Node{x, 0}}
	heap.Init(pq)

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(Node)
		curDist := cur.dist
		if dist[cur.to] < curDist {
			continue
		}
		for _, nt := range adjs[cur.to] {
			nextDist := curDist + nt.dist
			if nextDist < dist[nt.to] {
				dist[nt.to] = nextDist
				heap.Push(pq, Node{nt.to, nextDist})
			}
		}
	}

	var isCheck bool = false

	for i, v := range dist {
		if i == 0 {
			continue
		}
		if v == k {
			isCheck = true
			fmt.Println(i)
		}
	}

	if !isCheck {
		fmt.Println("-1")
	}
}
