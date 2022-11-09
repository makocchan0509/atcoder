package main

import "fmt"

func main() {
	var n, y int
	fmt.Scan(&n, &y)

	var ana, anb, anc = -1, -1, -1
	for a := 0; a < n; a++ {
		for b := 0; b < n; b++ {
			c := n - (a + b)
			if c < 0 && c > n {
				continue
			}

			if (10000*a)+(5000*b)+(1000*c) == y {
				fmt.Printf("%d %d %d\n", a, b, c)
				return
			}
		}
	}
	fmt.Printf("%d %d %d\n", ana, anb, anc)
}
