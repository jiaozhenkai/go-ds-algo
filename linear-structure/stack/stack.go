package stack

type Stack struct {
	elem []interface{}
	// len int32
	top int32
}

func NewStack(cap int32) *Stack {
	return &Stack{
		elem: make([]interface{}, 0, cap),
		// len:0,
		top:0,
	}
}

func (this *Stack)Push(elem interface{}) {

}

func (this *Stack)Pop() interface{} {
	var ret  interface{} = nil
	if this.IsEmpty() {
		return nil
	}


	return  ret
}

func (this *Stack)IsEmpty() bool {
	if this.top == 0 {
		return true
	}

	return false
}

