package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)

	var N, M, Q int
	fmt.Scan(&N, &M, &Q)

	graph := make([][]int, N)

	var u, v int
	for i := 0; i < M; i++ {
		fmt.Scan(&u, &v)
		graph[u-1] = append(graph[u-1], v)
		graph[v-1] = append(graph[v-1], u)
	}
	//fmt.Println(graph)

	splc := map[int]int{}
	var c int
	for i := 1; i <= N; i++ {
		fmt.Scan(&c)
		splc[i] = c
	}
	//fmt.Println(splc)

	for i := 0; i < Q; i++ {
		s1 := nextInt()
		s2 := nextInt()
		if s1 == 2 {
			fmt.Println(splc[s2])
			s3 := nextInt()
			splc[s2] = s3
		} else {
			fmt.Println(splc[s2])
			for _, v := range graph[s2-1] {
				cc := splc[s2]
				splc[v] = cc
			}
		}
	}
}
