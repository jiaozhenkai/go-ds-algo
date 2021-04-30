package example

import (
	"fmt"
	"github.com/jiaozhenkai/go-ds-algo/linear-structure/stack/arraystack"
)

var digit = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}

// BaseConvert 10 进制转换成任意进制 0 < base <= 16
func BaseConvert(num int64, base int64) string {
	if base > 16 || base < 2 {
		fmt.Println("base must great 2, less 16")
		return ""
	}

	stack := arraystack.NewStack()

	baseConvertRecursion(stack, num, base)

	var (
		ret string
		i   uint64
		l   uint64 = stack.Size()
	)

	for i = 0; i < l; i++ {
		e := stack.Pop().(string)
		ret += e
	}

	return ret
}

// 10 进制转换成任意进制 0 < base <= 16, 递归版本
func baseConvertRecursion(stack *arraystack.Stack, num int64, base int64) {
	if num > 0 {
		remainder := num % base
		num = num / base
		stack.Push(digit[remainder])
		baseConvertRecursion(stack, num, base)
	}
}

// 10 进制转换成任意进制 0 < base <= 16, 迭代版本
func baseConvertIter(stack *arraystack.Stack, num int64, base int64) {
	for num > 0 {
		remainder := num % base
		num = num / base
		stack.Push(digit[remainder])
	}
}
