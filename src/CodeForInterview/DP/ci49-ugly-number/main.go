package main

import "fmt"

// 将只包含质因子2 3和5的数称为丑数，求按从小到大的顺序的第n个丑数
// 迭代规则，n必须能够整除2 3 5其中之一，且整除后的结果为丑数
// 这种写法算出来的答案是正确的，但是超时了
func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n == 3 {
		return 3
	}
	if n == 4 {
		return 4
	}
	if n == 5 {
		return 5
	}
	mp := make(map[int]struct{})
	mp[1] = struct{}{}
	mp[2] = struct{}{}
	mp[3] = struct{}{}
	mp[4] = struct{}{}
	mp[5] = struct{}{}
	i := 6
	for count := 5; count < n; i++ {
		if i%2 == 0 {
			_, ok := mp[i/2]
			if ok {
				mp[i] = struct{}{}
				count++
				continue
			}
		}
		if i%3 == 0 {
			_, ok := mp[i/3]
			if ok {
				mp[i] = struct{}{}
				count++
				continue
			}
		}
		if i%5 == 0 {
			_, ok := mp[i/5]
			if ok {
				mp[i] = struct{}{}
				count++
				continue
			}
		}
	}
	return i - 1
}

// 三指针解法, p2 p3 p5都指向第一个丑数1，然后分别乘以2 3 5，取最小的结果作为下一个丑数
// 接下来将p2++指向下一个丑数
func nthUglyNumber1(n int) int {
	uglyNumArr := make([]int, n+1)
	uglyNumArr[1] = 1
	p2 := 1
	p3 := 1
	p5 := 1
	for i := 2; i < n+1; i++ {
		uglyNumArr[i] = minTriple(uglyNumArr[p2]*2, uglyNumArr[p3]*3, uglyNumArr[p5]*5)
		if uglyNumArr[i] == uglyNumArr[p2]*2 {
			p2++
		}
		if uglyNumArr[i] == uglyNumArr[p3]*3 {
			p3++
		}
		if uglyNumArr[i] == uglyNumArr[p5]*5 {
			p5++
		}
	}
	return uglyNumArr[n]
}

func minTriple(a, b, c int) int {
	if a > b {
		if b > c {
			return c
		}
		return b
	}
	if a > c {
		return c
	}
	return a
}

func main() {
	fmt.Println(nthUglyNumber(1352))
	fmt.Println(nthUglyNumber(1352))
}
