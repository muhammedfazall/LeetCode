func containsDuplicate(nums []int) bool {
    freq := make(map[int]int)

    for _,v := range nums{
        freq[v]++
        if freq[v] > 1{
            return true
        } 
    }

    return false
}