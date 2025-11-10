func twoSum(nums []int, target int) []int {
    m:= make(map[int]int)

    for i,v := range nums{
        diff := target - v;
        if value,ok := m[diff];ok{
            return []int{value,i}
        }
        m[v] = i
    }
    return []int{}
}