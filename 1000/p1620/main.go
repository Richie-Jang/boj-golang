package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var n, m int
var dic_eng map[string]int
var dic_int map[int]string

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var str string
	dic_eng = make(map[string]int, n+1)
	dic_int = make(map[int]string, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscanln(reader, &str)
		dic_eng[str] = i
		dic_int[i] = str
	}

	var k string
	int_sign := map[rune]bool{
		'1': true, '2': true, '3': true, '4': true, '5': true, '6': true,
		'7': true, '8': true, '9': true, '0': true,
	}
	for j := 1; j <= m; j++ {
		fmt.Fscanln(reader, &k)
		if _, ok := int_sign[rune(k[0])]; ok {
			v, _ := strconv.Atoi(k)
			fmt.Println(dic_int[v])
		} else {
			fmt.Println(dic_eng[k])
		}
	}

}
