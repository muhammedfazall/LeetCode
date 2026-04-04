func smallerNumbersThanCurrent(nums []int) []int {
    n := len(nums)
    result := make([]int,n)
    for i:= 0 ; i < n ; i ++ {
        count := 0
        for j:= 0 ; j < n ; j ++ {
            if j == i {
                continue
            } 
            if nums[j] < nums[i] {
                count += 1
            }
        }
        result[i] = count
    }
    return result
}