package main

import "fmt"

// 八皇后问题
func solveNQueens(n int) [][]string {
	columns := make([]int, n)
	res := make([][]string, 0)
	column := make(map[int]bool)
	leftup := make(map[int]bool)
	rightup := make(map[int]bool)
	backtrack(0, n, columns, &res, column, leftup, rightup)
	return res
}

// 对于八皇后问题，在左上角的满足“行值减去列值固定”，在右上角的满足“行值加上列值固定”
func backtrack(row, n int, columns []int, res *[][]string, column map[int]bool, leftup map[int]bool, rightup map[int]bool) {
	if row == n {
		*res = append(*res, intToStringSlice(columns))
		return
	}
	for i := 0; i < n; i++ {
		if column[i] {
			continue
		}
		leftupValue := row - i
		if leftup[leftupValue] {
			continue
		}
		rightupValue := row + i
		if rightup[rightupValue] {
			continue
		}
		columns[row] = i
		column[i] = true
		leftup[leftupValue], rightup[rightupValue] = true, true
		backtrack(row+1, n, columns, res, column, leftup, rightup)
		columns[row] = -1
		delete(column, i)
		delete(leftup, leftupValue)
		delete(rightup, rightupValue)
	}
}

func intToStringSlice(columns []int) []string {
	n := len(columns)
	ss := make([]string, 0)
	for i := 0; i < n; i++ {
		str := ""
		for j := 0; j < n; j++ {
			if j == columns[i] {
				str += "Q"
			} else {
				str += "."
			}
		}
		ss = append(ss, str)
	}
	return ss
}

func main() {
	fmt.Println(solveNQueens(4))
}
