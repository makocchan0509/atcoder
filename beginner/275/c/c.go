package main

import (
	"fmt"
	"strings"
)

func main() {
	var p string
	grid := make([][]string, 9)

	for i := 0; i < 9; i++ {
		fmt.Scan(&p)
		grid[i] = strings.Split(p, "")

	}
	ans := 0
	for i, ss := range grid {
		for j, _ := range ss {
			for x := 1; x+i < 9 && x+j < 9; x++ {
				for y := 0; y < x; y++ {
					if j+y >= 9 || i+y >= 9 || i+x-y <= 0 || j+x-y <= 0 {
						break
					}
					if grid[i][j+y] == "#" && grid[i+y][j+x] == "#" && grid[i+x-y][j] == "#" && grid[i+x][j+x-y] == "#" {
						ans++
					}
				}
			}
		}
	}
	fmt.Println(ans)
}
