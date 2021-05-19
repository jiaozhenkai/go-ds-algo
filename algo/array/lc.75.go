package array

func sortColors(nums []int) {
	n := len(nums)
	if n == 0 || n == 1 {
		return
	}

	count := make([]int, 3) // count[i] 表示 i 元素的个数。len(count) 表示元素的种数。
	for i := 0; i < n; i++ {
		count[nums[i]]++
	}

	var index int
	for i := 0; i < len(count); i++ {
		for j := 0; j < count[i]; j++ {
			nums[index] = i
			index++
		}
	}
}

func sortColors1(nums []int) {
	n := len(nums)
	if n == 0 || n == 1 {
		return
	}

	zero := -1 // [0, zero] 区间元素都为 0
	two := n   // [two, n-1] 区间元素都为 2

	for i := 0; i < two; {

		if nums[i] == 1 {
			i++
		} else if nums[i] == 0 {
			nums[zero+1], nums[i] = nums[i], nums[zero+1]
			zero++
			i++
		} else if nums[i] == 2 {
			nums[two-1], nums[i] = nums[i], nums[two-1]
			two--
		}
	}
}
