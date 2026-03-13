func countConsistentStrings(allowed string, words []string) int {
	m := make(map[rune]int)
	for _, v := range allowed {
		m[v]++
	}
	count := 0
	for i := 0; i < len(words); i++ {
    consistent := true
		for _, v := range words[i] {
			if _, ok := m[v]; !ok {
            consistent = false
			}
		}
        if consistent {
		count++
        }
	}

	return count
}