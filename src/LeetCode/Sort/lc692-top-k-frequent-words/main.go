package main

import (
	"container/heap"
	"fmt"
)

/**
 * @Author: luo
 * @Author: hustly123@gmail.com
 * @Date: 2021/5/20 9:12
 * @Desc: 前k个高频单词--给一非空的单词列表，返回前K个出现次数最多的单词
 */

type pair struct {
	key string
	val int
}

// 建立一个小根堆
type MinHeap []*pair

func (h MinHeap) Less(i, j int) bool {
	if h[i].val == h[j].val {
		return h[i].key > h[j].key
	}
	return h[i].val < h[j].val
}
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Len() int            { return len(h) }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(*pair)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func topKFrequent(words []string, k int) []string {
	mp := make(map[string]int)
	var (
		times, count int
		ok           bool
		x            *pair
	)
	ss := make([]string, k)
	for i := 0; i < len(words); i++ {
		if times, ok = mp[words[i]]; !ok { // 如果已经出现了
			mp[words[i]] = 1
		} else {
			mp[words[i]] = times + 1
		}
	}
	// 遍历mp中的每个单词
	h := make(MinHeap, 0)
	for key, val := range mp {
		x = &pair{key, val}
		if count < k { // 前k个元素建堆
			count++
			heap.Push(&h, x)
		} else { // 剩下的元素首先和堆顶元素比较，如果小于丢掉，大于弹出堆顶元素，同时插入该元素
			// 这里有点小bug需要注意。。。
			if h[0].val < x.val || (h[0].val == x.val && h[0].key > x.key) {
				h[0] = x
				heap.Fix(&h, 0)
			}
		}
	}
	// 依次弹出堆中元素并放入结果集合中
	for i := k - 1; i >= 0; i-- {
		ss[i] = heap.Pop(&h).(*pair).key
	}
	return ss
}

func main() {
	fmt.Println(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 3))
	fmt.Println(topKFrequent([]string{"aaa", "aa", "a"}, 1))
}
