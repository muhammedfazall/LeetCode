func maximumWealth(accounts [][]int) int {
	richest := 0
	for _, row := range accounts {
		w := 0
		for _, val := range row {
			w += val
		}
		if w > richest {
			richest = w
		}
	}
	return richest
}