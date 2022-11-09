package main

import (
	"fmt"
	"math"
)

func main() {
	var N, L int
	fmt.Scan(&N, &L)

	var x int

	hd := genBoolArr(false, L+1)

	for i := 0; i < N; i++ {
		fmt.Scan(&x)
		hd[x] = true
	}

	var t1, t2, t3 int
	fmt.Scan(&t1, &t2, &t3)

	cost := genIntArr(Pow(10, 10), L+1)
	cost[0] = 0

	//fmt.Println(cost)
	//fmt.Println(hd)
	for i := 1; i <= L; i++ {

		if i-1 >= 0 {
			cost[i] = min(cost[i], cost[i-1]+t1)
		}
		if i-2 >= 0 {
			cost[i] = min(cost[i], cost[i-2]+t1+t2)
		}
		if i-4 >= 0 {
			cost[i] = min(cost[i], cost[i-4]+t1+t2*3)
		}

		if hd[i] {
			cost[i] = cost[i] + t3
		}
	}
	fmt.Println(cost)
	ans := cost[L]
	for _, v := range []int{L - 1, L - 2, L - 3} {
		if v >= 0 {
			ans = min(ans, cost[v]+(t1/2)+(t2*(2*(L-v)-1))/2)
			fmt.Println(v, cost[v]+(t1/2)+(t2*(2*(L-v)-1))/2)
		}
	}
	fmt.Println(ans)
}

func Pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func genBoolArr(def bool, cnt int) []bool {
	arr := make([]bool, cnt)

	for i := 0; i < cnt; i++ {
		arr[i] = def
	}
	return arr
}

func genIntArr(def, cnt int) []int {
	arr := make([]int, cnt)

	for i := 0; i < cnt; i++ {
		arr[i] = def
	}
	return arr
}
