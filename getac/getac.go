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

func nextStr() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)

	var n, q, l, r int
	var s string

	fmt.Scan(&n, &q)
	fmt.Scan(&s)

	ns := make([]int, n+1)

	for i := 1; i < n; i++ {
		ss := s[i-1 : i+1]
		//		fmt.Println(ss)
		if ss == "AC" {
			ns[i+1] = ns[i] + 1
		} else {
			ns[i+1] = ns[i]
		}
	}
	//	fmt.Println(ns)
	for i := 0; i < q; i++ {

		fmt.Scan(&l, &r)
		//		l -= 1
		ans := ns[r] - ns[l]
		fmt.Println(ans)
	}
}
