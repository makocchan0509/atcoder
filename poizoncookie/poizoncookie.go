package main

import "fmt"

func main() {
	var A, B, C int
	fmt.Scan(&A, &B, &C)

	if (A + B + 1) >= C {
		fmt.Println(B + C)
	} else {
		fmt.Println(B + (A + B + 1))
	}
}
