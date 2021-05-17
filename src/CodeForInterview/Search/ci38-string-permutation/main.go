package main

import "fmt"

/**
 * @Author: luo
 * @Author: hustly123@gmail.com
 * @Date: 2021/5/17 15:08
 * @Desc: 给定一个字符串，打印出该字符串中字符的所有排列
 */

// 思路：用一个map统计每个字符在字符串中出现了多少次, 回溯，剪枝
var m map[byte]int

func permutation(s string) []string {
	p := []byte(s)
	m = make(map[byte]int)
	// 统计字符串每个字符出现的次数
	for i := 0; i < len(p); i++ {
		if v, ok := m[p[i]]; ok {
			m[p[i]] = v + 1
		} else {
			m[p[i]] = 1
		}
	}
	res := make([]byte, len(s))
	// set := make(map[string]struct{})
	// backtrace(res, 0, len(s), set)
	ss := make([]string, 0)
	// for k := range set {
	// 	ss = append(ss, k)
	// }
	backtrace(res, 0, len(s), &ss)
	return ss
}

func backtrace(res []byte, i int, n int, s *[]string) {
	if i == n { // 访问完毕
		str := string(res)
		*s = append(*s, str)
		return
	}
	for k, v := range m {
		if v > 0 {
			res[i] = k
			m[k] = v - 1
			backtrace(res, i+1, n, s)
			m[k] = v
		}
	}
}

func main() {
	fmt.Println(permutation("abc"))
}
