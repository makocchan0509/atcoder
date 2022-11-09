package main

import "fmt"

var ar, br, cr, dr []int
var N, M, Q int

func main() {

	fmt.Scan(&N, &M, &Q)

	ar = genIntArr(0, Q)
	br = genIntArr(0, Q)
	cr = genIntArr(0, Q)
	dr = genIntArr(0, Q)

	var a, b, c, d int
	for i := 0; i < Q; i++ {
		fmt.Scan(&a, &b, &c, &d)
		ar[i] = a
		br[i] = b
		cr[i] = c
		dr[i] = d
	}

	var A []int
	fmt.Println(rec(A))
}

func rec(A []int) int {

	if len(A) == N {
		return score(A)
	}

	result := 0
	a := 1

	if len(A) >= 1 {
		a = A[len(A)-1]
	}

	for i := a; i <= M; i++ {
		A = append(A, i)
		result = max(result, rec(A))
		A = pop(A)
	}
	return result
}

func pop(A []int) []int {
	if len(A) == 0 {
		return A
	}
	return A[:len(A)-1]
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func score(A []int) int {

	score := 0
	for i := 0; i < Q; i++ {
		//fmt.Println(br[i]-1, ar[i]-1)
		if A[br[i]-1]-A[ar[i]-1] == cr[i] {
			score += dr[i]
		}
	}
	return score
}

func genIntArr(def, cnt int) []int {
	arr := make([]int, cnt)

	for i := 0; i < cnt; i++ {
		arr[i] = def
	}
	return arr
}
