func missingNumber(nums []int) int {
    m := map[int]int{}
    n := len(nums)
        for i := 0 ; i <= n - 1 ; i++{
            m[nums[i]] = i 
        }

        for i := 0 ; i <= n ; i++{
           if _,ok := m[i] ; !ok {
            return i
           }    
        }
    return 0
}