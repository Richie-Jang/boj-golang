package main

import (
	"bufio"
	"fmt"
	"os"
)

var n int
var key string

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscanf(r, "%s\n", &key)
	fmt.Fscanf(r, "%d\n", &n)

	var stringFinder func(s string, cur, keycur int, isCont bool) bool
	stringFinder = func(s string, cur, keycur int, isCont bool) bool {
		if keycur == len(key) && isCont {
			return true
		}
		if isCont == false && cur == len(s) {
			return false
		}
		if isCont && cur == len(s) {
			if s[0] == key[keycur] {
				return stringFinder(s, 1, keycur+1, true)
			}
			return false
		}
		c := s[cur]
		k := key[keycur]
		if c != k {
			return stringFinder(s, cur+1, keycur, false)
		}
		return stringFinder(s, cur+1, keycur+1, true)
	}

	var s string
	var count = 0
	for i := 0; i < n; i++ {
		fmt.Fscanf(r, "%s\n", &s)
		if stringFinder(s, 0, 0, false) {
			count++
		}
	}
	fmt.Println(count)
}
