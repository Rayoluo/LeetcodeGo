package main

import (
	"fmt"
)

/**
 * @Author: luo
 * @Author: hustly123@gmail.com
 * @Date: 2021/7/1 10:16
 * @Desc: 两数之和
 */
func twoSum(nums []int, target int) []int {
	mp := make(map[int]int) // mp：value-pos键值对
	// mp := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if pos, exist := mp[target-nums[i]]; exist {
			return []int{pos, i}
		}
		mp[nums[i]] = i
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{3, 3}, 6)) // error
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
