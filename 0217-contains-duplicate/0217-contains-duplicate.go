func containsDuplicate(nums []int) bool {

    if len(nums) == 1 {
        return false
    }

    freq := make(map[int]int, len(nums))

    for _,v := range nums{
        freq[v]++
        if freq[v] > 1{
            return true
        } 
    }

    return false
}