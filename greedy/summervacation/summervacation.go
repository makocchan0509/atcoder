package main

import (
	"container/heap"
	"fmt"
)

func main() {
	var N, M int
	fmt.Scan(&N, &M)
	var a, b int
	ab := make([][]int, M+1)
	for i := 0; i <= M; i++ {
		ab[i] = []int{}
	}

	h := &intHeap{}
	heap.Init(h)

	for i := 0; i < N; i++ {
		fmt.Scan(&a, &b)

		if M < a {
			continue
		}
		ab[a] = append(ab[a], b)
	}

	var result int
	for _, v := range ab {

		for _, vv := range v {
			heap.Push(h, vv)
		}
		if h.Len() > 0 {
			result = result + heap.Pop(h).(int)
		}
	}

	fmt.Println(result)
}

func genIntArr(def, cnt int) []int {
	arr := make([]int, cnt)

	for i := 0; i < cnt; i++ {
		arr[i] = def
	}
	return arr
}

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
