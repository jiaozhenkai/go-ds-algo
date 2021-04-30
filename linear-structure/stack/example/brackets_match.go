package example

import (
	"github.com/jiaozhenkai/go-ds-algo/linear-structure/stack/arraystack"
)

func BracketsMatching(expr string) bool {
	stack := arraystack.NewStack()

	var (
		l int = len(expr)
	)

	for i := 0; i < l; i++ {
		ele := string(expr[i])
		switch ele {
		case "(", "[", "{":
			stack.Push(ele)
		case ")":
			if stack.IsEmpty() || stack.Pop() != "(" {
				return false
			}
		case "]":
			if stack.IsEmpty() || stack.Pop() != "[" {
				return false
			}
		case "}":
			if stack.IsEmpty() || stack.Pop() != "{" {
				return false
			}
		default:
		}
	}

	return stack.IsEmpty()
}
