package stack

import (
	"errors"
	"fmt"
)

type stack struct {
	slice []int
}

func Init() *stack {
	return &stack{make([]int, 0)}
}

func (stack *stack) Push(element int) {
	stack.slice = append(stack.slice, element)
}

func (stack *stack) Pop() error {
	length := len(stack.slice)

	if length == 0 {
		return errors.New("Empty Stack")
	}

	stack.slice = stack.slice[:length-1]
	return nil
}

func (stack *stack) Top() (int, error) {
	length := len(stack.slice)

	if length == 0 {
		return 0, errors.New("Empty Stack")
	}

	element := stack.slice[length-1]
	return element, nil
}

func (stack *stack) Print() {
	length := len(stack.slice)
	if length > 0 {
		fmt.Println(stack.slice)
	} else {
		fmt.Println("Empty Stack")
	}
}

func (stack *stack) Size() int {
	return len(stack.slice)
}

func (stack *stack) IsEmpty() bool {
	return len(stack.slice) == 0
}


