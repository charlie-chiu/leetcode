package median_of_two_sorted_arrays

// 16ms, 6MB on leetcode
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	var p1, p2 int
	var merged []int

	for i := 0; i < totalLen; i++ {
		if p1 >= len(nums1) {
			merged = append(merged, nums2[p2])
			p2++
			continue
		}
		if p2 >= len(nums2) {
			merged = append(merged, nums1[p1])
			p1++
			continue
		}

		if nums1[p1] < nums2[p2] {
			merged = append(merged, nums1[p1])
			p1++
		} else {
			merged = append(merged, nums2[p2])
			p2++
		}
	}

	if totalLen%2 == 0 {
		//even
		return float64(merged[totalLen/2]+merged[(totalLen/2)-1]) / 2
	} else {
		//odd
		return float64(merged[totalLen/2])
	}
}
