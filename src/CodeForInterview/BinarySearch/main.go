package main

import "fmt"

/**
 * @Author: luo
 * @Author: hustly123@gmail.com
 * @Date: 2021/6/2 16:38
 * @Desc: 旋转数组的最小数字， 二分查找
 */

func minArray(numbers []int) int {
	n := len(numbers) - 1
	for n > 0 && numbers[n] == numbers[0] {
		n--
	}
	if n == 0 {
		return numbers[0]
	}
	if numbers[n] > numbers[0] {
		return numbers[0]
	}
	l, r := 0, n
	for l < r {
		mid := l + (r-l)>>1
		if numbers[mid] >= numbers[0] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return numbers[l]
}

func main() {
	fmt.Println(minArray([]int{3, 4, 5, 1, 2}))
	fmt.Println(minArray([]int{2, 2, 2, 0, 1}))
}
