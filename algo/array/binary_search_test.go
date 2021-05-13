package array

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		arr    []int
		n      int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				arr:    []int{1, 2, 3, 4, 5},
				n:      5,
				target: 3,
			},
			want: 2,
		},
		{
			name: "test1",
			args: args{
				arr:    []int{1, 2, 3, 4, 5},
				n:      5,
				target: 6,
			},
			want: -1,
		},
		{
			name: "test2",
			args: args{
				arr:    []int{1},
				n:      1,
				target: 1,
			},
			want: 0,
		},
		{
			name: "test3",
			args: args{
				arr:    []int{},
				n:      0,
				target: 6,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.arr, tt.args.n, tt.args.target); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearch1(t *testing.T) {
	type args struct {
		arr    []int
		n      int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				arr:    []int{1, 2, 3, 4, 5},
				n:      5,
				target: 3,
			},
			want: 2,
		},
		{
			name: "test1",
			args: args{
				arr:    []int{1, 2, 3, 4, 5},
				n:      5,
				target: 6,
			},
			want: -1,
		},
		{
			name: "test2",
			args: args{
				arr:    []int{1},
				n:      1,
				target: 1,
			},
			want: 0,
		},
		{
			name: "test3",
			args: args{
				arr:    []int{},
				n:      0,
				target: 6,
			},
			want: -1,
		},
		{
			name: "test4",
			args: args{
				arr:    []int{1, 2, 3, 4, 5, 6},
				n:      6,
				target: 2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch1(tt.args.arr, tt.args.n, tt.args.target); got != tt.want {
				t.Errorf("BinarySearch1() = %v, want %v", got, tt.want)
			}
		})
	}
}
