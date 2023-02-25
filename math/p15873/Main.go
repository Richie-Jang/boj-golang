package main

import (
	"fmt"
	"strconv"
)

var s string

func loop(ss string) int {
	var left, right = "", ""
	left = left + string(ss[0])
	if ss[1] == '0' {
		left = left + string(ss[1])
		right = right + string(ss[2])
	} else {
		right = string(ss[1:])
	}
	i1, _ := strconv.Atoi(left)
	i2, _ := strconv.Atoi(right)
	return i1 + i2
}

func main() {
	fmt.Scanln(&s)
	l := len(s)
	switch l {
	case 4:
		fmt.Println("20")
	case 3:
		fmt.Println(loop(s))
	default:
		i1, _ := strconv.Atoi(string(s[0]))
		i2, _ := strconv.Atoi(string(s[1]))
		fmt.Println(i1 + i2)
	}
}
