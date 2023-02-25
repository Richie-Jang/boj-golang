package main

import (
	"bufio"
	"fmt"
	"os"
)

var t, n int
var larr []int
var rarr []int

func lowerBound(arr []int, target int) int {
	l := 0
	r := len(arr) - 1
	for l < r {
		m := (l + r) / 2
		if target <= arr[m] {
			r = m
		} else {
			l = m + 1
		}
	}
	return r
}

func upperBound(arr []int, target int) int {
	l := 0
	r := len(arr) - 1
	for l < r {
		m := (l + r) / 2
		if target < arr[m] {
			r = m
		} else {
			l = m + 1
		}
	}
	return r
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	fmt.Fscanln(rd, &t)
	fmt.Fscanln(rd, &n)
	larr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &larr[i])
	}
	fmt.Fscanln(rd)
	fmt.Fscanln(rd, &n)
	rarr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &rarr[i])
	}

	left_sum := make([]int, 0)
	for i := 0; i < len(larr); i++ {
		sum := larr[i]
		left_sum = append(left_sum, sum)
		for j := i + 1; j < len(larr); j++ {
			sum += larr[j]
			left_sum = append(left_sum, sum)
		}
	}
	right_map := make(map[int]int)
	for i := 0; i < len(rarr); i++ {
		sum := rarr[i]
		if _, ok := right_map[sum]; ok == false {
			right_map[sum] = 1
		} else {
			right_map[sum]++
		}
		for j := i + 1; j < len(rarr); j++ {
			sum += rarr[j]
			if _, ok := right_map[sum]; ok == false {
				right_map[sum] = 1
			} else {
				right_map[sum]++
			}
		}
	}

	ans := int64(0)

	for _, v := range left_sum {
		df := t - v
		vv, ok := right_map[df]
		if ok == false {
			continue
		}
		ans += int64(vv)
	}

	fmt.Println(ans)

}
