package main

import (
	"fmt"
)

/**
 * @Author: luo
 * @Author: hustly123@gmail.com
 * @Date: 2021/5/21 21:35
 * @Desc: 不相交的线
 */

type pair struct {
	i, j int
}

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	m := make(map[pair]int)
	if nums1[0] == nums2[0] {
		m[pair{0, 0}] = 1
	} else {
		m[pair{0, 0}] = 0
	}
	return f(nums1, nums2, len(nums1)-1, len(nums2)-1, m)
}

func f(nums1 []int, nums2 []int, i, j int, m map[pair]int) int {
	if i < 0 || j < 0 {
		return 0
	}
	if val, ok := m[pair{i, j}]; ok {
		return val
	}
	var ret int
	if nums1[i] == nums2[j] {
		ret = 1 + f(nums1, nums2, i-1, j-1, m)
	} else {
		ret = max(f(nums1, nums2, i, j-1, m), f(nums1, nums2, i-1, j, m))
	}
	m[pair{i, j}] = ret
	return ret
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func main() {
	fmt.Println(maxUncrossedLines([]int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2}))
	fmt.Println(maxUncrossedLines([]int{1, 3, 7, 1, 7, 5}, []int{1, 9, 2, 5, 1}))
	fmt.Println(maxUncrossedLines([]int{3}, []int{3, 3, 2}))
}
