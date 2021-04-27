package arrayqueue_test

import (
	"fmt"
	"github.com/jiaozhenkai/go-ds-algo/linear-structure/queue/arrayqueue"
	"strconv"
	"testing"
)

func Example() {
	q := arrayqueue.NewArrayQueue(5)
	for i := 0; i < 5; i++ {
		q.EnQueue(i)
	}
	fmt.Println(q)

	q.DeQueue()
	fmt.Println(q)

	fmt.Println(q.Len())
	fmt.Println(q.IsEmpty())

	// Output:
	// [0 1 2 3 4]
	// [1 2 3 4]
	// 4
	// false
}

func Test_1(t *testing.T) {
	q := arrayqueue.NewArrayQueue(3)
	for i := 0; i < 3; i++ {
		q.EnQueue(i)
	}
	q.DeQueue()
	q.DeQueue()
	fmt.Println(q.Front(), q.Back())

	q2 := arrayqueue.NewArrayQueue(0)
	fmt.Println(q2)

	q1 := arrayqueue.NewArrayQueue(3)
	for i := 3; i > 0; i-- {
		q1.EnQueue(i)
	}
	fmt.Println(q1.Front(), q1.Back(), q1.Len())
	fmt.Println(q1)
}

// test String
func Test_2 (t *testing.T) {
	q := arrayqueue.NewArrayQueue(3)
	for i := 0; i < 3; i++ {
		q.EnQueue(i)
	}
	fmt.Println(q)

	type person struct {
		 name string
	}
	q1 := arrayqueue.NewArrayQueue(3)
	for i := 0; i < 3; i++ {
		q1.EnQueue(person{name:strconv.Itoa(i)})
	}
	fmt.Println(q1)

	q2 := arrayqueue.NewArrayQueue(3)
	for i := 0; i < 3; i++ {
		q2.EnQueue([]int{i})
	}
	fmt.Println(q2)
}