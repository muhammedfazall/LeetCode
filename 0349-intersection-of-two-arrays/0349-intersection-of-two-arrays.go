func intersection(nums1 []int, nums2 []int) []int {
	m := map[int]bool{}
    inter := []int{}
	for _,v := range nums1{
        m[v] = true
    }
    for _,v := range nums2{
        if m[v] {
            inter = append(inter,v)
            delete(m,v)
        }
    }
    return  inter
}