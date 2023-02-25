package main

import (
	"fmt"
)

var (
	s, t       string
	slen, tlen int
)

func reversStr(s string) string {
	l := len(s)
	bs := make([]byte, l)
	half := l / 2
	if l%2 != 0 {
		bs[half] = s[half]
	}
	r := l - 1
	for i := 0; i < l/2; i++ {
		right := r - i
		bs[i] = s[right]
		bs[right] = s[i]
	}
	return string(bs)
}

type cs struct {
	hc, lc byte
}

func rec(cur string) bool {
	if cur == s {
		return true
	}

	curlen := len(cur)
	if curlen < slen {
		return false
	}

	c := cs{cur[0], cur[curlen-1]}

	/* switch {
	case c.hc == 'A' && c.lc == 'A':
		return rec(cur[:curlen-1])
	case c.hc == 'A' && c.lc == 'B':
		return false
	case c.hc == 'B' && c.lc == 'A':
		return rec(cur[:curlen-1]) || rec(reversStr(cur)[:curlen-1])
	default:
		return rec(reversStr(cur)[:curlen-1])
	} */

	switch c {
	case cs{'A', 'A'}:
		return rec(cur[:curlen-1])
	case cs{'A', 'B'}:
		return false
	case cs{'B', 'A'}:
		return rec(cur[:curlen-1]) || rec(reversStr(cur)[:curlen-1])
	default:
		return rec(reversStr(cur)[:curlen-1])
	}

}

func main() {
	fmt.Scanf("%s\n", &s)
	fmt.Scanf("%s\n", &t)

	slen = len(s)
	tlen = len(t)

	if rec(t) {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}

}
