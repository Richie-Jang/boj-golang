package main

import "fmt"

func moveSingle(_from, _to int) {
	fmt.Printf("%d %d\n", _from, _to)
}
func hanoi(count, _from, _to, _middle int) {
	if count == 1 {
		moveSingle(_from, _to)
		return
	}
	hanoi(count-1, _from, _middle, _to)
	moveSingle(_from, _to)
	hanoi(count-1, _middle, _to, _from)
}

func main() {
	var n int
	fmt.Scanln(&n)
	if n <= 20 {
		hanoi(n, 1, 3, 2)
	}
}
