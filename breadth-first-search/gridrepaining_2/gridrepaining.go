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
	a []address
}

func (q *queue) enQueue(a address) {
	q.a = append(q.a, a)
}

func (q *queue) deQueue() address {
	a := q.a[0]
	q.a = q.a[1:]
	return a
}

func (q *queue) len() int {
	return len(q.a)
}

func main() {

	var H, W int
	fmt.Scan(&H, &W)

	grid := make([][]string, H)

	var S string

	for i := 0; i < H; i++ {
		fmt.Scan(&S)
		grid[i] = strings.Split(S, "")
	}
	//fmt.Println(grid)

	ans := make([][]int, H)

	for i := 0; i < H; i++ {
		ans[i] = genIntArr(-1, W)
	}

	//↑←→↓
	xr := []int{0, -1, 1, 0}
	yr := []int{-1, 0, 0, 1}

	var q queue

	ans[0][0] = 0
	q.enQueue(address{0, 0})

	for q.len() > 0 {
		a := q.deQueue()
		//fmt.Println(a)
		for i := 0; i < 4; i++ {

			//fmt.Println(a.y+yr[i], a.x+xr[i])
			if a.y+yr[i] < 0 || a.y+yr[i] >= H || a.x+xr[i] < 0 || a.x+xr[i] >= W {
				continue
			}
			if grid[a.y+yr[i]][a.x+xr[i]] == "#" || ans[a.y+yr[i]][a.x+xr[i]] != -1 {
				continue
			}
			q.enQueue(address{a.x + xr[i], a.y + yr[i]})
			ans[a.y+yr[i]][a.x+xr[i]] = ans[a.y][a.x] + 1
		}
	}

	//fmt.Println(ans)

	if ans[H-1][W-1] == -1 {
		fmt.Println(-1)
	} else {

		point := 0
		for _, vv := range grid {
			for _, v := range vv {
				if v == "." {
					point += 1
				}
			}
		}

		fmt.Println(point - ans[H-1][W-1] - 1)
	}
}

func genIntArr(def, cnt int) []int {
	arr := make([]int, cnt)

	for i := 0; i < cnt; i++ {
		arr[i] = def
	}
	return arr
}
