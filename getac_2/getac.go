package main

import "fmt"

func main() {

	var N, Q, l, r int
	var S string
	fmt.Scan(&N, &Q)
	fmt.Scan(&S)

	cusum := make([]int, N+1)
	cusum[0] = 0
	for i := 1; i < N; i++ {
		if S[i-1:i+1] == "AC" {
			cusum[i+1] = cusum[i] + 1
		} else {
			cusum[i+1] = cusum[i]
		}
	}
	//fmt.Println(cusum)

	for i := 0; i < Q; i++ {
		fmt.Scan(&l, &r)
		fmt.Println(cusum[r] - cusum[l])
	}
}
