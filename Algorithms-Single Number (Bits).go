// Time: O(n)/O(n * p); Space: O(1)/O(p)
// Each element in nums appears exactly p times except for one element which appears once
func SingleNumber(nums []int, p int) int {
	if p&1 == 0 {
		var result int
		for i := range nums {
			result ^= nums[i]
		}

		return result
	}

	var mask int
	s := make([]int, p)

	for i := range nums {
		for k := p - 1; k > 0; k-- {
			s[k] ^= s[k-1] & nums[i]
		}
		s[0] ^= nums[i]

		mask = ^(s[p-1] & s[p-2])
		for k := range s {
			s[k] &= mask
		}
	}

	return s[0]
}
