package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 * @Author: luo
 * @Author: hustly123@gmail.com
 * @Date: 2021/6/15 20:56
 * @Desc: 给一个整数数组，请将该数组升序排列, 使用快排
 */

func sortArray(nums []int) []int {
	quicksort(nums, 0, len(nums)-1)
	return nums
}

func quicksort(nums []int, l, r int) {
	if l < r {
		pivot := partition(nums, l, r)
		quicksort(nums, l, pivot-1)
		quicksort(nums, pivot+1, r)
	}
}

func partition(nums []int, l, r int) (pivot int) {
	rand.Seed(time.Now().UnixNano())
	pos := l + rand.Intn(r-l+1)
	nums[pos], nums[r] = nums[r], nums[pos]
	area := l - 1
	for i := l; i < r; i++ {
		if nums[i] > nums[r] {
			continue
		}
		area++
		nums[area], nums[i] = nums[i], nums[area]
	}
	nums[area+1], nums[r] = nums[r], nums[area+1]
	return area + 1
}

func main() {
	fmt.Println(sortArray([]int{5, 1, 1, 2, 0, 0}))
	fmt.Println(sortArray([]int{5, 2, 3, 1}))
}
