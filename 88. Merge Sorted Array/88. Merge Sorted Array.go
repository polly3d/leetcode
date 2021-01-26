func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	t := m + n - 1
	for i >= 0 || j >= 0 {
		if i < 0 {
			nums1[t] = nums2[j]
			j--
		} else if j < 0 {
			nums1[t] = nums1[i]
			i--
		} else {
			if nums1[i] < nums2[j] {
				nums1[t] = nums2[j]
				j--
			} else {
				nums1[t] = nums1[i]
				i--
			}
		}
		t--
	}
}