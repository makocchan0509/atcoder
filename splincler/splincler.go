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
	var n, m, q int
	n = nextInt()
	m = nextInt()
	q = nextInt()

	u := make([][]int, n)
	//for i := 0; i < n; i++ {
	//	u[i] = make([]int, n)
	//}

	var u1, u2 int
	for i := 0; i < m; i++ {
		u1 = nextInt()
		u2 = nextInt()
		u[u1-1] = append(u[u1-1], u2)
		u[u2-1] = append(u[u2-1], u1)
	}
	//fmt.Println(u)

	var c int
	colors := make([]int, n)
	for i := 0; i < n; i++ {
		c = nextInt()
		colors[i] = c
	}
	//fmt.Println(colors)

	s := make([][]int, q)
	for i := 0; i < len(s); i++ {
		s1 := nextInt()
		if s1 == 1 {
			s[i] = []int{s1, nextInt()}
		} else {
			s[i] = []int{s1, nextInt(), nextInt()}
		}
	}
	//fmt.Println(s)

	for _, v := range s {

		if v[0] == 1 {
			fmt.Println(colors[v[1]-1])
			for _, vv := range u[v[1]-1] {
				colors[vv-1] = colors[v[1]-1]
			}
		} else {
			fmt.Println(colors[v[1]-1])
			colors[v[1]-1] = v[2]
		}
	}
}
