package main

import (
	"container/heap"
	"fmt"
)

/**
 * @Author: luo
 * @Author: hustly123@gmail.com
 * @Date: 2021/5/19 9:44
 * @Desc: 找出第k大的异或坐标值
 */

const MAXLEN = 1000

var dp [MAXLEN][MAXLEN]int

type maxHeap []int

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kthLargestValue(matrix [][]int, k int) int {
	m := len(matrix)    // 行
	n := len(matrix[0]) // 列
	// s := make(map[int]struct{}) // 集合用于去重
	h := make(maxHeap, 0) // 初始化堆的底层数组
	// 初始化
	dp[0][0] = matrix[0][0]
	// s[dp[0][0]] = struct{}{}
	heap.Push(&h, dp[0][0])
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] ^ matrix[0][j]
		// if _, ok := s[dp[0][j]]; !ok {
		// 	// 如果集合中不存在该元素则加入集合并插入堆中
		// 	s[dp[0][j]] = struct{}{}
		// 	heap.Push(&h, dp[0][j])
		// }
		heap.Push(&h, dp[0][j])
	}
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] ^ matrix[i][0]
		// if _, ok := s[dp[i][0]]; !ok {
		// 	// 如果集合中不存在该元素则加入集合并插入堆中
		// 	s[dp[i][0]] = struct{}{}
		// 	heap.Push(&h, dp[i][0])
		// }
		heap.Push(&h, dp[i][0])
	}
	// 计算每一个位置的异或值
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] ^ dp[i][j-1] ^ dp[i-1][j-1] ^ matrix[i][j]
			// if _, ok := s[dp[i][j]]; !ok {
			// 	// 如果集合中不存在该元素则加入集合并插入堆中
			// 	s[dp[i][j]] = struct{}{}
			// 	heap.Push(&h, dp[i][j])
			// }
			heap.Push(&h, dp[i][j])
		}
	}
	// debug
	// for i := 0; i < m; i++ {
	// 	for j := 0; j < n; j++ {
	// 		fmt.Println(dp[i][j])
	// 	}
	// }
	// fmt.Println(s)

	var ret interface{}
	for i := 1; i <= k; i++ {
		ret = heap.Pop(&h)
	}
	return ret.(int)
}

func main() {
	// fmt.Println(kthLargestValue([][]int{{5, 2}, {1, 6}}, 1))
	fmt.Println(kthLargestValue([][]int{{10, 9, 5}, {2, 0, 4}, {1, 0, 9}, {3, 4, 8}}, 10))
}
