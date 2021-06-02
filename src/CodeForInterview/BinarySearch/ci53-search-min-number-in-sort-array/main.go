package main

import "fmt"

/**
 * @Author: luo
 * @Author: hustly123@gmail.com
 * @Date: 2021/6/2 19:45
 * @Desc: 查找排序数组中数字出现的次数
 */

func main() {
	fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 6))
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	l, r := 0, len(nums)-1
	var count int
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			// to do
			// fmt.Println(mid)
			for i := mid; i >= l; i-- {
				if nums[i] == target {
					count++
				} else {
					break
				}
			}
			for i := mid + 1; i <= r; i++ {
				if nums[i] == target {
					count++
				} else {
					break
				}
			}
			break
		}
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}

	}
	return count
}
