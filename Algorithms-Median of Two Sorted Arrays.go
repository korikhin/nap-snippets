// Time: O(log(min(m, n))); Space: O(1)
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	left, right := 0, m

	for left <= right {
		partitionX := (left + right) / 2
		partitionY := (m+n+1)/2 - partitionX

		maxLeftX := math.Inf(-1)
		if partitionX != 0 {
			maxLeftX = float64(nums1[partitionX-1])
		}

		minRightX := math.Inf(1)
		if partitionX != m {
			minRightX = float64(nums1[partitionX])
		}

		maxLeftY := math.Inf(-1)
		if partitionY != 0 {
			maxLeftY = float64(nums2[partitionY-1])
		}

		minRightY := math.Inf(1)
		if partitionY != n {
			minRightY = float64(nums2[partitionY])
		}

		if maxLeftX <= minRightY && maxLeftY <= minRightX {
			if (m+n)&1 == 0 {
				return (max(maxLeftX, maxLeftY) + min(minRightX, minRightY)) / 2
			}
			return max(maxLeftX, maxLeftY)
		} else if maxLeftX > minRightY {
			right = partitionX - 1
		} else {
			left = partitionX + 1
		}
	}

	return 0
}
