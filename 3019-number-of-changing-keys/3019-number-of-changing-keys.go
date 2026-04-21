func countKeyChanges(s string) int {
count := 0    
lower := strings.ToLower(s)
    for i := 0 ; i < len(lower)-1 ; i ++ {
        if lower[i] != lower[i+1]{
            count++
        }
    }
    return count
}