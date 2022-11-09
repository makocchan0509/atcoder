package main

import (
	"fmt"
	"math"
)

func main() {

	var N int
	fmt.Scan(&N)

	var t, x, y int
	var nt, nx, ny int
	for i := 0; i < N; i++ {
		fmt.Scan(&t, &x, &y)

		t = int(math.Abs(float64(t - nt)))
		x = int(math.Abs(float64(x - nx)))
		y = int(math.Abs(float64(y - ny)))

		if t < x+y {
			fmt.Println("No")
			return
		}
		if !(t%2 == (x+y)%2) {
			fmt.Println("No")
			return
		}

		nt = t
		nx = x
		ny = y
	}
	fmt.Println("Yes")
}
