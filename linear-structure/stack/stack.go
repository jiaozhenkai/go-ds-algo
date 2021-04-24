package stack

import (
	"fmt"
	"math"
)

type Stack struct {
	elem []interface{}
	top  uint32 // top 指向下一个可以插入的元素位置
	cap  uint32
}

const maxCap = math.MaxUint32

func NewStack(cap uint32) *Stack {
	if cap > maxCap {
		fmt.Println("stack cap: 0 <= ca p<= math.MaxInt32")
		return nil
	}

	return &Stack{
		elem: make([]interface{}, 0, cap),
		top:  0,
		cap:  cap,
	}
}

func (this *Stack) Push(e interface{}) {
	if this.top == this.cap {
		panic("no enough cap to push elem")
	}

	this.elem = append(this.elem, e)
	this.top++
}

func (this *Stack) Pop() interface{} {
	if this.IsEmpty() {
		panic("pop elem form empty stack")
	}

	ret := this.elem[this.top-1]
	this.top--

	return ret
}

func (this *Stack) IsEmpty() bool {
	if this.top == 0 {
		return true
	}

	return false
}

func (this *Stack) Len() uint32 {
	return this.top
}
