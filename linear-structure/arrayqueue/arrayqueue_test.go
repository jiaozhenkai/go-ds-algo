package arrayqueue_test

import (
	"fmt"
	"testing"

	aq "github.com/jiaozhenkai/go-ds-algo/linear-structure/arrayqueue"
)

func Example() {
	q := aq.NewArrayQueue(5)
	for i := 0; i < 5; i++ {
		q.EnQueue(i)
	}
	q.ToString()

	q.DeQueue()
	q.ToString()

	fmt.Println(q.Len())
	fmt.Println(q.IsEmpty())

	// Output:
	// [ 0 1 2 3 4 ]
	// [ 1 2 3 4 ]
	// 4
	// false
}

func Test_1(t *testing.T) {
	q := aq.NewArrayQueue(3)
	for i := 0; i < 3; i++ {
		q.EnQueue(i)
	}
	q.DeQueue()
	q.DeQueue()
	q.ToString()
	fmt.Println(q.Front(), q.Back())

	q2 := aq.NewArrayQueue(0)
	q2.ToString()

	q1 := aq.NewArrayQueue(3)
	for i := 3; i > 0; i-- {
		q1.EnQueue(i)
	}
	q1.ToString()
	fmt.Println(q1.Front(), q1.Back(), q1.Len())
	fmt.Println(q1)
}
