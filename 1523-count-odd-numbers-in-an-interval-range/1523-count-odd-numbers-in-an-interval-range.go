func countOdds(low int, high int) int {
	if low%2 == 1 || high%2 == 1{
        return (high - low ) / 2 + 1
    }
    return (high - low )/2
}