func findWordsContaining(words []string, x byte) []int {
    indices := []int{}
    for i,w := range words {
        for _,v := range w{
            if byte(v) == x {
                indices = append(indices,i)
                break
            }
        }
    }
    return indices
}