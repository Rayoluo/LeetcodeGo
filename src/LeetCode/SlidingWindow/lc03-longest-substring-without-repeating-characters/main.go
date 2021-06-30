package main

import "fmt"

/**
 * @Author: luo
 * @Author: hustly123@gmail.com
 * @Date: 2021/6/15 20:07
 * @Desc: 无重复字符的最长字串的滑动窗口写法
 */

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	arr := []byte(s)
	m := make(map[byte]int)
	maxlen := 1
	left := 0
	m[arr[0]] = 0
	for i := 1; i < len(arr); i++ {
		if index, ok := m[arr[i]]; ok { // 包含重复的
			for j := left; j < index; j++ {
				delete(m, arr[j])
			}
			left = index + 1
		}
		m[arr[i]] = i
		maxlen = max(maxlen, i-left+1)
	}
	return maxlen
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abba"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
