func findDifference(nums1 []int, nums2 []int) [][]int {
    
    s1 := make(map[int]bool)
    s2 := make(map[int]bool)

    for _,v := range nums1{ s1[v] = true }
    for _,v := range nums2{ s2[v] = true }

    answer := make([][]int,2)

    for n := range s1{
        if !s2[n] {
            answer[0] = append(answer[0],n)
        }
    }

    for n := range s2{
        if !s1[n] {
            answer[1] = append(answer[1],n)
        }
    }

    return answer
}