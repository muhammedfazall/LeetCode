func dominantIndex(nums []int) int {
    largest := -1
    index := -1
    for i,v := range nums{
        if v > largest {
            largest = v
            index = i
        }
    }
    for i,v := range nums{
        if i != index && largest < 2*v{
            return -1
        }  
    }
    return index
}