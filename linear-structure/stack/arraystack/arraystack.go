package arraystack

import (
	"fmt"
	"strings"
)

type Stack struct {
	elem []interface{}
	top  uint64 // top 指向下一个可以插入的元素位置
	cap  uint64
}

func NewStack(cap uint64) *Stack {
	return &Stack{
		elem: make([]interface{}, cap, cap),
		top:  0,
		cap:  cap,
	}
}

func (this *Stack) Push(e interface{}) {
	if this.top >= this.cap {
		panic("no enough cap to push elem")
	}

	this.elem[this.top] = e
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

func (this *Stack) Top() interface{} {
	if this.IsEmpty() {
		return nil
	}

	return this.elem[this.top-1]
}

func (this *Stack) IsEmpty() bool {
	if this.top == 0 {
		return true
	}

	return false
}

func (this *Stack) Len() uint64 {
	return this.top
}

func (this *Stack)	String() string {
	return this.toString()
}

func (this *Stack) toString() string {
	if this.IsEmpty() {
		return "[]"
	}

	s := make([]string, 0, this.Len()+1)
	s = append(s, "[")

	var i uint64
	for i=0; i<this.top; i++{
		s = append(s, fmt.Sprintf("%v ", this.elem[i]))
	}

	ret := strings.Join(s, "")
	ret = strings.Trim(ret, " ")

	return ret + "]"
}