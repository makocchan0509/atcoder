package main

import "fmt"

func main() {

	var N, d int

	fmt.Scan(&N)

	m := map[int]int{}

	for i := 0; i < N; i++ {
		fmt.Scan(&d)
		if v, ok := m[d]; ok {
			m[d] = v + 1
		} else {
			m[d] = v
		}
	}
	fmt.Println(len(m))
}
