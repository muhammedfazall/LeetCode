func longestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	longest := s[0:1]
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			temp := s[i:j]
			if len(temp) > len(longest) && IsPalindrome(temp) {
				longest = temp
			}
		}
	}
	return longest
}

func IsPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}