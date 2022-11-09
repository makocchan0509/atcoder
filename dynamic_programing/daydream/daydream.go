package main

import "fmt"

// dream dreamer erase eraser

func main() {

	var S string
	fmt.Scan(&S)

	N := len(S)

	dp := genBoolArr(false, N+1)

	dp[0] = true

	for i := 1; i <= N; i++ {
		if i >= 5 && dp[i-5] {
			if S[i-5:i] == "dream" {
				dp[i] = true
			}
		}
		if i >= 7 && dp[i-7] {
			if S[i-7:i] == "dreamer" {
				dp[i] = true
			}
		}
		if i >= 5 && dp[i-5] {
			if S[i-5:i] == "erase" {
				dp[i] = true
			}
		}
		if i >= 6 && dp[i-6] {
			if S[i-6:i] == "eraser" {
				dp[i] = true
			}
		}
	}
	//fmt.Println(dp)
	if dp[N] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func genBoolArr(def bool, cnt int) []bool {
	arr := make([]bool, cnt)

	for i := 0; i < cnt; i++ {
		arr[i] = def
	}
	return arr
}
