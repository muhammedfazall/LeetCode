func subtractProductAndSum(n int) int {
    product := 1
    sum := 0
    l := 0

    for n > 0 {
        l = n % 10 
        product = product * l
        sum = sum + l
        n /= 10
    }

    return product - sum
}