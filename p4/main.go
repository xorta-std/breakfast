package main

import (
	"slices"
)

// 4. Median of Two Sorted Arrays

func getMedian(nums []int) float64 {
	m := len(nums) / 2
	b0 := len(nums) & 1
	if b0 == 0 {
		return float64(nums[m]+nums[m-1]) / 2
	} else {
		return float64(nums[m])
	}
}

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	slices.Sort(nums1)
	return getMedian(nums1)
}

// 31 times faster :-o
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	nums := make([]int, len(nums1)+len(nums2))

	i := 0
	j := 0

	for {
		if i < len(nums1) && j >= len(nums2) ||
			i < len(nums1) && nums1[i] < nums2[j] {
			nums[i+j] = nums1[i]
			i += 1
		} else if j < len(nums2) {
			nums[i+j] = nums2[j]
			j += 1
		} else {
			break
		}
	}

	return getMedian(nums)
}
