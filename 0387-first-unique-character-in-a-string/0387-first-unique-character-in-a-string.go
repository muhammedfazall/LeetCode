func firstUniqChar(s string) int {
    freq := make(map[byte]int)

    for i := range s{
        freq[s[i]]++
    }

    for i := range s{
        if freq[s[i]] == 1{
            return i
        }
    }
    return -1
}