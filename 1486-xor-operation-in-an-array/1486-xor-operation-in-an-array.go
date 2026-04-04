func xorOperation(n int, start int) int {
    nums := make([]int,n)
    result := 0
    for i := 0 ; i < n ; i++ {
        nums[i] = start + 2 * i
        result ^= nums[i]
    }
    return result
}