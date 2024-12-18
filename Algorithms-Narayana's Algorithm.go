// Time: O(n); Space: O(1); Sorting: Lexicographic
func NextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}

	i := len(nums) - 2
	for i >= 0 && nums[i+1] <= nums[i] {
		i--
	}

	if i >= 0 {
		j := len(nums) - 1
		for nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	for k, j := i+1, len(nums)-1; k < j; k, j = k+1, j-1 {
		nums[k], nums[j] = nums[j], nums[k]
	}
}
