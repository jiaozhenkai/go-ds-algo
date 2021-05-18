package array

import (
	"fmt"
	"testing"
)

func Test_sortColors(t *testing.T) {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}

func Test_sortColors1(t *testing.T) {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors1(nums)
	fmt.Println(nums)
}
