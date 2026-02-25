func longestPalindrome(s string) string {
	longest := ""
	for i := 0; i < len(s); i++ {
		odd := expand(s, i, i)
		even := expand(s, i, i+1)
		if len(odd) > len(longest) {
			longest = odd
		}
        if len(even) > len(longest) {
			longest = even
		}
	}
	return longest
}

func expand(s string, l, r int) string {
	for l >= 0 && r <= len(s)-1 && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}