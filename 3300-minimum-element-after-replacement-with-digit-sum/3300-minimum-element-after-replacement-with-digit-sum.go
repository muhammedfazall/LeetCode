func minElement(nums []int) int {

	getDigitSum := func(n int) int {
		s := 0
		for n > 0 {
			s += n % 10
			n /= 10
		}
        return s
	}

    min := getDigitSum(nums[0])
	for _, num := range nums[1:] {	
        sum := getDigitSum(num)
		if sum < min {
			min = sum
		}
	}
	return min
}