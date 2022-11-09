package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func main() {
	var s string
	fmt.Scan(&s)
	var dict []string
	arr := strings.Split(s, "")
	for i := 0; i < len(arr); {
		j := i + 1
		for j < len(arr) && unicode.IsLower(rune(arr[j][0])) {
			j++
		}
		dict = append(dict, s[i:j+1])
		i = j + 1
	}

	sort.SliceStable(dict, func(i, j int) bool { return strings.ToLower(dict[i]) < strings.ToLower(dict[j]) })
	//fmt.Println(dict)
	fmt.Println(strings.Join(dict, ""))
}
