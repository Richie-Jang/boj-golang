package main

import "fmt"

func countStr(s string) int {
	prev := ' '
	maxValue := 0
	count := 0
	for _, v := range s {
		if v != prev {
			prev = v
			if maxValue < count {
				maxValue = count
			}
			count = 1
		} else {
			count++
		}
	}
	if maxValue < count {
		maxValue = count
	}
	return maxValue
}

func main() {
	var s string
	for {
		_, err := fmt.Scanln(&s)
		if err != nil {
			break
		}
		fmt.Println(countStr(s))
	}
}
