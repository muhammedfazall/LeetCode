import (
    "slices"
)

func finalString(s string) string {

	str := make([]rune,0,len(s))

	for _, ch := range s {
		switch ch{
            case 'i' : 
                slices.Reverse(str)
            default  : 
            str = append(str,ch)
        }
	}
	return string(str)
}

// func finalString(s string) string {

// 	str := []byte(s)
// 	final := []byte{}

// 	for _, v := range str {
// 		if v == 'i' {
// 			reverse(final)
// 		} else {
// 			final = append(final, v)
// 		}
// 	}
// 	return string(final)
// }

// func reverse(s []byte) []byte {
// 	l := 0
// 	r := len(s) - 1
// 	for l < r {
// 		s[l], s[r] = s[r], s[l]
// 		l++
// 		r--
// 	}
// 	return s
// }