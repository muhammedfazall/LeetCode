func prefixCount(words []string, pref string) int {
	count := 0
	n := len(pref)
	for _, v := range words {
		if len(v) >= n {
			prefix := v[0:n]
			if prefix == pref {
				count++
			}
		}
	}
	return count
}