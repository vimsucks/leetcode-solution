package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	nums1Len := len(nums1)
	nums2Len := len(nums2)
	v1 := 0
	v2 := 0
	i1 := 0
	i2 := 0
	for i := 0; i <= totalLength/2; i++ {
		v1 = v2

		if nums1Len == 0 || i1 >= nums1Len {
			v2 = nums2[i2]
			i2++
			continue
		}

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
