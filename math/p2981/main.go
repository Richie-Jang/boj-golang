package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

var (
	n   int
	arr []int
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	br := bufio.NewReader(os.Stdin)
	fmt.Fscanf(br, "%d\n", &n)
	arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(br, "%d\n", &arr[i])
	}

	sort.Ints(arr)
	diff := arr[1] - arr[0]
	for i := 1; i < n-1; i++ {
		diff = gcd(diff, arr[i+1]-arr[i])
	}

	fmt.Println("Diff", diff)

	res := make([]int, 0, 10)
	count := 1
	for count*count <= diff {
		if diff%count == 0 {
			res = append(res, count)
			vv := diff / count
			if count != vv {
				res = append(res, vv)
			}
		}
		count++
	}
	bf := new(bytes.Buffer)
	sort.Ints(res)
	for _, a := range res {
		if a != 1 {
			bf.WriteString(fmt.Sprintf("%d ", a))
		}
	}
	fmt.Println(bf.String())
}
