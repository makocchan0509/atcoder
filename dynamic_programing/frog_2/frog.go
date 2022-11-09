package main

import (
	"fmt"
	"math"
)

var memo []int
var N int
var hr []int

func main() {

	fmt.Scan(&N)

	hr = make([]int, N)

	var h int
	for i := 0; i < N; i++ {
		fmt.Scan(&h)
		hr[i] = h
	}

	memo = genIntArr(-1, N)

	fmt.Println(rec(N - 1))
}

func rec(n int) int {

	if n == 0 {
		return 0
	}

	if memo[n] != -1 {
		return memo[n]
	}

	result := 9999999999999999

	if n-1 >= 0 {
		result = min(result, rec(n-1)+abs(hr[n-1]-hr[n]))
	}
	if n-2 >= 0 {
		result = min(result, rec(n-2)+abs(hr[n-2]-hr[n]))
	}
	memo[n] = result
	return result
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func genIntArr(def, cnt int) []int {
	arr := make([]int, cnt)

	for i := 0; i < cnt; i++ {
		arr[i] = def
	}
	return arr
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}
