func differenceOfSum(nums []int) int {
    elementSum := 0
    digitSum := 0
    for _,v := range nums{
        elementSum += v

        for v > 0 {
            digitSum += v % 10
            v /= 10 
        }
    }
    return elementSum - digitSum
}