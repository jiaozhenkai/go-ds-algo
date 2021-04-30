package example

import (
	"fmt"
	"github.com/jiaozhenkai/go-ds-algo/linear-structure/stack/arraystack"
	"testing"
)

func TestBaseConvert(t *testing.T) {
	type args struct {
		num  int64
		base int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{num: 23, base: 8},
			want: "27",
		},
		{
			name: "test2",
			args: args{num: 23, base: 16},
			want: "17",
		},
		{
			name: "test3",
			args: args{num: 180, base: 16},
			want: "B4",
		},
		{
			name: "test4",
			args: args{num: 180, base: 8},
			want: "264",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BaseConvert(tt.args.num, tt.args.base); got != tt.want {
				t.Errorf("BaseConvert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Example_baseConvertIter() {
	stack := arraystack.NewStack()
	type tests struct {
		num  int64
		base int64
	}

	exp := []tests{{num: 180, base: 8}, {num: 180, base: 16}, {num: 23, base: 8}, {num: 23, base: 16}}

	for _, v := range exp {
		baseConvertIter(stack, v.num, v.base)
		var i uint64
		var ret string
		l := stack.Size()
		for i = 0; i < l; i++ {
			e := stack.Pop().(string)
			ret += e
		}
		fmt.Println(ret)
	}

	// Output:
	// 264
	// B4
	// 27
	// 17

}
