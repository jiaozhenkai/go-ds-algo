package array

import "sort"

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

func findKthLargest1(nums []int, k int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}

	if numsLen == 1 {
		return nums[0]
	}

	target := numsLen - k

	return find(nums, 1, numsLen-1, nums[0], target)
}

func find(nums []int, left, right, pivot, target int) int {
	return 0
}
