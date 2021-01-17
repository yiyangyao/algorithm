package basictips

import "fmt"

// slice
func stack() {
	// create stack
	stack := make([]int, 0)
	// push
	stack = append(stack, 10)
	// pop
	v := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Printf("%v", v)
	// check is empty
	if len(stack) == 0 {
		fmt.Print("true")
	}
}
