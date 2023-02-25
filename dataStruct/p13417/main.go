package main

import (
	"bufio"
	"bytes"
	"container/list"
	"fmt"
	"os"
	"strings"
)

var t int
var buf bytes.Buffer

func solve(br *bufio.Reader, n int) {
	arr := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(br, "%s", &arr[i])
	}
	fmt.Fscanln(br)

	mylist := list.New()
	var curFront string
	mylist.PushFront(arr[0])
	curFront = arr[0]
	for _, v := range arr[1:] {
		comp := strings.Compare(v, curFront)
		if comp <= 0 {
			mylist.PushFront(v)
			curFront = v
		} else {
			mylist.PushBack(v)
		}
	}
	for v := mylist.Front(); v != nil; v = v.Next() {
		aa, _ := v.Value.(string)
		buf.WriteString(aa)
	}
	buf.WriteString("\n")
}

func main() {
	br := bufio.NewReader(os.Stdin)
	fmt.Fscanln(br, &t)

	for i := 1; i <= t; i++ {
		var n int
		fmt.Fscanln(br, &n)
		solve(br, n)
	}

	fmt.Print(buf.String())
}
