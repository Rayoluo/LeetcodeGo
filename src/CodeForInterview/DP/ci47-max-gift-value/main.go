package main

import (
	"fmt"
	"math"
)

// 礼物的最大价值：最基本的背包问题
func maxValue(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([]int, n+1)
	dp[0] = math.MinInt32
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[j] = max(dp[j-1], dp[j]) + grid[i-1][j-1]
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxValue([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}))
}
