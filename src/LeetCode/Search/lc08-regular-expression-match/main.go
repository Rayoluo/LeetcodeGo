package main

import "fmt"

/**
 * @Author: luo
 * @Author: hustly123@gmail.com
 * @Date: 2021/5/16 16:09
 * @Desc: 正则表达式匹配
 */

func isMatch(s string, p string) bool {
	// return backtrack(s, p, 0, 0)
}

func backtrace(s, p string, l1, l2 int) bool {
	if p[l2] != '*' { // 当前模式串为普通字符或者.通配符
		if l2+1 >= len(p) || p[l2+1] != '*' { // 模式串下一个字符不是*
			if p[l2] == '.' || p[l2] == s[l1] {
				return backtrace(s, p, l1+1, l2+1)
			}
			return false
		}
		// 模式串下一个字符为*
		// -- 如果当前模式串字符为普通字符
	}

}

func main() {
	fmt.Println(isMatch("aab", "c*a*b"))
	fmt.Println(isMatch("mississippi", "mis*is*p*."))
	fmt.Println(isMatch("a", "ab*"))
	fmt.Println(isMatch("aad", ".*")) // 预期结果为true
	fmt.Println(isMatch("bbbba", ".*a*a"))
}
