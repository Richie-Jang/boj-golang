package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Oper int

const (
	ADD Oper = iota
	MINUS
	MUL
	DIV
)

func compute(a, b int, op Oper) int {
	var res int
	switch op {
	case ADD:
		res = a + b
	case MINUS:
		res = a - b
	case MUL:
		res = a * b
	case DIV:
		res = a / b
	}
	return res
}

func main() {

	var s string
	var valCount int
	var a, b int
	var op Oper
	var needStop bool = false
	br := bufio.NewReader(os.Stdin)
	for !needStop {
		_, er := fmt.Fscanln(br, &s)
		if er != nil {
			break
		}

		ss := strings.TrimSpace(s)
		switch ss {
		case "+":
			op = ADD
		case "-":
			op = MINUS
		case "*":
			op = MUL
		case "/":
			op = DIV
		case "=":
			{
				fmt.Println(a)
				needStop = true
			}
		default:
			if valCount == 0 {
				valCount++
				a, _ = strconv.Atoi(ss)
			} else {
				b, _ = strconv.Atoi(ss)
				a = compute(a, b, op)
			}
		}
	}

}
