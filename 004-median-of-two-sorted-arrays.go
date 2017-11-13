package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	nums1Len := len(nums1)
	nums2Len := len(nums2)
	// v1 and v2 store the median(s)
	v1 := 0
	v2 := 0
	// index of nums1/2 during the loop
	i1 := 0
	i2 := 0
	for i := 0; i <= totalLength/2; i++ { // loop until middle is reached
		// make v1 stores v2's old value
		v1 = v2

		// assign v2 the smaller one between nums1[i1] and nums2[i2]
		// , and plus 1 corresponding i1 or i2

		// if nums1 is empty or its iteration finished
		//, then only nums2 will be iterated
		if nums1Len == 0 || i1 >= nums1Len {
			v2 = nums2[i2]
			i2++
			continue
		}

		// also
		if nums2Len == 0 || i2 >= nums2Len {
			v2 = nums1[i1]
			i1++
			continue
		}

		if nums1[i1] < nums2[i2] {
			v2 = nums1[i1]
			i1++
		} else {
			v2 = nums2[i2]
			i2++
		}
	}

	// if total length of nums1 and nums2 is odd
	// , return v2 which is just the mid value
	if totalLength&1 != 0 {
		return float64(v2)
	}
	return float64(v1)/2 + float64(v2)/2
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2, 4}
	println(findMedianSortedArrays(nums1, nums2))
	nums1 = []int{1, 3}
	nums2 = []int{2}
	println(findMedianSortedArrays(nums1, nums2))
	nums1 = []int{}
	nums2 = []int{2, 4}
	println(findMedianSortedArrays(nums1, nums2))
	nums1 = []int{}
	nums2 = []int{1}
	println(findMedianSortedArrays(nums1, nums2))
	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	println(findMedianSortedArrays(nums1, nums2))
}
