func concatWithReverse(nums []int) []int {
	n := len(nums)
	ans := make([]int, 0, n*2)
	ans = append(ans, nums...)

	for  i := n - 1; i >= 0; i-- {
		ans = append(ans, nums[i])
	}
	return ans
}