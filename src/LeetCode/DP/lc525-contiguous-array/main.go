package main

import "fmt"

/**
 * @Author: luo
 * @Author: hustly123@gmail.com
 * @Date: 2021/6/3 17:13
 * @Desc: 给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。哈希表+前缀和
 */

func findMaxLength(nums []int) int {
	// 哈希表存指定前缀和值出现的最小下标
	mp := map[int]int{0: -1}
	sum := make([]int, len(nums))
	var counter int
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			if nums[i] == 0 {
				sum[i] = -1
			} else {
				sum[i] = 1
			}
		} else {
			if nums[i] == 0 {
				sum[i] = sum[i-1] - 1
			} else {
				sum[i] = sum[i-1] + 1
			}
		}
		if v, exist := mp[sum[i]]; exist {
			counter = max(counter, i-v)
		} else {
			mp[sum[i]] = i
		}
	}
	return counter
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(findMaxLength([]int{0, 1}))
	fmt.Println(findMaxLength([]int{0, 1, 0}))
	// fmt.Println(findMaxLength([]int{0, 1}))

}
