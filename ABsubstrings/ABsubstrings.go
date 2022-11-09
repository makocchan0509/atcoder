package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	var S string
	var t, p1, p2, p3 int

	for i := 0; i < N; i++ {
		fmt.Scan(&S)
		for i := 0; i+1 < len(S); {
			if S[i:i+2] == "AB" {
				t++
			}
			i = i + 1
		}

		p := S[0:1]
		s := S[len(S)-1:]

		if p == "B" && s == "A" {
			p1++
		} else if p == "B" {
			p2++
		} else if s == "A" {
			p3++
		}
	}

	//fmt.Printf("%d,%d,%d,%d\n", t, p1, p2, p3)

	var ans int
	if p2 == 0 && p3 == 0 {
		ans = t + max(p1-1, 0)
	} else {
		ans = t + p1 + min(p2, p3)
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
