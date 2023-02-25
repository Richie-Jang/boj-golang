package main

import "fmt"

var n int

func main() {
	fmt.Scanf("%d\n", &n)
	inps := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &inps[i])
	}
	var isOkay bool = true
	fc := inps[0][0]
	for i := 0; i < n; i++ {
		if fc != inps[i][0] {
			isOkay = false
			break
		}
	}
	if isOkay {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
}
