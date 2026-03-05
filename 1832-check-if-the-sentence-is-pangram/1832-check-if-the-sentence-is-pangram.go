func checkIfPangram(sentence string) bool {
    seen := make(map[rune]bool,len(sentence))

    for _,ch := range sentence {
        seen[ch] = true
    }

    if len(seen) == 26 {
        return true
    }
    return false
}