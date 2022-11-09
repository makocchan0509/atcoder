package main

import "fmt"

var ln, sn []int

func main() {
	var N, X int

	fmt.Scan(&N, &X)

	ln = genIntArr(1, N+1)
	sn = genIntArr(1, N+1)

	for i := 1; i < N+1; i++ {
		ln[i] = ln[i-1]*2 + 3
		sn[i] = sn[i-1]*2 + 1
	}

	//fmt.Println(ln)
	//fmt.Println(sn)

	fmt.Println(recber(N, X))

}

func recber(N, X int) int {

	if N == 0 {
		return 1
	}
	//fmt.Printf("%d:%d\n", N, X)
	if X == 1 {
		return 0
	} else if X <= ln[N-1]+1 {
		return recber(N-1, X-1)
	} else if X == ln[N-1]+2 {
		return sn[N-1] + 1
	} else if X <= ln[N-1]*2+2 {
		return recber(N-1, X-ln[N-1]-2) + sn[N-1] + 1
	} else {
		return sn[N-1]*2 + 1
	}
}

func genIntArr(def, cnt int) []int {
	arr := make([]int, cnt)

	for i := 0; i < cnt; i++ {
		arr[i] = def
	}
	return arr
}
