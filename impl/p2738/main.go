package main

import "fmt"

var n, m int
var a1 [100][100]int

func main() {
	fmt.Scanln(&n, &m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&a1[i][j])
		}
	}
	tmp := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&tmp)
			a1[i][j] += tmp
		}
	}

	// print
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j == m-1 {
				fmt.Print(a1[i][j])
			} else {
				fmt.Print(a1[i][j], " ")
			}
		}
		fmt.Println()
	}

}
