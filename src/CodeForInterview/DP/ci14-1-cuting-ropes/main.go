package main

import "fmt"

func cuttingRope(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			if j == i-1 {
				dp[i] = max(dp[i], j*1)
			} else {
				dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
			}
		}
	}
	return dp[n]
}

func max(m, n int) int {
	if m > n {
		return m
	} else {
		return n
	}
}

func main() {
	fmt.Println(cuttingRope(10))
}
