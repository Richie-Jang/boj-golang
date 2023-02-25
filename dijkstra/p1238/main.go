package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

type Node struct {
	to, dist int
}

type NodeHeap []Node

func (h NodeHeap) Len() int {
	return len(h)
}

func (h NodeHeap) Less(i, j int) bool {
	return h[i].dist < h[j].dist
}

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodeHeap) Push(ele any) {
	*h = append(*h, ele.(Node))
}

func (h *NodeHeap) Pop() any {
	old := *h
	nn := len(old)
	ele := old[nn-1]
	*h = old[0 : nn-1]
	return ele
}

var (
	n, m, x int
	adjs    map[int][]Node = make(map[int][]Node)
)

const IMAX = 999_999_999

func dijkstra(from int) []int {
	dest := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i != from {
			dest[i] = IMAX
		}
	}

	nh := &NodeHeap{Node{from, 0}}
	heap.Init(nh)

	for nh.Len() > 0 {
		cur := heap.Pop(nh).(Node)
		if dest[cur.to] < cur.dist {
			continue
		}

		for _, nx := range adjs[cur.to] {
			nextTo := nx.to
			nextDist := cur.dist + nx.dist
			if nextDist < dest[nextTo] {
				dest[nextTo] = nextDist
				newV := Node{nextTo, nextDist}
				heap.Push(nh, newV)
			}
		}

	}
	return dest
}

func main() {
	inp := os.Stdin
	br := bufio.NewReader(inp)

	defer inp.Close()

	fmt.Fscanf(br, "%d %d %d\n", &n, &m, &x)
	for i := 0; i < m; i++ {
		a, b, c := 0, 0, 0
		fmt.Fscanf(br, "%d %d %d\n", &a, &b, &c)
		nodes, ok := adjs[a]
		if !ok {
			adjs[a] = []Node{{b, c}}
		} else {
			adjs[a] = append(nodes, Node{b, c})
		}
	}

	xToOthers := dijkstra(x)
	for i := 1; i <= n; i++ {
		if i != x {
			shorts := dijkstra(i)
			xToOthers[i] += shorts[x]
		}
	}

	sort.Slice(xToOthers, func(i, j int) bool {
		return xToOthers[i] > xToOthers[j]
	})

	fmt.Println(xToOthers[0])
}
