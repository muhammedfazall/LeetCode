func numberGame(nums []int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		if min != i {
			nums[i], nums[min] = nums[min], nums[i]
		}
	}

	for i := 0; i < n; i++ {
		if i%2 != 0 {
			nums[i], nums[i-1] = nums[i-1], nums[i]
		}
	}
	return nums
}