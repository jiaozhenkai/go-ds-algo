package array

// 左闭右闭区间
func BinarySearch(arr []int, n int, target int) int {
	if n == 0 {
		return -1
	}

	// 定义闭区间 [left, right]，那么之后就是在这个闭区间中找 target
	left := 0
	right := n - 1

	for left <= right { // left == right 表示数组中还有一个元素
		mid := left + (right-left)>>1
		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target { // target 更大，要去右边找
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// 左闭右开区间
func BinarySearch1(arr []int, n int, target int) int {
	if n == 0 {
		return -1
	}

	// 定义闭区间 [left, right)，那么之后就是在这个闭区间中找 target
	left := 0
	right := n

	for left < right { // left == right-1 时表示数组中还有一个元素
		// mid := left + (right-left)>>1
		mid := left + (right-left-1)>>1
		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target { // target 更大，要去右边找
			left = mid + 1
		} else {
			right = mid
		}
	}
	return -1
}
