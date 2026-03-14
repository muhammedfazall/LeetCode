func areOccurrencesEqual(s string) bool {
	m := make(map[rune]int)

	for _, ch := range s {
		m[ch]++
	}

    if len(m) == 1 {
        return true
    }

    var count int
	for _,v := range m {
		count = v
	}

    for _,v := range m{
        if v != count{
            return false
        }
    }
	return true
}