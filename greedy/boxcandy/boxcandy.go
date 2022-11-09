package main

import (
	"fmt"
	"math"
)

func main() {
	var N, x, a int
	fmt.Scan(&N, &x)

	ar := genIntArr(0, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&a)
		ar[i] = a
	}

	var result int

	if ar[0] > x {
		result += abs(ar[0] - x)
		ar[0] = ar[0] - (ar[0] - x)
	}
	for i := 0; i+1 < N; i++ {
		if ar[i]+ar[i+1] > x {
			result += abs((ar[i] + ar[i+1]) - x)
			ar[i+1] = ar[i+1] - (ar[i] + ar[i+1] - x)
		}
	}
	//fmt.Println(ar)
	fmt.Println(result)
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
