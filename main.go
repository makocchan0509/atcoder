package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	arr := strings.Split(s, "")
	for i := 0; i < len(arr); {
		j := i
		for j < len(arr) && arr[i] == arr[j] {
			j++
		}
		fmt.Printf("(%s,%d)\n", arr[i], j-i)
		i = j
	}
}
