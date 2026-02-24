func firstPalindrome(words []string) string {
    for _,v := range words{
        r := []rune(v)

        for i,j := 0, len(r)-1 ; i < j ; i,j = i+1,j-1{
            r[i],r[j] = r[j],r[i]
        }

        if string(r) == v{
            return v
        }
    }
    return ""
}