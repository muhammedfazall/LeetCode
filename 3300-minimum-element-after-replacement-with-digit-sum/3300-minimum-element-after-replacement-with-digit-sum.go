func minElement(nums []int) int {
	min := math.MaxInt
	for _, num := range nums {
		sum := 0
		for num > 0 {
			sum += num % 10
			num /= 10
		}
		if sum < min {
			min = sum
		}
	}
	return min
}