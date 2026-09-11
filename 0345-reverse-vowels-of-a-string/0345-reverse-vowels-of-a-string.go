func reverseVowels(s string) string {
    str := []byte(s)
    l := 0
    r := len(s) - 1

    for l < r {        
        for l < r && !isVowel(str[l]) {
            l++
        }

        for l < r && !isVowel(str[r]) {
            r--
        }

        if l < r {
            str[l],str[r] = str[r],str[l]
            l++
            r--
        }
    }
    return string(str)
}

func isVowel(b byte) bool {
    switch b {
        case 'a','e','i','o','u','A','E','I','O','U':
            return true
        default :
            return false
    }
}