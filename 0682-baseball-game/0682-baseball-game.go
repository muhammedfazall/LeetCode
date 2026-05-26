func calPoints(operations []string) int {
    var stack []int

    for _, op := range operations {
        switch op {
        case "C":
            if len(stack) > 0 {
                stack = stack[:len(stack)-1]
            }
        case "D":
            if len(stack) > 0 {
                lastScore := stack[len(stack)-1]
                stack = append(stack, lastScore*2)
            }
        case "+":
            if len(stack) >= 2 {
                last := stack[len(stack)-1]
                secondLast := stack[len(stack)-2]
                stack = append(stack, last+secondLast)
            }
        default:
            score, _ := strconv.Atoi(op)
            stack = append(stack, score)
        }
    }

    totalSum := 0
    for _, score := range stack {
        totalSum += score
    }

    return totalSum
}