func addDigits(num int) int {
	if num == 0 {
        return num
    }

    result := num % 9
    if result == 0{
        return 9
    }else{
        return result
    }
}