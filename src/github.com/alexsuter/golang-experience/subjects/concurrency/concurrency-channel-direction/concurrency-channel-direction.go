package main

import (
  "fmt"
  "time"
)

func main() {
    fruits := make(chan string)

	go produce(fruits, "Apple")
	go consume(fruits)
	go consume(fruits)
	go produce(fruits, "Banana")
	go produce(fruits, "Orange")
	go produce(fruits, "Mango")
	time.Sleep(1 * time.Second)
	go consume(fruits)
	go consume(fruits)

	time.Sleep(2 * time.Second)
}

func produce(fruits chan<- string, fruit string) {
    fmt.Println("Send ", fruit)
    fruits <- fruit
}

func consume(fruits <-chan string) {
    fruit := <- fruits
    fmt.Println("Recive ", fruit)
}