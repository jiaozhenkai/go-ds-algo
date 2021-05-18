package array

import (
	"sort"
)

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	m := make(map[int]struct{})

	for _, v := range nums {
		m[v] = struct{}{}
	}

	var i int
	tmp := make([]int, len(m))
	for k, _ := range m {
		tmp[i] = k
		i++
	}

	sort.Ints(tmp)
	for j := 0; j < len(tmp); j++ {
		nums[j] = tmp[j]
	}

	return len(m)
}

func removeDuplicates1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	slow := 1

	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}
