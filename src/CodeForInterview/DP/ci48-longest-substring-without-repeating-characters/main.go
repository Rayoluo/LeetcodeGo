package main

import "fmt"

// 最长不含重复字符的子字符串
// 动态规划的写法, dp[i]表示以第i个字符结尾的最长不含重复字符的子字符串的长度
// 则有下列推导关系： x表示第i个字符最近一次出现的索引位置
// dp[j] = dp[j-1]+1 if j-dp[j-1]>x || x < 0
// dp[j] = 1+dp[j-1]-(x-j+dp[j-1]+1) = j-x ,x >= j - dp[j-1]
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([]int, len(s))
	mp := make(map[byte]int)
	chs := []byte(s)
	dp[0] = 1
	mp[chs[0]] = 0
	for i := 1; i < len(chs); i++ {
		if index, ok := mp[chs[i]]; !ok {
			dp[i] = dp[i-1] + 1
		} else if index < i-dp[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = i - index
		}
		mp[chs[i]] = i
	}
	max := dp[0]
	for i := 1; i < len(chs); i++ {
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
