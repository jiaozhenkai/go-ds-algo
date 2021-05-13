package array

func moveZeroes(nums []int) {
	n := len(nums)
	if n == 0 || n == 1 {
		return
	}

	pos := 0 //下一个应该被非 0 元素填充的位置，如果没有非 0 元素，则不做任何操作。
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[pos] = nums[i]
			pos++
		}
	}

	if pos == 0 {
		return
	} // 表示没有非 0 元素

	for p := pos; p < n; p++ {
		nums[p] = 0
	}
}

func moveZeroes1(nums []int) {
	n := len(nums)
	if n == 0 || n == 1 {
		return
	}

	pos := 0
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[i], nums[pos] = nums[pos], nums[i]
			pos++
		}
	}

}
