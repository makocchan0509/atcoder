package main

import "fmt"

var N, W int
var vr, wr []int

func main() {
	var v, w int

	fmt.Scan(&N, &W)

	vr = genIntArr(0, N)
	wr = genIntArr(0, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&v, &w)
		vr[i] = v
		wr[i] = w
	}

	fmt.Println(rec(0, 0, 0))
}

func rec(n, w, v int) int {

	if n == N {
		return v
	}
	result := 0
	result = max(result, rec(n+1, w, v))

	if w+wr[n] <= W {
		result = max(result, rec(n+1, w+wr[n], v+vr[n]))
	}
	return result
}

func genIntArr(def, cnt int) []int {
	arr := make([]int, cnt)

	for i := 0; i < cnt; i++ {
		arr[i] = def
	}
	return arr
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
