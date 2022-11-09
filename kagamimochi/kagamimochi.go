package main

import "fmt"

func main() {
	var n int
	var d int
	dmap := map[int]int{}
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&d)

		v, ok := dmap[d]
		if ok {
			dmap[d] = v + 1
		} else {
			dmap[d] = 1
		}
	}
	fmt.Println(len(dmap))
}
