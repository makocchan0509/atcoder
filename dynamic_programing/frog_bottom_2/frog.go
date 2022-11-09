package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)

	hr := make([]int, N)

	var h int
	for i := 0; i < N; i++ {
		fmt.Scan(&h)
		hr[i] = h
	}

	dp := genIntArr(999999999999999, N)
	dp[0] = 0
	for i, _ := range hr {

		if i-1 >= 0 {
			dp[i] = min(dp[i], dp[i-1]+abs(hr[i]-hr[i-1]))
		}
		if i-2 >= 0 {
			dp[i] = min(dp[i], dp[i-2]+abs(hr[i]-hr[i-2]))
		}
	}
	//fmt.Println(dp)
	fmt.Println(dp[N-1])
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
