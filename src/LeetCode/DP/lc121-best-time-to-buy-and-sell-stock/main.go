package main

/**
 * @Author: luo
 * @Author: hustly123@gmail.com
 * @Date: 2021/6/10 11:22
 * @Desc: 买卖股票的最大时机
 */
func maxProfit(prices []int) int {
	dp := make([]int, len(prices))
	minPrice := prices[0]
	for i := 1; i < len(dp); i++ {
		dp[i] = prices[i] - minPrice
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}
	maxVal := 0
	for i := 1; i < len(dp); i++ {
		if dp[i] > maxVal {
			maxVal = dp[i]
		}
	}
	return maxVal
}
