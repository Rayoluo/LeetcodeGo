package main

import "fmt"

/**
 * @Author: luo
 * @Author: hustly123@gmail.com
 * @Date: 2021/5/18 9:56
 * @Desc: 机器人的运动范围
 */

var count int

type pair struct {
	x, y int
}

var visited map[pair]struct{}

func movingCount(m int, n int, k int) int {
	count = 0
	visited = make(map[pair]struct{})
	backtrace(0, 0, m, n, k)
	return count
}

// 参数说明 row表示机器人当前所在行，col表示机器人当前所在列
func backtrace(row int, col int, m int, n int, k int) {
	// 如果已经访问过了，直接返回
	if _, ok := visited[pair{row, col}]; ok {
		return
	}
	visited[pair{row, col}] = struct{}{}
	rc := sumrc(row, col)
	// 行坐标和列坐标数位之和大于k则直接返回
	if rc > k {
		return
	}
	count++
	// 向上
	if row-1 > 0 {
		backtrace(row-1, col, m, n, k)
	}
	// 向下
	if row+1 < m {
		backtrace(row+1, col, m, n, k)
	}
	// 向左
	if col-1 > 0 {
		backtrace(row, col-1, m, n, k)
	}
	// 向右
	if col+1 < n {
		backtrace(row, col+1, m, n, k)
	}
}

// 计算当前所在行列的坐标的数位之和
func sumrc(row, col int) int {
	sum := 0
	for row != 0 {
		sum += row % 10
		row /= 10
	}
	for col != 0 {
		sum += col % 10
		col /= 10
	}
	return sum
}

func main() {
	fmt.Println(movingCount(3, 2, 17))
}
