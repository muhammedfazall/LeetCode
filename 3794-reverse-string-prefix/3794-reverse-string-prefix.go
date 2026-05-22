func reversePrefix(s string, k int) string {
	start := 0
	end := k - 1
	bytes := []byte(s)
	for start < end {
		bytes[start], bytes[end] = bytes[end], bytes[start]
		start++
		end--
	}
	return string(bytes)
}