package main

import (
	"fmt"
	"math"
)

var hr []int
var memo []int

func main() {
	var N int
	fmt.Scan(&N)

	var h int
	hr = make([]int, N)
	memo = genIntArr(-1, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&h)
		hr[i] = h
	}

	fmt.Println(rec(N - 1))
}

func rec(N int) int {

	if N == 0 {
		return 0
	}
	if memo[N] != -1 {
		return memo[N]
	}

	result := 9999999999

	if N-1 >= 0 {
		result = min(result, rec(N-1)+abs(hr[N]-hr[N-1]))
	}
	if N-2 >= 0 {
		result = min(result, rec(N-2)+abs(hr[N]-hr[N-2]))
	}

	memo[N] = result
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

func genIntArr(def, cnt int) []int {
	arr := make([]int, cnt)

	for i := 0; i < cnt; i++ {
		arr[i] = def
	}
	return arr
}
