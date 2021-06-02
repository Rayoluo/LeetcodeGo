package main

import "fmt"

/**
 * @Author: luo
 * @Author: hustly123@gmail.com
 * @Date: 2021/6/2 14:02
 * @Desc: 矩阵中的路径
 */

type pair struct {
	x, y int
}

// board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	visited := make(map[pair]struct{})
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			visited[pair{x, y}] = struct{}{}
			if backtrace(x, y, m, n, board, word, 0, visited) {
				return true
			}
			delete(visited, pair{x, y})
		}
	}
	return false
}

func main() {
	board := [][]byte{{'A', 'B'}, {'C', 'D'}}
	fmt.Println(exist(board, "ABCD"))
	board = [][]byte{{'a', 'b', 'c', 'e'}, {'s', 'f', 'c', 's'}, {'a', 'd', 'e', 'f'}}
	fmt.Println(exist(board, "abcced"))
	board = [][]byte{{'a'}}
	fmt.Println(exist(board, "ab"))
	board = [][]byte{{'a'}}
	fmt.Println(exist(board, "a"))
}

// x,y表示当前位置, board为矩阵，word为路径字符串，cur为路径字符串访问index
func backtrace(x int, y int, m int, n int, board [][]byte, word string, cur int, visited map[pair]struct{}) bool {
	if cur == len(word) {
		return true
	}
	if board[x][y] != ([]byte(word))[cur] {
		return false
	}
	if x-1 >= 0 {
		if _, ok := visited[pair{x - 1, y}]; !ok { // (x-1, y)没有被访问
			visited[pair{x - 1, y}] = struct{}{}
			flag := backtrace(x-1, y, m, n, board, word, cur+1, visited)
			delete(visited, pair{x - 1, y})
			if flag {
				return true
			}
		}
	}
	if x+1 <= m-1 {
		if _, ok := visited[pair{x + 1, y}]; !ok { // (x+1, y)没有被访问
			visited[pair{x + 1, y}] = struct{}{}
			flag := backtrace(x+1, y, m, n, board, word, cur+1, visited)
			delete(visited, pair{x + 1, y})
			if flag {
				return true
			}
		}
	}
	if y-1 >= 0 {
		if _, ok := visited[pair{x, y - 1}]; !ok { // (x, y-1)没有被访问
			visited[pair{x, y - 1}] = struct{}{}
			flag := backtrace(x, y-1, m, n, board, word, cur+1, visited)
			delete(visited, pair{x, y - 1})
			if flag {
				return true
			}
		}
	}
	if y+1 <= n-1 {
		if _, ok := visited[pair{x, y + 1}]; !ok { // (x, y+1)没有被访问
			visited[pair{x, y + 1}] = struct{}{}
			flag := backtrace(x, y+1, m, n, board, word, cur+1, visited)
			delete(visited, pair{x, y + 1})
			if flag {
				return true
			}
		}
	}
	if cur == len(word)-1 {
		return true
	}
	return false
}
