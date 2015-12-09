package main

import (
	"fmt"
)

func main() {
	chanApple := make(chan string)

	select {
	case apple := <-chanApple:
		fmt.Println("Apple received: ", apple)
	default:
		fmt.Println("Nothing received")
	}

	select {
	case chanApple <- "apple":
		fmt.Println("Apple sent")
	default:
		fmt.Println("Nothing sent")
	}

}
