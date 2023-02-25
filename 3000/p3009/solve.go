package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x int
	y int
}

func searchLastPoint(points []Point) {
	var xmin, ymin, xmax, ymax int

	xmin = points[0].x
	ymin = points[0].y
	xmax = points[0].x
	ymax = points[0].y
	for _, p := range points {
		if p.x > xmax {
			xmax = p.x
		}
		if p.x < xmin {
			xmin = p.x
		}

		if p.y > ymax {
			ymax = p.y
		}
		if p.y < ymin {
			ymin = p.y
		}
	}

	a1 := Point{xmin, ymin}
	a2 := Point{xmin, ymax}
	a3 := Point{xmax, ymin}
	a4 := Point{xmax, ymax}
	ps := []Point{a1, a2, a3, a4}
	points_set := make(map[Point]bool)

	for _, i := range points {
		points_set[i] = true
	}

	for _, p := range ps {
		_, exists := points_set[p]
		if !exists {
			fmt.Printf("%d %d\n", p.x, p.y)
		}
	}
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var a, b int
	points := make([]Point, 0)
	for i := 1; i <= 3; i++ {
		fmt.Fscanf(reader, "%d %d\n", &a, &b)
		points = append(points, Point{a, b})
	}

	searchLastPoint(points)
}
