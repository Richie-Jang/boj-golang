package main

import (
	"fmt"
)

var (
	t         int
	asc_table map[rune]int
)

func solve(str string, check_map map[rune]bool) {
	for _, v := range str {
		check_map[v] = true
	}
	var sum = 0
	for k, v := range check_map {
		if !v {
			vv, _ := asc_table[k]
			sum += vv
		}
	}
	fmt.Println(sum)
}

func main() {
	fmt.Scanln(&t)
	var v string
	asc_table = make(map[rune]int)
	check_map := make(map[rune]bool)
	for b := 'A'; b <= 'Z'; b++ {
		asc_table[b] = int(b)
		check_map[b] = false
	}
	for i := 0; i < t; i++ {
		fmt.Scanln(&v)
		for k, _ := range check_map {
			check_map[k] = false
		}
		solve(v, check_map)
	}
}
