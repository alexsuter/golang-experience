package main

import (
	"fmt"
	"time"
)

func main() {
	chanApple := make(chan string)
	chanBanana := make(chan string)

	go func() {
		for {
			time.Sleep(time.Second * 1)
			chanApple <- "Apple"
		}
	}()
	go func() {
		for {
			time.Sleep(time.Second * 2)
			chanBanana <- "Banana"
		}
	}()

	for {
		select {
		case apple := <-chanApple:
			fmt.Println("Fruit: ", apple)
		case banana := <-chanBanana:
			fmt.Println("Fruit: ", banana)
		}
	}
}