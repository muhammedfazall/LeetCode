func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
    count := 0
    for i:=0 ; i <= len(hours)-1 ; i++ {
        if hours[i] >= target {
            count++
        } 
    }
    return count
}