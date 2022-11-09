package main

import (
	"fmt"
	"strings"
)

type address struct {
	x int
	y int
}

type queue struct {
	q []address
}

func main() {
	var N, W int
	fmt.Scan(&N, &W)

	ss := make([][]string, N)
	ii := genIntTwo(-1, N, W)

	// ↑,←,→,↓
	xr := [4]int{0, -1, 1, 0}
	yr := [4]int{-1, 0, 0, 1}

	var s string
	for i := 0; i < N; i++ {
		fmt.Scan(&s)
		ss[i] = strings.Split(s, "")
	}

	var Q queue
	Q.enqueue(address{0, 0})
	ii[0][0] = 0

	for Q.len() > 0 {
		a := Q.dequeue()

		for i := 0; i < 4; i++ {
			dx := a.x + xr[i]
			dy := a.y + yr[i]
			if dx < 0 || dy < 0 || dx >= W || dy >= N {
				continue
			}
			if ss[dy][dx] == "#" {
				continue
			}
			if ii[dy][dx] != -1 {
				continue
			}
			Q.enqueue(address{dx, dy})
			ii[dy][dx] = ii[a.y][a.x] + 1
		}
	}

	if ii[N-1][W-1] == -1 {
		fmt.Println(-1)
		return
	}

	white := 0
	for _, vv := range ss {
		for _, v := range vv {
			if v == "." {
				white += 1
			}
		}
	}

	fmt.Println(white - ii[N-1][W-1] - 1)
}

func genIntTwo(d, n, w int) [][]int {
	ii := make([][]int, n)
	for i := 0; i < n; i++ {
		ii[i] = genIntArr(-1, w)
	}
	return ii
}

func genIntArr(def, cnt int) []int {
	arr := make([]int, cnt)

	for i := 0; i < cnt; i++ {
		arr[i] = def
	}
	return arr
}

func (qu *queue) dequeue() address {
	a := qu.q[0]
	qu.q = qu.q[1:]
	return a
}

func (qu *queue) enqueue(a address) {
	qu.q = append(qu.q, a)
}

func (qu *queue) len() int {
	return len(qu.q)
}
