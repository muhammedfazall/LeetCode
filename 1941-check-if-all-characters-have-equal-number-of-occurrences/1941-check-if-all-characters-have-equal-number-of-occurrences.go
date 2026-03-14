func areOccurrencesEqual(s string) bool {
	m := make(map[rune]int)
	count := 0
	for _, ch := range s {
		m[ch]++
		count = m[ch]
	}

	for k,_ := range m {
		if m[k] != count {
			return false
		}
	}
	return true
}