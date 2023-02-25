package main

import (
    "bufio"
    "fmt"
    "os"
)

var (
    n,m int
    arr []int
    res int
)

func solve(cur int) {
    if cur >= n {
        return
    }
    sum := 0
    for i := cur; i < n; i++ {
        sum += arr[i]
        if sum == m {
            res++
            break
        }
    }
    solve(cur+1)
}

func main() {
    br := bufio.NewReader(os.Stdin)
    fmt.Fscanf(br, "%d %d\n", &n, &m)
    arr = make([]int, n)
    for i :=0; i<n; i++ {
        fmt.Fscanf(br, "%d", &arr[i])
    }
    fmt.Fscanln(br)

    solve(0)
    fmt.Println(res)
}