func mostWordsFound(sentences []string) int {
	max := 0

    for _,s := range sentences {
        count := 1

        for i := 0; i < len(s) ; i++{
            if s[i] == ' '{
                count ++
            }
        }
        if count > max {
            max = count
        }
    }
    return max
}



// func mostWordsFound(sentences []string) int {
// 	wordCount := map[string]int{}
// 	for _,s := range sentences {
// 		count := 0
// 		for _,ch := range s {
// 			if string(ch) == " " {
// 				count++
// 			}
// 		}
// 		wordCount[s] = count + 1
// 	}

// 	max := 0
// 	for v := range wordCount {
// 		if wordCount[v] > max {
// 			max = wordCount[v]
// 		}
// 	}
// 	return max
// }