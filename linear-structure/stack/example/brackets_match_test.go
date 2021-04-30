package example

import "testing"

func TestBracketsMatching(t *testing.T) {
	type args struct {
		expr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name:"test1",
			args:args{expr: "a+b"},
			want: true,
		},
		{
			name:"test2",
			args:args{expr: "(a+b"},
			want: false,
		},
		{
			name:"test3",
			args:args{expr: "a+b)"},
			want: false,
		},
		{
			name:"test4",
			args:args{expr: "(a+b)"},
			want: true,
		},
		{
			name:"test5",
			args:args{expr: "(a+b)*c-((c-d)+{c*d})"},
			want: true,
		},
		{
			name:"test6",
			args:args{expr: "(a+b)*c-(c-d)+{c*d})"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BracketsMatching(tt.args.expr); got != tt.want {
				t.Errorf("BracketsMatching() = %v, want %v", got, tt.want)
			}
		})
	}
}
