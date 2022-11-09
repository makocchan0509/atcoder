package main

import "fmt"

func main() {
	var A, B, C, K int

	fmt.Scan(&A, &B, &C, &K)

	if K%2 == 0 {
		fmt.Println(A - B)
	} else {
		fmt.Println(B - A)
	}
	//for i := 0; i < K; i++ {
	//
	//	AA := B + C
	//	BB := A + C
	//	CC := A + B
	//	fmt.Println(AA - BB)
	//	A = AA
	//	B = BB
	//	C = CC
	//}
}
