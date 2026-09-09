func minElement(nums []int) int {
	min := nums[0]
	for i, num := range nums {
		sum := 0
		for num > 0 {
			sum += num % 10
			num /= 10
		}
		nums[i] = sum
		if sum < min {
			min = sum
		}
	}
	return min
}