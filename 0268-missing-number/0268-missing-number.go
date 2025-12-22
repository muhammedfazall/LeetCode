func missingNumber(nums []int) int {
	n := len(nums)
	actualS := 0
	expectedS := n * (n + 1) / 2
	for _, v := range nums {
		actualS += v
	}
	return expectedS - actualS
}