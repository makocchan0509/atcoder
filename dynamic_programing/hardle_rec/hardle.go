package main

import (
	"fmt"
	"math"
)

var N, L int
var hd []bool
var t1, t2, t3 int
var cost []int
var done []bool

func main() {

	fmt.Scan(&N, &L)

	var x int
	hd = genBoolArr(false, L+1)
	for i := 0; i < N; i++ {
		fmt.Scan(&x)
		hd[x] = true
	}
	fmt.Scan(&t1, &t2, &t3)

	cost = genIntArr(Pow(10, 10), L+1)
	cost[0] = 0

	done = genBoolArr(false, L+1)

	ans := rec(L)

	//fmt.Println(cost)
	for _, v := range []int{L - 1, L - 2, L - 3} {
		if v >= 0 {
			//ans = min(ans, cost[v]+(t1/2)+(t2*(2*(L-v)-1))/2)
			ans = min(ans, cost[v]+(t1/2)+(t2*(2*(L-v)-1))/2)
			//fmt.Println(v, cost[v]+(t1/2)+(t2*(2*(L-v)-1))/2)
		}
	}
	fmt.Println(ans)
}

func rec(l int) int {

	if done[l] {
		return cost[l]
	}
	if l == 0 {
		return 0
	}

	if l-1 >= 0 {
		cost[l] = min(cost[l], rec(l-1)+t1)
	}
	if l-2 >= 0 {
		cost[l] = min(cost[l], rec(l-2)+t1+t2)
	}
	if l-4 >= 0 {
		cost[l] = min(cost[l], rec(l-4)+t1+t2*3)
	}
	if hd[l] {
		cost[l] = cost[l] + t3
	}
	done[l] = true

	return cost[l]
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

func genIntArr(def, cnt int) []int {
	arr := make([]int, cnt)

	for i := 0; i < cnt; i++ {
		arr[i] = def
	}
	return arr
}
func genBoolArr(def bool, cnt int) []bool {
	arr := make([]bool, cnt)

	for i := 0; i < cnt; i++ {
		arr[i] = def
	}
	return arr
}
