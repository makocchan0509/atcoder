package main

import "fmt"

var aq, bq, cq, dq []int
var N, M, Q int

func main() {

	var a, b, c, d int

	fmt.Scan(&N, &M, &Q)

	aq = genIntArr(0, Q)
	bq = genIntArr(0, Q)
	cq = genIntArr(0, Q)
	dq = genIntArr(0, Q)

	for i := 0; i < Q; i++ {
		fmt.Scan(&a, &b, &c, &d)

		aq[i] = a
		bq[i] = b
		cq[i] = c
		dq[i] = d

		aq[i]--
		bq[i]--
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

	if len(A) != 0 {
		a = A[len(A)-1]
	}

	for i := a; i <= M; i++ {
		A = append(A, i)
		result = max(result, rec(A))
		A = pop(A)
	}
	return result
}

func score(A []int) int {
	var score int
	for i := 0; i < Q; i++ {
		if A[bq[i]]-A[aq[i]] == cq[i] {
			score += dq[i]
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

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func pop(A []int) []int {
	if len(A) == 0 {
		return A
	}
	A = A[:len(A)-1]
	return A
}
