package main

import (
	"stack"
	"fmt"
)

func main() {
	myStack := Init()

	fmt.Println("Size =", myStack.Size())
	fmt.Println("IsEmpty =", myStack.IsEmpty())
	myStack.Print()
	myStack.Top()

	myStack.Push(42)
	myStack.Push(77)
	myStack.Push(1)
	fmt.Println("Size =", myStack.Size())
	fmt.Println("IsEmpty =", myStack.IsEmpty())
	myStack.Print()
}
