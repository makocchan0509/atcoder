package main

import "fmt"

func main() {

	var N, Y int
	fmt.Scan(&N, &Y)

	for x := 0; x <= N; x++ {
		for y := 0; y <= N; y++ {
			if x+y > N {
				break
			}
			z := N - (x + y)
			ans := 10000*x + 5000*y + 1000*z
			if ans == Y {
				fmt.Printf("%d %d %d\n", x, y, z)
				return
			}
		}
	}
	fmt.Printf("%d %d %d\n", -1, -1, -1)
}
