package array

import (
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}
	sort.Ints(nums1)
}

func merge1(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	index1 := 0
	index2 := 0

	for index1 < m && index2 < n {
		if nums1[index1] < nums2[index2] {
			sorted = append(sorted, nums1[index1])
			index1++
		} else {
			sorted = append(sorted, nums2[index2])
			index2++
		}
	}

	if index1 == m {
		sorted = append(sorted, nums2[index2:]...)
	}

	if index2 == n {
		sorted = append(sorted, nums1[index1:]...)
	}

	copy(nums1, sorted)
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	curr := m + n - 1 // current insert position
	index1 := m - 1   // [0, index1] not compare
	index2 := n - 1   // [0, index2] not compare

	for index1 > -1 && index2 > -1 {
		if nums2[index2] > nums1[index1] {
			nums1[curr] = nums2[index2]
			curr--
			index2--
		} else {
			nums1[curr] = nums1[index1]
			curr--
			index1--
		}
	}

	if index1 == -1 {
		for i := index2; i > -1; i-- {
			nums1[curr] = nums2[i]
			curr--
		}
	}

	if index2 == -1 {
		for i := index1; i > -1; i-- {
			nums1[curr] = nums1[i]
		}
	}
}

//func helper(nums1 []int, index int) {
//	for i
//}
