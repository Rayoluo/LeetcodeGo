package main

import (
	"fmt"
	"math"
)

func dicesProbability(n int) []float64 {
	// 声明状态表示数组,dp[n][s]表示n个骰子得到s点的次数
	dp := make([]int, 6*n+1)
	res := make([]float64, 0)
	// 对代码进行初始化
	for i := 1; i <= 6; i++ {
		dp[i] = 1
	}
	// k表示骰子的个数 2个骰子的点数范围是2-12
	for k := 2; k <= n; k++ {
		// j表示k个骰子的总点数
		for j := 6 * k; j >= k; j-- {
			// m表示最后一个骰子的点数 范围是1-6
			dp[j] = 0
			for m := 1; m <= 6; m++ {
				if j-m >= k-1 && j-m <= (k-1)*6 {
					dp[j] += dp[j-m]
				}
			}
		}
	}
	sum := math.Pow(6, float64(n))
	// fmt.Println(dp)
	for i := n; i <= n*6; i++ {
		res = append(res, float64(dp[i])/sum)
	}
	return res
}

func main() {
	fmt.Println(dicesProbability(2))
}
