package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	n           int
	inp         string
	left, right []byte
)

func readString(br *bufio.Scanner) string {
	br.Scan()
	return strings.TrimSpace(br.Text())
}

func stackPush(st *[]byte, v byte) {
	*st = append(*st, v)
}

func stackPop(st *[]byte) (byte, bool) {
	vst := *st
	l := len(vst)
	if l == 0 {
		return byte(0), false
	}
	v := vst[l-1]
	*st = vst[0 : l-1]
	return v, true
}

var reg *regexp.Regexp = regexp.MustCompile(`P\s+(\w)`)

func parsing(s string) {
	l := len(s)
	if l > 1 {
		addC := reg.FindStringSubmatch(s)[1]
		stackPush(&left, addC[0])
	} else {
		switch s {
		case "L":
			v, ok := stackPop(&left)
			if ok {
				stackPush(&right, v)
			}
		case "B":
			stackPop(&left)
		case "D":
			v, ok := stackPop(&right)
			if ok {
				stackPush(&left, v)
			}
		}
	}
}

func reverse(sl *[]byte) []byte {
	vsl := *sl
	size := len(vsl)

	if size <= 1 {
		return vsl
	}

	isEven := size%2 == 0
	stopNum := size / 2
	if isEven {
		stopNum--
	}
	res := make([]byte, size)
	for i := 0; i <= stopNum; i++ {
		ri := (size - 1) - i
		a, b := vsl[i], vsl[ri]
		res[i] = b
		res[ri] = a
	}
	return res
}

func convertStr(v []byte, isRev bool) string {
	var res string
	size := len(v)
	if size == 0 {
		return ""
	}
	if size == 1 {
		res = string(v)
	} else {
		if isRev {
			res = string(reverse(&v))
		} else {
			res = string(v)
		}
	}
	return res
}

func main() {
	br := bufio.NewScanner(os.Stdin)
	left = make([]byte, 0)
	right = make([]byte, 0)
	inp = readString(br)
	for _, c := range inp {
		stackPush(&left, byte(c))
	}

	nc := readString(br)
	n, _ = strconv.Atoi(nc)
	for i := 0; i < n; i++ {
		parsing(readString(br))
		//fmt.Println(left, right)
	}

	fmt.Println(convertStr(left, false) + convertStr(right, true))
}
