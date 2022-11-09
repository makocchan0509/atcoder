package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func main() {

	var S string
	fmt.Scan(&S)

	arr := strings.Split(S, "")
	var dict []string

	for i := 0; i < len(arr); {
		j := i + 1
		for j < len(arr) && unicode.IsLower(rune(arr[j][0])) {
			j++
		}
		dict = append(dict, S[i:j+1])
		i = j + 1
	}
	sort.Slice(dict, func(i, j int) bool { return strings.ToLower(dict[i]) < strings.ToLower(dict[j]) })

	s := strings.Join(dict, "")
	fmt.Println(s)
}
