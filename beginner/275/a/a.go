package main

import "fmt"

func main() {
	var N, H int

	fmt.Scan(&N)

	high := 0
	ans := 0
	for i := 1; i <= N; i++ {
		fmt.Scan(&H)
		if high < H {
			ans = i
			high = H
		}
	}

	fmt.Println(ans)
}
