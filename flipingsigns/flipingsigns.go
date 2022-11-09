package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)

	min := 10000000000
	var cnt_neg int
	var A int
	var sum int

	for i := 0; i < N; i++ {
		fmt.Scan(&A)

		if A < 0 {
			cnt_neg++
		}
		a := int(math.Abs(float64(A)))
		if min > a {
			min = a
		}
		sum = sum + a
	}

	if cnt_neg%2 == 0 {
		fmt.Println(sum)
	} else {
		fmt.Println(sum - (min * 2))
	}
}
