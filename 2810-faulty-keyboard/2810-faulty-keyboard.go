func finalString(s string) string {
    st := []byte(s)
    str := []byte{}

    for _,v := range st{
        if v == 'i'{
            reverse(str)
        } else {
        str = append(str,v)
        }
    }
    return string(str)
}

func reverse(s []byte) []byte {
    l := 0
    r := len(s) - 1
    for l < r {
        s[l],s[r] = s[r],s[l]
        l++
        r--
    }
    return s
}