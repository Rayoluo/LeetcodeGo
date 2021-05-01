package ci51

func reversePairs(nums []int) int {
	if len(nums) == 0 || nums == nil {
		return 0
	}
	arr := make([]int, len(nums))
	return divide(nums, 0, len(nums)-1, arr)
}

func divide(nums []int, l int, r int, arr []int) int {
	if l == r {
		return 0
	}
	mid := l + (r-l)>>1
	return divide(nums, l, mid, arr) + divide(nums, mid+1, r, arr) + conquer(nums, l, r, arr)
}

func conquer(nums []int, l int, r int, arr []int) (count int) {
	var (
		mid, ll, rr, pos int
	)
	mid = l + (r-l)>>1
	ll, rr, pos, count = l, mid+1, l, 0
	// arr = make([]int, r-l+1)
	for ll <= mid && rr <= r {
		if nums[ll] <= nums[rr] {
			arr[pos] = nums[ll]
			pos++
			ll++
		} else {
			arr[pos] = nums[rr]
			pos++
			rr++
			count += mid - ll + 1
		}
	}
	for ll <= mid {
		arr[pos] = nums[ll]
		pos++
		ll++
	}
	for rr <= r {
		arr[pos] = nums[rr]
		pos++
		rr++
	}
	// 将排好序的数组拷贝回原数组
	for i := l; i <= r; i++ {
		nums[i] = arr[i]
	}
	return
}
