package main

import "fmt"

var (
	n   int
	str string
)

func makePiArr() []int {
	slen := len(str)
	res := make([]int, slen)
	j := 0
	for i := 1; i < slen; i++ {
		for j > 0 && str[i] != str[j] {
			j = res[j-1]
		}
		if str[i] == str[j] {
			res[i] = j + 1
			j++
		} else {
			res[i] = 0
		}
	}
	return res
}

func main() {
	/* fmt.Scanf("%d\n", &n)
	fmt.Scanf("%s\n", &str)
	*/
	n = 7
	str = "abadabc"

	pi := makePiArr()
	fmt.Println(pi)
	fmt.Println(n - pi[n-1])

}
