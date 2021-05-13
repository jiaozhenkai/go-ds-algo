package array

import (
	"fmt"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes1(nums)
	fmt.Println(nums)

	nums1 := []int{2, 1}
	moveZeroes1(nums1)
	fmt.Println(nums1)

}
