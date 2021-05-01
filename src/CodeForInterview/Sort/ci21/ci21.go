package ci21

func exchange(nums []int) []int {
	less := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] % 2 != 0 {
			less++
			// swap(nums, less, i)
			nums[less], nums[i] = nums[i], nums[less]
		}
	}
	return nums
}


