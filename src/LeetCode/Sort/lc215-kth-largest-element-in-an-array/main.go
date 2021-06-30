package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 * @Author: luo
 * @Author: hustly123@gmail.com
 * @Date: 2021/6/15 10:17
 * @Desc: 查找数组中第k大的数字，k总是有效地，且k的值在1到数组长度之间
 */

func findKthLargest(nums []int, k int) int {
	return find(nums, 0, len(nums)-1, k)
}

func find(nums []int, l, r, k int) int {
	pivot := partition(nums, l, r)
	// 大于nums[pivot]的元素有k-1个，则nums[pivot]刚好是第k大元素
	if r-pivot == k-1 {
		return nums[pivot]
	}
	// 大于nums[pivot]的元素大于k-1个，则继续在右边区间中查找
	if r-pivot > k-1 {
		return find(nums, pivot+1, r, k)
	} else {
		return find(nums, l, pivot-1, k-r+pivot-1)
	}
}

func partition(nums []int, l, r int) (pivot int) {
	// 随机选择一个数字进行比较
	var (
		index, area, i int
	)
	rand.Seed(time.Now().UnixNano())
	index = l + rand.Intn(r-l+1)
	nums[index], nums[r] = nums[r], nums[index]
	area = l - 1
	for i = l; i < r; i++ {
		if nums[i] > nums[r] {
			continue
		}
		area++
		nums[i], nums[area] = nums[area], nums[i]
	}
	nums[r], nums[area+1] = nums[area+1], nums[r]
	pivot = area + 1
	return
}

func main() {
	nums := []int{3, 2, 1, 3, 3, 4}
	fmt.Println(findKthLargest(nums, 6))
}
