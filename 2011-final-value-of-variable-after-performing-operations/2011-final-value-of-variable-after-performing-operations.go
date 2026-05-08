func finalValueAfterOperations(operations []string) int {
    X := 0
    for _,v := range operations {
        if v == "--X" || v == "X--" {
          X-- 
        }
        if v == "++X" || v == "X++" {
          X++
        }
    }
    return X
}