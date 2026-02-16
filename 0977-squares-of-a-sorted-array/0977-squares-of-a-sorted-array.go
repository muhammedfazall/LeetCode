func sortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int,n)
	left := 0
	right := len(nums) - 1
	p := n - 1

	for left <= right {
		r := nums[right] * nums[right]
		l := nums[left] * nums[left]
		if l > r {
			result[p] = l
			left++
		} else {
			result[p] = r
			right--
        }
		p--
	}
	return result
}