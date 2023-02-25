package main

import (
	"fmt"
	"sort"
	"strings"
)

var vowels = []string{"a", "e", "i", "o", "u"}

var arr []string

func isVowel(b string) bool {
	res := false
	for _, vv := range vowels {
		if b == vv {
			res = true
			break
		}
	}
	return res
}

func recusive(pos, aNum, bNum, limitLen int, buf []string) {
	if len(buf) >= limitLen {
		if aNum >= 1 && bNum >= 2 {
			fmt.Println(strings.Join(buf, ""))
		}
		return
	}

	for cc := pos; cc < len(arr); cc++ {
		v := arr[cc]
		buf = append(buf, v)
		if isVowel(v) {
			recusive(cc+1, aNum+1, bNum, limitLen, buf)
		} else {
			recusive(cc+1, aNum, bNum+1, limitLen, buf)
		}
		buf = buf[0 : len(buf)-1]
	}

}

func main() {

	var l, c int
	fmt.Scanln(&l, &c)
	arr = make([]string, c)
	for i := 0; i < c; i++ {
		fmt.Scanf("%s", &arr[i])
	}

	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	buf := make([]string, 0)
	recusive(0, 0, 0, l, buf)

}
