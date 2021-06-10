package main

/**
 * @Author: luo
 * @Author: hustly123@gmail.com
 * @Date: 2021/6/10 12:38
 * @Desc: 买卖股票的最佳时机-II：可以尽可能地完成更多的交易
 */

func maxProfit(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}
