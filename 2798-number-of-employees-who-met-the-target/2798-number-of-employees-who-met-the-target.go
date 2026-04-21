func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
    n := len(hours)
    count := 0
    for i:=0 ; i <= n-1 ; i++ {
        if hours[i] >= target {
            count++
        } 
    }
    return count
}