package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var N int
	var s string

	fmt.Scan(&N)

	m := map[string]int{}

	for i := 0; i < N; i++ {
		fmt.Scan(&s)
		ss := strings.Split(s, "")
		sort.Strings(ss)
		t := strings.Join(ss, "")

		if _, ok := m[t]; ok {
			m[t] = m[t] + 1
		} else {
			m[t] = 1
		}
	}

	var ans int
	for _, v := range m {
		ans = ans + (v * (v - 1) / 2)
	}
	fmt.Println(ans)
}
