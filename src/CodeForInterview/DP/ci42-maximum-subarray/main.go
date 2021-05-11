package main

import (
	"fmt"
	"math"
)

// 连续子数组的最大和：思想很简单，dp[i]表示以i作为最后一个元素的连续子数组的和，最后取最大值即可
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(dp); i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = nums[i] + dp[i-1]
		}
	}
	max := math.MinInt32
	for i := 0; i < len(dp); i++ {
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

// 空间复杂度太大，优化版本
func maxSubArray1(nums []int) int {
	dp := nums[0]
	max := dp
	for i := 1; i < len(nums); i++ {
		if dp < 0 {
			dp = nums[i]
		} else {
			dp = nums[i] + dp
		}
		if dp > max {
			max = dp
		}
	}
	return max
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
