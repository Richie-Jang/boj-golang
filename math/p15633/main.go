package main

import "fmt"

var n int

func main() {

	fmt.Scanf("%d\n", &n)

	if n == 1 {
		fmt.Println(1*5 - 24)
		return
	}

	arr := []int{1, n}
	visited := make([]bool, 10_001)
	visited[1] = true
	visited[n] = true

	for i := 2; i < n; i++ {
		v := n % i
		if v == 0 && !visited[i] {
			g := n / i
			arr = append(arr, i)
			visited[i] = true
			if !visited[g] {
				arr = append(arr, g)
				visited[g] = true
			}
		}
	}

	var sum = 0
	for _, v := range arr {
		sum += v
	}
	res := sum*5 - 24
	fmt.Println(res)
}
