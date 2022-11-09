package main

import "fmt"

// レベルNバーガー B N-1バーガ P N-1バーガ B
// レベルNのパティ数 2(N-1バーガのパティ数) + 1
// レベルNの層数　2(N-1バーガの層数) + 3
// レベル0はパティ1

func main() {

	var N, X int
	fmt.Scan(&N, &X)

	p := genIntArr(0, N+1)
	t := genIntArr(0, N+1)

	p[0] = 1
	t[0] = 1

	for i := 1; i <= N; i++ {
		p[i] = p[i-1]*2 + 1
		t[i] = t[i-1]*2 + 3
	}

	fmt.Println(rec(N, X, p, t))
}

func rec(N int, X int, p []int, t []int) int {

	if N == 0 {
		return 1
	}

	if X == 1 {
		return 0
	} else if X <= t[N-1]+1 {
		return rec(N-1, X-1, p, t)
	} else if X <= t[N-1]+2 {
		return p[N-1] + 1
	} else if X <= t[N-1]*2+2 {
		return rec(N-1, X-t[N-1]-2, p, t) + p[N-1] + 1
	} else {
		return p[N-1]*2 + 1
	}
}

func genIntArr(def, cnt int) []int {
	arr := make([]int, cnt)

	for i := 0; i < cnt; i++ {
		arr[i] = def
	}
	return arr
}
