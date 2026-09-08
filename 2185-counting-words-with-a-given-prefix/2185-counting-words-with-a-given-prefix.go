func prefixCount(words []string, pref string) int {
	count := 0
	for _, word := range words {
		if strings.HasPrefix(word,pref){
            count++
        }
	}
	return count
}

// func prefixCount(words []string, pref string) int {
// 	count := 0
// 	n := len(pref)
// 	for _, Word := range words {
// 		if len(Word) >= n {
// 			prefix := Word[0:n]
// 			if prefix == pref {
// 				count++
// 			}
// 		}
// 	}
// 	return count
// }