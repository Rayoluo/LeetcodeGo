package ci45

import (
	"bytes"
	"sort"
	"strconv"
)

// 题目描述：如何将数组进行排序使得组合出来的数字最小，并返回组合出来的最小字符串
func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		str1, str2:= strconv.Itoa(nums[i]), strconv.Itoa(nums[j])
		s1 := str1 + str2
		s2 := str2 + str1
		if s1 < s2 {
			return true
		}
		return false
	})
	var buf bytes.Buffer
	for i := 0; i < len(nums); i++ {
		buf.WriteString(strconv.Itoa(nums[i]))
	}
	return buf.String()
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
