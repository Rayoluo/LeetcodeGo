package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	// 对区间按照第一个元素的大小进行排序，如果区间的第一个元素相等，则按照第二个元素大小排序
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		} else {
			return intervals[i][1] < intervals[j][1]
		}
	})
	// 双指针
	var (
		i, j, t int
		ret     [][]int
	)
	ret = make([][]int, 0)
	for i = 0; i < len(intervals); {
		t = intervals[i][1]
		j = i + 1
		for j < len(intervals) && intervals[j][0] <= t {
			t = max(t, intervals[j][1])
			j++
		}
		ret = append(ret, []int{intervals[i][0], t})
		i = j
	}

	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(merge([][]int{[]int{1, 4}, []int{1, 3}}))
	fmt.Println(merge([][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}}))
	fmt.Println(merge([][]int{[]int{1, 4}, []int{5, 6}}))
}
