package sort

import "math/rand"

func Quicksort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	quicksort(arr, 0, len(arr)-1)
}

func quicksort(arr []int, l int, r int) {
	if l < r {
		swap(arr, l+rand.Intn(r-l+1), r)
		pivot := partition(arr, l, r)
		quicksort(arr, l, pivot-1)
		quicksort(arr, pivot+1, r)
	}
}

func partition(arr []int, l int, r int) int {
	less := l - 1
	// arr[r]作为枢纽元
	for i := l; i <= r; i++ {
		if arr[i] <= arr[r] {
			less++
			swap(arr, less, i)
		}
	}
	return less
}

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
