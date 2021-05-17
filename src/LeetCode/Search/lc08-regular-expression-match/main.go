package main

import "fmt"

/**
 * @Author: luo
 * @Author: hustly123@gmail.com
 * @Date: 2021/5/17 14:21
 * @Desc: 动态规划之正则表达式
 */
type pair struct {
	i, j int
}

var memo map[pair]bool

func isMatch(s string, p string) bool {
	memo = make(map[pair]bool)
	return backtrace(s, 0, p, 0)
}

func backtrace(s string, i int, p string, j int) bool {
	if v, ok := memo[pair{i, j}]; ok {
		return v
	}
	slen, plen := len(s), len(p)
	var res bool
	// base case1: 模式串**匹配完了**，如果字符串刚好访问完则返回true，否则返回false
	if j == plen {
		res = i == slen
		memo[pair{i, j}] = res
		return res
	}
	// base case2: 字符串访问完了，则普通字符（或者.）和*必须成对出现，没有**的写法
	if i == slen {
		if (plen-j)%2 == 1 {
			res = false
			// return false
			memo[pair{i, j}] = res
			return res
		}
		for ; j < plen-1; j += 2 {
			if p[j+1] != '*' {
				res = false
				// return false
				memo[pair{i, j}] = res
				return res
			}
		}
		res = true
		// return true
		memo[pair{i, j}] = res
		return res
	}

	// 当前访问的字符匹配上, 即字符相同或者模式串字符为.
	if s[i] == p[j] || p[j] == '.' {
		if j+1 < plen && p[j+1] == '*' {
			res = backtrace(s, i+1, p, j) || backtrace(s, i, p, j+2)
			// return backtrace(s, i+1, p, j) || backtrace(s, i, p, j+2)
			memo[pair{i, j}] = res
			return res
		} else {
			res = backtrace(s, i+1, p, j+1)
			// return backtrace(s, i+1, p, j+1)
			memo[pair{i, j}] = res
			return res
		}
	} else { // 当前访问的字符不匹配
		if j+1 < plen && p[j+1] == '*' {
			res = backtrace(s, i, p, j+2)
			// return backtrace(s, i, p, j+2)
			memo[pair{i, j}] = res
			return res
		} else {
			res = false
			// return false
			memo[pair{i, j}] = res
			return res
		}
	}
}

func main() {
	fmt.Println(isMatch("aab", "c*a*b"))
}
