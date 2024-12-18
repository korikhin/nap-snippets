// Time: O(k), k - number of 1-bits; Space: O(1)
func HammingWeight(num uint32) int {
	var w uint8
	for num > 0 {
		num &= num - 1
		w++
	}

	return int(w)
}
