package array

func removeElement(nums []int, val int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	pos := 0 // pos 指向下一个可以被不等于 val 的值覆盖的元素。
	for i := 0; i < n; i++ {
		if nums[i] != val {
			nums[pos] = nums[i]
			pos++
		}
	}

	return pos
}
