func countDigits(num int) int {
	count := 0
    temp := num

    if temp < 0 {
        temp = -temp
    }
    
	for temp > 0 {
		n := temp % 10

		if n > 0 && num%n == 0 {
			count++
		}
		temp = temp / 10
	}
	return count
}