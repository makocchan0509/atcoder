package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	var n int
	fmt.Scan(&n)
	var smap map[string]int
	var s string
	smap = map[string]int{}
	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		arr := strings.Split(s, "")
		sort.Strings(arr)
		ss := strings.Join(arr, "")

		smap[ss] += 1
	}

	result := 0
	for _, v := range smap {
		result += v * (v - 1) / 2
	}
	fmt.Println(result)
}
