package sort

import (
	"math/rand"
	"testing"
)

func TestQuickSort(t *testing.T) {
	// 生成一个1000以内的随机数表示待排序数组有多少个元素
	count := rand.Intn(1000)
	nums := make([]int, count)
	// 生成随机数数组
	for i := 0; i < count; i++ {
		nums[i] = rand.Intn(10000)
	}
	Quicksort(nums)
	for i := 0; i < count-1; i++ {
		if nums[i] > nums[i+1] {
			t.Error("Fuck!")
		}
	}
}
