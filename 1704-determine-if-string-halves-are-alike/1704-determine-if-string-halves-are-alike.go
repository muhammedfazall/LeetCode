func isVowel(s rune) bool {
	switch s{
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}

func halvesAreAlike(s string) bool {
    l := 0
	r := len(s) - 1
	rightVowelCount := 0
	leftVowelCount := 0

	for l < r {
		if isVowel(rune(s[l])) {
			leftVowelCount++
		}
		if isVowel(rune(s[r])) {
			rightVowelCount++
		}
        l++
        r--
	}
    return rightVowelCount == leftVowelCount
}
