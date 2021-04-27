package arraystack_test

import (
	"fmt"
	"github.com/jiaozhenkai/go-ds-algo/linear-structure/stack/arraystack"
	"testing"
)

func ExampleStack_IsEmpty() {
	s := arraystack.NewStack(3)
	s.Push(1)
	fmt.Println(s.IsEmpty())
	s.Pop()
	fmt.Println(s.IsEmpty())

	s1 := arraystack.NewStack(0)
	fmt.Println(s1.IsEmpty())

	// Output:
	// false
	// true
	// true
}

func ExampleStack_Len() {
	s := arraystack.NewStack(0)
	fmt.Println(s.Len())

	s1 := arraystack.NewStack(3)
	s1.Push(1)
	fmt.Println(s1.Len())

	// Output:
	// 0
	// 1
}

func ExampleStack_Pop() {
	// panic
	//s1 := stack.NewStack(0)
	//s1.Pop()

	s := arraystack.NewStack(3)
	for i := 0; i < 3; i++ {
		s.Push(i)
	}

	e := s.Pop()
	fmt.Println(e)
	e1 := s.Pop()
	fmt.Println(e1)
	e2 := s.Pop()
	fmt.Println(e2)

	// Output:
	// 2
	// 1
	// 0
}

func ExampleStack_Push() {
	// panic
	//s := stack.NewStack(0)
	//s.Push(1)

	s := arraystack.NewStack(3)
	for i := 0; i < 3; i++ {
		s.Push(i)
	}

	for i := 0; i < 3; i++ {
		fmt.Println(s.Pop())
	}

	// Output:
	// 2
	// 1
	// 0
}

func TestStack_Top(t *testing.T) {
	s := arraystack.NewStack(3)

	s.Push(1)
	s.Push(2)

	if s.Top().(int) != 2 {
		t.Errorf("get %v, want %v", s.Top(), 2)
	}
}


func TestString(t *testing.T) {
	s := arraystack.NewStack(3)
	s.Push(1)
	s.Push(2)
	fmt.Println(s)
}