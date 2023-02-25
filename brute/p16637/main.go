package main

import (
	"bufio"
	"fmt"
	"os"
)

type sign int

const (
	PLUS sign = iota
	MINUS
	MUL
)

var (
	n   int
	str string
	res int = -1 * (1 << 30)
)

func compute(a, b int, s sign) int {
	res := 0
	switch s {
	case PLUS:
		res = a + b
	case MUL:
		res = a * b
	case MINUS:
		res = a - b
	}
	return res
}

func checkSign(v byte) sign {
	r := PLUS
	switch v {
	case '+':
		r = PLUS
	case '-':
		r = MINUS
	case '*':
		r = MUL
	}
	return r
}

func solve(idx, curValue int) {
	if idx >= n {
		if res < curValue {
			res = curValue
		}
		return
	}

	sg := PLUS
	if idx > 0 {
		sg = checkSign(str[idx-1])
	}

	if idx+2 < n {
		a := compute(int(str[idx]-'0'), int(str[idx+2]-'0'), checkSign(str[idx+1]))
		solve(idx+4, compute(curValue, a, sg))
	}

	solve(idx+2, compute(curValue, int(str[idx]-'0'), sg))
}

func main() {
	br := bufio.NewReader(os.Stdin)
	fmt.Fscanf(br, "%d\n", &n)
	fmt.Fscanln(br, &str)
	solve(0, 0)
	fmt.Println(res)
}
