package main

import "fmt"

func main() {
	var data = []int{1, 2, 3, 4}

	var selection = 2
	var loop func(int, int)
	var comb = make([]int, selection)

	loop = func(ci, di int) {
		if ci == selection {
			for j := 0; j < len(comb); j++ {
				fmt.Print(comb[j], " ")
			}
			fmt.Println()
			return
		}

		if di >= len(data) {
			return
		}

		comb[ci] = data[di]
		fmt.Printf("ci:%d [%d] di:%d [%d]\n", ci, di, comb[ci], data[di])
		loop(ci+1, di+1)
		loop(ci, di+1)
	}

	loop(0, 0)

}
