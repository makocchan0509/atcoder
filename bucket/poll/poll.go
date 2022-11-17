package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type aggreVote struct {
	key   string
	total int
}

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	var S string

	fmt.Fscan(in, &N)

	m := make(map[string]int, 0)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &S)
		m[S] += 1
	}
	var vote []aggreVote

	ma := -1
	for k, v := range m {
		ma = max(ma, v)
		vote = append(vote, aggreVote{key: k, total: v})
	}
	//fmt.Println(vote)
	vote = sortSlice(vote)
	//fmt.Println(vote)
	for i := 0; i < len(vote); i++ {
		if vote[i].total == ma {
			fmt.Fprintln(out, vote[i].key)
		} else {
			break
		}
	}
}

func sortSlice(arr []aggreVote) []aggreVote {
	sort.SliceStable(arr, func(i, j int) bool { return arr[i].key < arr[j].key })
	sort.SliceStable(arr, func(i, j int) bool { return arr[i].total > arr[j].total })
	return arr
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
