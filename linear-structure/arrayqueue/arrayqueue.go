package arrayqueue

import "fmt"

type ArrayQueue struct {
	elem  []interface{}
	len   uint64
	cap   uint64
	front uint64
	back  uint64
}

func NewArrayQueue(cap uint64) *ArrayQueue {
	return &ArrayQueue{
		elem:  make([]interface{}, cap+1, cap+1),
		len:   0,
		cap:   cap + 1,
		front: 0,
		back:  0,
	}
}

// Front return first element of queue, if queue is empty return nil
func (this *ArrayQueue) Front() interface{} {
	if this.IsEmpty() {
		return nil
	}

	ret := this.elem[this.front]
	return ret
}

// Back return last element of queue, if queue is empty return nil
func (this *ArrayQueue) Back() interface{} {
	if this.IsEmpty() {
		return nil
	}

	if this.back != 0 {
		return this.elem[this.back-1]
	}

	ret := this.elem[this.cap-1]
	return ret
}

// Enqueue put elem to queue tail
func (this *ArrayQueue) EnQueue(e interface{}) {
	if this.IsFull() {
		panic("no enough cap to push elem")
	}

	this.elem[this.back] = e
	this.len++
	this.back = (this.back + 1) % this.cap
}

// DeQueue delete elem from queue head
func (this *ArrayQueue) DeQueue() interface{} {
	if this.IsEmpty() {
		panic("pop elem form empty queue")
	}

	ret := this.elem[this.front]
	this.front = (this.front + 1) % this.cap
	this.len--

	return ret
}

func (this *ArrayQueue) Len() uint64 {
	return this.len
}

func (this *ArrayQueue) IsEmpty() bool {
	if this.front == this.back {
		return true
	}
	return false
}

func (this *ArrayQueue) IsFull() bool {
	if ((this.back + 1) % this.cap) == this.front {
		return true
	}

	return false
}

// ToString print queue elem from front to back
func (this *ArrayQueue) ToString() {
	if this.IsEmpty() {
		fmt.Println("[ ]")
		return
	}

	var i uint64 = this.front
	fmt.Printf("%v", "[ ")
	for i != this.back {
		fmt.Printf("%v ", this.elem[i])
		i = (i + 1) % this.cap
	}

	fmt.Println("]")
}
