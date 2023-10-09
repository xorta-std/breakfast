package main

import (
	"math/rand"
	"sort"
	"testing"
)

type Data struct {
	Nums1 []int
	Nums2 []int
}

func generateRandomTestCases(nums1Len int, nums2Len int) Data {
	nums1 := make([]int, nums1Len)
	nums2 := make([]int, nums2Len)

	for i := 0; i < nums1Len; i++ {
		nums1[i] = rand.Intn(100_000)
	}

	for i := 0; i < nums2Len; i++ {
		nums2[i] = rand.Intn(100_000)
	}
	sort.Ints(nums1)
	sort.Ints(nums2)

	return Data{
		Nums1: nums1,
		Nums2: nums2,
	}
}

var testCase Data

func getTestCase() Data {
	testCase = generateRandomTestCases(rand.Intn(10_000), rand.Intn(10_000))
	return testCase
}

func BenchmarkFindMedianSortedArrays1(b *testing.B) {
	data := getTestCase()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findMedianSortedArrays1(data.Nums1, data.Nums2)
	}
}

func BenchmarkFindMedianSortedArrays2(b *testing.B) {
	data := getTestCase()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findMedianSortedArrays2(data.Nums1, data.Nums2)
	}
}
