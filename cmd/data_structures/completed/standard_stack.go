package main

import (
	"fmt"
)

type Stack struct {
	items []interface{}
}

func (s *Stack) AddToStack(input interface{}) {
	s.items = append(s.items, input)
}

func (s *Stack) RemoveFromStack() (interface{}, bool) {
	if len(s.items) == 0 {
		return nil, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}
func main() {
	var stack Stack

	stack.AddToStack("ABC")
	stack.AddToStack(123)
	stack.AddToStack(25.5)

	for len(stack.items) > 0 {
		val, boolean := stack.RemoveFromStack()
		if boolean {
			fmt.Println(val)
		}
	}

}
