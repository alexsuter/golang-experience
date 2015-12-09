package main

import (
	"fmt"
	"time"
)

func main() {

	fruits := make(chan string, 1000)

	for i := 1; i <= 10; i++ {
		time.Sleep(time.Millisecond * 600)
		go eater(i, fruits)
	}

	go picker("Apple", fruits, 5)
	time.Sleep(time.Millisecond * 600)
	go picker("Orange", fruits, 3)
	time.Sleep(time.Millisecond * 600)
	go picker("Strawberry", fruits, 2)
	

	time.Sleep(time.Second * 360)
}

func picker(fruit string, fruits chan<- string, timeToPick int) {
	for {
		fmt.Println("Pick ", fruit)
		fruits <- fruit
		time.Sleep(time.Duration(timeToPick) * time.Second)
	}
}

func eater(i int, fruits <-chan string) {
	for fruit := range fruits {
		fmt.Println("Esser ", i, "Mhmmm ", fruit)
		time.Sleep(time.Second * 8)
	}
}
