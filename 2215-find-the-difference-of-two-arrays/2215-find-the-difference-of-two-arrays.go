func findDifference(nums1 []int, nums2 []int) [][]int {
    
    s1 := make(map[int]bool)
    s2 := make(map[int]bool)

    for _,v := range nums1{ s1[v] = true }
    for _,v := range nums2{ s2[v] = true }

    var diff1 []int
    var diff2 []int

    for n := range s1{
        if !s2[n] {
            diff1 = append(diff1,n)
        }
    }

    for n := range s2{
        if !s1[n] {
            diff2 = append(diff2,n)
        }
    }

    return [][]int{diff1,diff2}
}