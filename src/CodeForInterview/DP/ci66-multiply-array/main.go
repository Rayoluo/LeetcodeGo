package main

import "fmt"

// 构建乘积数组
func constructArr(a []int) []int {
	n := len(a)
	if a == nil || n == 0 {
		return nil
	}
	left := make([]int, n)
	right := make([]int, n)
	left[0] = 1
	right[n-1] = 1
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * a[i-1]
		right[n-i-1] = right[n-i] * a[n-i]
	}
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = left[i] * right[i]
	}
	return ret
}

func main() {
	fmt.Println(constructArr([]int{1, 2, 3, 4, 5}))
}
