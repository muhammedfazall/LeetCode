func xorOperation(n int, start int) int {
    result := 0
    for i := 0 ; i < n ; i++ {
        result ^= start + 2 * i
    }
    return result
}



//  // with slice
//     result := 0
//     for i := 0 ; i < n ; i++ {
//         nums[i] = start + 2 * i
//         result ^= nums[i]
//     }
//  // with two loop
//     result := 0
//     for _,v := range nums{
//         result ^= v
//     }
//     return result